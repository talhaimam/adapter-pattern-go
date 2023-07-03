package integration1

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

func (studio *Integration1) FetchClasses(url url.URL) {
	fmt.Println("Integration1 FetchClasses")
}

func (studio *Integration1) FetchStudios(url url.URL) {
	fmt.Println("Integration1 FetchStudios")
}
