
#  Golang Restful API Messaging using GORM ORM (MySQL) Gorilla Mux

  

##  Getting Started

 
###  Folder Structure

This is my folder structure under my `$GOPATH` or `$HOME/your_username/go`.

```

.

|-- bin

+-- src

| |-- .env

| |-- main.go

| +-- controllers

| | |-- usersController.go

| +-- models

| | |-- base.go

| | |-- user.go

| +-- routes

| | |-- api.go

| +-- utils

| | |-- utils.go

```

  

##  Download the packages used to create this rest API

Run the following Golang commands to install all the necessary packages. These packages will help you set up a web server, ORM for interacting with your db, mysql driver for db connection and load your .env file or setting up environment variables respectively.

```

go get

```

###  Running documentation locally (Only documentation of packages your have installed)

For offline documentation on the following packages run `godoc -http :6060` and then visit `http://localhost:6060`. Note that you can change the port to your preferred port number.

  

##  Setting configuration file

Create a .env file in the root of the project and set the following parameters

  

```

base_url=http://base_url.com base url

  

db_name=db name

db_user=db username

db_pass=db password

db_type=db type eg mysql

db_host=db host

db_port=db port

charset=utf8

parse_time=True

timezone=Asia%2FJakarta format time

web_port=5900 your web service's port

prefix=/api/v1 your prefix

  

one_signal_key= one signal key

one_signal_server= one signal server

  
  
  

waave_cell_name= your wave cell name

wave_cell_url= your wave cell url

wave_cell_key= your wave cell key

  

mailgun_domain= your mailgun domain

mailgun_private_api= your mailgun private key

mailgun_sender_email= your mailgun sender email

  

ONE_SIGNAL_USER_KEY= your one signal user key

ONE_SIGNAL_APP_KEY= your one signal app key

ONE_SIGNAL_APP_ID= your one signal app id

```