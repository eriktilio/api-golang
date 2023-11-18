# Use a imagem oficial do Go como base
FROM golang:1.21-alpine3.18

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /go/src/api-golang

# Copie os arquivos do projeto para o contêiner
COPY . .

# Baixe as dependências
RUN go mod init api-golang && go get ./...

# Exponha a porta em que a aplicação será executada
EXPOSE 8080

# Comando para executar a aplicação quando o contêiner for iniciado
CMD ["go","run","src/main.go"]
