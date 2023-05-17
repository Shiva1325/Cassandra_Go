# Cassandra_Go

-> This Project is implemented in GoLang and uses grpc.
-> The model is built for 3 server architecture and assumes that the server is always reliable.
-> In order to test Please comment and uncomment some lines in main.go file and and doing this will automatically loads the test data that is present in Test folder. Each of these files contain 30,000 records of data. Which part to comment is clearly stated in main.go file.

--> First please use "make server_main" in order to build server_main.go build file.
--> Then please use "make client_main" in order to build client_main.go build file.

------> For starting server:


--> In order to start all servers atonce: please give command "make server"
--> In order to start server1: please give command "make server1"
--> In order to start server2: please give command "make server2"
--> In order to start server3: please give command "make server3"


------> For starting client: 

--> First use "make client" command to start client side
--> use "c 0.0.0.0:1200x" {where x= 1, 2, 3} for connecting to server 1, 2 or 3
--> use "a" to write data
--> use "g" to get data

------> For get and set operations

--> if you want to write data: please give a name and country as prompted. country data is passed to hash function which generates partition key and based on this partition key, the record in save in a particular node.

--> if you want to get data: please give the country name and based on this key all the users with the given country will be printed.


-------> For using test data: Please comment some lines highlighted in main.go file


----> when made any changes in client side code, please build the code client_main.go by using command "make client_main" and then start client.

----> When made changes on server side, please build the server using "make server_main" and start all the servers again.