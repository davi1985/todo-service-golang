package todo

import (
	"github.com/gin-gonic/gin"
	"todo-api/internal/services"
	"todo-api/pkg/utils"
)

// GetTodos retrieves all todos
// @Summary Get all todos
// @Description Retrieves a list of all todos from the database
// @Tags todos
// @Accept  json
// @Produce json
// @Success 200 {array} object "List of todos"
// @Failure 500 {object} object "Internal server error"
// @Router /todos [get]
func GetTodos(service services.TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := service.GetAll()
		if err != nil {
			utils.InternalServerError(c, "Failed to get todos", err.Error())
			return
		}
		
		utils.OK(c, todos)
	}
}
