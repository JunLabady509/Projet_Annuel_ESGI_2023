package handlers

import "github.com/labstack/echo"

func auth(username, password string, ctx echo.Context) (bool, error) {
	if username == "joe" && password == "secret" {
		return true, nil
	}
	return false, nil
}
