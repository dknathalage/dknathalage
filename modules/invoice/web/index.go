package web

import "github.com/labstack/echo/v4"

func Index(c echo.Context) error {
	return c.HTML(200, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Invoice</title>
	</head>
	<body>
		<h1>Login</h1>
	</body>
	</html>
	`)
}
