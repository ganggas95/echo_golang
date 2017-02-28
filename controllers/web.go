package controllers

import (
	"Echo/models"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Web struct {
}

func (w *Web) Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "index.html", "helloWorld")
	}
}

func (w *Web) Login() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "admin_login.html", "helloWorld")
	}
}

func (w *Web) Auth(db *gorm.DB, secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var jsonAuth AuthJson
		jsonAuth.Username = c.FormValue("username")
		jsonAuth.Password = c.FormValue("password")
		user, _ := models.AuthUser(db, jsonAuth.Username, jsonAuth.Password)
		//log.Println(user)
		if user != nil {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["name"] = user.Username
			claims["exp"] = time.Now().Add(time.Hour * 20).Unix()
			t, err := token.SignedString([]byte(secret))
			if err != nil {
				log.Println(err)
				return echo.ErrUnauthorized
			}
			cookies := &http.Cookie{
				Expires: time.Now().Add(time.Hour * 20),
				Name:    "app_token",
				Value:   t,
			}
			c.SetCookie(cookies)
			return c.JSON(http.StatusOK, H{
				"token": t,
			})
		}
		return c.JSON(http.StatusUnauthorized, H{
			"error_msg": "Unauthorized",
		})
	}
}
