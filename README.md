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
