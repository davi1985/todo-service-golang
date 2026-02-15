package todo

import (
	"github.com/gin-gonic/gin"
	"todo-api/internal/models"
	"todo-api/internal/services"
	"todo-api/pkg/utils"
)

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
