package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder struct {
	projectKey string
	customerId string
	client     *Client
}

/**
*	Retrieves the recently modified active Cart of a Customer with [CartOrigin](ctp:api:type:CartOrigin) `Customer`. If no active Cart exists, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
*	To ensure the Cart is up-to-date with current values (such as Prices and Discounts), use the [Recalculate](ctp:api:type:CartRecalculateAction) update action.
*
 */
func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder) Get() *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	return &ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/carts/customer-id=%s", rb.projectKey, rb.customerId),
		client: rb.client,
	}
}

/**
*	Checks if a Cart of a Customer exists. Returns a `200 OK` status if the Cart exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder) Head() *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodHead {
	return &ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodHead{
		url:    fmt.Sprintf("/%s/carts/customer-id=%s", rb.projectKey, rb.customerId),
		client: rb.client,
	}
}
