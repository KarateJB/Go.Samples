## Install packages

```s
$ go get github.com/99designs/gqlgen
$ go run github.com/99designs/gqlgen init
Creating gqlgen.yml
Creating graph/schema.graphqls
Creating server.go
Generating...
go: downloading github.com/stretchr/testify v1.7.1
go: downloading github.com/andreyvit/diff v0.0.0-20170406064948-c7f18ee00883
go: downloading github.com/arbovm/levenshtein v0.0.0-20160628152529-48b4e1c0c4d0
go: downloading github.com/dgryski/trifles v0.0.0-20200323201526-dd97f9abfb48
go: downloading gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c

Exec "go run ./server.go" to start GraphQL server
```


## Structure

```
├── go.mod                
├── go.sum                
├── gqlgen.yml            
├── graph                 
|  ├── generated          
|  |  └── generated.go    
|  ├── model              
|  |  └── models_gen.go   
|  ├── resolver.go        
|  ├── schema.graphqls    
|  └── schema.resolvers.go
└── server.go
```

## go generate

```s
go get github.com/99designs/gqlgen@v0.17.9
go get github.com/99designs/gqlgen/internal/imports@v0.17.9
go get github.com/99designs/gqlgen/codegen/config@v0.17.9
go get github.com/99designs/gqlgen/internal/imports@v0.17.9
go run github.com/99designs/gqlgen generate
```

or

```s
go generate ./...
```



## Update Resolver

> Resolver is a collection of functions that generate response for a GraphQL query.

---
`graph\resolver.go`

Update following code,

```go
type Resolver struct {
	todos []*model.Todo
}
```

to

```go
type Resolver struct {
	todo  *model.Todo
	todos []*model.Todo
}
```

And then run `go generate`. The updated resolver will make a new response function.

---
`graph\schema.resolvers.go`

```go
func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
  panic(fmt.Errorf("not implemented"))
}
```


## Reference

- [Don’t eagerly fetch the user](https://gqlgen.com/getting-started/#dont-eagerly-fetch-the-user)
- [Regenerating resolvers from schema removes comments of resolver methods #1069](https://github.com/99designs/gqlgen/issues/1069)
