package appMiddleware

import (
	"echo-api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func DummyBasicAuth(username, password string, c echo.Context) (bool, error) {
	if username == "root" && password == "root" {
		return true, nil
	}
	return false, nil
}

func MakePersonBasicAuth(m appModel.PersonModel) middleware.BasicAuthValidator {
	return func(username, password string, c echo.Context) (bool, error) {
		_, err := m.GetByEmailAndPassword(username, password)
		if err == nil {
			return true, nil
		}
		return false, nil
	}
}
