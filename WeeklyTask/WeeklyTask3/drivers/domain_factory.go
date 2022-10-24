package drivers

import (
	userDomain "weakly-task-3/businesses/users"
	userDB "weakly-task-3/drivers/mysql/users"

	blogsDomain "weakly-task-3/businesses/blogs"
	blogsDB "weakly-task-3/drivers/mysql/blogs"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewBlogsRepository(DB *gorm.DB) blogsDomain.Repository {
	return blogsDB.NewMysqlRepositroy(DB)
}
