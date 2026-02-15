package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"todo-api/internal/services"
	"todo-api/pkg/utils"
)

func GetTodo(service services.TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			utils.HandleIDError(c, err)
			return
		}
		
		todo, err := service.GetByID(id)
		if err != nil {
			utils.NotFound(c, "Todo not found", err.Error())
			return
		}
		
		utils.OK(c, todo)
	}
}
