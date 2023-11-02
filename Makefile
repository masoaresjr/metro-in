# Defina as variáveis para o nome do banco de dados e a senha
DB_NAME=seu_banco_de_dados
DB_PASSWORD=sua_senha

# Comando para construir o projeto Go
build:
    "go build -o myapp

# Comando para criar e executar o contêiner MySQL
run-mysql:
    docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD=$(DB_PASSWORD) -p 3306:3306 mysql:latest

# Comando para estabelecer a conexão com o MySQL
connect:
    go run main.go

# Comando para limpar (parar e remover) o contêiner MySQL
clean-mysql:
    docker stop mysql-container
    docker rm mysql-container

# Comando para limpar (remover) o binário Go
clean:
    rm -f myapp

.PHONY: build run-mysql connect clean-mysql clean
