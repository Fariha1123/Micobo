# Micobo


# Run Server 
### The server is connected with psql so it needs database connection parameters (DBNAME, DBPASS, DBUSER) so you would have to provide it as environment variable as below:
#### export DBNAME="database name"
#### export DBPASS="database password"
#### export DBUSER="database username" 
#### go run `ls *.go | grep -v _test.go`  --- this will run the server at port 8080 and you will be good to send requests from postman


# RUN test files
### Command: go test -v
