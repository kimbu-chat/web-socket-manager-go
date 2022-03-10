## Setup
1. Install go1.17
2. Install postgres
3. Set database env variables: DB_HOST, DB_USER, DB_PASSWORD
4. Create database
```
PGPASSWORD=$DB_PASSWORD createdb --username=$DB_USER golightweb
```

## Run
```
go run cmd/golightweb/main.go
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
sql-migrate up
```
