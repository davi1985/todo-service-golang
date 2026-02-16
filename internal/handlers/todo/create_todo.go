package todo

import (
	"github.com/gin-gonic/gin"
	"todo-api/internal/models"
	"todo-api/internal/services"
	"todo-api/pkg/utils"
)

// CreateTodo creates a new todo
// @Summary Create a new todo
// @Description Creates a new todo with the provided title and description
// @Tags todos
// @Accept  json
// @Produce json
// @Param todo body object true "Todo data"
// @Success 201 {object} object "Todo created successfully"
// @Failure 400 {object} object "Invalid request body or validation error"
// @Failure 500 {object} object "Internal server error"
// @Router /todos [post]
func CreateTodo(service services.TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todo models.Todo
		
		if err := c.ShouldBindJSON(&todo); err != nil {
			utils.HandleJSONError(c, err)
			return
		}
		
		if err := service.Create(&todo); err != nil {
			utils.BadRequest(c, "Failed to create todo", err.Error())
			return
		}
		
		utils.Created(c, todo)
	}
}
