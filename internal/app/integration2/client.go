package integration2

import (
	"fmt"
	"net/url"
)

func (studio *Integration2) FetchClasses(url url.URL) {
	fmt.Println("Integration2 FetchClasses")
}

func (studio *Integration2) FetchStudios(url url.URL) {
	fmt.Println("Integration2 FetchStudios")
}
