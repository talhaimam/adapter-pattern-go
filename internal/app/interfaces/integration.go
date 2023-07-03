package interfaces

import "net/url"

type Integration interface {
	FetchClasses(url url.URL)
	FetchStudios(url url.URL)
}
