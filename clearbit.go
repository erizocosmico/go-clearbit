package clearbit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// API is a service to retrieve data from the Clearbit enrichment API.
type API interface {
	// Person returns information about a person given their email. If the second
	// parameter is true, it will use the streaming API, waiting until the request
	// completes. If the streaming is not activated, the resultant person may be nil.
	// In this case, you will have to either try again in a bit or setup a webhook flow.
	Person(email string, streaming bool) (*Person, error)

	// Company returns information about a company given their domain. If the second
	// parameter is true, it will use the streaming API, waiting until the request
	// completes. If the streaming is not activated, the resultant person may be nil.
	// In this case, you will have to either try again in a bit or setup a webhook flow.
	Company(domain string, streaming bool) (*Company, error)
}

type api struct {
	key string
}

// New creates a new instance of the API given an API Key.
func New(key string) API {
	return &api{
		key: key,
	}
}

func (a *api) Person(email string, streaming bool) (*Person, error) {
	var p Person
	ok, err := a.request(apiURL(findPerson, email, streaming), &p)
	if err != nil || !ok {
		return nil, err
	}

	return &p, nil
}

func (a *api) Company(domain string, streaming bool) (*Company, error) {
	var c Company
	ok, err := a.request(apiURL(findCompany, domain, streaming), &c)
	if err != nil || !ok {
		return nil, err
	}

	return &c, nil
}

// ErrNotFound means the entity was not found on clearbit.
var ErrNotFound = errors.New("entity not found on clearbit")

func (a *api) request(url string, into interface{}) (bool, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.SetBasicAuth(a.key, "")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return false, err
	}

	switch resp.StatusCode {
	case http.StatusNotFound:
		return false, ErrNotFound
	case http.StatusAccepted:
		return false, nil
	}

	if err := json.NewDecoder(resp.Body).Decode(into); err != nil {
		return false, err
	}

	return true, nil
}

type apiMethod struct {
	subdomain string
	path      string
	arg       string
}

type apiMethodKind int

const (
	findPerson apiMethodKind = 1 << iota
	findCompany
)

var methods = map[apiMethodKind]*apiMethod{
	findPerson: &apiMethod{
		subdomain: "person",
		path:      "people",
		arg:       "email",
	},
	findCompany: &apiMethod{
		subdomain: "company",
		path:      "companies",
		arg:       "domain",
	},
}

const (
	streamingURL = "https://%s-stream.clearbit.com/v1/%s/%s/%s"
	baseURL      = "https://%s.clearbit.com/v2/%s/find?%s=%s"
)

func apiURL(kind apiMethodKind, arg string, streaming bool) string {
	m := methods[kind]
	if streaming {
		return fmt.Sprintf(streamingURL, m.subdomain, m.path, m.arg, arg)
	}

	return fmt.Sprintf(baseURL, m.subdomain, m.path, m.arg, arg)
}
