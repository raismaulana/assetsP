package entity

import (
	"time"

	"github.com/raismaulana/assetsP/application/apperror"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID          int64     `gorm:"primary_key:auto_increment;column:id_user"`
	Name        string    `gorm:"type:varchar(100) not null"`
	Email       string    `gorm:"type:varchar(100) not null"`
	Password    string    `gorm:"type:varchar(100) not null"`
	Role        string    `gorm:"type:varchar(20) not null;default:peasant"`
	ActivatedAt null.Time `gorm:"default:null"`
	CreatedAt   time.Time `` //
	UpdatedAt   time.Time `` //
}

type UserRequest struct {
	Name     string ``
	Email    string ``
	Password string ``
	Role     string ``
}

func NewUser(req UserRequest) (*User, error) {
	if req.Name == "" {
		return nil, apperror.NameMustNotEmpty
	}
	if req.Email == "" {
		return nil, apperror.EmailMustNotEmpty
	}
	if req.Password == "" {
		return nil, apperror.EmailMustNotEmpty
	}
	if req.Role == "" {
		req.Role = "peasant"
	}
	obj := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
	return &obj, nil
}
