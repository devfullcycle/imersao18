# Imersão Full Stack & FullCycle - Sistema de Venda de Ingressos

## Descrição

Repositório da API feita em Golang (Venda de ingressos)

## Rodar a aplicação

Dentro da pasta `golang` execute o comando abaixo para rodar os containers `Docker`:
```
docker compose up
```

Quando os containers estiverem prontos, precisamos acessar o container do `golang` e executar a aplicação:

```
// entrar no container:
docker compose exec golang sh

// instalar as dependências:
go mod tidy

// executar a aplicação:
go run cmd/events/main.go
```

### Para Windows 

Lembrar de instalar o WSL2 e Docker. Vejo o vídeo: [https://www.youtube.com/watch?v=btCf40ax0WU](https://www.youtube.com/watch?v=btCf40ax0WU) 

Siga o guia rápido de instalação: [https://github.com/codeedu/wsl2-docker-quickstart](https://github.com/codeedu/wsl2-docker-quickstart)