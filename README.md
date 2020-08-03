# Golang Restful API using GORM ORM (MySQL) Gorilla Mux

## Getting Started

### Folder Structure
This is my folder structure under my `$GOPATH` or `$HOME/your_username/go`.
```
.
|-- bin
+-- src
|   +-- github.com
|   |   +-- gamorvi
|   |   |   +-- restapi
|   |   |   |   |-- .env
|   |   |   |   |-- main.go
|   |   |   |   +-- controllers
|   |   |   |   |   |-- usersController.go
|   |   |   |   +-- models
|   |   |   |   |   |-- base.go
|   |   |   |   |   |-- user.go
|   |   |   |   +-- routes
|   |   |   |   |   |-- api.go
|   |   |   |   +-- utils
|   |   |   |   |   |-- utils.go
```
Ensure you create the `gamorvi` directory in your `github.com` directory. `cd` into the `gamorvi` directory before `git clone https://github.com/gamorvi/restful-api-with-golang.git`

## Download the packages used to create this rest API
Run the following Golang commands to install all the necessary packages. These packages will  help you set up a web server, ORM for interacting with your db, mysql driver for db connection and load your .env file or setting up environment variables respectively.
```
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
go get -u github.com/joho/godotenv
```
### Running documentation locally (Only documentation of packages your have installed)
For offline documentation on the following packages run `godoc -http :6060` and then visit `http://localhost:6060`. Note that you can change the port to your preferred port number.

## Setting configuration file
Create a .env file in the root of the project and set the following parameters

```
db_name = database_name     # Name of database
db_user = user              # Database username
db_pass = secret            # Database password
db_type = mysql             # MySQL driver
db_host = localhost         # Database host
db_port = 3306              # Database port
charset = utf8              # Database charset
parse_time = True           # Database parse time
web_port = 8085             # Port to serve api
prefix = /api/v1            # API route sub route prefix
```

## Database Table Creation Statement
Use the following DDL (Data Definition Language) to create the users table.

``` SQL
CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;
```
## Running the project

`go run *.go`

## API Endpoints & Usage

* GET `api/v1/users` retrieve all users
* GET `api/v1/users/1` retrieve user with id = 1
* POST `api/v1/users` create a new user
* PUT `api/v1/users/1` update the record with id = 1
* DELETE `api/v1/users/1` delete the user with id = 1

### To create a new user

1. POST `api/v1/users`

```
{
	"Name": "Joe Bloke",
	"Email": "joe.bloke@fake-domain.com"
}
```

*** Output ***

```
{
    "message": "success",
    "status": true,
    "user": {
        "ID": 1,
        "CreatedAt": "2019-05-06T00:54:22.09382+01:00",
        "UpdatedAt": "2019-05-06T00:54:22.09382+01:00",
        "DeletedAt": null,
        "Name": "Joe Bloke",
        "Email": "joe.bloke@fake-domain.com"
    }
}
```

2. Get `api/v1/users/1`

*** Output ***

```
{
    "data": {
        "ID": 1,
        "CreatedAt": "2019-05-05T23:54:22Z",
        "UpdatedAt": "2019-05-05T23:54:22Z",
        "DeletedAt": null,
        "Name": "Joe Bloke",
        "Email": "joe.bloke@fake-domain.com"
    },
    "message": "success",
    "status": true
}
```
