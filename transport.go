package circleci

import (
	"errors"
	"net/http"
)

type Transport struct {
	Config *Config
	Base   *http.Transport
}

var _ http.RoundTripper = (*Transport)(nil)

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Config == nil {
		return nil, errors.New("circleci: Transport's Config is nil.")
	}

	token, err := t.Config.Token()
	if err != nil {
		return nil, err
	}

	req2 := cloneRequest(req)
	req2.Header["Accept"] = []string{"application/json"}
	q := req2.URL.Query()
	q.Set("circle-token", token)
	req2.URL.RawQuery = q.Encode()

	res, err := t.base().RoundTrip(req2)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Transport) base() *http.Transport {
	if t.Base != nil {
		return t.Base
	}

	return http.DefaultTransport
}

func cloneRequest(r *http.Request) *http.Request {
	r2 := new(http.Request)
	*r2 = *r

	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}

	return r2
}
