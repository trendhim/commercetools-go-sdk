package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyProductSelectionsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductSelectionsRequestMethodGetInput
}

func (r *ByProjectKeyProductSelectionsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductSelectionsRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyProductSelectionsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Add(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyProductSelectionsRequestMethodGet) Expand(v []string) *ByProjectKeyProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsRequestMethodGet) Sort(v []string) *ByProjectKeyProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsRequestMethodGet) Limit(v int) *ByProjectKeyProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsRequestMethodGet) Offset(v int) *ByProjectKeyProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsRequestMethodGet) WithTotal(v bool) *ByProjectKeyProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsRequestMethodGet) Where(v []string) *ByProjectKeyProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsRequestMethodGet) WithQueryParams(input ByProjectKeyProductSelectionsRequestMethodGetInput) *ByProjectKeyProductSelectionsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductSelectionsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductSelectionsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductSelectionsRequestMethodGet) Execute(ctx context.Context) (result *ProductSelectionPagedQueryResponse, err error) {
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
		if err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj

	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
