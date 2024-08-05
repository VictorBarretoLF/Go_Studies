Claro! Vou adicionar a seção sobre `go mod init` e explicar como executar migrações com o [Tern](https://github.com/jackc/tern).

---

# HELPME.md

## Executando o Contêiner Docker

Para acessar o shell do contêiner Docker da aplicação, execute o seguinte comando:

```bash
docker exec -it <app-name> bash
```

Substitua `<app-name>` pelo nome ou ID do seu contêiner Docker.

## Configurando o Módulo Go

### Inicializando um Módulo Go com `go mod init`

O comando `go mod init` é utilizado para inicializar um novo módulo Go. Ele cria um arquivo `go.mod` no diretório atual, que serve como o ponto central para o gerenciamento de dependências do projeto. Aqui está o que ele faz:

-   Cria um arquivo `go.mod` no diretório atual.
-   Define o nome do módulo, que normalmente é o caminho do repositório (por exemplo, `github.com/username/projectname`).

Para inicializar um módulo Go, use:

```bash
go mod init <module-name>
```

Substitua `<module-name>` pelo nome desejado para o módulo.

### Organizando o Módulo Go com `go mod tidy`

O comando `go mod tidy` é usado para limpar e organizar as dependências do módulo Go. Ele realiza as seguintes ações:

-   Adiciona todas as dependências necessárias que estão faltando no arquivo `go.mod`.
-   Remove dependências não utilizadas do arquivo `go.mod`.
-   Garante que o arquivo `go.sum` contenha verificações criptográficas precisas para todas as dependências listadas.

Este comando ajuda a manter seu projeto Go limpo e bem organizado, assegurando que apenas as dependências necessárias estejam listadas nos arquivos de configuração do módulo.

Para executar o `go mod tidy`, use:

```bash
go mod tidy
```

### Testando o Módulo Go com `go test ./...`

O comando `go test ./...` executa todos os testes do seu projeto Go. Aqui está uma explicação detalhada:

-   `go test` é o comando usado para rodar testes em Go.
-   `./...` é um padrão que indica para o Go procurar por testes recursivamente em todos os subdiretórios a partir do diretório atual.

Este comando é útil para garantir que todos os pacotes e suas respectivas funções de teste (`*_test.go`) sejam executados, ajudando a verificar se o código está funcionando conforme o esperado.

Para executar todos os testes do seu projeto Go, use:

```bash
go test ./...
```

## Executando Migrations com o Tern

[Tern](https://github.com/jackc/tern) é uma ferramenta para gerenciar migrações de banco de dados com SQL. Aqui estão os passos básicos para usar o Tern:

### Instalando o Tern

Primeiro, instale o Tern:

```bash
go install github.com/jackc/tern/v2@latest
```

### Inicializando Migrações

Para inicializar um novo diretório de migrações, use:

```bash
mkdir -p ./internal/store/pgstore  
tern init ./internal/store/pgstore/migrations
```

### Criar nova Migração

```bash
tern new --migrations ./internal/store/pgstore/migrations create_rooms_table
tern new --migrations ./internal/store/pgstore/migrations create_messages_table
```

### Aplicando Migrações

Para aplicar migrações ao banco de dados, use:

```bash
tern migrate -m ./migrations
```

Certifique-se de que o diretório `./migrations` contém os arquivos de migração SQL.

#### Caso o banco de dados não tenha sido criado, executar os seguintes passos

Para entrar no container do banco de dados

```bash
docker exec -it db-postgres bash
```

Acessar o postgres

```bash
psql -U postgres
```

Listar banco de dados criados

```bash
\l
```

Criar o banco de dados

```bash
CREATE DATABASE "go-socket-db";
```

Listar os nomes dos bancos de dados

```bash
SELECT datname FROM pg_database;
```

## Exemplo Completo

Para configurar, organizar as dependências, rodar os testes e aplicar migrações, siga os passos abaixo:

1. Acesse o contêiner Docker:

    ```bash
    docker exec -it <app-name> bash
    ```

2. Inicialize o módulo Go (se ainda não tiver feito):

    ```bash
    go mod init <module-name>
    ```

3. Organize as dependências do módulo Go:

    ```bash
    go mod tidy
    ```

4. Execute todos os testes:

    ```bash
    go test ./...
    ```

5. Inicialize as migrações (se ainda não tiver feito):

    ```bash
    tern init
    ```

6. Aplique as migrações ao banco de dados:

    ```bash
    tern migrate -m ./migrations
    ```
