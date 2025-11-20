# Task Manager Go

Um gerenciador de tarefas simples construído com Go, utilizando Gin para o framework web, GORM para ORM e PostgreSQL como banco de dados.

## Funcionalidades

- **Registro e Login de Usuários**: Criação de contas e autenticação via JWT.
- **Gerenciamento de Tarefas**: Criar, atualizar status (pendente, fazendo, concluída) e deletar tarefas.
- **Perfil do Usuário**: Visualizar informações do perfil e lista de tarefas associadas.
- **Autenticação JWT**: Proteção de rotas sensíveis com middleware de autenticação.

## Tecnologias Utilizadas

- **Go 1.25.3**: Linguagem de programação.
- **Gin**: Framework web para Go.
- **GORM**: ORM para Go.
- **PostgreSQL**: Banco de dados relacional.
- **JWT**: Para autenticação de usuários.
- **Docker**: Para containerização da aplicação

## Pré-requisitos

- Go 1.25.3 ou superior
- PostgreSQL
- Docker (opcional, para execução via container)

## Instalaçãoo

1. Clone o repositório:
   ```bash
   git clone https://github.com/rlevidev/taskmanager-go.git
   cd taskmanager-go
   ```

2. Instale as dependências:
   ```bash
   go mod download
   ```

3. Configure o banco de dados:
   - Crie um banco de dados PostgreSQL.
   - Copie o arquivo `.env.example` para `.env` e configure as variáveis de ambiente:
     ```
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=your_user
     DB_PASSWORD=your_password
     DB_NAME=taskmanager
     JWT_SECRET=your_jwt_secret
     ```

4. Execute as migrações (se aplicável) ou inicie a aplicação.

## Executando a Aplicação

### Localmente

1. Certifique-se de que o PostgreSQL está rodando.
2. Execute o comando:
   ```bash
   go run src/cmd/main.go
   ```

A aplicação estará disponível em `http://localhost:8080`.

### Com Docker

1. Construa e execute com Docker Compose:
   ```bash
   docker-compose up --build
   ```

## API Endpoints

### Autenticação

#### Registrar Usuário
- **POST** `/api/v1/users/register`
- **Corpo da Requisição**:
  ```json
  {
    "name": "Nome do Usuário",
    "email": "usuario@example.com",
    "password": "senha123"
  }
  ```
- **Resposta**:
  ```json
  {
    "message": "Usuário criado com sucesso"
  }
  ```

#### Login
- **POST** `/api/v1/users/login`
- **Corpo da Requisição**:
  ```json
  {
    "email": "usuario@example.com",
    "password": "senha123"
  }
  ```
- **Resposta**:
  ```json
  {
    "token": "jwt_token_aqui"
  }
  ```

### Perfil do Usuário (Requer Autenticação)

#### Obter Informações do Perfil
- **GET** `/api/v1/users/profile/info`
- **Cabeçalhos**:
  ```
  Authorization: Bearer <token_jwt>
  Content-Type: application/json
  ```
- **Resposta**:
  ```json
  {
    "user_id": "b82d5b6a-4fd2-4e3f-8e7b-182b3d8b83d3",
    "user_name": "rlevidev",
    "user_create_at": "2025-10-20T22:00:00Z",
    "tasks": [
      {
        "task_id": "f6b08a52-52c9-4b3b-9c4e-b9a821b9dbf",
        "task_title": "Estudar GoLang",
        "task_status": "pending"
      },
      {
        "task_id": "8d2019a7-122a-44d5-8a48-2090e72a6d42",
        "task_title": "Testar API com Postman",
        "task_status": "done"
      }
    ]
  }
  ```

#### Criar Tarefa
- **POST** `/api/v1/users/profile/createtask`
- **Cabeçalhos**:
  ```
  Authorization: Bearer <token_jwt>
  Content-Type: application/json
  ```
- **Corpo da Requisição**:
  ```json
  {
    "task_title": "Estudar GoLang",
    "task_description": "Estudar documentação do JWT"
  }
  ```
