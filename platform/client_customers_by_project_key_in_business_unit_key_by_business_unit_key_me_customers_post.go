package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestMethodPost struct {
	body    MyBusinessUnitAssociateDraft
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	The My Business Unit endpoint does not support assigning existing Customers to a Business Unit.
*	Associates with the `UpdateAssociates` [Permission](ctp:api:type:Permission) can use this endpoint to create a new Customer and associate it with the Business Unit.
*	If the required [Permission](/projects/associate-roles#permission) is missing, an [AssociateMissingPermission](/errors#associatemissingpermission) error is returned.
*
 */
func (rb *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestMethodPost) Execute(ctx context.Context) (result *CustomerSignInResult, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
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
	case 201:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
