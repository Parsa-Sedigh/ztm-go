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
## 107-005 JSON Endpoints
## 108-006 Server CLI
## 109-007 Testing JSON API
## 110-008 Protocol Buffers
## 111-009 gRPC Data Processing Functions
## 112-010 gRPC Endpoints
## 113-011 gRPC Client