package main

import (
	"net/http"

	"go-cicd/todos"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	t := todos.TodosHandler{}
	t.Initialize()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello go2")
	})

	e.GET("/todos", t.GetAllTodos)
	e.GET("/todos/:id", t.GetTodos)

	e.Logger.Fatal(e.Start(":1234"))
}
