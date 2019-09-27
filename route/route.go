package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/letanthang/my_framework/config"
	"github.com/letanthang/my_framework/db/types"
	"github.com/letanthang/my_framework/handlers"
)

func All(e *echo.Echo) {
	Public(e)
	Staff(e)
}

func Public(e *echo.Echo) {
	publicRoute := e.Group("/v1/public")
	publicRoute.GET("/student", handlers.GetAllStudent)
	publicRoute.PATCH("/student/simple", handlers.GetStudent)
	publicRoute.PATCH("/student", handlers.SearchStudent)
	publicRoute.GET("/health", handlers.CheckHealth)
	publicRoute.GET("/student/group/last_name", handlers.GroupStudent)
}

func Staff(e *echo.Echo) {
	staffRoute := e.Group("/v1/staff")
	config := middleware.JWTConfig{
		Claims:     &types.MyClaims{},
		SigningKey: []byte(config.Config.Encryption.JWTSecret),
	}
	staffRoute.Use(middleware.JWTWithConfig(config))
	staffRoute.POST("/student", handlers.AddStudent)
	staffRoute.DELETE("/student", handlers.DeleteStudent)
}
