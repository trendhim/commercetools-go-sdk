// Automatically generated, do not edit

package commercetools

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// ReviewURLPath is the commercetools API path.
const ReviewURLPath = "reviews"

// ReviewCreate creates a new instance of type Review
func (client *Client) ReviewCreate(ctx context.Context, draft *ReviewDraft, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	err = client.Create(ctx, ReviewURLPath, params, draft, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewQuery allows querying for type Review
func (client *Client) ReviewQuery(ctx context.Context, input *QueryInput) (result *ReviewPagedQueryResponse, err error) {
	err = client.Query(ctx, ReviewURLPath, input.toParams(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewDeleteWithKey for type Review
func (client *Client) ReviewDeleteWithKey(ctx context.Context, key string, version int, dataErasure bool, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("reviews/key=%s", key)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewGetWithKey for type Review
func (client *Client) ReviewGetWithKey(ctx context.Context, key string, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("reviews/key=%s", key)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewUpdateWithKeyInput is input for function ReviewUpdateWithKey
type ReviewUpdateWithKeyInput struct {
	Key     string
	Version int
	Actions []ReviewUpdateAction
}

// ReviewUpdateWithKey for type Review
func (client *Client) ReviewUpdateWithKey(ctx context.Context, input *ReviewUpdateWithKeyInput, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("reviews/key=%s", input.Key)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewDeleteWithID for type Review
func (client *Client) ReviewDeleteWithID(ctx context.Context, id string, version int, dataErasure bool, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	params.Set("version", strconv.Itoa(version))
	params.Set("dataErasure", strconv.FormatBool(dataErasure))
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("reviews/%s", id)
	err = client.Delete(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewGetWithID for type Review
func (client *Client) ReviewGetWithID(ctx context.Context, id string, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("reviews/%s", id)
	err = client.Get(ctx, endpoint, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReviewUpdateWithIDInput is input for function ReviewUpdateWithID
type ReviewUpdateWithIDInput struct {
	ID      string
	Version int
	Actions []ReviewUpdateAction
}

// ReviewUpdateWithID for type Review
func (client *Client) ReviewUpdateWithID(ctx context.Context, input *ReviewUpdateWithIDInput, opts ...RequestOption) (result *Review, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}
	endpoint := fmt.Sprintf("reviews/%s", input.ID)
	err = client.Update(ctx, endpoint, params, input.Version, input.Actions, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
