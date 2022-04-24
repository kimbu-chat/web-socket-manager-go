## Setup

1. Install go1.18.0
2. Run postgres instance:
    - Clone services repository <https://github.com/kimbu-chat/services>
    - Run postgres in context of services repository: `docker-compose up -d postgres`
    - Create database in postgres: `docker exec postgres createdb -U sa websocketmanager`
3. Copy file `.env.sample` to `.env`
4. Run migrations with env variables from .env file: `export $(cat .env | xargs); sql-migrate up`

## Run

```
go run cmd/websocketmanager/main.go
```

## Migrations

### Install migrations tool

```
go install github.com/rubenv/sql-migrate/sql-migrate@latest
```

### Create

```
sql-migrate new MIGRATION_NAME
```

### Run

```
export $(cat .env | xargs); sql-migrate up
```

## Linter

### Install

```
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
```

Or <https://golangci-lint.run/usage/install/#local-installation>

### Use

```
golangci-lint run
```

## Swagger

After API changes call swag to regenerate documentation

At first you need to download swag tool (v1.8.1) <https://github.com/swaggo/swag#getting-started>

Regenerate documentation

```
swag init -g internal/config/routes/routes.go -ot go
```
