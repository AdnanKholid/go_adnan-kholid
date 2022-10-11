package drivers

import (
	userDB "echo-api"
	userDomain "echo-api"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlRepository(conn)
}
