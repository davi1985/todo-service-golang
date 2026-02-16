package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"todo-api/internal/services"
	"todo-api/pkg/utils"
)

// GetTodo retrieves a single todo by ID
// @Summary Get todo by ID
// @Description Retrieves a specific todo by its ID
// @Tags todos
// @Accept  json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} object "Todo details"
// @Failure 400 {object} object "Invalid ID format"
// @Failure 404 {object} object "Todo not found"
// @Router /todos/{id} [get]
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
