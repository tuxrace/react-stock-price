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
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Trend  string `json:"trend"`
}

// GetStocks function
func GetStocks(c *fiber.Ctx) {
	db := database.DB
	var stocks []Stocks
	db.Find(&stocks)
	c.JSON(stocks)
}
