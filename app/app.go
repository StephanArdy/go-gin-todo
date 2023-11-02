package app

import (
	"fmt"
	"go-gin-todo/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h *Handler) GetTodoById(c *gin.Context) {
	todoId := c.Param("id")
	
	var todos models.Todo

	if err := h.DB.Find(&todos,"id=?" ,todoId).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.HTML(http.StatusOK, "detailTodo.html",gin.H{
		"title": todos.Title,
		"data":    todos,
	})
}

func (h *Handler) CreateTodo(c *gin.Context) {
	var todos models.Todo

	if err := c.ShouldBind(&todos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	todos.CreatedAt = time.Now()

	if err := h.DB.Create(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create todo",
			"error":   err.Error(),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, fmt.Sprint("/todo"))
}

func (h *Handler) UpdateTodo(c *gin.Context) {
	todoId := c.Param("id")

	var todos models.Todo

	if err := h.DB.Find(&todos,"id=?",todoId).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var reqTodo = todos

	c.Bind(&reqTodo)

	h.DB.Model(&todos).Where("id=?", todoId).Updates(reqTodo)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprint("/todo"))
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	todoId := c.Param("id")

	var todos models.Todo

	if err := h.DB.Find(&todos,"id=?" ,todoId).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	h.DB.Delete(&todos)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprint("/todo"))

}
