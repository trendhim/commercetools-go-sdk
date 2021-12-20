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

type ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves an import summary for the given import sink.
*
*	The import summary is calculated on demand.
*
 */
func (rb *ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestMethodGet) Execute(ctx context.Context) (result *ImportSummary, err error) {
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
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}