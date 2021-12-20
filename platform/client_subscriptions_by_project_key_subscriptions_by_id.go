// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeySubscriptionsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves the representation of a subscription by its id.
 */
func (rb *ByProjectKeySubscriptionsByIDRequestBuilder) Get() *ByProjectKeySubscriptionsByIDRequestMethodGet {
	return &ByProjectKeySubscriptionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/subscriptions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeySubscriptionsByIDRequestBuilder) Post(body SubscriptionUpdate) *ByProjectKeySubscriptionsByIDRequestMethodPost {
	return &ByProjectKeySubscriptionsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/subscriptions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeySubscriptionsByIDRequestBuilder) Delete() *ByProjectKeySubscriptionsByIDRequestMethodDelete {
	return &ByProjectKeySubscriptionsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/subscriptions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}