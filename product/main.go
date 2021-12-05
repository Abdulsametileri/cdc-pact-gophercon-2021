package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	if err := makeRequest(3000); err != nil {
		log.Fatal(err)
	}
}

func makeRequest(port int) error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/products/1/discount?rate=30", port))
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
