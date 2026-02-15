package todo

import (
	"github.com/gin-gonic/gin"
	"todo-api/internal/services"
	"todo-api/pkg/utils"
)

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
