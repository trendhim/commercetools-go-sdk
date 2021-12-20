// Generated file, please do not change!!!
package importapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves the ImportOperation of a given ID.
*
 */
func (rb *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) Execute(ctx context.Context) (result *ImportOperation, err error) {
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
	case 404, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}