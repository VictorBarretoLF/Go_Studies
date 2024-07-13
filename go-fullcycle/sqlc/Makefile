createmigration:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql:3306)/courses" -verbose up

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql:3306)/courses" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql:3306)/courses" -verbose down

.PHONY: migrate createmigration migratedown