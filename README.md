# go-api

API REST em Go (gin) com cadastro e login de usuários, usando autenticação JWT via cookie.

## Tecnologias

- [Go](https://go.dev/) 1.26
- [Gin](https://github.com/gin-gonic/gin) — framework HTTP
- [PostgreSQL](https://www.postgresql.org/) — banco de dados (driver `lib/pq`)
- [golang-jwt](https://github.com/golang-jwt/jwt) — geração e validação de tokens
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — hash de senhas
- [validator](https://github.com/go-playground/validator) — validação das requisições

## Estrutura

```
cmd/api/              → ponto de entrada (main.go)
internal/
  configuracao/       → leitura das variáveis de ambiente (.env)
  inicializacao/      → criação do usuário admin ao subir a aplicação
  roteador/           → definição das rotas
  middleware/         → middleware de autenticação (cookie + JWT)
  manipulador/        → handlers HTTP e formato padrão de resposta
  servico/            → regras de negócio
  repositorio/        → acesso ao banco de dados
  modelo/             → entidades (Usuario, Role, ClaimsUsuario)
  requisicao/         → structs de entrada das requisições
```

## Como rodar

### 1. Pré-requisitos

- Go 1.26 ou superior
- PostgreSQL rodando

### 2. Banco de dados

Crie o banco e a tabela de usuários:

```sql
CREATE DATABASE go_api;

\c go_api

CREATE TABLE usuarios (
    id    UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    senha TEXT NOT NULL,
    role  TEXT NOT NULL DEFAULT 'user' CHECK (role IN ('user', 'admin'))
);
```

### 3. Variáveis de ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
# Servidor
PORTA=8080

# Banco de dados
HOST_BD=localhost
PORTA_BD=5432
USUARIO_BD=postgres
SENHA_BD=sua_senha_do_banco
NOME_BD=go_api

# JWT — use um valor longo e aleatório (ex.: openssl rand -base64 32)
SEGREDO_JWT=troque_por_um_segredo_forte

# Usuário admin criado automaticamente na inicialização
EMAIL_ADMIN=admin@exemplo.com
SENHA_ADMIN=troque_esta_senha
```

| Variável | Obrigatória | Padrão | Descrição |
|---|---|---|---|
| `PORTA` | não | `8080` | Porta do servidor HTTP |
| `HOST_BD` | não | `localhost` | Host do PostgreSQL |
| `PORTA_BD` | não | `5432` | Porta do PostgreSQL |
| `USUARIO_BD` | **sim** | — | Usuário do banco |
| `SENHA_BD` | **sim** | — | Senha do banco |
| `NOME_BD` | sim | — | Nome do banco |
| `SEGREDO_JWT` | sim | — | Chave usada para assinar os tokens |
| `EMAIL_ADMIN` | **sim** | — | E-mail do admin criado na inicialização |
| `SENHA_ADMIN` | **sim** | — | Senha do admin (mínimo 6 caracteres) |

> ⚠️ O `.env` está no `.gitignore` e **nunca** deve ser commitado.

### 4. Executar

```bash
go mod download
go run ./cmd/api
```

A API sobe em `http://localhost:8080`.

## Endpoints

### `POST /auth/register`

Cadastra um novo usuário (sempre com role `user`).

```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email": "usuario@exemplo.com", "senha": "123456"}'
```

Resposta `201 Created`:

```json
{
  "data": "c0a8012e-5b1f-4c7e-9d3a-2f6e8b1a4d90",
  "erro": "",
  "isSuccess": true
}
```

### `POST /auth/login`

Autentica o usuário e define o cookie `auth` com o token JWT (válido por 24h).

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "usuario@exemplo.com", "senha": "123456"}' \
  -c cookies.txt
```

Resposta `200 OK`:

```json
{
  "data": null,
  "erro": "",
  "isSuccess": true
}
```

### Formato de erro

Todas as respostas de erro seguem o mesmo formato:

```json
{
  "data": null,
  "erro": "Credenciais inválidas",
  "isSuccess": false
}
```

## Validações

- `email`: obrigatório e em formato de e-mail válido
- `senha`: obrigatória, com no mínimo 6 caracteres
