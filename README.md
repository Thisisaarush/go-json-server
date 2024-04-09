# GO RSS SERVER

## Prerequisites

- Start postgres as a local service through pgadmin or use postgres Docker image
- Add a .env file in the root directory (see .env.example)

## Installation

```bash
go mod tidy
go mod vendor
```

## Run the server

```bash
go build && ./go-rss # for linux or mac
go build && ./go-rss.exe # for windows
```

## Common commands for sqlc and goose library

```bash
sqlc generate # generate go code from sql schema
goose postgres postgres://postgres:PASSWORD@localhost:5432/DATABASE up # run migrations
goose -dir ./sql/schema postgres postgres://postgres:PASSWORD@localhost:5432/DATABASE up # run migrations if above command fails to find migration files
```
