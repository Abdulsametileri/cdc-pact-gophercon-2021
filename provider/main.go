package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func main() {
	products := Products{}
	products.InitProducts()

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Get("/products/:id/discount", func(c *fiber.Ctx) error {
		fmt.Println("İstek geldi")

		discountRate, _ := strconv.ParseFloat(c.Query("rate"), 64)
		productID, _ := c.ParamsInt("id", 0)

		product := products[productID]

		discountedPrice := products[productID].Price - (products[productID].Price*discountRate)/100
		product.Price = discountedPrice

		return c.JSON(product)
	})

	log.Println("Provider Service Listening :3000")
	log.Fatal(app.Listen(":3000"))
}
