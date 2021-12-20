// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	It is used to calculate the [taxPortions](/../api/projects/carts#taxedprice) field in a Cart or Order.
 */
type SubRate struct {
	// Name of the SubRate.
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type TaxCategory struct {
	// Unique ID of the TaxCategory.
	ID string `json:"id"`
	// Current version of the TaxCategory.
	Version int `json:"version"`
	// Date and time (UTC) the TaxCategory was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the TaxCategory was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Name of the TaxCategory.
	Name string `json:"name"`
	// Description of the TaxCategory.
	Description *string `json:"description,omitempty"`
	// Tax rates and subrates of states and countries. Each TaxRate in the array has a unique ID assigned by the platform.
	Rates []TaxRate `json:"rates"`
	// User-defined unique identifier for the TaxCategory.
	Key *string `json:"key,omitempty"`
}

type TaxCategoryDraft struct {
	// Name of the TaxCategory.
	Name string `json:"name"`
	// Description of the TaxCategory.
	Description *string `json:"description,omitempty"`
	// Tax rates and subrates of states and countries.
	Rates []TaxRateDraft `json:"rates,omitempty"`
	// User-defined unique identifier for the TaxCategory.
	Key *string `json:"key,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [TaxCategory](ctp:api:type:TaxCategory).
*
 */
type TaxCategoryPagedQueryResponse struct {
	// Number of results requested in the query request.
	Limit int `json:"limit"`
	// Offset supplied by the client or the server default.
	// It is the number of elements skipped, not a page number.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/contract#queries).
	Total *int `json:"total,omitempty"`
	// [TaxCategories](ctp:api:type:TaxCategory) matching the query.
	Results []TaxCategory `json:"results"`
}

/**
*	[Reference](/../api/types#reference) to a [TaxCategory](ctp:api:type:TaxCategory).
*
 */
type TaxCategoryReference struct {
	// Unique ID of the referenced [TaxCategory](ctp:api:type:TaxCategory).
	ID string `json:"id"`
	// Contains the representation of the expanded TaxCategory. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for TaxCategory.
	Obj *TaxCategory `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryReference) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "tax-category", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](/../api/types#resourceidentifier) to a [TaxCategory](ctp:api:type:TaxCategory).
*
 */
type TaxCategoryResourceIdentifier struct {
	// Unique ID of the referenced [TaxCategory](ctp:api:type:TaxCategory). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced [TaxCategory](ctp:api:type:TaxCategory). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "tax-category", Alias: (*Alias)(&obj)})
}

type TaxCategoryUpdate struct {
	// Expected version of the TaxCategory on which the changes should be applied. If the expected version does not match the actual version, a 409 Conflict will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the TaxCategory.
	Actions []TaxCategoryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TaxCategoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias TaxCategoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type TaxCategoryUpdateAction interface{}

func mapDiscriminatorTaxCategoryUpdateAction(input interface{}) (TaxCategoryUpdateAction, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "addTaxRate":
		obj := TaxCategoryAddTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := TaxCategoryChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeTaxRate":
		obj := TaxCategoryRemoveTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "replaceTaxRate":
		obj := TaxCategoryReplaceTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := TaxCategorySetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := TaxCategorySetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type TaxRate struct {
	// Present if the TaxRate is part of a [TaxCategory](ctp:api:type:TaxCategory).
	// Absent for external TaxRates in [LineItem](ctp:api:type:LineItem), [CustomLineItem](ctp:api:type:CustomLineItem), and [ShippingInfo](ctp:api:type:ShippingInfo).
	ID *string `json:"id,omitempty"`
	// Name of the TaxRate.
	Name string `json:"name"`
	// Tax rate. If subrates are used, the amount must be the sum of all subrates.
	Amount float64 `json:"amount"`
	// If `true`, tax is included in [Prices](ctp:api:type:Price) and the `taxedPrice` is present on [LineItems](ctp:api:type:LineItem). In this case, the platform calculates the `totalNet` price based on the TaxRate.
	IncludedInPrice bool `json:"includedInPrice"`
	// Country in which the tax rate is applied in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format.
	Country string `json:"country"`
	// State within the country, such as Texas in the United States.
	State *string `json:"state,omitempty"`
	// Used to calculate the [taxPortions](/../api/projects/carts#taxedprice) field in a Cart or Order. It is useful if the total tax of a country (such as the US) is a combination of multiple taxes (such as state and local taxes).
	SubRates []SubRate `json:"subRates,omitempty"`
}

type TaxRateDraft struct {
	// Name of the TaxRate.
	Name string `json:"name"`
	// Tax rate.
	// Must be supplied if no `subRates` are specified.
	// If `subRates` are specified, this field can be omitted or it must be the sum of amounts of all `subRates`.
	Amount *float64 `json:"amount,omitempty"`
	// Set to `true`, if tax should be included in [Prices](ctp:api:type:Price) and the `taxedPrice` should be present on [Line Items](ctp:api:type:LineItem). In this case, the platform calculates the `totalNet` price based on the TaxRate.
	IncludedInPrice bool `json:"includedInPrice"`
	// Country in which the tax rate is applied in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format.
	Country string `json:"country"`
	// State within the country, such as Texas in the United States.
	State *string `json:"state,omitempty"`
	// Used to calculate the [taxPortions](/../api/projects/carts#taxedprice) field in a Cart or Order. It is useful if the total tax of a country (such as the US) is a combination of multiple taxes (such as state and local taxes).
	SubRates []SubRate `json:"subRates,omitempty"`
}

type TaxCategoryAddTaxRateAction struct {
	// Value to append to the `rates` array.
	TaxRate TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryAddTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryAddTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategoryChangeNameAction struct {
	// New value to set. Must not be empty.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type TaxCategoryRemoveTaxRateAction struct {
	// ID of the TaxRate to remove.
	TaxRateId string `json:"taxRateId"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryRemoveTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryRemoveTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategoryReplaceTaxRateAction struct {
	// ID of the TaxRate to replace.
	TaxRateId string `json:"taxRateId"`
	// New TaxRate to replace with.
	TaxRate TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryReplaceTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReplaceTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "replaceTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategorySetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type TaxCategorySetKeyAction struct {
	// Value to set. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategorySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}