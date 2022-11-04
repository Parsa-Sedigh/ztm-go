# 09 - Project MailingList Microservice

## 103-001 Project Setup
After gcc and protobuf compiler(protoc) are installed, we have to install protobuff codegen tools.

To initialize the project, run in the `mailinglist` directory:
```shell
go mod init mailinglist
```
which is gonna create the go.mod file

## 104-002 Creating Database Tables
We're gonna create a database layer. Create `mdb` folder short for mail database.

Now that we've written functionality to create the DB tables and to convert query data into a go data structure, we can move on to implementing DB CRUD operations.

## 105-003 Implementing CRUD Operations
With db.Exec() , we don't need to run any Close() operation because it just runs at one time and it's finished. However, with db.Query() , it keeps it open so we can
keep reading more rows. So we need to do rows.Close() once the function completes, that way it frees the DB which can be done with `defer` keyword.
## 106-004 JSON Data Processing Functions
Create `jsonapi` directory. Before written our handler functions for the API, we need to create a few support functions for transforming data and for returning 
appropriate headers.
## 107-005 JSON Endpoints
## 108-006 Server CLI
We're gonna create a CLI app to run our server. Create `server` directory.

When you're creating microservices, it helps a lot to read configuration from the environment. Environment files are easy to check into source control and easy to replicate
across multiple servers.

Run:
```shell
go mod tidy
```
to download deps that you're using.

Then run:
```shell
go run ./server
```
in the project's root directory.

## 109-007 Testing JSON API
You can use thunder client vscode extension to test the api.

If we don't see an ip address when our server in this case is running, it means our ip is: 127.0.0.1 and port is 8080(in this case!)

When sending req, ensure this header is set: `Content-Type: application/json`.

## 110-008 Protocol Buffers
gRPC is a way to do remote procedure calls and it uses protocol buffers.

Create a folder named `proto`.

After writing the proto file, run the protoc command in the root of the project and that command is written in README file of `mailinglist` project.

The generated code provide functions for creating a gRPC server using the messages that we specified. 
Now we need to create the grpc api server.

## 111-009 gRPC Data Processing Functions
Create `grpcapi` folder

## 112-010 gRPC Endpoints
Now that we have data conversion functions, we can start creating endpoints for our grpc api.
## 113-011 gRPC Client
Create a folder called `client`. The client is gonna be it's own application, so we use `package main` for it.

To send the reqs, first run the server in one terminal window and in another one, run the client app:
```shell
go run ./server
go run ./client
```

If you updated the code on server, rerun the `go run ./server` command.