package models

import (
	"golang.org/x/crypto/bcrypt"
)

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Profile   Profile
	Comment   []Comment
	Bookmark  []Bookmark
}

type Profile struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `json:"userId"`
	Avatar  string `json:"avatar"`
	Bio     string `json:"bio"`
	Gender  string `json:"gender"`
	Country string `json:"country"`
}

type Bookmark struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint `json:"userId"`
	MovieID     uint `json:"movieId"`
	Is_favorite bool `json:"is_favorite"`
	Is_watched  bool `json:"is_watched"`
}
