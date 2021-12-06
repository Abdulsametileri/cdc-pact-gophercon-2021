package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	if err := makeRequest(3000, 1); err != nil {
		log.Fatal(err)
	}
}

func makeRequest(port, productID int) error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/products/%d/discount?rate=30", port, productID))
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
