# TODO API Architecture

## Layer Communication Diagram

[TODO API Architecture v2](https://www.figma.com/online-whiteboard/create-diagram/ca9b7c5b-600e-4c79-94f1-5f198d743870?utm_source=other&utm_content=edit_in_figjam&oai_id=&request_id=ad8545f9-6447-4b5c-807d-cba00725a446)

## Project Structure

```
todo-api/
├── cmd
│   └── server
│       └── main.go              # Entry point (35 lines)
├── internal
│   ├── config
│   │   └── config.go            # Database configuration
│   ├── database
│   │   └── database.go          # SQLite connection
│   ├── handlers
│   │   └── todo                 # Handlers separated by action
│   │       ├── create_todo.go
│   │       ├── delete_todo.go
│   │       ├── get_todo.go
│   │       ├── get_todos.go
│   │       └── update_todo.go
│   ├── models
│   │   └── todo.go              # Data model
│   ├── repositories
│   │   └── todo_repository.go   # Data access layer
│   ├── server
│   │   └── server.go            # Server configuration
│   └── services
│       └── todo_service.go      # Business logic
├── migrations
│   └── 001_create_todos_table.sql # Migration script
├── pkg
│   └── utils
│       └── response.go          # HTTP response utilities
├── go.mod                       # Project dependencies
├── go.sum                       # Dependencies checksum
└── data/
    └── todos.db                 # SQLite database (auto-created)

15 directories, 16 files
```

## Request Flow

### 1. Presentation Layer (Handlers)

- **Location**: `internal/handlers/todo/`
- **Responsibility**: Handle HTTP requests and responses
- **Files**:
  - `get_todos.go` - List all TODOs
  - `get_todo.go` - Get TODO by ID
  - `create_todo.go` - Create new TODO
  - `update_todo.go` - Update existing TODO
  - `delete_todo.go` - Delete TODO

### 2. Service Layer (Business Logic)

- **Location**: `internal/services/todo_service.go`
- **Responsibility**: Validations and business rules
- **Validations**:
  - Title: 3-100 characters
  - Description: max 500 characters
  - ID: positive
  - Field normalization (trim)

### 3. Repository Layer (Data Access)

- **Location**: `internal/repositories/todo_repository.go`
- **Responsibility**: Database access
- **Operations**: CRUD with prepared statements
- **SQL**: Optimized queries with indexes

### 4. Database Layer

- **Location**: `internal/database/database.go`
- **Responsibility**: Connection and pool with SQLite
- **Features**: Connection pooling, migrations

## Dependencies and Injection

### Initialization Flow

```
main.go → server.NewServer()
         → config.NewConfig()
         → database.NewConnection()
         → runMigrations()
         → repositories.NewTodoRepository()
         → services.NewTodoService()
         → setupRoutes()
```

### Dependency Injection

- **Server** injects **Service**
- **Service** injects **Repository**
- **Repository** injects **Database**
- **Handlers** receive **Service** via parameter

## Utilities

### Response Utils

- **Location**: `pkg/utils/response.go`
- **Functions**:
  - `BadRequest()`, `NotFound()`, `InternalServerError()`
  - `OK()`, `Created()`, `Message()`
  - `HandleIDError()`, `HandleJSONError()`

## Data Model

### Todo Model

- **Location**: `internal/models/todo.go`
- **Fields**: ID, Title, Description, Completed, CreatedAt, UpdatedAt
- **Tags**: JSON for API, DB for database, validate for validation

## Configuration

### Database Config

- **Location**: `internal/config/config.go`
- **Path**: `data/todos.db` (auto-created)
- **Driver**: `modernc.org/sqlite`

## Architectural Patterns

### 1. Repository Pattern

- Data access separation
- Interface for testability
- Concrete implementation with SQLite

### 2. Service Layer Pattern

- Isolated business logic
- Centralized validations
- Repository orchestration

### 3. Dependency Injection

- Manual dependency construction
- Facilitates unit testing
- Low coupling between layers

### 4. Handler Pattern

- Pure functions as handlers
- Dependency injection via parameter
- Standardized responses with utils

## API Endpoints

```
GET    /api/v1/todos        → todo.GetTodos()
GET    /api/v1/todos/:id    → todo.GetTodo()
POST   /api/v1/todos        → todo.CreateTodo()
PUT    /api/v1/todos/:id    → todo.UpdateTodo()
DELETE /api/v1/todos/:id    → todo.DeleteTodo()
GET    /health              → health check
```

## Technical Features

### Performance

- Connection pooling (25 connections)
- Database indexes (title, completed)
- Prepared statements

### Security

- SQL injection prevention
- Input validation
- CORS configured

### Maintainability

- Modular and separated code
- Essential logs only
- Clear and conventional structure
