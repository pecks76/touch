#HOW-TO

##Run tests: 
* *go test ./...* 

##Run in Docker: 
* *docker compose up --build*

##Run locally 
* need a MySQL database running on 3306 
* need to create the schema using the script(s) in /sql directory 
* create user as per connection string in connection.go 
* adjust connection string accordingly (i.e. run on localhost)
* *go run main.go*

##Generate mocks: 
add a line to the thing you want mocked, e.g: 
- //go:generate mockgen -source=account_persistence.go -destination=mock/account_persistence.go
* *go generate -v ./service/receipt*

#TODOs
* Testing is not complete
    - would be nice to have some REST level
    - some services missing

* Database connection needs to change if you run locally
    - find a better way to do this

* Flyway sometimes fails to connect to MySQL
    - think it needs longer timeout / wait time 
  
* Error handling could be better 