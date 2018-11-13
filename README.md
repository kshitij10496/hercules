# hercules

**hercules**, the Roman God of Power, will provide you with the API for information related to IIT Kharagpur's academic life.

![artwork](images/artwork.jpg "Hide of the Nemean lion worn by Hercules")

# Motivation

- No more ad-hoc scrapping code.
- No more JSON files in projects.
- Single source of truth for all the projects.
- Building stuff should be a fun and creative process and shouldn't be hampered by the difficulty in the procurement of data.

# Implementation

- I've tried to use microservices as extensively as possible without trying to enfore the design pattern.
Each service is responsible for a logical data entity and the corresponding table in the DB.  

- Inter service communication is currently carried over HTTP via `SendToService` function.  

- The `common` package provides the minimal scafolding which should be share across all the services.
Essentially, it houses the data types, custom errors, database queries and the `server` interface.
  
Current List of Endpoints:

### service-course
- [x] [`/course/timetable/{code}`](https://hercules-10496.herokuapp.com/api/v1/course/timetable/MA61023): Timetable of a course given the course code.
- [x] [`/course/info/department/{code}`](https://hercules-10496.herokuapp.com/api/v1/course/info/department/MA): List of all the courses offered by a department.
- [x] [`/course/info/faculty`](https://hercules-10496.herokuapp.com/api/v1/course/info/faculty?name=Pratima%20Panigrahi&dept=MA): List of all the courses offered by a faculty member.

### service-department
- [x] [`/department/info/all`](https://hercules-10496.herokuapp.com/api/v1/department/info/all): List of all the departments.

### service-faculty
- [x] [`/faculty/info/all`](https://hercules-10496.herokuapp.com/api/v1/faculty/info/all): List all the faculty members at IIT Kharagpur.
- [x] [`/faculty/info/{code}`](https://hercules-10496.herokuapp.com/api/v1/faculty/info/MA): List all the faculty members of a particular department.
- [x] [`/faculty/timetable`](https://hercules-10496.herokuapp.com/api/v1/faculty/timetable?name=Pratima%20Panigrahi&dept=MA): Timetable of a faculty member


## Roadmap

After laying down the basic infrastructure for the project, I plan to develop the API on a *needs* basis.   
This would mean that new endpoints would be developed on based on the use-cases around the problem statement.  
For the time being, I have a shortlist of projects which I think would benefit the most from the API provided by `hercules`.  

October 2019: Port `wimp` to use `hercules`.   
November 2019: Port `mcmp` to `hercules`.  
December 2019: Port `Kronos` to `hercules`.  
December 2019: Port `gyft` to `hercules`.  

### Contributing

I have commented a lot of possible improvements to the API as `TODO` in the code which can be good place to start.   
At the same time, new/parallel feature requests are welcomed and I would try my best to accomodate time to discuss them through.

Apart from contributing LoC, you can help us by discussing some of the techincal problems we are currently facing:
- [ ] Keep the DB updated: Schedule the scrapping of data from the source links and triage.
- [ ] Testing handlers, unit testing models and functional testing services.
- [ ] Deploy services after writing `docker-compose`.


### Local environment setup

You can develop on and test hercules using Docker. Make sure that docker and
docker-compose are installed on your computer before following the steps below.

1. Start the app and DB containers

    ```sh
    cd dev-env
    docker-compose up --build -d
    ```

2. Restore the initial data to your DB container's database

    ```sh
    cd dev-env/init-data
    ./restore_init_data.sh
    ```

Now, you should be able to access the API. You can test this by visiting:
`http://localhost:8070/api/v1/course/timetable/ME30005` or running:

```sh
curl http://localhost:8070/api/v1/course/timetable/ME30005
```

This setup uses [fresh](https://github.com/pilu/fresh) to rebuild the server
binary and restart the server whenever files are edited on the host machine. So,
you can edit the files on your host machine and save them. Wait for a while, and
your new code will be built and served at the same location.

You can look at the app logs using `docker logs -f devenv_hercules_api_1`.

In order to know the requirements and setting up the project and database locally, visit [Development Workflow wiki](https://github.com/kshitij10496/hercules/wiki/Development-Workflow)
