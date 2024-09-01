##HOW-TO

###Run tests: 
go test ./...

###Run in docker: 
docker compose up --build

##TODOs

* Testing is not complete
    - would be nice to have some REST level

* Database connection needs to change if you run locally
    - find a better way to do this

* Flyway sometimes fails to connect to MySQL
    - think it needs longer timeout / wait time 