# Demo

To begin, launch the server by running `go run ./cmd/srv/...` (this will initiate a server listening on port 8080).

Next, to run the tests, simply execute the `make` command.

The `make` command operates in two phases. 

Initially, it conducts a test run that records the HTTP interactions with the server. 
Please note that deliberate 5-second sleeps are inserted for each HTTP call on the server during this phase.

Subsequently, the command re-runs the tests utilizing the recorded interactions.
