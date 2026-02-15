# TODO API Service

A clean and scalable REST API for managing TODO items built with Go and SQLite.

## 🚀 Features

- ✅ **REST API** with full CRUD operations
- ✅ **SQLite Database** with migrations and indexes
- ✅ **Clean Architecture** with separated layers
- ✅ **Input Validation** with business rules
- ✅ **Standardized Responses** with proper HTTP status codes
- ✅ **Graceful Shutdown** with signal handling
- ✅ **Connection Pooling** for performance
- ✅ **CORS Support** for web applications

## 📋 Requirements

- Go 1.21 or higher
- SQLite (included with modernc.org/sqlite driver)

## 🛠️ Installation

1. Clone the repository:
```bash
git clone https://github.com/davi1985/todo-service-golang.git
cd todo-service-golang
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8082`

## 🏗️ Architecture

This project follows clean architecture principles with clear separation of concerns:

```
cmd/server/main.go              # Entry point (35 lines)
├── internal/
│   ├── config/                 # Database configuration
│   ├── database/               # SQLite connection
│   ├── handlers/todo/          # HTTP handlers (separated by action)
│   ├── models/                 # Data models
│   ├── repositories/           # Data access layer
│   ├── server/                 # Server configuration
│   └── services/               # Business logic
├── pkg/utils/                  # HTTP response utilities
├── migrations/                 # Database migrations
└── .spec/                      # Architecture documentation
```

### Layer Communication

[View Architecture Diagram](https://www.figma.com/online-whiteboard/create-diagram/ca9b7c5b-600e-4c79-94f1-5f198d743870)

## 📚 API Documentation

### Base URL
```
http://localhost:8082/api/v1
```

### Endpoints

#### Get all TODOs
```http
GET /api/v1/todos
```

**Response:**
```json
[
  {
    "id": 1,
    "title": "Buy groceries",
    "description": "Milk, eggs, bread",
    "completed": false,
    "created_at": "2026-02-15T10:30:00Z",
    "updated_at": "2026-02-15T10:30:00Z"
  }
]
```

#### Get TODO by ID
```http
GET /api/v1/todos/{id}
```

#### Create TODO
```http
POST /api/v1/todos
Content-Type: application/json

{
  "title": "New task",
  "description": "Optional description",
  "completed": false
}
```

#### Update TODO
```http
PUT /api/v1/todos/{id}
Content-Type: application/json

{
  "title": "Updated task",
  "description": "Updated description",
  "completed": true
}
```

#### Delete TODO
```http
DELETE /api/v1/todos/{id}
```

#### Health Check
```http
GET /health
```

**Response:**
```json
{
  "status": "ok",
  "service": "todo-api"
}
```

## 🎯 Validation Rules

- **Title**: Required, 3-100 characters
- **Description**: Optional, max 500 characters
- **Completed**: Boolean, defaults to `false`
- **ID**: Positive integer for update/delete operations

## 🚨 Error Responses

All errors follow a consistent format:

```json
{
  "error": "Error description",
  "details": "Additional error details"
}
```

### Common HTTP Status Codes

- `200 OK` - Successful request
- `201 Created` - Resource created successfully
- `400 Bad Request` - Validation error or invalid input
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## 🛠️ Development

### Project Structure

```
todo-api/
├── cmd/server/main.go              # Application entry point
├── internal/
│   ├── config/config.go            # Database configuration
│   ├── database/database.go        # SQLite connection and pooling
│   ├── handlers/todo/              # HTTP handlers separated by action
│   │   ├── get_todos.go
│   │   ├── get_todo.go
│   │   ├── create_todo.go
│   │   ├── update_todo.go
│   │   └── delete_todo.go
│   ├── models/todo.go              # Todo data model
│   ├── repositories/todo_repository.go  # Data access layer
│   ├── server/server.go            # Server setup and routing
│   └── services/todo_service.go    # Business logic layer
├── pkg/utils/response.go           # HTTP response utilities
├── migrations/001_create_todos_table.sql  # Database schema
└── .spec/architecture-diagram.md   # Architecture documentation
```

### Running Tests

```bash
go test ./...
```

### Building for Production

```bash
go build -o todo-api cmd/server/main.go
./todo-api
```

### Environment Variables

- `GIN_MODE`: Set to `release` for production (default: `debug`)

## 📊 Database Schema

```sql
CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL CHECK (length(title) >= 3),
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX idx_todos_title ON todos(title);
CREATE INDEX idx_todos_completed ON todos(completed);

-- Trigger for automatic updated_at
CREATE TRIGGER update_todos_updated_at 
    AFTER UPDATE ON todos
    FOR EACH ROW
    BEGIN
        UPDATE todos SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
    END;
```

## 🔧 Technologies Used

- **Go** - Programming language
- **Gin Framework** - HTTP web framework
- **SQLite** - Database engine
- **modernc.org/sqlite** - Pure Go SQLite driver
- **FigJam** - Architecture diagrams

## 📝 Design Patterns

- **Repository Pattern** - Data access abstraction
- **Service Layer Pattern** - Business logic encapsulation
- **Dependency Injection** - Manual dependency construction
- **Handler Pattern** - HTTP request handling

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- [Architecture Documentation](.spec/architecture-diagram.md)
- [Interactive Diagram](https://www.figma.com/online-whiteboard/create-diagram/ca9b7c5b-600e-4c79-94f1-5f198d743870)
- [GitHub Repository](https://github.com/davi1985/todo-service-golang)

---

**Built with ❤️ using Go and clean architecture principles**
