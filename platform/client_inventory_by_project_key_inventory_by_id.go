// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInventoryByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyInventoryByIDRequestBuilder) Get() *ByProjectKeyInventoryByIDRequestMethodGet {
	return &ByProjectKeyInventoryByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/inventory/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInventoryByIDRequestBuilder) Post(body InventoryEntryUpdate) *ByProjectKeyInventoryByIDRequestMethodPost {
	return &ByProjectKeyInventoryByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventory/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInventoryByIDRequestBuilder) Delete() *ByProjectKeyInventoryByIDRequestMethodDelete {
	return &ByProjectKeyInventoryByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/inventory/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}