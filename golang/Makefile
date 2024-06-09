.PHONY: all clean swag init run

# Diretórios
SWAGGER_DIR = ./docs
CMD_DIR = ./cmd/events
HANDLER_DIR = ./internal/events/infra/http

# Nome do binário
BINARY_NAME = events

# Variáveis
DB_USER = test_user
DB_PASSWORD = test_password
DB_HOST = localhost
DB_PORT = 3306
DB_NAME = test_db

all: swag build run

# Limpa os arquivos gerados
clean:
	rm -rf $(SWAGGER_DIR)
	rm -f $(BINARY_NAME)

# Gera os arquivos do Swagger
swag:
	swag init --output docs --dir ./cmd/events,./internal/events/infra/http,./internal/events/usecase

# Inicializa o banco de dados (se necessário)
init:
	mysql -u$(DB_USER) -p$(DB_PASSWORD) -h$(DB_HOST) -P$(DB_PORT) -e "CREATE DATABASE IF NOT EXISTS $(DB_NAME);"

# Compila o código
build:
	go build -o $(BINARY_NAME) $(CMD_DIR)/main.go

# Executa o servidor
run: build
	./$(BINARY_NAME)

# Comando para facilitar o desenvolvimento (gera swagger, compila e executa)
dev: swag build run
