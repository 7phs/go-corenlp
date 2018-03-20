package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type coreNlpClient struct {
	ServerParameter

	client *http.Client
}

func CoreNlpClient(parameter ServerParameter) *coreNlpClient {
	return &coreNlpClient{
		ServerParameter: parameter,
		client:          &http.Client{},
	}
}

func (o *coreNlpClient) makeUrl(properties interface{}) (string, error) {
	buf, err := json.Marshal(properties)
	if err != nil {
		return "", errors.Wrap(err, "failed to encode properties to json")
	}

	parameters := url.Values{}
	parameters.Add("properties", string(buf))

	u := &url.URL{
		Scheme:   o.Scheme(),
		Host:     fmt.Sprintf("%s:%d", o.Host(), o.Port()),
		Path:     "/",
		RawQuery: parameters.Encode(),
	}

	return u.String(), nil
}

func (o *coreNlpClient) request(properties interface{}, text io.Reader) (*http.Request, error) {
	requestUrl, err := o.makeUrl(properties)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request to CoreNLP server")
	}

	return http.NewRequest("POST", requestUrl, text)
}

func (o *coreNlpClient) DoString(properties interface{}, text string) (*http.Response, error) {
	return o.Do(properties, bytes.NewReader([]byte(text)))
}

func (o *coreNlpClient) Do(properties interface{}, text io.Reader) (*http.Response, error) {
	req, err := o.request(properties, text)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request to CoreNLP server")
	}

	resp, err := o.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute request to CoreNLP server")
	}

	return resp, nil
}
