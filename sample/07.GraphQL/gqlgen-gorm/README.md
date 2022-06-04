# GraphQL: gqlgen + GORM

This is a sample code of creating a GraphQL server with [gqlgen](https://github.com/99designs/gqlgen), [GORM](https://github.com/go-gorm/ gorm) and PostgreSQL.


## Install packages

```s
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

- [Don’t eagerly fetch the user](https://gqlgen.com/getting-started/#dont-eagerly-fetch-the-user)
- [Regenerating resolvers from schema removes comments of resolver methods #1069](https://github.com/99designs/gqlgen/issues/1069)
- [Custom Directive: Supported locations](https://www.apollographql.com/docs/apollo-server/schema/creating-directives#supported-locations)
- [Issue: Add field info to directive context](https://github.com/99designs/gqlgen/issues/1084)
