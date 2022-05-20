## curl samples

### GET

`To get all TODOs`

```s
$ curl http://localhost:8001/api/todo
```

`Get a TODO by Id`

```s
$ curl http://localhost:8001/api/todo/aa3cdd2f-17b9-4f43-9eb0-af56b42908c5
```

`Search TODOs by Title and IsDone`

```s
$ curl http://localhost:8001/api/todo/search?title=Task%20A
$ curl http://localhost:8001/api/todo/search?title=Task&isDone=true
$ curl http://localhost:8001/api/todo/search?title=B&isDone=true
```


### POST

```s
$ curl -X POST "http://localhost:8001/api/todo" --include --header "Content-Type: application/json" --data '{ "title": "Task D", "isDone": false }'
```

### PUT

```s
$ curl -X PUT "http://localhost:8001/api/todo" --include --header "Content-Type: application/json" --data '{"id": "aa3cdd2f-17b9-4f43-9eb0-af56b42908c5", "title": "Task AAA", "isDone": true}'

$ curl -X PUT "http://localhost:8001/api/todo" --include --header "Content-Type: application/json" --data '{"id": "cca89c32-a0d9-43c9-84e2-ae1224c5d755", "title": "Task CCC", "isDone": false}'
```

### DELETE

```s
$ curl -X DELETE "http://localhost:8001/api/todo" --include --header "Content-Type: application/json" --data '{"id": "bbf5d05c-c442-4869-8326-ab5cfa832f6a", "title": "Task B", "isDone": true}'

$ curl -X DELETE "http://localhost:8001/api/todo" --include --header "Content-Type: application/json" --data '{"id": "bbf5d05c-c442-4869-8326-ab5cfa832f6a"}'
```


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
2022/05/02 01:53:36 Generate swagger docs....
2022/05/02 01:53:36 Generate general API Info, search dir:./
2022/05/02 01:53:37 create docs.go at  docs/docs.go
2022/05/02 01:53:37 create swagger.json at  docs/swagger.json
2022/05/02 01:53:37 create swagger.yaml at  docs/swagger.yaml```
```


## Reference

- [General API Info](https://github.com/swaggo/swag#general-api-info)
- [API Operation](https://github.com/swaggo/swag#api-operation)
  - [Param Type](https://github.com/swaggo/swag#param-type)
  - [Param samples](https://github.com/swaggo/swag#attribute)
  - [Data Type](https://github.com/swaggo/swag#data-type)




