package todos

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Todo struct {
	Id       uint64 `json:"id"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	CreateAt string `json:"create_at"`
}

type TodosHandler struct {
	todo []Todo
}

func (t *TodosHandler) Initialize() {
	t.todo = append(t.todo, Todo{
		Id:       0,
		Title:    "read book",
		Detail:   "read golang book",
		CreateAt: "20/20/2020",
	}, Todo{
		Id:       1,
		Title:    "learning golang",
		Detail:   "learn golang from go tour",
		CreateAt: "20/20/2020",
	})
}

func (t *TodosHandler) GetAllTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, t.todo)
}

func (t *TodosHandler) GetTodos(c echo.Context) error {
	id, error := strconv.ParseUint(c.Param("id"), 10, 64)

	if error != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	todoIndex := -1

	for i := 0; i < len(t.todo); i++ {
		if t.todo[i].Id == id {
			todoIndex = i
		}
	}

	if todoIndex == -1 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, t.todo[todoIndex])
}
