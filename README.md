# Coding Sample - Anthony Anderson

This is my coding sample writing a simple REST server in golang. There is a front end application that can be ran to use a search bar to display results.


## Packages and Technologies needed

This program was created and ran through VScode so if using VScode you will need the Go extension

* Needs a gcc compiler installed to use the sqlite-3 driver for querying the database
* Need to install a couple packages:
    * Use command: "go get github.com/gorilla/mux"
    * Use command: "go get github.com/mattn/go-sqlite3"
* To start the server, use the command: "go run server.go" within the project directory or use the full path for the server.go file if not in the directory
* To view the front end, you need something to run the html file, i.e. Live server in VScode
* The pregenerated chinook database file is within the project