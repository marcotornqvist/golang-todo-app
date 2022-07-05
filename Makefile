migrateup:
		migrate -database ${POSTGRESQL_URL} -path db/migrations up

migratedown:
		migrate -database ${POSTGRESQL_URL} -path db/migrations down