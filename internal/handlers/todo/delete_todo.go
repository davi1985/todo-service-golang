package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"todo-api/internal/services"
	"todo-api/pkg/utils"
)

func DeleteTodo(service services.TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			utils.HandleIDError(c, err)
			return
		}
		
		if err := service.Delete(id); err != nil {
			utils.NotFound(c, "Failed to delete todo", err.Error())
			return
		}
		
		utils.Message(c, "Todo deleted successfully")
	}
}
