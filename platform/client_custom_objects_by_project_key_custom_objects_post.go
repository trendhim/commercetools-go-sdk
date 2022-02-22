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

type ByProjectKeyCustomObjectsRequestMethodPost struct {
	body    CustomObjectDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomObjectsRequestMethodPostInput
}

func (r *ByProjectKeyCustomObjectsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomObjectsRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyCustomObjectsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomObjectsRequestMethodPost) Expand(v []string) *ByProjectKeyCustomObjectsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsRequestMethodPost) WithQueryParams(input ByProjectKeyCustomObjectsRequestMethodPostInput) *ByProjectKeyCustomObjectsRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomObjectsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyCustomObjectsRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates a new custom object or updates an existing custom object.
*	If an object with the given container/key exists,
*	the object will be replaced with the new value and the version is incremented.
*	If the request contains a version and an object with the given container/key exists then the version
*	must match the version of the existing object. Concurrent updates for the same custom object still can result
*	in a Conflict (409) even if the version is not provided.
*	Fields with null values will not be saved.
*
 */
func (rb *ByProjectKeyCustomObjectsRequestMethodPost) Execute(ctx context.Context) (result *CustomObject, err error) {
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
	case 200, 201:
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
