package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	productID := 1
	discountRate := 30

	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/products/%d/discount?rate=%d", productID, discountRate))
	if err != nil {
		log.Fatal("unable to send the request")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("unable to read the body")
	}

	fmt.Println(string(body))
}
