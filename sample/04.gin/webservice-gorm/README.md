
# curl samples

## api/user

```s
// Create a user
$ curl -X POST "http://localhost:8001/api/user" --include --header "Content-Type: application/json" --data '{"Id": "JB", "Name": "JB Lin"}'

// Query the user
$ curl --include http://localhost:8001/api/user/JB

// Update the user
$ curl -X PUT "http://localhost:8001/api/user" --include --header "Content-Type: application/json" --data '{"id": "JB", "name": "Bon Jovi"}'

// Delete the user by Id
$ curl -X DELETE "http://localhost:8001/api/user" --include --header "Content-Type: application/json" --data '{"id": "JB"}'
```

## api/todo

```s
// Get the TODO by its Id
$ curl --include http://localhost:8001/api/todo/a2a25eb2-9bf7-4cb6-b9db-ece05b78e975

// Create a TODO
$ curl -X POST "http://localhost:8001/api/todo" --include --header "Content-Type: application/json" \
  --data '{ "title": "XXXX", "isDone": false, "userId": "JB", "todoExt": {"description": "YYYY", "priorityId": 2}, "tags": [{"id": "6aee5542-3f70-4cbc-ab05-fd020285f021"}, {"id": "dcc5a568-ae07-4600-9055-97eb129f319c"}] }'

// Update the TODO
curl -X PUT "http://localhost:8001/api/todo" --include --header "Content-Type: application/json" \
  --data '{ "id": "af4f750b-895f-4be0-99c4-38a5fedf642c", "title": "ZZZZ", "isDone": true, "userId": "JB", "todoExt": {"id": "af4f750b-895f-4be0-99c4-38a5fedf642c", "description": "WWWW", "priorityId": 3}, "tags": [{"id": "6aee5542-3f70-4cbc-ab05-fd020285f021"}] }'

// DELETe the TODO
$ curl -X DELETE "http://localhost:8001/api/todo" --include --header "Content-Type: application/json" \
  --data '{"id": "af4f750b-895f-4be0-99c4-38a5fedf642c"}'
```

## api/todos

```s
// Get all TODOs
$ curl --include http://localhost:8001/api/todos

// Search TODOs by Title and IsDone
$ curl --include "http://localhost:8001/api/todo/search?title=Task%20A"
$ curl --include "http://localhost:8001/api/todo/search?title=Task&isDone=true"
$ curl --include "http://localhost:8001/api/todo/search?title=B&isDone=true"

// Delete TODOs by their Id
$ curl --include -X DELETE "http://localhost:8001/api/todos" --include --header "Content-Type: application/json" \
  --data '[{"id": "26447744-68f5-4bfa-8126-e11637bd1533"}, {"id": "dd7b1e23-542f-4beb-9bb6-96435d8e5305"}]'
```

***
# Swagger

## Install swag

> GitHub: [swaggo/swag](https://github.com/swaggo/swag)

Under project root path:

```s
$ go get -u github.com/swaggo/swag/cmd/swag

# 1.16 or newer
$ go install github.com/swaggo/swag/cmd/swag@latest
```


## Install 

> GitHub: [gin-swagger](https://github.com/swaggo/gin-swagger)

Under project root path:

```s
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
```


## swag init

```s
# Default: main.go
$ swag init

# Or specified file name
$ swag init -g|--generalInfo server.go
$ swag init  --parseDependency --parseInternal -g server.go
```


## Reference

- [General API Info](https://github.com/swaggo/swag#general-api-info)
- [API Operation](https://github.com/swaggo/swag#api-operation)
  - [Param Type](https://github.com/swaggo/swag#param-type)
  - [Param samples](https://github.com/swaggo/swag#attribute)
  - [Data Type](https://github.com/swaggo/swag#data-type)




