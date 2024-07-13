package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{"1", "Clean Room", false},
	{"2", "Read Book", false},
	{"3", "Record Video", false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)

}

func getTodoByID(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	todo.Completed = !todo.Completed

	c.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")

}
