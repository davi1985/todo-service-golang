package server

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"todo-api/internal/config"
	"todo-api/internal/database"
	"todo-api/internal/handlers/todo"
	"todo-api/internal/repositories"
	"todo-api/internal/services"
)


type Server struct {
	config  *config.Config
	db      *database.DB
	service services.TodoService
	router  *gin.Engine
}

func NewServer() (*Server, error) {
	cfg := config.NewConfig()

	db, err := database.NewConnection(&cfg.Database)
	if err != nil {
		return nil, err
	}
	
	if err := runMigrations(db); err != nil {
		return nil, err
	}
	
	repo := repositories.NewTodoRepository(db)
	service := services.NewTodoService(repo)
	
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	
	setupRoutes(r, service)
	
	return &Server{
		config:  cfg,
		db:      db,
		service: service,
		router:  r,
	}, nil
}

func (s *Server) Start(port string) error {
	log.Printf("TODO API listening on %s", port)
	return s.router.Run(port)
}

func (s *Server) Close() error {
	return s.db.Close()
}

func runMigrations(db *database.DB) error {
	migrationSQL, err := ioutil.ReadFile("migrations/001_create_todos_table.sql")
	if err != nil {
		return err
	}
	
	return db.RunMigration(string(migrationSQL))
}

func setupRoutes(r *gin.Engine, service services.TodoService) {
	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	api := r.Group("/api/v1")
	{
		todos := api.Group("/todos")
		{
			todos.GET("", todo.GetTodos(service))
			todos.GET("/:id", todo.GetTodo(service))
			todos.POST("", todo.CreateTodo(service))
			todos.PUT("/:id", todo.UpdateTodo(service))
			todos.DELETE("/:id", todo.DeleteTodo(service))
		}
	}
	
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"service": "todo-api",
		})
	})
}
