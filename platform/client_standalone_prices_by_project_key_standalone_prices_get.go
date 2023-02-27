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

type ByProjectKeyStandalonePricesRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyStandalonePricesRequestMethodGetInput
}

func (r *ByProjectKeyStandalonePricesRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyStandalonePricesRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyStandalonePricesRequestMethodGetInput) Values() url.Values {
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
			values.Set(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyStandalonePricesRequestMethodGet) Expand(v []string) *ByProjectKeyStandalonePricesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyStandalonePricesRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyStandalonePricesRequestMethodGet) Sort(v []string) *ByProjectKeyStandalonePricesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyStandalonePricesRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyStandalonePricesRequestMethodGet) Limit(v int) *ByProjectKeyStandalonePricesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyStandalonePricesRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyStandalonePricesRequestMethodGet) Offset(v int) *ByProjectKeyStandalonePricesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyStandalonePricesRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyStandalonePricesRequestMethodGet) WithTotal(v bool) *ByProjectKeyStandalonePricesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyStandalonePricesRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyStandalonePricesRequestMethodGet) Where(v []string) *ByProjectKeyStandalonePricesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyStandalonePricesRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyStandalonePricesRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyStandalonePricesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyStandalonePricesRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyStandalonePricesRequestMethodGet) WithQueryParams(input ByProjectKeyStandalonePricesRequestMethodGetInput) *ByProjectKeyStandalonePricesRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyStandalonePricesRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyStandalonePricesRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyStandalonePricesRequestMethodGet) Execute(ctx context.Context) (result *StandalonePricePagedQueryResponse, err error) {
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
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