- **Resposta**:
  ```json
  {
    "task_id": "f6b08a52-52c9-4b3b-9c4e-b9a821b9dbf3",
    "task_title": "Estudar GoLang",
    "task_description": "Estudar documentação do JWT",
    "task_status": "pending",
    "user_id": "b82d5b6a-4fd2-4e3f-8e7b-182b3d8b83d3",
    "task_created_at": "2025-10-20T22:00:00Z"
  }
  ```

#### Atualizar Status da Tarefa para "Fazendo"
- **PUT** `/api/v1/users/profile/:task_id/doingtask`
- **Cabeçalhos**:
  ```
  Authorization: Bearer <token_jwt>
  ```

#### Finalizar Tarefa
- **PUT** `/api/v1/users/profile/:task_id/finishtask`
- **Cabeçalhos**:
  ```
  Authorization: Bearer <token_jwt>
  ```

#### Deletar Tarefa
- **DELETE** `/api/v1/users/profile/:task_id/deletetask`
- **Cabeçalhos**:
  ```
  Authorization: Bearer <token_jwt>
  ```

## Estrutura do Projeto

```
taskmanager-go/
├── src/
│   ├── cmd/
│   │   └── main.go              # Ponto de entrada da aplicação
│   ├── config/
│   │   ├── database/
│   │   │   └── database.go      # Configuração do banco de dados
│   │   ├── middleware/
│   │   │   └── auth_middleware.go # Middleware de autenticação
│   │   ├── resterr/
│   │   │   └── rest_err.go      # Tratamento de erros REST
│   │   └── validation/
│   │       └── user_validate.go # Validações de usuário
│   ├── controllers/
│   │   ├── request/
│   │   │   ├── task_request.go  # Estruturas de requisição para tarefas
│   │   │   └── user_request.go  # Estruturas de requisição para usuários
│   │   ├── response/
│   │   │   ├── profile_response.go # Respostas do perfil
│   │   │   └── user_response.go # Respostas de usuário
│   │   ├── task_create.go       # Controller para criação de tarefas
│   │   ├── task_update.go       # Controller para atualização de tarefas
│   │   ├── user_create.go       # Controller para criação de usuários
│   │   ├── user_login.go        # Controller para login
│   │   └── user_profile.go      # Controller para perfil do usuário
│   ├── models/
│   │   ├── task_domain.go       # Modelo de domínio para tarefas
│   │   └── user_domain.go       # Modelo de domínio para usuários
│   ├── routes/
│   │   └── routes.go            # Definição das rotas
│   └── services/
│       ├── jwt_service.go       # Serviço de JWT
│       ├── task_create.go       # Serviço para criação de tarefas
│       ├── task_update.go       # Serviço para atualização de tarefas
│       ├── user_create.go       # Serviço para criação de usuários
│       ├── user_login.go        # Serviço para login
│       └── user_profile.go      # Serviço para perfil do usuário
├── collections/                 # Coleções do Postman para testes
├── .env.example                 # Exemplo de variáveis de ambiente
├── docker-compose.yml           # Configuração do Docker Compose
├── Dockerfile                   # Dockerfile para a aplicação
├── go.mod                       # Dependências do Go
├── go.sum                       # Sums das dependências
└── README.md                    # Este arquivo
```

## Testando a API

Você pode usar as coleções do Postman localizadas na pasta `collections/` para testar os endpoints da API.

Exemplos de arquivos:
- `Create_User.http`
- `Login_User.http`
- `Get_Profile.http`
- `Create_Task.http`
- `Doing_Task.http`
- `Finish_Task.http`
- `Delete_Task.http`

## Contribuição

1. Fork o projeto.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`).
4. Push para a branch (`git push origin feature/nova-feature`).
5. Abra um Pull Request.

## Licença

Este projeto está licenciado sob a MIT License - veja o arquivo LICENSE para detalhes.
