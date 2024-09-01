##DECISIONS

* clients are only an id - real world would have more info
* there's no validation on the JSON incoming, there should be

* haven't used gin or router for now, may change later - just using net/http for time being 
* directory structure, was domain but changed to by function

* the investments are ints, currency not included for now.

* split the structure, by receipt (which is, incoming messages) and account (real world)
** this determines the data structure, wanted to think about client - pots - accounts i.e. what we need to end up with
** then client links to deposit, and receipts are recorded as instructions

* separate dir for messages, doesn't want to be in domain

* use DI to make testing easier
* used mockgen for mocking





