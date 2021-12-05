package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

var products = Products{}

func main() {
	products.InitProducts()

	err := startServer(3000)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Provider Service Listening :3000")
}

func startServer(port int) error {
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
	err := app.Listen(fmt.Sprintf(":%d", port))
	return err
}