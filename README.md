# Simple TODO List Backend Service

## RESTful API

```
-------------------------------------------------------------------------
Todo List Route
-------------------------------------------------------------------------
[Method]  [Route]
GET       /todos                  Fetch all todo tasks
GET       /todos/:id              Fetch a todo task
POST      /todo                   Create a new todo task
PUT       /todos/:id              Update a todo task information (all)
DELETE    /todos/:id              Delete a todo task
```

## Model

```
-------------------------------------------------------------------------
Todo Task
-------------------------------------------------------------------------
[Attribute]  [Type]
id           integer   
title        varchar(100)
description  varchar(200)
completed    boolean
created_at   timestamp
```


## Technologies
- [Golang](https://go.dev/) - An open source programming language designed for building simple, fast, and reliable software.
- [Gin](https://github.com/gin-gonic/gin) - A high performance go web framework.
- [MySQL](https://www.mysql.com/) - An open-source relational database management system to store `Todo Task` data.
- [Swag](https://github.com/swaggo/swag) - Automatically generate RESTful API documentation with Swagger 2.0 for Go.

## Prerequisite
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Start and Terminate the Service
Run the simple todo list service.
```bash
$ docker-compose up -d --build 
```

Shutdown the simple todo list service.
```bash
$ docker-compose down -v
```

## Usage
After spinning up the simple todo list service by `docker-compose`, you are able to interact with the service at `http://localhost:8080/swagger/index.html`.

![Alt text](screenshots/swagger-doc-home.png?raw=true)

## Examples

### Create a Task (POST /todo)
1. Hit the `Try it out` button to input your HTTP request data.

![Alt text](screenshots/post-try-button.png?raw=true)

2. Hit the `Execute` button to send the request. Swagger will automatically generate the `cURL` command amd execute it.

![Alt text](screenshots/post-execute-button.png?raw=true)

3. Service response.

![Alt text](screenshots/post-reply.png?raw=true)

### Get All Tasks (GET /todos)
Follow the process above to send GET http request.

![Alt text](screenshots/get-all.png?raw=true)

The execution process is similar for other APIs (i.e. Update, Delete). You could also call the APIs without Swagger UI. Instead, you could use `Postman` or other API testing tools. Swagger is one convenient way to test and document the APIs. 