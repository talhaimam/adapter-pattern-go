package main

import (
	"fmt"
	"net/url"
)

type IntegrationAdapter struct {
	integration Integration
}

func NewIntegrationAdapter(integration Integration) *IntegrationAdapter {
	return &IntegrationAdapter{integration}
}

func (adapter *IntegrationAdapter) FetchClasses(url url.URL) {
	adapter.integration.FetchClasses(url)
}

func (adapter *IntegrationAdapter) FetchStudios(url url.URL) {
	adapter.integration.FetchStudios(url)
}

func main() {
	fmt.Println("Hello World")
}
