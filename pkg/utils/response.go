package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func BadRequest(c *gin.Context, message, details string) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Error:   message,
		Details: details,
	})
}

func NotFound(c *gin.Context, message, details string) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Error:   message,
		Details: details,
	})
}

func InternalServerError(c *gin.Context, message, details string) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Error:   message,
		Details: details,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Message(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func HandleIDError(c *gin.Context, err error) {
	BadRequest(c, "Invalid ID format", "ID must be a number")
}

func HandleJSONError(c *gin.Context, err error) {
	BadRequest(c, "Invalid JSON format", err.Error())
}