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

type ByProjectKeyShippingMethodsByIDRequestMethodPost struct {
	body    ShippingMethodUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShippingMethodsByIDRequestMethodPostInput
}

func (r *ByProjectKeyShippingMethodsByIDRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShippingMethodsByIDRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyShippingMethodsByIDRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyShippingMethodsByIDRequestMethodPost) Expand(v []string) *ByProjectKeyShippingMethodsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsByIDRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsByIDRequestMethodPost) WithQueryParams(input ByProjectKeyShippingMethodsByIDRequestMethodPostInput) *ByProjectKeyShippingMethodsByIDRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShippingMethodsByIDRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyShippingMethodsByIDRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyShippingMethodsByIDRequestMethodPost) Execute(ctx context.Context) (result *ShippingMethod, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
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
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 409, 400, 401, 403, 500, 502, 503:
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