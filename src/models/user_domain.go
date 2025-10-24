package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserDomain struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewUserDomain(
	Email string,
	Name string,
	Password string,
) *UserDomain {
	return &UserDomain{
		ID:       uuid.New().String(),
		Email:    Email,
		Name:     Name,
		Password: Password,
	}
}

func (ud *UserDomain) EncryptPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ud.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ud.Password = string(hashedPassword)
	return nil
}

func (ud *UserDomain) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ud.Password), []byte(password))
	return err == nil
}
