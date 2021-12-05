package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:3000/products/1/discount?rate=30")
	if err != nil {
		log.Fatal("unable to send the request")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("unable to read the body")
	}

	fmt.Println(string(body))
}
