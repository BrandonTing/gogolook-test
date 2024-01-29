package server

import (
	"net/http"

	"gogolook-test/cmd/web"
	"gogolook-test/internal/schema"
	"gogolook-test/internal/tasks"

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

	taskGroup := e.Group("/tasks")
	taskGroup.GET("", tasks.GetTasksHandler)
	taskGroup.POST("", tasks.SetTasksHandler)
	taskGroup.PUT("/:id", tasks.UpdateTasksHandler)
	taskGroup.DELETE("/:id", tasks.RemoveTasksHandler)

	webGroup := e.Group("/web")
	webGroup.GET("", echo.WrapHandler(http.HandlerFunc(web.TasksHomeHandler)))
	webGroup.POST("/task/new", echo.WrapHandler(http.HandlerFunc(web.NewTaskHandler)))
	webGroup.PUT("/task/status/update", echo.WrapHandler(http.HandlerFunc(web.UpdateTaskStatusHandler)))
	webGroup.DELETE("/task/delete", echo.WrapHandler(http.HandlerFunc(web.DeleteTaskHandler)))

	e.GET("/", s.HelloWorldHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
