# Sobre o Kong

O Kong é plataforma de serviços que ajudam a conectar e gerenciar APIs e microsserviços em escala.
O Kong Gateway (API Gateway da Kong Inc.) é apenas um dos serviços fornecidos pelo Kong Inc. Ele é dividido em 2 partes: Community e Enterprise. A versão Community é gratuita e open-source, enquanto a Enterprise é paga e possui mais funcionalidades.

## Manuseio do Kong

Existem 2 modos para usar o Kong Gateway: com banco de dados e sem banco de dados. 

No modo com banco de dados, é possível usar o Kong Manager (UI) para gerenciar as APIs, plugins, etc. Além disso, é possível usar o Admin API (porta 8001, via REST) para fazer as mesmas operações que o Kong Manager.

No modo sem banco de dados, o Kong Gateway rodará em modo declarativo, ou seja, ele irá ler um arquivo de configuração (YAML) e aplicar as configurações nele contidas.
O Admin API e o Kong Manager estarão disponíveis apenas em modo de leitura, ou seja, não será possível fazer alterações via estas interfaces, apenas via arquivo de configuração e é possível usar o Kong Deck para fazer alterações.


## Executando o Kong Gateway

No projeto existem 2 modos de execução do Kong Gateway:

* Com banco de dados (Postgres) - docker-compose.with-db.yaml
* Sem banco de dados - docker-compose.with-dbless.yaml

O arquivo de configuração sem banco de dados possui um volume para o arquivo de configuração (kong.yml).


## Kong Deck

O Estado é um conjunto de configurações de Kong que é a fonte da verdade. O deck pegará o estado e fará chamadas à API Admin para Kong para combinar a configuração armazenada no banco de dados do Kong com o estado. Isso também é conhecido como estado alvo ou estado desejado.

### Verificar se o Deck está conectado ao Kong Gateway

Para verificar se o Deck está conectado ao Kong Gateway, execute o seguinte comando:

```bash
docker run --add-host host.docker.internal:host-gateway --network host kong/deck:v1.37.0 gateway ping --kong-addr http://host.docker.internal:8001  
```

### Fazer dump do Kong Gateway via Deck

Para fazer dump do Kong Gateway via Deck, execute o seguinte comando:

```bash
docker run --add-host host.docker.internal:host-gateway --network host kong/deck:v1.37.0 gateway dump --kong-addr http://host.docker.internal:8001  
```

Isto vai gerar o resultado do YAML de todas as configurações do Kong Gateway no console


## APIs da Imersão

### API dos partners (Nest.js)

Temos 2 API para os partners:

* Endereço original - http://localhost:3000 - Partner 1
* Endereço no Kong Gateway http://host.docker.internal:8000/partner1 - Partner 1

Apenas o endpoint de reserva de ingresso está protegido por autenticação key-auth.
Para fazer a chamada é necessário passar o header `X-Api-Token` com o valor `123`.

Arquivo `api.http` tem o teste da API.

* Endereço original - http://localhost:3001 - Partner 2
* Endereço no Kong Gateway http://host.docker.internal:8000/partner2 - Partner 2

Apenas o endpoint de reserva de ingresso está protegido por autenticação key-auth.
Para fazer a chamada é necessário passar o header `X-Api-Token` com o valor `000`.

## API de vendas (Golang)

* Endereço original - http://localhost:8080 - API de vendas
* Endereço no Kong Gateway http://host.docker.internal:8000/golang - API de vendas

Todos os endpoints estão protegidos por autenticação key-auth.
Para fazer a chamada é necessário passar o header `X-Api-Token` com o valor `890`.

## Front-end (Next.js)

* Endereço original http://localhost:3002 - Front-end
* Endereço no Kong Gateway http://host.docker.internal:8000/nextjs - Front-end





