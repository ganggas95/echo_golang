package controllers

import (
	"Echo/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type H map[string]interface{}
type UserJson struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
type AuthJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ConfigClaims struct {
	User  string `json:"user"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func SignUp(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		c.Bind(&user)
		if ok := models.CheckEmailAndUsername(db, user.Username, user.Email); ok {
			return c.JSON(http.StatusConflict, H{
				"user_conflict": 409,
			})
		}
		err := models.SaveUser(db, &user)
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": 201,
			})
		} else {
			return err
		}
	}
}

func GetUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		temp, _ := strconv.Atoi(c.Param("id"))
		id := int64(temp)
		user := models.GetUser(db, id)
		if user != nil {
			var userjson UserJson
			userjson.Id = user.IdUser
			userjson.Email = user.Email
			userjson.Username = user.Username
			return c.JSON(http.StatusOK, userjson)
		}
		return c.JSON(http.StatusNotFound, H{
			"error_code": "not found",
		})
	}
}

func GetUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := models.GetUsers(db)
		users := make([]UserJson, len(u))
		for k, _ := range u {
			users[k].Id = u[k].IdUser
			users[k].Email = u[k].Email
			users[k].Username = u[k].Username
		}
		return c.JSON(http.StatusOK, users)
	}
}

func DeleteUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		temp, _ := strconv.Atoi(c.Param("id"))
		id := int64(temp)
		err := models.DeleteUser(db, id)
		if err != nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}

func EditUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		if err := c.Bind(&user); err != nil {
			return err
		}
		err := models.PutUser(db, user)
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"updated": 201,
			})
		} else {
			return err
		}
	}
}

func Login(db *gorm.DB, secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		var jsonAuth AuthJson
		c.Bind(&jsonAuth)
		user, _ := models.AuthUser(db, jsonAuth.Username, jsonAuth.Password)
		if user != nil {
			configClaims := &ConfigClaims{
				user.Username,
				true,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 20).Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodRS512, configClaims)
			t, err := token.SignedString(secret)
			log.Println(t)
			log.Println(err)
			if err != nil {
				return echo.ErrUnauthorized
			}
			return c.JSON(http.StatusOK, H{
				"token": t,
			})
		}
		fmt.Println(user)
		return c.JSON(http.StatusUnauthorized, H{
			"error_msg": "Unauthorized",
		})
	}
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*ConfigClaims)
	name := claims.User
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

/*
func GetTokenSource(src string, publicKey *rsa.PublicKey) *jwt.Token {
	authToken, err := jwt.Parse(src, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil
	}
	if ok := authToken.Valid; ok {
		return authToken
	}
	return nil


}

*/
