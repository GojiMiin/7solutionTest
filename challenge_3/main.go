package main

import (
	"challenge_3/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	meatHandler := handler.MeatHandler{}
	e.GET("/beef/summary", meatHandler.GetSummary)
	e.Logger.Fatal(e.Start(":8080"))
}
