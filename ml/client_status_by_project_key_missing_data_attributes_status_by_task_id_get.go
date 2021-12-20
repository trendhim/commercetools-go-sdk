// Generated file, please do not change!!!
package ml

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMissingDataAttributesStatusByTaskIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyMissingDataAttributesStatusByTaskIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyMissingDataAttributesStatusByTaskIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMissingDataAttributesStatusByTaskIdRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMissingDataAttributesStatusByTaskIdRequestMethodGet) Execute(ctx context.Context) (result *MissingDataTaskStatus, err error) {
	queryParams := url.Values{}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}