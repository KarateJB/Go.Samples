# GraphQL: gin + gqlgen + GORM

This is a sample code of creating a RESTful API and GraphQL server with
- [gin](https://github.com/gin-gonic/gin)
- [gqlgen](https://github.com/99designs/gqlgen)
- [GORM](https://github.com/go-gorm/ gorm) and PostgreSQL.


## Install packages

```s
$ go get github.com/gin-gonic/gin
$ go get github.com/99designs/gqlgen
$ go run github.com/99designs/gqlgen init
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


## Reference

- [Setup Gin and gqlgen together](https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/gin.md)
