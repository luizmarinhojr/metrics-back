# Metricas-back

Esta é uma API de métricas para a aplicação Metrics. A chave da API é armazenada através de cookie, portanto, para conseguir a chave da API basta que faça o login no endpoint "login".

## Endpoints

### Autenticação

#### `POST /api/v1/login`

Autentica um usuário e define o cookie de sessão com o token JWT.

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

**Responses:**

*   `200 OK`: Login bem-sucedido
*   `400 Bad Request`: Erro de validação
*   `401 Unauthorized`: Não autorizado

#### `POST /api/v1/signup`

Cria um novo usuário com os dados fornecidos.

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

**Responses:**

*   `201 Created`: ID do usuário criado
*   `400 Bad Request`: Erro de validação
*   `500 Internal Server Error`: Erro interno do servidor

#### `GET /api/v1/validate-token`

Valida o token JWT contido no cookie.

**Responses:**

*   `200 OK`: Token válido

#### `GET /api/v1/logout`

Remove o cookie de sessão para efetuar o logout. Requer autenticação.

**Responses:**

*   `200 OK`: Logout bem-sucedido

### Brokers

#### `GET /api/v1/brokers`

Retorna a lista de todos os brokers. Requer autenticação.

**Responses:**

*   `200 OK`: Lista de brokers
*   `401 Unauthorized`: Não autorizado

#### `GET /api/v1/broker/id`

Retorna um broker específico pelo ID. Requer autenticação.

**Request Body:**

```json
{
  "id": 1
}
```

**Responses:**

*   `200 OK`: Broker encontrado
*   `400 Bad Request`: Erro de validação
*   `401 Unauthorized`: Não autorizado
*   `404 Not Found`: Broker não encontrado

#### `GET /api/v1/broker/name`

Retorna um broker específico pelo nome. Requer autenticação.

**Request Body:**

```json
{
  "nome": "Nome do Broker"
}
```

**Responses:**

*   `200 OK`: Broker encontrado
*   `400 Bad Request`: Erro de validação
*   `401 Unauthorized`: Não autorizado
*   `404 Not Found`: Broker não encontrado

#### `POST /api/v1/broker`

Cria um novo broker com os dados fornecidos. Requer autenticação.

**Request Body:**

```json
{
  "nome": "Nome do Broker",
  "user_id": 1
}
```

**Responses:**

*   `201 Created`: ID do broker criado
*   `400 Bad Request`: Erro de validação
*   `401 Unauthorized`: Não autorizado
*   `500 Internal Server Error`: Erro interno do servidor

### Métricas

#### `GET /api/v1/metrics`

Retorna a lista de todas as métricas.

**Responses:**

*   `200 OK`: Lista de métricas

#### `GET /api/v1/metric/id`

Retorna uma métrica específica pelo ID. Requer autenticação.

**Request Body:**

```json
{
  "id": 1
}
```

**Responses:**

*   `200 OK`: Métrica encontrada
*   `400 Bad Request`: Erro de validação
*   `401 Unauthorized`: Não autorizado
*   `404 Not Found`: Métrica não encontrada

#### `POST /api/v1/metric`

Cria uma nova métrica com os dados fornecidos. Requer autenticação.

**Request Body:**

```json
{
  "data": "2024-07-26",
  "leads_recebidos": 10,
  "ligacoes": 20,
  "agendamentos": 5,
  "visitas": 3,
  "propostas": 2,
  "negociacoes": 1,
  "vendas": 1,
  "captacoes": 4,
  "espontaneo": 1
}
```

**Responses:**

*   `201 Created`: ID da métrica criada
  *   `400 Bad Request`: Erro de validação
  *   `401 Unauthorized`: Não autorizado
  *   `500 Internal Server Error`: Erro interno do servidor

## Como Executar

1.  Clone o repositório:

    ```bash
    git clone https://github.com/luizmarinhojr/metrics-back.git
    ```

2.  Instale as dependências:

    ```bash
    go mod tidy
    ```

3.  Execute a aplicação:

    ```bash
    go run cmd/main.go
    ```
