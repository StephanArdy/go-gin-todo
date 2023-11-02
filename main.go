package main

import (
	"go-gin-todo/app"
	"go-gin-todo/db"
	"go-gin-todo/models"
	"html/template"

	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.Default()

	conn := db.InitDB()

	router.LoadHTMLGlob("templates/*")

	handler := app.New(conn)

	router.GET("/tasks/:id", handler.GetTodoById)
	router.POST("/tasks", handler.CreateTodo)
	router.POST("/updateTasks/:id", handler.UpdateTodo)
	router.POST("/deleteTasks/:id", handler.DeleteTodo)

	router.GET("/todo", func(c *gin.Context) {
		var todos []models.Todo
		handler.DB.Find(&todos)
		render(c, "todo.html", gin.H{"data": todos})
	})

	router.Run()
}

func render(c *gin.Context, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles("templates/" + templateName)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := tmpl.Execute(c.Writer, data); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
}
