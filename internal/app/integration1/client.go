package integration1

import (
	"fmt"
	"net/url"
)

func (studio *Integration1) FetchClasses(url url.URL) {
	fmt.Println("Integration1 FetchClasses")
}

func (studio *Integration1) FetchStudios(url url.URL) {
	fmt.Println("Integration1 FetchStudios")
}
