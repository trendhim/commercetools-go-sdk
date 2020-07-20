// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ExtensionURLPath is the commercetools API path.
const ExtensionURLPath = "extensions"

// ExtensionCreate creates a new instance of type Extension
func (client *Client) ExtensionCreate(ctx context.Context, draft *ExtensionDraft, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, ExtensionURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionQuery allows querying for type Extension
func (client *Client) ExtensionQuery(ctx context.Context, input *QueryInput) (result *ExtensionPagedQueryResponse, err error) {
	err = client.Query(ctx, ExtensionURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionDeleteWithKey for type Extension
func (client *Client) ExtensionDeleteWithKey(ctx context.Context, key string, version int, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("extensions/key=%s", key)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionGetWithKey Retrieves the representation of an extension by its key.
func (client *Client) ExtensionGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("extensions/key=%s", key)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionUpdateWithKeyInput is input for function ExtensionUpdateWithKey
type ExtensionUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ExtensionUpdateAction
}

// ExtensionUpdateWithKey for type Extension
func (client *Client) ExtensionUpdateWithKey(ctx context.Context, input *ExtensionUpdateWithKeyInput, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("extensions/key=%s", input.Key)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionDeleteWithID for type Extension
func (client *Client) ExtensionDeleteWithID(ctx context.Context, id string, version int, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))

	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("extensions/%s", id)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionGetWithID Retrieves the representation of an extension by its id.
func (client *Client) ExtensionGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("extensions/%s", id)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ExtensionUpdateWithIDInput is input for function ExtensionUpdateWithID
type ExtensionUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ExtensionUpdateAction
}

// ExtensionUpdateWithID for type Extension
func (client *Client) ExtensionUpdateWithID(ctx context.Context, input *ExtensionUpdateWithIDInput, opts ...RequestOption) (result *Extension, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("extensions/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
