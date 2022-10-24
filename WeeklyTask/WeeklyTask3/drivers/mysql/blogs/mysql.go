package blogs

import (
	blogs "weakly-task-3/businesses/blogs"

	"gorm.io/gorm"
)

type blogsRepository struct {
	DB *gorm.DB
}

func NewMysqlRepositroy(DB *gorm.DB) blogs.Repository {
	return &blogsRepository{
		DB: DB,
	}
}

func (br *blogsRepository) GetAll() []blogs.Domain {
	var rec []Blogs

	br.DB.Find(&rec)

	blogsDomain := []blogs.Domain{}

	for _, blogs := range rec {
		blogsDomain = append(blogsDomain, blogs.ToDomain())
	}

	return blogsDomain
}

func (br *blogsRepository) GetByID(id string) blogs.Domain {
	var rec Blogs

	br.DB.First(&rec, "id =?", id)

	return rec.ToDomain()
}

func (br *blogsRepository) Create(blogsDomain *blogs.Domain) blogs.Domain {
	rec := FromDomain(blogsDomain)

	result := br.DB.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (br *blogsRepository) Update(id string, blogsDomain *blogs.Domain) blogs.Domain {
	rec := br.GetByID(id)

	updateBlogs := FromDomain(&rec)

	updateBlogs.Title = blogsDomain.Title
	updateBlogs.Description = blogsDomain.Title
	updateBlogs.CategoryID = blogsDomain.CategoryID

	br.DB.Save(&updateBlogs)

	return updateBlogs.ToDomain()
}

func (br *blogsRepository) Delete(id string) bool {
	rec := br.GetByID(id)

	deleteBlogs := FromDomain(&rec)

	result := br.DB.Delete(&deleteBlogs)

	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func (br *blogsRepository) GetByTitle(title string) blogs.Domain {
	var rec Blogs

	br.DB.Where("title=?", title).Find(&rec)

	return rec.ToDomain()
}

func (br *blogsRepository) GetByCategory(idCategory string) []blogs.Domain {
	var rec []Blogs
	blogsDomain := []blogs.Domain{}
	br.DB.Where("category_id=?", idCategory).Find(&rec)
	for _, blogs := range rec {
		blogsDomain = append(blogsDomain, blogs.ToDomain())
	}
	return blogsDomain
}
