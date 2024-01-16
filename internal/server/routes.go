package server

import (
	"net/http"

	"gogolook-test/cmd/web"
	"gogolook-test/internal/tasks"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/js/*", echo.WrapHandler(fileServer))

	e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	taskStore := tasks.CreateTaskStore()
	e.GET("/tasks", taskStore.GetTasks)
	e.POST("/tasks", taskStore.SetTasks)
	e.PUT("/tasks/:id", taskStore.UpdateTasks)
	e.DELETE("/tasks/:id", taskStore.RemoveTasks)

	e.GET("/", s.HelloWorldHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
