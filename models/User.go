package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	IdUser         int64  `json:"id" gorm:"column:id_user;primary_key;AUTO_INCREMENT"`
	Email          string `json:"email" gorm:"column:email;unique"`
	Username       string `json:"username" gorm:"column:username"`
	Password       string `json:"password" gorm:"-"`
	HashedPassword []byte `json:"-" gorm:"column:password"`
}

func GetUsers(db *gorm.DB) []*User {
	var user []*User
	db.Find(&user)
	if db.Error != nil {
		panic(db.Error)
	}
	return user
}

func GetUser(db *gorm.DB, id int64) *User {
	var user User
	db.First(&user, id)
	if db.Error != nil {
		return nil
	}
	return &user
}

func PutUser(db *gorm.DB, user User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.HashedPassword = hashedPassword
	db.Save(&user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func DeleteUser(db *gorm.DB, id int64) error {
	user := User{IdUser: id}
	db.Delete(&user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func SaveUser(db *gorm.DB, user *User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.HashedPassword = hashedPassword
	db.Create(user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
func CheckEmailAndUsername(db *gorm.DB, username, email string) bool {
	var user []User
	db.Where("username = ? OR email = ?", username, email).Find(&user)
	fmt.Println(len(user))
	if len(user) > 0 {
		return true
	}
	return false
}
func AuthUser(db *gorm.DB, username, password string) (*User, error) {
	var user User
	db.Where("username = ?", username).Find(&user)

	if db.RecordNotFound() {
		return nil, db.Error
	} else {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err != nil {
			return nil, err
		} else {

			return &user, nil
		}
	}

}
