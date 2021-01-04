# GOLANG

step1 : 

    1) run mongo

    a)  using docker image directly 
    docker run --name mongodb -p 27017:27017 mongo 

    b) using yml-file
    docker-compose -f mongo.yml up -d


# notes : 
just uncomment functions in main functions to toggle for diffrent functionality


# Mongo-Express
http://localhost:8081/

# Link
https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/

Bson and crud

https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial-part-1-connecting-using-bson-and-crud-operations


https://github.com/tfogo/mongodb-go-tutorial/blob/master/main.go
https://github.com/ruanbekker/code-examples/blob/master/mongodb/golang/examples.go
https://github.com/do-community/tasker

complex example

https://github.com/rosspatil/Golang-Mongo-Driver-Example/blob/master/mongo-exmple.go

# Imp

It is best practice to keep a client that is connected to MongoDB around so that the application can make use of connection pooling - you don't want to open and close a connection for each query. However, if your application no longer requires a connection, the connection can be closed with client.Disconnect() like so


