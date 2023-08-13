# Description

Simple application that implements the following concepts in Go:
- restful api
- crud
- mysql
- jwt (cookie and auth header)
- authentication
- authorization
- bcrypt
- transactions
- env

# Init

The main.go file starts the app and loads the env config, the priority is dev.env (lowest), .env (higher) and your os env variables (highest).

After the config is loaded the DB is initialized, this creates the DB instance (connection pool) that will be used for all future requests.

A transaction middleware is added to all requests, the expectation is that the backend serves a SASS like product that needs transactionality on all it's controllers.

An auth middleware is added to all requests, it takes the JWT from the request (cookie or auth header) and checks if it's valid extracting the user id and tying it into the context.

The final part of the setup is registering routes.

# Request Processing

A pod like structure is used for organizing concepts, each concept consists of a model, repository, controller and optionaly a validator. 

The controller has a RegisterEndpoints method that sets up the proper request filtering, the filters are in charge of checking the users access level.

Here is an example of registering routes of a user controller:

```go
func RegisterEndpoints(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/users/current", security.LoggedInFilter, current)
}
```

Controllers use repositories to load and manipulate models, the models directly map to DB tables. 

All queries are defined in the repositories and all repository functions require a transaction pointer to be passed in (this is obtained from the gin.Context using the transaction.utils.go)

Transaction filter that is attached to all requests will commit transactions only if there were no panics and if the final status of the request was 200 or 201.

# DB

The DB credentials are set in the dev.env file, if your local ENV has these flags defined they will override the dev.env file (.env file will also override it). 

The default credentials are root and root for the password. The DB structure can be loaded from create.sql file.

# API

The api represents a real world case of an application that supports user registration, login and retrival of the current logged in user. 

It also supports a restful tenant based API for lifts (as in going to the gym and lifting and then recording the lifts), the user has only ever access to their own lifts.

Following endpoints are implemented:
```
POST   /signup
POST   /login
GET    /users/current
GET    /lifts
GET    /lifts/:id
POST   /lifts
PUT    /lifts/:id
DELETE /lifts/:id
```

Example of a POST /signup request:
```json
{
    "name": "Dusan",
    "email": "dusan@dusan.com",
    "password": "dusan"
}
```


Example of a POST /login request:
```json
{
    "email": "dusan@dusan.com",
    "password": "dusan"
}
```

Example of a POST /lifts request (the id will be generated and returned, and the user id is going to be set to current users user id):
```json
{
    "id": 0,
    "userId": 0,
    "name": "Bench",
    "liftDate": "2023-02-02",
    "weight": 20000,
    "reps": 2
}
```



