package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"todo-api/internal/services"
	"todo-api/pkg/utils"
)

// DeleteTodo deletes a todo by ID
// @Summary Delete a todo
// @Description Deletes an existing todo by its ID
// @Tags todos
// @Accept  json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} object "Todo deleted successfully"
// @Failure 400 {object} object "Invalid ID format"
// @Failure 404 {object} object "Todo not found"
// @Router /todos/{id} [delete]
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
