# Instant Messaging System

This is a simple instant messaging system implemented in Go (Golang) using an HTTP server and an RPC server. The system allows users to send and receive messages by using the Pull API.

## Features

- HTTP server for handling message retrieval via a RESTful API.
- RPC server for pulling messages via remote procedure calls.
- Messages are stored in a MySQL database.

## Prerequisites

Before running the messaging system, make sure you have the following dependencies installed:

- Go (Golang): [https://golang.org/doc/install](https://golang.org/doc/install)
- MySQL: [https://dev.mysql.com/downloads/installer/](https://dev.mysql.com/downloads/installer/)

## Installation

1. Clone the repository:

```shell
git clone https://github.com/your-username/instant-messaging-system.git
cd instant-messaging-system
```

2. Set up the MySQL database:
- Start the MySQL server.
- Create a new database using the MySQL client. For example:
```shell
mysql -u your-username -p
CREATE DATABASE InstantMessagingDB;
```

3. Configure the database connection:
- Open the httpserver/main.go and rpcserver/main.go files.
- In each file, update the MySQL connection string with your MySQL username, password, and database details:
```go
db, err := sql.Open("mysql", "your-username:your-password@tcp(localhost:3306)/InstantMessagingDB")
```

4. Build and run the servers:
- HTTP server:
```shell   
cd httpserver
go build
./httpserver
```
- RPC server:
```shell
cd rpcserver
go build
./rpcserver
```

5. The HTTP server will be running on http://localhost:8080 and the RPC server on localhost:1234. You can now use the Pull API to fetch messages from the servers.

## API Endpoints

### HTTP Server
- GET /messages: Retrieve all messages stored in the database.
### RPC Server
- PullMessages(): Fetch messages via remote procedure calls (RPC).