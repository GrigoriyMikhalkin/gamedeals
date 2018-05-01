package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"gamedeals/database"
	"gamedeals/games"
)

func main() {
	e := echo.New()

	database.InitDB()
	games.InitDB()
	defer database.DB.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "assets")

	games.CreateGameEndpoints(e)
	games.CreateGameEditionsEndpoints(e)
	games.CreateStoresEndpoints(e)
	games.CreateStoreGamePagesEndpoints(e)

	e.GET("/search", games.GetSearchResult)

	err := database.DB.DB().Ping()
	fmt.Println(err)

	e.Start(":8000")
}
