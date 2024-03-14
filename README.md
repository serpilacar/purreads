# Purreads 🐈‍⬛

## What is it?

Purreads is a simple Goodreads clone with a very limited functionality. I initiated it to try out GraphQL with Go.

With Purreads you can create a new ~~user~~ cat, create a new book or add a book to a cats' library as an "in progress" or a "read" book.

Built with:

- [GORM](https://gorm.io/gorm)
- [PostgreSQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)
- [gqlgen](https://github.com/99designs/gqlgen)

## Schema

GraphQL schema can be found in [graph/schema.graphqls](graph/schema.graphqls) and resolvers are in [graph/schema.resolvers.go](graph/schema.resolvers.go).

## Usage

Start server by running:

```sh
$ go run server.go
```

GraphQL playground will be available at port 8080.

### Updating schema

After updating the GraphQL schema or operations, don't forget to run generate command:

```sh
$ go generate ./...
```

This command will update [generated code](graph/generated.go) according to your schema if necessary.
