```bash
    docker exec -it go_app_banco_de_dados bash
```

Tidy Your Module
Run the go mod tidy command to clean up the module dependencies, ensuring that all necessary dependencies are properly noted and unused ones are removed:

https://github1s.com/devfullcycle/goexpert/blob/main/6-Banco-de-Dados/1/main.go

https://github1s.com/devfullcycle/imersao18/blob/main/golang/cmd/events/main.go

```bash
    go mod tidy
```

```bash
    go test ./...
```

```bash
    docker compose exec mysql bash
    mysql -uroot -p goexpert
```

```sql
    create table products (id varchar(255), name varchar(80), price decimal, primary key (id));
```