package main

import (
	"context"
	"net/http"

	"github.com/Sourjaya/todo_htmx/dto"
	"github.com/Sourjaya/todo_htmx/templates"
	"github.com/Sourjaya/todo_htmx/templates/components"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func filterByID(todos []*dto.TodoCardDto, id string) (out []*dto.TodoCardDto) {
	for _, todo := range todos {
		if todo.ID != id {
			out = append(out, todo)
		}
	}
	return out
}

func findByID(todos []*dto.TodoCardDto, id string) *dto.TodoCardDto {
	for _, todo := range todos {
		if todo.ID == id {
			return todo
		}
	}
	return nil
}

func main() {
	e := echo.New()

	todos := []*dto.TodoCardDto{{
		ID:      uuid.New().String(),
		Content: "1st todo",
	},
		{
			ID:      uuid.New().String(),
			Content: "2nd todo",
		},
	}

	e.PUT("/todos/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
		}
		todo := findByID(todos, id)
		if todo == nil {
			return echo.NewHTTPError(http.StatusNotFound, "Todo card not found")
		}
		todo.Content = c.FormValue("edit-todo-input")
		component := components.TodoCard(*todo)
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.DELETE("/todos/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
		}
		todos = filterByID(todos, id)
		component := components.TodoCards(todos)
		return component.Render(context.Background(), c.Response().Writer)
	})

	// component.Render(context.Background(), os.Stdout)
	e.GET("/", func(c echo.Context) error {
		component := templates.Index(todos)
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.POST("/todos", func(c echo.Context) error {
		text := c.FormValue("add-todo-input")
		if text == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid text")
		}
		todos = append(todos, &dto.TodoCardDto{ID: uuid.New().String(),
			Content: text,
		})
		component := components.TodoCardsWithBtn(todos)
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/components", func(c echo.Context) error {
		t := c.QueryParam("type")
		id := c.QueryParam("id")
		switch t {
		case "add-todo":
			component := components.AddTodoInput()
			return component.Render(context.Background(), c.Response().Writer)
		case "add-todo-btn":
			component := components.AddTodoButton()
			return component.Render(context.Background(), c.Response().Writer)
		case "edit-todo-input":
			todo := findByID(todos, id)
			component := components.EditTodoInput(todo)
			return component.Render(context.Background(), c.Response().Writer)
		case "edit-todo-btn":
			todo := findByID(todos, id)
			component := components.TodoCard(*todo)
			return component.Render(context.Background(), c.Response().Writer)
		case "view-todo":
			todo := findByID(todos, id)
			component := components.View_Todo(todo)
			return component.Render(context.Background(), c.Response().Writer)
		}
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid element")
	})
	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":3000"))
}
