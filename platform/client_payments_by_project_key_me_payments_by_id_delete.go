// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyMePaymentsByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMePaymentsByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyMePaymentsByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMePaymentsByIDRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyMePaymentsByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyMePaymentsByIDRequestMethodDelete) Version(v int) *ByProjectKeyMePaymentsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyMePaymentsByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyMePaymentsByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyMePaymentsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyMePaymentsByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyMePaymentsByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyMePaymentsByIDRequestMethodDeleteInput) *ByProjectKeyMePaymentsByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMePaymentsByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyMePaymentsByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMePaymentsByIDRequestMethodDelete) Execute(ctx context.Context) (result *MyPayment, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
