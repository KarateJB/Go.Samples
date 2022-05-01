## curl samples

### GET

```s
$ curl http://localhost:8001/api/todo

$ curl http://localhost:8001/api/todo/aa3cdd2f-17b9-4f43-9eb0-af56b42908c5
```

### POST

```s
$ curl -X POST "http://localhost:8001/api/todo/create" --include --header "Content-Type: application/json" --data '{ "title": "Task D", "isDone": false }'
```

### PUT

```s
$ curl -X PUT "http://localhost:8001/api/todo/edit" --include --header "Content-Type: application/json" --data '{"id": 
"aa3cdd2f-17b9-4f43-9eb0-af56b42908c5", "title": "Task AAA", "isDone": true}'

$ curl -X PUT "http://localhost:8001/api/todo/edit" --include --header "Content-Type: application/json" --data '{"id": "cca89c32-a0d9-43c9-84e2-ae1224c5d755", "title": "Task CCC", "isDone": false}'
```

### DELETE

```s
$ curl -X DELETE "http://localhost:8001/api/todo/remove" --include --header "Content-Type: application/json" --data '{"id": 
"bbf5d05c-c442-4869-8326-ab5cfa832f6a", "title": "Task B", "isDone": true}'

$ curl -X DELETE "http://localhost:8001/api/todo/remove" --include --header "Content-Type: application/json" --data '{"id": "bbf5d05c-c442-4869-8326-ab5cfa832f6a"}'
```