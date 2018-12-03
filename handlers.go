package main

import "github.com/labstack/echo"

func assignHandlers(e *echo.Echo) {
	e.GET("/testdetails", testDetails)
	return

}
