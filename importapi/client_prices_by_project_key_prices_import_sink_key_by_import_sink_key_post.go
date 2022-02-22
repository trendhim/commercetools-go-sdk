// Generated file, please do not change!!!
package importapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestMethodPost struct {
	body    PriceImportRequest
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates import request for creating new prices or updating existing ones.
 */
func (rb *ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestMethodPost) Execute(ctx context.Context) (result *ImportResponse, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 201:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
