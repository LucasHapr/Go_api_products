# 🛍️ Go API Products

Uma API RESTful desenvolvida em Go para gerenciamento de produtos, seguindo os princípios de Clean Architecture e utilizando PostgreSQL como banco de dados.

## 📋 Descrição

Esta API fornece endpoints completos para operações CRUD (Create, Read, Update, Delete) de produtos, implementada com boas práticas de desenvolvimento, separação de responsabilidades e arquitetura em camadas.

## 🚀 Tecnologias

- **Go 1.25.3** - Linguagem de programação
- **Gin Framework** - Web framework HTTP
- **PostgreSQL 12** - Banco de dados relacional
- **Docker & Docker Compose** - Containerização
- **lib/pq** - Driver PostgreSQL para Go

## 🏗️ Arquitetura

O projeto segue o padrão **Clean Architecture** com as seguintes camadas:

```
go-api/
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── controller/
│   └── product_controller.go   # Handlers HTTP
├── usecase/
│   └── product_usecase.go      # Regras de negócio
├── repository/
│   └── product_repository.go   # Acesso aos dados
├── model/
│   ├── product.go              # Entidade Product
│   └── response.go             # Estruturas de resposta
├── db/
│   └── conn.go                 # Conexão com o banco de dados
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── go.sum
```

### Camadas da Aplicação

- **Controller**: Recebe requisições HTTP e retorna respostas
- **UseCase**: Contém a lógica de negócio da aplicação
- **Repository**: Gerencia a persistência de dados no banco
- **Model**: Define as estruturas de dados
- **DB**: Gerencia a conexão com o banco de dados

## 📦 Instalação

### Pré-requisitos

- Go 1.25.3 ou superior
- Docker e Docker Compose
- PostgreSQL 12 (caso não use Docker)

### 1. Clone o repositório

```bash
git clone https://github.com/LucasHapr/Go_api_products.git
cd Go_api_products
```

### 2. Instale as dependências

```bash
go mod download
```

### 3. Configure o banco de dados

#### Usando Docker Compose (Recomendado)

```bash
docker-compose up -d
```

#### Configuração Manual

Se preferir não usar Docker, configure o PostgreSQL localmente:

```sql
-- Crie o banco de dados
CREATE DATABASE postgres;

-- Crie a tabela de produtos
CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);
```

### 4. Execute a aplicação

```bash
cd cmd
go run main.go
```

Ou usando Docker:

```bash
docker build -t go-api .
docker run -p 8080:8080 go-api
```

A API estará disponível em: `http://localhost:8080`

## 🔌 Endpoints da API

### Base URL
```
http://localhost:8080
```

### Listar todos os produtos

```http
GET /products
```

**Resposta de Sucesso (200 OK):**
```json
[
  {
    "id_product": 1,
    "name": "Notebook",
    "price": 3500.00
  },
  {
    "id_product": 2,
    "name": "Mouse",
    "price": 50.00
  }
]
```

### Buscar produto por ID

```http
GET /products/:id
```

**Parâmetros:**
- `id` (path) - ID do produto

**Resposta de Sucesso (200 OK):**
```json
{
  "id_product": 1,
  "name": "Notebook",
  "price": 3500.00
}
```

**Resposta de Erro (404 Not Found):**
```json
{
  "Message": "Product not found"
}
```

### Criar novo produto

```http
POST /products
```

**Body:**
```json
{
  "name": "Teclado Mecânico",
  "price": 450.00
}
```

**Resposta de Sucesso (201 Created):**
```json
{
  "id_product": 3,
  "name": "Teclado Mecânico",
  "price": 450.00
}
```

### Atualizar produto

```http
PUT /products/:id
```

**Parâmetros:**
- `id` (path) - ID do produto

**Body:**
```json
{
  "name": "Teclado Mecânico RGB",
  "price": 550.00
}
```

**Resposta de Sucesso (200 OK):**
```json
{
  "id_product": 3,
  "name": "Teclado Mecânico RGB",
  "price": 550.00
}
```

### Deletar produto

```http
DELETE /products/:id
```

**Parâmetros:**
- `id` (path) - ID do produto

**Resposta de Sucesso (200 OK):**
```json
{
  "Message": "Product deleted successfully"
}
```

## ⚙️ Configuração

As configurações do banco de dados estão localizadas em `db/conn.go`:

```go
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "1234"
    dbname   = "postgres"
)
```
## 🐳 Docker

### Construir a imagem

```bash
docker build -t go-api-products .
```

### Executar o container

```bash
docker run -p 8080:8080 go-api-products
```

### Usar Docker Compose

```bash
# Iniciar serviços
docker-compose up -d

# Parar serviços
docker-compose down

# Ver logs
docker-compose logs -f
```

## 📝 Estrutura do Banco de Dados

### Tabela: product

| Campo        | Tipo         | Descrição              |
|--------------|--------------|------------------------|
| id           | SERIAL       | Chave primária         |
| product_name | VARCHAR(255) | Nome do produto        |
| price        | DECIMAL(10,2)| Preço do produto       |

## 👤 Autor

**LucasHapr**

- GitHub: [@LucasHapr](https://github.com/LucasHapr)
- Repositório: [Go_api_products](https://github.com/LucasHapr/Go_api_products)

