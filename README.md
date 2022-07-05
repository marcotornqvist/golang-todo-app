# **Golang Todo App**

Testing CRUD functionality with Golang, PostgreSQL and Chi router.

### **Requirements**

- Golang & Docker
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

### **Usage**

Clone the repository:

```
git clone https://github.com/marcotornqvist/golang-todo-app.git
```

Environment Variables - Create a .env file and change the DB connection values $PG_USER, $PG_PASSWORD and $PG_DB to your own database connection values below:

```
POSTGRES_USER=$PG_USER
POSTGRES_PASSWORD=$PG_PASSWORD
POSTGRES_DB=$PG_DB
POSTGRES_URL=localhost # (or "host.docker.internal" if running on with docker-compose up)
POSTGRES_PORT=5432
```

### **Migration**

Remember to migrate the database to the latest version using the instructions below. Follow the instructions below also when making changes to the database in the future. Read more on how to use golang-migrate [here](https://github.com/golang-migrate/migrate).

**Creates two new migration files in folder db/migrations:**

Run command below to create new migration files (Remember to change "migration_name" below to something more fitting):

```
migrate create -ext sql -dir db/migrations -seq migration_name
```

**Run command below to export PostgreSQL database URL (Remember to change the DB connection values $PG_USER, $PG_PASSWORD and $PG_DB to your own database connection values below):**

```
export POSTGRESQL_URL="postgres://$PG_USER:$PG_PASSWORD@localhost:5432/$PG_DB?sslmode=disable"
```

**Run command below to migrate database to latest version:**

```
make migrateup
```

or

```
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```

**Run command below to start the project:**

```
go run main.go
```
