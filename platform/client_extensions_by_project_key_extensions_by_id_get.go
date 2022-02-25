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

type ByProjectKeyExtensionsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyExtensionsByIDRequestMethodGetInput
}

func (r *ByProjectKeyExtensionsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyExtensionsByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyExtensionsByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyExtensionsByIDRequestMethodGet) Expand(v []string) *ByProjectKeyExtensionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyExtensionsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyExtensionsByIDRequestMethodGet) WithQueryParams(input ByProjectKeyExtensionsByIDRequestMethodGetInput) *ByProjectKeyExtensionsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyExtensionsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyExtensionsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves the representation of an extension by its id.
 */
func (rb *ByProjectKeyExtensionsByIDRequestMethodGet) Execute(ctx context.Context) (result *Extension, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
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
	if err != nil {
		return nil, err
	}
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
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
