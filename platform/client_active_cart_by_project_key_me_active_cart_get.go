// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMeActiveCartRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyMeActiveCartRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyMeActiveCartRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMeActiveCartRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMeActiveCartRequestMethodGet) Execute(ctx context.Context) (result *Cart, err error) {
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
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}