//go:build provider
// +build provider

package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
)

type Settings struct {
	Host            string
	ProviderName    string
	BrokerBaseURL   string
	BrokerUsername  string
	BrokerPassword  string
	ConsumerName    string
	ConsumerVersion string
	ConsumerTag     string
	ProviderVersion string
}

func (s *Settings) getPactURL(useLocal bool) string {
	var pactURL string

	if useLocal {
		pactURL = "../product/pacts/productservice-campaignservice.json"
		return pactURL
	}

	if s.ConsumerVersion == "" {
		pactURL = fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/latest/master.json", s.BrokerBaseURL, s.ProviderName, s.ConsumerName)
	} else {
		pactURL = fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/version/%s.json", s.BrokerBaseURL, s.ProviderName, s.ConsumerName, s.ConsumerVersion)
	}

	return pactURL
}

func (s *Settings) create() {
	s.Host = "127.0.0.1"
	s.ProviderName = "CampaignService"
	s.ConsumerName = "ProductService"
	s.BrokerBaseURL = "http://localhost"
	s.ConsumerTag = "master"
	s.ProviderVersion = "1.0.3"
	s.ConsumerVersion = "1.0.7"
}

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	go startServer(port)

	settings := Settings{}
	settings.create()

	pact := dsl.Pact{
		Host:                     settings.Host,
		Provider:                 settings.ProviderName,
		Consumer:                 settings.ConsumerName,
		DisableToolValidityCheck: true,
	}

	verifyRequest := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://%s:%d", settings.Host, port),
		ProviderVersion: settings.ProviderVersion,
		BrokerUsername:  settings.BrokerUsername,
		BrokerPassword:  settings.BrokerPassword,
		Tags:            []string{settings.ConsumerTag},
		PactURLs:        []string{settings.getPactURL(false)},
		StateHandlers: map[string]types.StateHandler{
			"i get new product price with specified discount rate": func() error {
				return nil
			},
			"i get campaign not found error when the product has no discount": func() error {
				return nil
			},
			"i get the product does not exist": func() error {
				delete(products, 3)
				return nil
			},
		},
		PublishVerificationResults: true,
		FailIfNoPactsFound:         true,
	}

	verifyResponses, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatal(err)
	}

	pp.Println(len(verifyResponses), "pact tests run")
}
