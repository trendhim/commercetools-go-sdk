package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost struct {
	body    CartUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPostInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPostInput) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Updates a cart in the store specified by {storeKey}.
*	If the cart exists in the commercetools project but does not have the store field,
*	or the store field references a different store, this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost) Execute(ctx context.Context) (result *Cart, err error) {
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
