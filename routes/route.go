package routes

import (
	"Echo/controllers"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type H map[string]interface{}

var secret_key string

func InitRoute(e *echo.Echo, GORM *gorm.DB, secret string) {
	web := controllers.Web{}
	secret_key = secret
	admin := controllers.Admin{}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Static("/static", "static")
	e.GET("/", web.Index())
	e.GET("/admin/login", web.Login())
	e.POST("/admin/auth", web.Auth(GORM, secret))
	e.POST("/user/sign_up", controllers.SignUp(GORM))

	// Restricted group
	r := e.Group("/admin/web")
	r.Use(middleware.JWT([]byte(secret)))
	r.GET("", admin.Index())
	r.PUT("/user/edit", controllers.EditUser(GORM))
	r.GET("/user/:id", controllers.GetUser(GORM))
	r.DELETE("/user/delete/:id", controllers.DeleteUser(GORM))
	r.GET("/user/all", controllers.GetUsers(GORM))
}
