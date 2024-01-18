package server

import (
	"net/http"

	"gogolook-test/cmd/web"
	"gogolook-test/internal/schema"
	"gogolook-test/internal/storage"
	"gogolook-test/internal/tasks"

	"github.com/a-h/templ"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Validator = &schema.Validator{Validator: validator.New()}
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/js/*", echo.WrapHandler(fileServer))

	e.GET("/web", echo.WrapHandler(templ.Handler(web.HelloForm())))
	e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	storage.SetupStore()
	e.GET("/tasks", tasks.GetTasksHandler)
	e.POST("/tasks", tasks.SetTasksHandler)
	e.PUT("/tasks/:id", tasks.UpdateTasksHandler)
	e.DELETE("/tasks/:id", tasks.RemoveTasksHandler)

	e.GET("/", s.HelloWorldHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
