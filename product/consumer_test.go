//go:build consumer
// +build consumer

package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pact-foundation/pact-go/dsl"
	"net/http"
	"strconv"
	"testing"
)

func createPact() (pact *dsl.Pact, cleanUp func()) {
	pact = &dsl.Pact{
		Host:                     "localhost",
		Consumer:                 "ProductService",
		Provider:                 "CampaignService",
		DisableToolValidityCheck: true,
		PactFileWriteMode:        "merge",
		LogDir:                   "./pacts/logs",
	}

	cleanUp = func() { pact.Teardown() }

	return pact, cleanUp
}

func Test_IGetNewProductPriceWithSpecifiedDiscountRate(t *testing.T) {
	pact, cleanUp := createPact()
	defer cleanUp()

	const productIDWithDiscount = 1
	const discountRate = 30

	pact.
		AddInteraction().
		Given("i get new product price with specified discount rate").
		UponReceiving("A request for campaign").
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.String(fmt.Sprintf("/products/%d/discount", productIDWithDiscount)),
			Query: map[string]dsl.Matcher{
				"rate": dsl.Like(strconv.Itoa(discountRate)),
			},
		}).
		WillRespondWith(dsl.Response{
			Status: http.StatusOK,
			Headers: dsl.MapMatcher{
				fiber.HeaderContentType: dsl.String(fiber.MIMEApplicationJSON),
			},
			Body: dsl.StructMatcher{
				"id":    dsl.Integer(),
				"price": dsl.Decimal(),
				"name":  dsl.Like(""),
			},
		})
	err := pact.Verify(func() error {
		return makeRequest(pact.Server.Port, productIDWithDiscount, discountRate)
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_IGetCampaignNotFoundErrorWhenTheProductHasNoDiscount(t *testing.T) {
	pact, cleanUp := createPact()
	defer cleanUp()

	const productIDWithoutDiscount = 2
	const discountRate = 30

	pact.
		AddInteraction().
		Given("i get campaign not found error when the product has no discount").
		UponReceiving("A request for campaign without discounted product").
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.String(fmt.Sprintf("/products/%d/discount", productIDWithoutDiscount)),
			Query: map[string]dsl.Matcher{
				"rate": dsl.Like(strconv.Itoa(discountRate)),
			},
		}).
		WillRespondWith(dsl.Response{
			Status: http.StatusNotAcceptable,
			Headers: dsl.MapMatcher{
				fiber.HeaderContentType: dsl.String(fiber.MIMEApplicationJSON),
			},
			Body: dsl.StructMatcher{
				"message": dsl.Like("No campaign found for this product"),
			},
		})
	err := pact.Verify(func() error {
		return makeRequest(pact.Server.Port, productIDWithoutDiscount, discountRate)
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_IGetProductDoesNotExist(t *testing.T) {
	pact, cleanUp := createPact()
	defer cleanUp()

	const notExistProductID = 3
	const discountRate = 30

	pact.
		AddInteraction().
		Given("i get the product does not exist").
		UponReceiving("A request for campaign with nonexist product").
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.String(fmt.Sprintf("/products/%d/discount", notExistProductID)),
			Query: map[string]dsl.Matcher{
				"rate": dsl.Like(strconv.Itoa(discountRate)),
			},
		}).
		WillRespondWith(dsl.Response{
			Headers: map[string]dsl.Matcher{
				fiber.HeaderContentType: dsl.String(fiber.MIMETextPlainCharsetUTF8),
			},
			Status: http.StatusNotFound,
		})
	err := pact.Verify(func() error {
		return makeRequest(pact.Server.Port, notExistProductID, discountRate)
	})

	if err != nil {
		t.Fatal(err)
	}
}
