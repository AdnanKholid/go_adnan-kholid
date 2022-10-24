package blogs

import "time"

type Domain struct {
	ID          uint
	Title       string
	Description string
	CategoryID  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(blogsDomain *Domain) Domain
	Update(id string, blogDomain *Domain) Domain
	Delete(id string) bool
	GetByCategory(idCategory string) []Domain
	GetByTitle(title string) Domain
}

type Repository interface {
	GetAll() []Domain
	GetByID(id string) Domain
	Create(blogsDomain *Domain) Domain
	Update(id string, blogDomain *Domain) Domain
	Delete(id string) bool
	GetByTitle(title string) Domain
	GetByCategory(idCategory string) []Domain
}
