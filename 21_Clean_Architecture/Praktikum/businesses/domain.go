package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Email     string
	Password  string
}

type Usecase interface {
	GetAll() []Domain
	Create(userDomain *Domain) Domain
}

type Repository interface {
	GetAll() []Domain
	Create(userDomain *Domain) Domain
}
