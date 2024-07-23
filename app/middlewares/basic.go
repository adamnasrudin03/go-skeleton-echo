package middlewares

import (
	"crypto/subtle"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// BasicAuth function basic auth
func BasicAuth(basicUsername, basicPassword string) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte(basicUsername)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(basicPassword)) == 1 {
			return true, nil
		}
		return false, nil
	})
}
