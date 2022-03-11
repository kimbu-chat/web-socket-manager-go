## Setup

1. Install go1.17
2. Run postgres instance:
    - Clone services repository <https://github.com/kimbu-chat/services>
    - Run postgres in context of services repository: `docker-compose up -d postgres`
    - Create database in postgres: `docker exec postgres createdb -U sa websocketmanager`
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
