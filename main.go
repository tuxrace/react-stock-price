package main

import (
	"fmt"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/tuxrace/react-stock-price/database"
	"github.com/tuxrace/react-stock-price/stocks"
)

func index(c *fiber.Ctx) {
	c.Send("Index page")
}

func startDB() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "stocks.db")

	if err != nil {
		panic("Error connect")
	}

	fmt.Println("Connected")
	database.DB.AutoMigrate(&stocks.Stocks{})
}

func boostrapDB() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "stocks.db")

	if err != nil {
		panic("Error bootstrap")
	}

	db := database.DB
	var item stocks.Stocks
	item.Symbol = "D05"
	item.Price = "20.1"
	item.Trend = "Up"
	db.Create(&item)
	fmt.Println("Bootstrap done")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", index)

	app.Get("/api/stocks", stocks.GetStocks)
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	startDB()
	setupRoutes(app)
	boostrapDB()
	app.Listen(3001)
	defer database.DB.Close()
}
