package urier

import (
	"net/url"
)

type uri struct {
	url *url.URL
}

func Build() *uri {
	return &uri{
		url: &url.URL{},
	}
}

func (t *uri) Scheme(v UriScheme) *uri {
	t.url.Scheme = v.String()
	return t
}

func (t *uri) Host(v string) *uri {
	t.url.Host = v
	return t
}
