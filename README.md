# API Golang

Chamadas do Next.js para o Golang

## GET http://localhost:8000/events 

Listar os eventos

```json
[
        {
            "id": 1,
            "name": "Queen Celebration In Concert e Orquestra",
            "location": "Espaço Unimed - Rua Tagipuru, 795 - Barra Funda - São Paulo - SP",
            "organization": "Eda Shows e Eventos Ltda",
            "rating": "L", // L - Livre, L10 - 10 anos, L12 - 12 anos, L14 - 14 anos, L16 - 16 anos, L18 - 18 anos
            "date": "2021-01-01T21:00:00",
            "image_url": "http://localhost:8001/images/1.jpg", //deixar sempre o endereço das imagem em localhost:8001, vou criar um fake server para as imagens
            "created_at": "2021-01-01T00:00:00",
            "updated_at": "2021-01-01T00:00:00",
            "tickets": [
                {
                    "id": 1,
                    "name": "Inteira",
                    "price": 100.00,
                },
                {
                    "id": 2,
                    "type": "Meia",
                    "price": 50.00,
                }
            ]
        }
    ]
]
```

## GET http://localhost:8000/events/:eventId

Capturar um evento

O mesmo response acima, só que com um evento específico

## GET http://localhost:8000/events/:eventId/spots

Listar os assentos de um evento

```json
[
        {
            "id": 1,
            "name": "Queen Celebration In Concert e Orquestra",
            "location": "Espaço Unimed - Rua Tagipuru, 795 - Barra Funda - São Paulo - SP",
            "organization": "Eda Shows e Eventos Ltda",
            "rating": "L", // L - Livre, L10 - 10 anos, L12 - 12 anos, L14 - 14 anos, L16 - 16 anos, L18 - 18 anos
            "date": "2021-01-01T21:00:00",
            "image_url": "http://localhost:8001/images/1.jpg", //deixar sempre o endereço das imagem em localhost:8001, vou criar um fake server para as imagens
            "created_at": "2021-01-01T00:00:00",
            "updated_at": "2021-01-01T00:00:00",
            "tickets": [
                {
                    "id": 1,
                    "name": "Inteira",
                    "price": 100.00,
                },
                {
                    "id": 2,
                    "type": "Meia",
                    "price": 50.00,
                }
            ],
            "spots": [
                {"id": 1, "name": "A1", "status": "available", "event_id": 1},
                {"id": 1, "name": "A2", "status": "sold", "event_id": 1},
                ...
                {"id": 1, "name": "B1", "status": "available", "event_id": 1},
            ]
        }
    ]
]
```

## POST http://localhost:8000/checkout

Realizar a compra de um ingresso

- Request 
```json
{
    "event_id": 1,
    "spots": ["A1", "A2"],
    "ticket_id": 1,
    "card_hash": "hash_do_cartao",
    "email": "test@test.com"
}
```

- Response
```json
{
    "id": 1,
    "event_id": 1,
    "spots": ["A1", "A2"],
    "ticket_id": 1,
    "email": "test@test.com",
    "created_at": "2021-01-01T00:00
}
```

Chamadas do Golang para o Nest.js via API Gateway

## API 1

### GET http://localhost:9000/api1/spots/:spotId

Capturar um assento

```json
{
    "id": 1,
    "name": "A1",
    "status": "available", //available, sold
    "event_id": 1
}
```

### POST http://localhost:9000/api1/spots/:spotId/reserve

Reservar um assento

- Request
```json
{

    "name": "A1",
    "spots": ["A1", "A2"],
    "ticket_id": 10, //na API 1, o ticket_id = 10 é o ticket inteiro, o ticket_id = 20 é o ticket meia
    "event_id": 1
}
```

- Response
```json
{
    "id": 1,
    "spots": ["A1", "A2"],
    "ticket_id": 10,
    "status": "reserved",
    "event_id": 1
}
```

## API 2

### GET http://localhost:9000/api2/lugar/:lugarId

Capturar um lugar

```json
{
    "id": 1,
    "lugares": ["A1", "A2"],
    "tipo_ingresso": "inteira", //ou meia
    "estado": "disponivel", //disponivel, reservado
    "evento_id": 1
}
```

### POST http://localhost:9000/api2/lugar/:lugarId/reservar

Reservar um lugar

- Request
```json
{
    "lugares": ["A1", "A2"],
    "tipo_ingresso": "inteira", //ou meia
    "evento_id": 1
}
```

- Response
```json
{
    "id": 1,
    "lugares": ["A1", "A2"],
    "tipo_ingresso": "inteira", //ou meia
    "estado": "reservado", //disponivel, reservado
    "evento_id": 1
}
```

Dentro da API Gateway, configurar um token de autenticação forever, quando a API Gateway chamar a API 1 e API 2, enviar o token de autenticação como:
```
Authorization: Bearer <token>
```
