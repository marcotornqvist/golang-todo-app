# **Golang Todo App**

Testing CRUD functionality with Golang, PostgreSQL and Chi router.

### **Requirements**

- Golang & Docker
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)

### **Usage**

Run the command below to clone the repository:

```
git clone https://github.com/marcotornqvist/golang-todo-app.git
```

Environment Variables - Create a .env file and copy & paste the values below and change $PG_USER, $PG_PASSWORD and $PG_DB values to your own database connection values below:

```
POSTGRES_USER=$PG_USER
POSTGRES_PASSWORD=$PG_PASSWORD
POSTGRES_DB=$PG_DB
POSTGRES_URL=localhost
POSTGRES_PORT=5432
```

### **Migration**

Remember to migrate the database to the latest version using the instructions below. Follow the instructions below also when making changes to the database in the future. Read more on how to use golang-migrate [here](https://github.com/golang-migrate/migrate).

Run command below to create new migration files (Remember to change "migration_name" below to something more fitting):

```
migrate create -ext sql -dir db/migrations -seq migration_name
```

Run command below to export PostgreSQL database URL so that it can be accessed by Makefile for instance (Remember to change the DB connection values $PG_USER, $PG_PASSWORD and $PG_DB to your own database connection values below):

```
export POSTGRESQL_URL="postgres://$PG_USER:$PG_PASSWORD@localhost:5432/$PG_DB?sslmode=disable"
```

Run command below to migrate database to latest version:

```
make migrateup
```

or

```
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```

### **Start the application**

Run command below to start the project without Docker:

```
go run main.go
```

### **Docker**

Remember to set POSTGRES_URL value in .env file to "host.docker.internal" if running application with Docker.

```
POSTGRES_URL=host.docker.internal
```

Run command below to start the project with Docker:

```
docker-compose up
```
