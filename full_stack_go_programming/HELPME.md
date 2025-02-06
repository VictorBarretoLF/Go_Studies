# HELPME.md

## Comandos Docker

### Subir os containers em segundo plano
```bash
docker compose up -d
```

### Acessar o shell do container
```bash
docker exec -it <app-name> bash
```

### Rebuildar a imagem Docker sem cache
```bash
docker-compose build --no-cache
```

## Gerenciamento de Dependências

### Limpar e organizar dependências do módulo Go
```bash
go mod tidy
```

## Testes

### Executar todos os testes
```bash
go test ./...
```