
# curl samples

## api/user

```s
// Create a user
curl -X POST "http://localhost:8080/api/user" --include --header "Content-Type: application/json" --data '{"Id": "JB", "Name": "JB Lin"}'

// Query the user
curl --include http://localhost:8080/api/user/JB

// Update the user
curl -X PUT "http://localhost:8080/api/user" --include --header "Content-Type: application/json" --data '{"id": "JB", "name": "Bon Jovi"}'

// Delete the user by Id
curl -X DELETE "http://localhost:8080/api/user" --include --header "Content-Type: application/json" --data '{"id": "JB_1654521635"}'
```

## api/todo

```s
// Get the TODO by its Id
curl --include http://localhost:8080/api/todo/0bf19501-4385-4b70-9cfa-a31ca6ab47c1

// Create a TODO
curl -X POST "http://localhost:8080/api/todo" --include --header "Content-Type: application/json" \
  --data '{ "title": "XXXX", "isDone": false, "userId": "JB", "todoExt": {"description": "YYYY", "priorityId": 2}, "tags": [{"id": "6aee5542-3f70-4cbc-ab05-fd020285f021"}, {"id": "dcc5a568-ae07-4600-9055-97eb129f319c"}] }'

// Update the TODO
curl -X PUT "http://localhost:8080/api/todo" --include --header "Content-Type: application/json" \
  --data '{ "id": "8fdf7e08-a065-433d-a31b-dcdcb8186ba6", "title": "ZZZZ", "isDone": true, "userId": "JB", "todoExt": {"id": "8fdf7e08-a065-433d-a31b-dcdcb8186ba6", "description": "WWWW", "priorityId": 3}, "tags": [{"id": "6aee5542-3f70-4cbc-ab05-fd020285f021"}] }'

// DELETe the TODO
curl -X DELETE "http://localhost:8080/api/todo" --include --header "Content-Type: application/json" \
  --data '{"id": "ff0409da-e55f-4542-a11c-e33c0618f312"}'
```

## api/todos

```s
// Get all TODOs
curl --include http://localhost:8080/api/todos

// Search TODOs by Title and IsDone
curl --include "http://localhost:8080/api/todos/search?title=XX"
curl --include "http://localhost:8080/api/todos/search?title=Task&isDone=true"
curl --include "http://localhost:8080/api/todos/search?title=B&isDone=true"

// Delete TODOs by their Id
$ curl --include -X DELETE "http://localhost:8080/api/todos" --include --header "Content-Type: application/json" \
  --data '[{"id": "64caa6c9-0fa1-4290-b394-60f37137b756"}, {"id": "c9c278a6-5733-4c08-86f6-f9cf5f41c93d"}]'
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




