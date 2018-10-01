# hercules

# Motivation

- No more scrapping code.
- No more JSON files in projects.
- Single source of truth for all the projects.
- Building stuff should be a fun and creative process and shouldn't be hampered by the difficulty in procurement of resources.

hercules, the Greek God of Power, will provide you with the API for information related to IIT Kharagpur's academic life.

# Implementation

- [ ] [Discussion]: Decide whether to use Relational/Non-relational DB here.
- [ ] [Need Advice]: Keep the DB updated: Schedule the scrapping of data from the source links and triage.
  
Current List of Endpoints:

- [x] `/departments/info`: List of all the departments.
- [x] `/courses/info`: List of all the courses.
- [x] `/faculty/info`: List of all the faculty members.
- [x] `/faculty/timetable`: List of a faculty member's daily academic timetable.

## Roadmap


After laying down the basic infrastructure for the project, I want to develop the API
on a needs basis. This would mean that new endpoints would be developed on based on the use-case.
A lot of project hopping would take place to remove the previous ad-hoc methods.
At the same time, new/parallel feature requests are welcomed and I would try my best to accomodate time to discuss them through.

October 2019: Port `wimp` to use `hercules`.
November 2019: Port `Kronos` to `hercules`.
December 2019: Port `gyft` to `hercules`.
