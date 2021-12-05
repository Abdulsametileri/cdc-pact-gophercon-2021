package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pact-foundation/pact-go/dsl"
	"net/http"
	"testing"
)

func TestConsumer(t *testing.T) {
	pact := &dsl.Pact{
		Host:                     "localhost",
		Consumer:                 "ProductService",
		Provider:                 "CampaignService",
		DisableToolValidityCheck: true,
	}
	defer pact.Teardown()

	pact.
		AddInteraction().
		Given("i get new product price with specified discount rate").
		UponReceiving("A request for campaign").
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.String("/products/1/discount"),
			Query: map[string]dsl.Matcher{
				"rate": dsl.Like("30"),
			},
		}).
		WillRespondWith(dsl.Response{
			Status: http.StatusOK,
			Headers: dsl.MapMatcher{
				fiber.HeaderContentType: dsl.String(fiber.MIMEApplicationJSON),
			},
			Body: dsl.StructMatcher{
				"id":    dsl.Like(1),
				"price": dsl.Like(70),
				"name":  dsl.Like(""),
			},
		})

	err := pact.Verify(func() error {
		return makeRequest(pact.Server.Port)
	})

	if err != nil {
		t.Fatal(err)
	}

	/*
		   curl -X PUT \
	       http://localhost/pacts/provider/CampaignService/consumer/ProductService/version/1.0.0 \
		  -H "Content-Type: application/json" \
		  -d @/Users/abdulsamet.ileri/Desktop/personal/cdc-pact-gophercon-2021/product/pacts/productservice-campaignservice.json
	*/
}
