```bash
docker compose exec -it sqlc bash
```

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

```bash
migrate create -ext=sql -dir=sql/migrations -seq init
```

```bash
migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql:3306)/courses" -verbose up
```

```bash
docker compose exec -it mysql bash
```

```bash
mysql -uroot -p courses
root
```

```bash
show tables;
```

```bash
desc courses
```

```bash
desc categories
```

```bash
migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql:3306)/courses" -verbose down
```

```bash
show tables;
```
