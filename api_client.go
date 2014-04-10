package paymill

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

import _ "crypto/sha512"

const APIScheme string = "https"
const APIBase string = "api.paymill.com"
const APIVersion string = "v2"

type ApiClient struct {
	Token string
}

func NewApiClient(token string) (c *ApiClient) {
	if strings.Trim(token, " ") == "" {
		return nil
	}

	c = &ApiClient{
		Token: token,
	}
	return
}

func UrlFor(entity string, data url.Values) url.URL {
	url := url.URL{Scheme: APIScheme, Host: APIBase}
	path := fmt.Sprintf("/%s/%s", APIVersion, entity)

	url.Path = path
	url.RawQuery = data.Encode()

	return url
}

func (c *ApiClient) doRequest(resource string, method string, urlData url.Values, formData url.Values) (resp *http.Response, body []byte) {
	http_client := &http.Client{}

	// This can be wrapped in a method
	var req *http.Request
	url := UrlFor(resource, urlData)
	req, err := http.NewRequest(method, url.String(), strings.NewReader(formData.Encode()))
	if err != nil {
		panic(err)
	}

	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	req.SetBasicAuth(c.Token, "")

	resp, err = http_client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return
}
