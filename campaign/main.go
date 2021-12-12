package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

var products = Products{}

func main() {
	err := startServer(3000)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Provider Service Listening :3000")
}

func startServer(port int) error {
	products.InitProducts()

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Get("/products/:id/discount", func(c *fiber.Ctx) error {
		productID, _ := c.ParamsInt("id", 0)
		discountRate, _ := strconv.ParseFloat(c.Query("rate"), 64)

		product, ok := products[productID]
		if !ok {
			return c.SendStatus(http.StatusNotFound)
		}

		if !product.HasCampaign {
			return c.
				Status(http.StatusNotAcceptable).
				JSON(map[string]interface{}{
					"message": "No campaign found for this product",
				})
		}

		discountedPrice := products[productID].Price - (products[productID].Price*discountRate)/100
		product.Price = discountedPrice

		return c.JSON(product)
	})
	err := app.Listen(fmt.Sprintf(":%d", port))
	return err
}
