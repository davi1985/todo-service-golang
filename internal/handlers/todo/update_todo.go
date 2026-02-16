package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"todo-api/internal/models"
	"todo-api/internal/services"
	"todo-api/pkg/utils"
)

// UpdateTodo updates an existing todo
// @Summary Update a todo
// @Description Updates an existing todo with the provided ID
// @Tags todos
// @Accept  json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body models.Todo true "Updated todo data"
// @Success 200 {object} models.Todo "Todo updated successfully"
// @Failure 400 {object} object "Invalid ID format or request body"
// @Failure 500 {object} object "Internal server error"
// @Router /todos/{id} [put]
func UpdateTodo(service services.TodoService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			utils.HandleIDError(c, err)
			return
		}
		
		var todo models.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			utils.HandleJSONError(c, err)
			return
		}
		
		todo.ID = id
		
		if err := service.Update(&todo); err != nil {
			utils.BadRequest(c, "Failed to update todo", err.Error())
			return
		}
		
		utils.OK(c, todo)
	}
}
