# API Golang

Este é um projeto simples de API em Go com estrutura modular.

## Estrutura do Projeto

```bash
api-golang/
|-- src/
| |-- main.go
| |-- api/
|     |-- routes/
|         |-- routes.go
|     |-- controllers/
|         |-- controller.go
```

## Como Executar

Siga os seguintes passos para executar o sistema:

1. Certifique-se de ter o Go instalado no seu sistema.

2. Execute o seguinte comando para instalar as dependências do projeto:

   ```bash
   go mod init api-golang && go get ./...
   ```

3. Navegue até o diretório do projeto no terminal.

4. Execute o seguinte comando para iniciar o servidor:

   ```bash
   go run src/main.go
   ```

Abra seu navegador e vá para http://localhost:8080. Você deverá ver a mensagem "Bem-vindo à minha API Go!".

## Subir Container Docker

1. Crie a imagem 'api-golang' com o comando:
   ```bash
   docker build -t api-golang .
   ```
2. Criando o container 'api-golang' apartir da imagem, na porta 8080:
   ```bash
   docker run -d -p 8080:8080 --name api-golang api-golang
   ```

## Sobre as dependências

- gorilla/mux: Uma poderosa biblioteca de roteamento para Go.

## Contribuindo

Sinta-se à vontade para contribuir! Abra problemas (issues) ou envie pull requests.

## Licença

Este projeto é distribuído sob a licença MIT.
