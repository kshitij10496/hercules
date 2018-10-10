import json
import requests
from requests_html import HTML, HTMLSession

URL = 'http://www.iitkgp.ac.in/facultylist'
RECORDS = 674

def main():
    session = HTMLSession()
    try:
        res = session.get(URL)    
    except requests.ConnectionError:
        print("failed to connect")

    if res.status_code != 200:
        print("Unable to fetch data")
        return

    data = res.json()

    # Check for updates
    # Remember to update RECORDS manually.
    if data['recordsTotal'] != RECORDS:
        print("New faculty memebers have joined IIT KGP")

    # Parse the 'faculty' key of each object to obtain the name.
    faculty = data['data']
    for i, f in enumerate(faculty):
        html = HTML(html=f['faculty'])
        f['faculty'] = html.text
        f['code'] = f.pop('dept_code')

    # Write the final data to JSON file.
    # TODO: Use absolute path here
    with open('../faculty_directory.json', 'w') as f:
        json.dump(faculty, f)

if __name__ == '__main__':
    main()