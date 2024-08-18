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
