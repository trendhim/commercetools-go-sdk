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

type ByProjectKeyPaymentsByIDRequestMethodPost struct {
	body    PaymentUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyPaymentsByIDRequestMethodPostInput
}

func (r *ByProjectKeyPaymentsByIDRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyPaymentsByIDRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyPaymentsByIDRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyPaymentsByIDRequestMethodPost) Expand(v []string) *ByProjectKeyPaymentsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentsByIDRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyPaymentsByIDRequestMethodPost) WithQueryParams(input ByProjectKeyPaymentsByIDRequestMethodPostInput) *ByProjectKeyPaymentsByIDRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyPaymentsByIDRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyPaymentsByIDRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyPaymentsByIDRequestMethodPost) Execute(ctx context.Context) (result *Payment, err error) {
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