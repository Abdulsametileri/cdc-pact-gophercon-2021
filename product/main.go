package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	if err := makeRequest(3000, 1, 30); err != nil {
		log.Fatal(err)
	}
}

func makeRequest(port, productID, discountRate int) error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/products/%d/discount?rate=%d", port, productID, discountRate))
	if err != nil {
		return errors.New("unable to send the request")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("unable to read the body")
	}

	fmt.Println(string(body))

	return nil
}
