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

```s
go get github.com/99designs/gqlgen@v0.17.9
go get github.com/99designs/gqlgen/internal/imports@v0.17.9
go get github.com/99designs/gqlgen/codegen/config@v0.17.9
go get github.com/99designs/gqlgen/internal/imports@v0.17.9
go run github.com/99designs/gqlgen generate
```