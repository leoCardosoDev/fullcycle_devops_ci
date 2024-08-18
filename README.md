# Criando um Workflow CI
Claro! Para criar um ambiente de desenvolvimento Docker com Go e SQLite3, você pode seguir os passos abaixo. Vamos criar um Dockerfile e um arquivo `docker-compose.yml` para configurar o ambiente. O Dockerfile instalará o Go e o SQLite3 e o `docker-compose.yml` facilitará a execução do contêiner.

### 1. Dockerfile

Crie um arquivo chamado `Dockerfile` com o seguinte conteúdo:

```Dockerfile
# Use uma imagem base com Go
FROM golang:1.20 AS build

# Instale SQLite3
RUN apt-get update && \
    apt-get install -y sqlite3 libsqlite3-dev

# Crie um diretório para a aplicação
WORKDIR /app

# Copie o código-fonte da aplicação para o contêiner
COPY . .

# Compile o código Go
RUN go build -o main .

# Use uma imagem base mais leve para o ambiente de execução
FROM debian:bullseye-slim

# Instale SQLite3 no ambiente de execução
RUN apt-get update && \
    apt-get install -y sqlite3 libsqlite3-dev && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Crie um diretório para a aplicação
WORKDIR /app

# Copie o binário compilado do estágio de build
COPY --from=build /app/main .

# Defina o comando padrão a ser executado quando o contêiner iniciar
CMD ["./main"]
```

### 2. docker-compose.yml

Crie um arquivo chamado `docker-compose.yml` com o seguinte conteúdo:

```yaml
version: '3.8'

services:
  go-app:
    build: .
    container_name: go_app
    volumes:
      - .:/app
    ports:
      - "8080:8080"
    environment:
      - ENV=development
```

### 3. Configuração do Go

Certifique-se de que seu projeto Go tenha um arquivo `main.go` e qualquer outro arquivo necessário para sua aplicação. Aqui está um exemplo simples de um `main.go` que conecta ao SQLite3:

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Abra a conexão com o banco de dados SQLite3
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Crie uma tabela de exemplo
    sqlStmt := `
    CREATE TABLE IF NOT EXISTS example (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT
    );
    `
    _, err = db.Exec(sqlStmt)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Database initialized and ready.")
}
```

### Comandos para executar o Docker

1. **Construir a imagem Docker**:

   ```bash
   docker-compose build
   ```

2. **Iniciar o contêiner**:

   ```bash
   docker-compose up
   ```

   Para rodar o contêiner em segundo plano, use:

   ```bash
   docker-compose up -d
   ```

3. **Parar e remover o contêiner**:

   ```bash
   docker-compose down
   ```

Isso criará um ambiente de desenvolvimento com Go e SQLite3 pronto para uso. Se precisar de ajustes adicionais ou tiver alguma dúvida, é só me avisar!
