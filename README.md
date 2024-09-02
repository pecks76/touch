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

* Flyway sometimes fails to connect to MySQL when running in Docker
  - think it needs longer timeout / wait time

* Error handling could be better

#Decisions

* clients are only an id - real world would have more info
* there's no validation on the JSON incoming, there should be

* haven't used gin or router for now, may change later - just using net/http for time being
* directory structure, was domain but changed to by function

* reading an object returns if object doesn't exist, but with id of zero (as opposed to returning an error)

* the investments are ints, currency not included for now.

* split the structure, by receipt (which is, incoming messages) and account (real world)
  ** this determines the data structure, wanted to think about client - pots - accounts i.e. what we need to end up with
  ** then client links to deposit, and receipts are recorded as instructions

* separate dir for messages, doesn't want to be in domain

* use DI to make testing easier
* used mockgen for mocking

* generally try to log errors and keep going, rather than bailing out

