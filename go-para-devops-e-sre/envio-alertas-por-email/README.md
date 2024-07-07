# Documentação

## Executando Comandos Docker

Para acessar o container do banco de dados da aplicação Go em execução, use o seguinte comando:

```bash
docker exec -it go_app_banco_de_dados bash
```

Este comando abrirá um shell bash no container Docker especificado.

## Limpando Dependências do Módulo Go

Para limpar as dependências do módulo, garantindo que todas as dependências necessárias sejam corretamente anotadas e as não utilizadas sejam removidas, execute:

```bash
go mod tidy
```

Para mais detalhes, você pode consultar as implementações de exemplo:

-   [Exemplo 1](https://github1s.com/devfullcycle/goexpert/blob/main/6-Banco-de-Dados/1/main.go)
-   [Exemplo 2](https://github1s.com/devfullcycle/imersao18/blob/main/golang/cmd/events/main.go)

## Executando Testes

Para executar todos os testes no seu projeto Go, use:

```bash
go test ./...
```

Este comando executará todos os arquivos de teste no diretório atual e seus subdiretórios.

## Acessando o Banco de Dados MySQL no Docker

Para acessar o banco de dados MySQL em execução em um container Docker, use:

```bash
docker compose exec mysql bash
```

Uma vez no shell do container, conecte-se ao banco de dados MySQL com o comando:

```bash
mysql -uroot -p
```

Depois de inserir a senha, selecione o banco de dados `goexpert`:

```sql
USE goexpert;
```

## Criando a Tabela de Produtos

Para criar uma tabela de produtos no banco de dados, use o seguinte comando SQL:

```sql
CREATE TABLE products (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(80),
    price DECIMAL
);
```

---

Esta documentação fornece um guia passo a passo para acessar o banco de dados no Docker, limpar as dependências do módulo Go, executar testes e criar uma tabela no MySQL.
