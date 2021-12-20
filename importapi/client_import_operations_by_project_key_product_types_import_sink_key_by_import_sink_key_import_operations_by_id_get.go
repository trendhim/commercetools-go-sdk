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

type ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves the import operation with the given id.
*
 */
func (rb *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) Execute(ctx context.Context) (result *ImportOperation, err error) {
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