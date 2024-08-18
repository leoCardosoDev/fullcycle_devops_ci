# Etapa de construção e execução
FROM golang:1.23-bullseye AS build

# Defina o diretório de trabalho
WORKDIR /app

# Copie o arquivo go.mod e go.sum e faça o download das dependências
COPY go.mod go.sum ./
RUN go mod download

# Copie o código-fonte da aplicação para o contêiner
COPY . .

# Compile o código Go
RUN go build -o main .

# Etapa de execução
FROM golang:1.23-bullseye

# Defina o diretório de trabalho
WORKDIR /app

# Copie o binário compilado do estágio de build
COPY --from=build /app/main .

# Defina o comando padrão a ser executado quando o contêiner iniciar
CMD ["./main"]
