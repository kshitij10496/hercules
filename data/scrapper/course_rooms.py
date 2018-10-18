import sys
import json
import requests
import argparse

from requests_html import HTML, HTMLSession

URL = "https://erp.iitkgp.ac.in/Acad/timetable_track.jsp?action=second&dept={0}"

# TODO: Create a URL Builder rathen than using string formatting
# Add query string parameters
#   - action = second
#   - dept = <dept_code>

# Add Form data
#   - dept = <dept_code>
#   - for_semester = AUTUMN/SPRING
#   - for_session = 2018-2019
#

# A global variable to improve the logging of missed courses.
# TODO: Find a better logging solution
COUNTER = 0

class Course:
    '''
    code, string    : Course code, possibly, these are unique.
    name, string    : Name of the course.
    profs, list     : List of profs taking this course.
    credits, int    : Credits associated.
    slots, list     : List of allocated slots.
    rooms, list     : List of alloted rooms.
    '''

    def __init__(self, code, name, profs, credits, slots, rooms):
        self.code = code
        self.name = name
        self.profs = profs
        self.credits = credits
        self.slots = slots
        self.rooms = rooms
    
    def __str__(self):
        return 'Course(code={!r}, name={!r}, profs={!r}, credits={!r}, slots={!r}, rooms={!r})'.format(self.code, self.name, self.profs, self.credits, self.slots, self.rooms)

class DepartmentCourses:
    def __init__(self, dept, courses):
        self.dept = dept
        self.courses = courses

    def __str__(self):
        return 'DepartmentCourses(dept={0}, courses={1})'.format(self.dept, self.courses)

class CourseEncoder(json.JSONEncoder):
    def default(self, o):
        if isinstance(o, DepartmentCourses):
            return o.__dict__

        if isinstance(o, Course):
            # TODO: Prettify JSON encoding by adding indentation and new lines.
            return o.__dict__
        raise TypeError("Object of type '{type} is not JSON Serializable".format(o.__class__.__name__))

# This function parses a table row (list of cells) and returns a new course.
def parse_table_row(cells):
    course_code = cells[0].text
    course_name = cells[1].text

    profs = cells[2].text.split(",")
    # remove whitespaces, if any
    # ensure no duplicates because I accept that people make mistaeks.
    # ensure that the name is in standard form
    course_profs = list(set(prof.strip().title() for prof in profs))

    course_credits = int(cells[4].text)

    # remove whitespaces, if any
    # ensure no duplicates so that someone's lack of attention doesn't screw my code.
    slots= (cells[5].text).split(",")
    course_slots = list(set(slot.strip() for slot in slots)) 

    course_rooms = []
    if len(cells) == 7:
        # remove whitespaces, if any
        # ensure no duplicates since I prefer minimalism.
        rooms = (cells[6].text).split(",")
        course_rooms = list(set(room.strip() for room in rooms))

    course = Course(
                    code=course_code,
                    name=course_name,
                    profs=course_profs,
                    credits=course_credits,
                    slots=course_slots,
                    rooms=course_rooms
                )
    return course

# Given a department code and the session_id from ERP, it returns a list of
# department subjects.
def department_subjects_list(dept_code, session_id):
    print("FETCHING DEPARTMENT SUBJECTS LIST: ", dept_code)
    # Add department code to the base URL
    url = URL.format(dept_code)
    cookies = {"JSESSIONID": session_id}

    session = HTMLSession()
    try:
        res = session.post(url, cookies=cookies)
    except requests.ConnectionError:
        print("Failed to connect")

    if res.status_code != 200:
        print("Unable to fetch data")
        return

    data = res.html

    # The table containing the information is at 4th index.
    tables = data.find("table")
    if tables:
        table = tables[4]

    table_rows = table.find('tr')
    courses = []
    for row in table_rows:
        cells = row.find('td')
        # Logging table headers and some courses which are not in standard format.
        # 2018 Autumn: 13 courses are logged.
        if len(cells) < 6:
            global COUNTER
            COUNTER += 1
            print(COUNTER, '\t'.join([cell.text for cell in cells]))
            continue

        course = parse_table_row(cells)
        courses.append(course)

    print("TOTAL COURSES: ", len(courses))
    return courses


def main():
    # 1. Obtain all department codes
    # 2. Get individual department courses
    # 3. Concatenate all the courses
    # 4. Encode the resulting data into a JSON
    # 5. Store it in a JSON file
    parser = argparse.ArgumentParser(description='Scrape department timetables')

    parser.add_argument("-i", "--input-file", dest="input_file", required=True,
            help="path to the input file containing department codes")

    parser.add_argument("-o", "--output-file", dest="output_file", required=True,
            help="path to the output json file")

    parser.add_argument("--jsession-id", dest="jsession_id", required=True,
            help="JSESSIONID extracted from the ERP")

    args = parser.parse_args()

    JSESSIONID = args.jsession_id
    INPUT_FILE = args.input_file
    OUTPUT_FILE = args.output_file

    departments = []
    with open(INPUT_FILE, "r") as f:
        for line in f.readlines():
            if line.strip().startswith("#"):
                continue
            department = line.strip('\n')
            departments.append(department)

    print(departments)

    all_courses = []
    for dept in departments:
        courses = department_subjects_list(dept, JSESSIONID)
        if courses is None:
            print("Cannot load courses for {0}".format(dept))
            continue
        # TODO: Find an efficient way to do this using itertools, maybe
        dept_courses = DepartmentCourses(dept, courses)
        all_courses.append(dept_courses)

    print("TOTAL COURSES:", len(all_courses))
    # Encode data and store it in a JSON file
    with open(OUTPUT_FILE, 'w') as f:
        json.dump(all_courses, f, cls=CourseEncoder)
    
if __name__ == "__main__":
    main()
