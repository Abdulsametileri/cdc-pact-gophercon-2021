//go:build provider
// +build provider

package main

import (
	"fmt"
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

// Local pact file or remote based urls (Pact Broker)
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
	s.ProviderVersion = "1.0.0"
	s.ConsumerVersion = "1.0.0"
}

func TestProvider(t *testing.T) {
	// TODO: Implement me
}
