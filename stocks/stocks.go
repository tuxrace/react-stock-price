package stocks

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tuxrace/react-stock-price/database"
)

// Stocks Struct
type Stocks struct {
	gorm.Model
	SerialNumber string `json:"serialNumber"`
	Brand        string `json:"brand"`
	ModelName    string `json:"model"`
	Status       string `json:"status"`
	DateBought   string `json:"dateBought"`
}

// GetStocks function
func GetStocks(c *fiber.Ctx) {
	db := database.DB
	var stocks []Stocks
	db.Find(&stocks)
	c.JSON(stocks)
}
