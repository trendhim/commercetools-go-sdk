// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Attribute struct {
	Name string `json:"name"`
	// A valid JSON value, based on an AttributeDefinition.
	Value interface{} `json:"value"`
}

// CategoryOrderHints is something
type CategoryOrderHints map[string]string
type FacetResult interface{}

func mapDiscriminatorFacetResult(input interface{}) (FacetResult, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "filter":
		obj := FilteredFacetResult{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "range":
		obj := RangeFacetResult{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "terms":
		obj := TermFacetResult{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type FacetResultRange struct {
	From         float64 `json:"from"`
	FromStr      string  `json:"fromStr"`
	To           float64 `json:"to"`
	ToStr        string  `json:"toStr"`
	Count        int     `json:"count"`
	ProductCount *int    `json:"productCount,omitempty"`
	Total        float64 `json:"total"`
	Min          float64 `json:"min"`
	Max          float64 `json:"max"`
	Mean         float64 `json:"mean"`
}

type FacetResultTerm struct {
	Term         interface{} `json:"term"`
	Count        int         `json:"count"`
	ProductCount *int        `json:"productCount,omitempty"`
}

// FacetResults is something
type FacetResults map[string]FacetResult
type FacetTypes string

const (
	FacetTypesTerms  FacetTypes = "terms"
	FacetTypesRange  FacetTypes = "range"
	FacetTypesFilter FacetTypes = "filter"
)

type FilteredFacetResult struct {
	Count        int  `json:"count"`
	ProductCount *int `json:"productCount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj FilteredFacetResult) MarshalJSON() ([]byte, error) {
	type Alias FilteredFacetResult
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "filter", Alias: (*Alias)(&obj)})
}

type Product struct {
	// The unique ID of the product.
	ID string `json:"id"`
	// The current version of the product.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-specific unique identifier for the product.
	// *Product keys are different from product variant keys.*
	Key         *string              `json:"key,omitempty"`
	ProductType ProductTypeReference `json:"productType"`
	// The product data in the master catalog.
	MasterData  ProductCatalogData    `json:"masterData"`
	TaxCategory *TaxCategoryReference `json:"taxCategory,omitempty"`
	State       *StateReference       `json:"state,omitempty"`
	// Statistics about the review ratings taken into account for this product.
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
}

type ProductCatalogData struct {
	Published        bool        `json:"published"`
	Current          ProductData `json:"current"`
	Staged           ProductData `json:"staged"`
	HasStagedChanges bool        `json:"hasStagedChanges"`
}

type ProductData struct {
	Name               LocalizedString     `json:"name"`
	Categories         []CategoryReference `json:"categories"`
	CategoryOrderHints *CategoryOrderHints `json:"categoryOrderHints,omitempty"`
	Description        *LocalizedString    `json:"description,omitempty"`
	Slug               LocalizedString     `json:"slug"`
	MetaTitle          *LocalizedString    `json:"metaTitle,omitempty"`
	MetaDescription    *LocalizedString    `json:"metaDescription,omitempty"`
	MetaKeywords       *LocalizedString    `json:"metaKeywords,omitempty"`
	MasterVariant      ProductVariant      `json:"masterVariant"`
	Variants           []ProductVariant    `json:"variants"`
	SearchKeywords     SearchKeywords      `json:"searchKeywords"`
}

type ProductDraft struct {
	// A predefined product type assigned to the product.
	// All products must have a product type.
	ProductType ProductTypeResourceIdentifier `json:"productType"`
	Name        LocalizedString               `json:"name"`
	// Human-readable identifiers usually used as deep-link URLs for the product.
	// A slug must be unique across a project, but a product can have the same slug for different languages.
	// Slugs have a maximum size of 256.
	// Valid characters are: alphabetic characters (`A-Z, a-z`), numeric characters (`0-9`), underscores (`_`) and hyphens (`-`).
	Slug LocalizedString `json:"slug"`
	// User-specific unique identifier for the product.
	Key         *string          `json:"key,omitempty"`
	Description *LocalizedString `json:"description,omitempty"`
	// Categories assigned to the product.
	Categories         []CategoryResourceIdentifier `json:"categories,omitempty"`
	CategoryOrderHints *CategoryOrderHints          `json:"categoryOrderHints,omitempty"`
	MetaTitle          *LocalizedString             `json:"metaTitle,omitempty"`
	MetaDescription    *LocalizedString             `json:"metaDescription,omitempty"`
	MetaKeywords       *LocalizedString             `json:"metaKeywords,omitempty"`
	// The master product variant.
	// Required if the `variants` array has product variants.
	MasterVariant *ProductVariantDraft `json:"masterVariant,omitempty"`
	// An array of related product variants.
	Variants       []ProductVariantDraft          `json:"variants,omitempty"`
	TaxCategory    *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	SearchKeywords *SearchKeywords                `json:"searchKeywords,omitempty"`
	State          *StateResourceIdentifier       `json:"state,omitempty"`
	// If `true`, the product is published immediately.
	Publish *bool `json:"publish,omitempty"`
}

type ProductPagedQueryResponse struct {
	Limit   int       `json:"limit"`
	Count   int       `json:"count"`
	Total   *int      `json:"total,omitempty"`
	Offset  int       `json:"offset"`
	Results []Product `json:"results"`
}

type ProductProjection struct {
	// The unique ID of the Product.
	ID string `json:"id"`
	// The current version of the Product.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-specific unique identifier of the Product.
	Key         *string              `json:"key,omitempty"`
	ProductType ProductTypeReference `json:"productType"`
	Name        LocalizedString      `json:"name"`
	Description *LocalizedString     `json:"description,omitempty"`
	Slug        LocalizedString      `json:"slug"`
	// References to categories the product is in.
	Categories         []CategoryReference   `json:"categories"`
	CategoryOrderHints *CategoryOrderHints   `json:"categoryOrderHints,omitempty"`
	MetaTitle          *LocalizedString      `json:"metaTitle,omitempty"`
	MetaDescription    *LocalizedString      `json:"metaDescription,omitempty"`
	MetaKeywords       *LocalizedString      `json:"metaKeywords,omitempty"`
	SearchKeywords     *SearchKeywords       `json:"searchKeywords,omitempty"`
	HasStagedChanges   *bool                 `json:"hasStagedChanges,omitempty"`
	Published          *bool                 `json:"published,omitempty"`
	MasterVariant      ProductVariant        `json:"masterVariant"`
	Variants           []ProductVariant      `json:"variants"`
	TaxCategory        *TaxCategoryReference `json:"taxCategory,omitempty"`
	State              *StateReference       `json:"state,omitempty"`
	// Statistics about the review ratings taken into account for this product.
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
}

type ProductProjectionPagedQueryResponse struct {
	Limit   int                 `json:"limit"`
	Count   int                 `json:"count"`
	Total   *int                `json:"total,omitempty"`
	Offset  int                 `json:"offset"`
	Results []ProductProjection `json:"results"`
}

type ProductProjectionPagedSearchResponse struct {
	Limit   int                 `json:"limit"`
	Count   int                 `json:"count"`
	Total   *int                `json:"total,omitempty"`
	Offset  int                 `json:"offset"`
	Results []ProductProjection `json:"results"`
	Facets  FacetResults        `json:"facets"`
}

type ProductReference struct {
	// Unique ID of the referenced resource.
	ID  string   `json:"id"`
	Obj *Product `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductReference) MarshalJSON() ([]byte, error) {
	type Alias ProductReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product", Alias: (*Alias)(&obj)})
}

type ProductResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ProductResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product", Alias: (*Alias)(&obj)})
}

type ProductUpdate struct {
	Version int                   `json:"version"`
	Actions []ProductUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type ProductUpdateAction interface{}

func mapDiscriminatorProductUpdateAction(input interface{}) (ProductUpdateAction, error) {

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
	case "addAsset":
		obj := ProductAddAssetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addExternalImage":
		obj := ProductAddExternalImageAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addPrice":
		obj := ProductAddPriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addToCategory":
		obj := ProductAddToCategoryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addVariant":
		obj := ProductAddVariantAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssetName":
		obj := ProductChangeAssetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssetOrder":
		obj := ProductChangeAssetOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeMasterVariant":
		obj := ProductChangeMasterVariantAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ProductChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changePrice":
		obj := ProductChangePriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeSlug":
		obj := ProductChangeSlugAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "legacySetSku":
		obj := ProductLegacySetSkuAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "moveImageToPosition":
		obj := ProductMoveImageToPositionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "publish":
		obj := ProductPublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAsset":
		obj := ProductRemoveAssetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeFromCategory":
		obj := ProductRemoveFromCategoryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeImage":
		obj := ProductRemoveImageAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removePrice":
		obj := ProductRemovePriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeVariant":
		obj := ProductRemoveVariantAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "revertStagedChanges":
		obj := ProductRevertStagedChangesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "revertStagedVariantChanges":
		obj := ProductRevertStagedVariantChangesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetCustomField":
		obj := ProductSetAssetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetCustomType":
		obj := ProductSetAssetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetDescription":
		obj := ProductSetAssetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetKey":
		obj := ProductSetAssetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetSources":
		obj := ProductSetAssetSourcesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetTags":
		obj := ProductSetAssetTagsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAttribute":
		obj := ProductSetAttributeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAttributeInAllVariants":
		obj := ProductSetAttributeInAllVariantsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCategoryOrderHint":
		obj := ProductSetCategoryOrderHintAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := ProductSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDiscountedPrice":
		obj := ProductSetDiscountedPriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setImageLabel":
		obj := ProductSetImageLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ProductSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaDescription":
		obj := ProductSetMetaDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaKeywords":
		obj := ProductSetMetaKeywordsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaTitle":
		obj := ProductSetMetaTitleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPrices":
		obj := ProductSetPricesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setProductPriceCustomField":
		obj := ProductSetProductPriceCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setProductPriceCustomType":
		obj := ProductSetProductPriceCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setProductVariantKey":
		obj := ProductSetProductVariantKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSearchKeywords":
		obj := ProductSetSearchKeywordsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSku":
		obj := ProductSetSkuAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTaxCategory":
		obj := ProductSetTaxCategoryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionState":
		obj := ProductTransitionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "unpublish":
		obj := ProductUnpublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ProductVariant struct {
	ID                    int                         `json:"id"`
	Sku                   *string                     `json:"sku,omitempty"`
	Key                   *string                     `json:"key,omitempty"`
	Prices                []Price                     `json:"prices,omitempty"`
	Attributes            []Attribute                 `json:"attributes,omitempty"`
	Price                 *Price                      `json:"price,omitempty"`
	Images                []Image                     `json:"images,omitempty"`
	Assets                []Asset                     `json:"assets,omitempty"`
	Availability          *ProductVariantAvailability `json:"availability,omitempty"`
	IsMatchingVariant     *bool                       `json:"isMatchingVariant,omitempty"`
	ScopedPrice           *ScopedPrice                `json:"scopedPrice,omitempty"`
	ScopedPriceDiscounted *bool                       `json:"scopedPriceDiscounted,omitempty"`
}

type ProductVariantAvailability struct {
	IsOnStock         *bool                                 `json:"isOnStock,omitempty"`
	RestockableInDays *int                                  `json:"restockableInDays,omitempty"`
	AvailableQuantity *int                                  `json:"availableQuantity,omitempty"`
	Channels          *ProductVariantChannelAvailabilityMap `json:"channels,omitempty"`
}

type ProductVariantChannelAvailability struct {
	IsOnStock         *bool `json:"isOnStock,omitempty"`
	RestockableInDays *int  `json:"restockableInDays,omitempty"`
	AvailableQuantity *int  `json:"availableQuantity,omitempty"`
}

// ProductVariantChannelAvailabilityMap is something
type ProductVariantChannelAvailabilityMap map[string]ProductVariantChannelAvailability
type ProductVariantDraft struct {
	Sku        *string      `json:"sku,omitempty"`
	Key        *string      `json:"key,omitempty"`
	Prices     []PriceDraft `json:"prices,omitempty"`
	Attributes []Attribute  `json:"attributes,omitempty"`
	Images     []Image      `json:"images,omitempty"`
	Assets     []AssetDraft `json:"assets,omitempty"`
}

type RangeFacetResult struct {
	Ranges []FacetResultRange `json:"ranges"`
}

// MarshalJSON override to set the discriminator value
func (obj RangeFacetResult) MarshalJSON() ([]byte, error) {
	type Alias RangeFacetResult
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "range", Alias: (*Alias)(&obj)})
}

type SearchKeyword struct {
	Text             string           `json:"text"`
	SuggestTokenizer SuggestTokenizer `json:"suggestTokenizer,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SearchKeyword) UnmarshalJSON(data []byte) error {
	type Alias SearchKeyword
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.SuggestTokenizer != nil {
		var err error
		obj.SuggestTokenizer, err = mapDiscriminatorSuggestTokenizer(obj.SuggestTokenizer)
		if err != nil {
			return err
		}
	}
	return nil
}

// SearchKeywords is something
type SearchKeywords map[string][]SearchKeyword
type SuggestTokenizer interface{}

func mapDiscriminatorSuggestTokenizer(input interface{}) (SuggestTokenizer, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "custom":
		obj := CustomTokenizer{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "whitespace":
		obj := WhitespaceTokenizer{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CustomTokenizer struct {
	Inputs []string `json:"inputs"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomTokenizer) MarshalJSON() ([]byte, error) {
	type Alias CustomTokenizer
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "custom", Alias: (*Alias)(&obj)})
}

type Suggestion struct {
	// The suggested text.
	Text string `json:"text"`
}

// SuggestionResult is something
type SuggestionResult map[string][]Suggestion
type TermFacetResult struct {
	DataType TermFacetResultType `json:"dataType"`
	Missing  int                 `json:"missing"`
	Total    int                 `json:"total"`
	Other    int                 `json:"other"`
	Terms    []FacetResultTerm   `json:"terms"`
}

// MarshalJSON override to set the discriminator value
func (obj TermFacetResult) MarshalJSON() ([]byte, error) {
	type Alias TermFacetResult
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "terms", Alias: (*Alias)(&obj)})
}

type TermFacetResultType string

const (
	TermFacetResultTypeText     TermFacetResultType = "text"
	TermFacetResultTypeDate     TermFacetResultType = "date"
	TermFacetResultTypeTime     TermFacetResultType = "time"
	TermFacetResultTypeDatetime TermFacetResultType = "datetime"
	TermFacetResultTypeBoolean  TermFacetResultType = "boolean"
	TermFacetResultTypeNumber   TermFacetResultType = "number"
)

type WhitespaceTokenizer struct {
}

// MarshalJSON override to set the discriminator value
func (obj WhitespaceTokenizer) MarshalJSON() ([]byte, error) {
	type Alias WhitespaceTokenizer
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "whitespace", Alias: (*Alias)(&obj)})
}

type ProductAddAssetAction struct {
	VariantId *int       `json:"variantId,omitempty"`
	Sku       *string    `json:"sku,omitempty"`
	Staged    *bool      `json:"staged,omitempty"`
	Asset     AssetDraft `json:"asset"`
	// Position of the new asset inside the existing list (from `0` to the size of the list)
	Position *int `json:"position,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

type ProductAddExternalImageAction struct {
	VariantId *int    `json:"variantId,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	Image     Image   `json:"image"`
	Staged    *bool   `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddExternalImageAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddExternalImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addExternalImage", Alias: (*Alias)(&obj)})
}

type ProductAddPriceAction struct {
	VariantId *int       `json:"variantId,omitempty"`
	Sku       *string    `json:"sku,omitempty"`
	Price     PriceDraft `json:"price"`
	Staged    *bool      `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddPriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPrice", Alias: (*Alias)(&obj)})
}

type ProductAddToCategoryAction struct {
	Category  CategoryResourceIdentifier `json:"category"`
	OrderHint *string                    `json:"orderHint,omitempty"`
	Staged    *bool                      `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddToCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddToCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addToCategory", Alias: (*Alias)(&obj)})
}

type ProductAddVariantAction struct {
	Sku        *string      `json:"sku,omitempty"`
	Key        *string      `json:"key,omitempty"`
	Prices     []PriceDraft `json:"prices,omitempty"`
	Images     []Image      `json:"images,omitempty"`
	Attributes []Attribute  `json:"attributes,omitempty"`
	Staged     *bool        `json:"staged,omitempty"`
	Assets     []Asset      `json:"assets,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductAddVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addVariant", Alias: (*Alias)(&obj)})
}

type ProductChangeAssetNameAction struct {
	VariantId *int            `json:"variantId,omitempty"`
	Sku       *string         `json:"sku,omitempty"`
	Staged    *bool           `json:"staged,omitempty"`
	AssetId   *string         `json:"assetId,omitempty"`
	AssetKey  *string         `json:"assetKey,omitempty"`
	Name      LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangeAssetNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeAssetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetName", Alias: (*Alias)(&obj)})
}

type ProductChangeAssetOrderAction struct {
	VariantId  *int     `json:"variantId,omitempty"`
	Sku        *string  `json:"sku,omitempty"`
	Staged     *bool    `json:"staged,omitempty"`
	AssetOrder []string `json:"assetOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangeAssetOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeAssetOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetOrder", Alias: (*Alias)(&obj)})
}

type ProductChangeMasterVariantAction struct {
	VariantId *int    `json:"variantId,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	Staged    *bool   `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangeMasterVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeMasterVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMasterVariant", Alias: (*Alias)(&obj)})
}

type ProductChangeNameAction struct {
	Name   LocalizedString `json:"name"`
	Staged *bool           `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ProductChangePriceAction struct {
	// ID of the [Price](#price)
	PriceId string     `json:"priceId"`
	Price   PriceDraft `json:"price"`
	Staged  *bool      `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangePriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangePriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePrice", Alias: (*Alias)(&obj)})
}

type ProductChangeSlugAction struct {
	// Every slug must be unique across a project, but a product can have the same slug for different languages.
	// Allowed are alphabetic, numeric, underscore (`_`) and hyphen (`-`) characters.
	// Maximum size is `256`.
	Slug   LocalizedString `json:"slug"`
	Staged *bool           `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductChangeSlugAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSlug", Alias: (*Alias)(&obj)})
}

type ProductLegacySetSkuAction struct {
	Sku       *string `json:"sku,omitempty"`
	VariantId int     `json:"variantId"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductLegacySetSkuAction) MarshalJSON() ([]byte, error) {
	type Alias ProductLegacySetSkuAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "legacySetSku", Alias: (*Alias)(&obj)})
}

type ProductMoveImageToPositionAction struct {
	VariantId *int    `json:"variantId,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	// The URL of the image
	ImageUrl string `json:"imageUrl"`
	Position int    `json:"position"`
	Staged   *bool  `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductMoveImageToPositionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductMoveImageToPositionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "moveImageToPosition", Alias: (*Alias)(&obj)})
}

type ProductPublishAction struct {
	Scope *ProductPublishScope `json:"scope,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductPublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "publish", Alias: (*Alias)(&obj)})
}

type ProductRemoveAssetAction struct {
	VariantId *int    `json:"variantId,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	Staged    *bool   `json:"staged,omitempty"`
	AssetId   *string `json:"assetId,omitempty"`
	AssetKey  *string `json:"assetKey,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

type ProductRemoveFromCategoryAction struct {
	Category CategoryResourceIdentifier `json:"category"`
	Staged   *bool                      `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemoveFromCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveFromCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeFromCategory", Alias: (*Alias)(&obj)})
}

type ProductRemoveImageAction struct {
	VariantId *int    `json:"variantId,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	// The URL of the image.
	ImageUrl string `json:"imageUrl"`
	Staged   *bool  `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemoveImageAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeImage", Alias: (*Alias)(&obj)})
}

type ProductRemovePriceAction struct {
	// ID of the [Price](#price)
	PriceId string `json:"priceId"`
	Staged  *bool  `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemovePriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemovePriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePrice", Alias: (*Alias)(&obj)})
}

type ProductRemoveVariantAction struct {
	ID     *int    `json:"id,omitempty"`
	Sku    *string `json:"sku,omitempty"`
	Staged *bool   `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRemoveVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeVariant", Alias: (*Alias)(&obj)})
}

type ProductRevertStagedChangesAction struct {
}

// MarshalJSON override to set the discriminator value
func (obj ProductRevertStagedChangesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertStagedChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "revertStagedChanges", Alias: (*Alias)(&obj)})
}

type ProductRevertStagedVariantChangesAction struct {
	VariantId int `json:"variantId"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductRevertStagedVariantChangesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertStagedVariantChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "revertStagedVariantChanges", Alias: (*Alias)(&obj)})
}

type ProductSetAssetCustomFieldAction struct {
	VariantId *int        `json:"variantId,omitempty"`
	Sku       *string     `json:"sku,omitempty"`
	Staged    *bool       `json:"staged,omitempty"`
	AssetId   *string     `json:"assetId,omitempty"`
	AssetKey  *string     `json:"assetKey,omitempty"`
	Name      string      `json:"name"`
	Value     interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

type ProductSetAssetCustomTypeAction struct {
	VariantId *int    `json:"variantId,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	Staged    *bool   `json:"staged,omitempty"`
	AssetId   *string `json:"assetId,omitempty"`
	AssetKey  *string `json:"assetKey,omitempty"`
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// If set, the custom fields are set to this new value.
	Fields *interface{} `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

type ProductSetAssetDescriptionAction struct {
	VariantId   *int             `json:"variantId,omitempty"`
	Sku         *string          `json:"sku,omitempty"`
	Staged      *bool            `json:"staged,omitempty"`
	AssetId     *string          `json:"assetId,omitempty"`
	AssetKey    *string          `json:"assetKey,omitempty"`
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetDescription", Alias: (*Alias)(&obj)})
}

type ProductSetAssetKeyAction struct {
	VariantId *int    `json:"variantId,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	Staged    *bool   `json:"staged,omitempty"`
	AssetId   string  `json:"assetId"`
	// User-defined identifier for the asset.
	// If left blank or set to `null`, the asset key is unset/removed.
	AssetKey *string `json:"assetKey,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

type ProductSetAssetSourcesAction struct {
	VariantId *int          `json:"variantId,omitempty"`
	Sku       *string       `json:"sku,omitempty"`
	Staged    *bool         `json:"staged,omitempty"`
	AssetId   *string       `json:"assetId,omitempty"`
	AssetKey  *string       `json:"assetKey,omitempty"`
	Sources   []AssetSource `json:"sources"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

type ProductSetAssetTagsAction struct {
	VariantId *int     `json:"variantId,omitempty"`
	Sku       *string  `json:"sku,omitempty"`
	Staged    *bool    `json:"staged,omitempty"`
	AssetId   *string  `json:"assetId,omitempty"`
	AssetKey  *string  `json:"assetKey,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetTagsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetTags", Alias: (*Alias)(&obj)})
}

type ProductSetAttributeAction struct {
	VariantId *int    `json:"variantId,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	Name      string  `json:"name"`
	// If the attribute exists and the value is omitted or set to `null`, the attribute is removed.
	// If the attribute exists and a value is provided, the new value is applied.
	// If the attribute does not exist and a value is provided, it is added as a new attribute.
	Value  interface{} `json:"value,omitempty"`
	Staged *bool       `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttribute", Alias: (*Alias)(&obj)})
}

type ProductSetAttributeInAllVariantsAction struct {
	Name string `json:"name"`
	// The same update behavior as for Set Attribute applies.
	Value  interface{} `json:"value,omitempty"`
	Staged *bool       `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetAttributeInAllVariantsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAttributeInAllVariantsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttributeInAllVariants", Alias: (*Alias)(&obj)})
}

type ProductSetCategoryOrderHintAction struct {
	CategoryId string  `json:"categoryId"`
	OrderHint  *string `json:"orderHint,omitempty"`
	Staged     *bool   `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetCategoryOrderHintAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetCategoryOrderHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCategoryOrderHint", Alias: (*Alias)(&obj)})
}

type ProductSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
	Staged      *bool            `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ProductSetDiscountedPriceAction struct {
	PriceId    string                `json:"priceId"`
	Staged     *bool                 `json:"staged,omitempty"`
	Discounted *DiscountedPriceDraft `json:"discounted,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetDiscountedPriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetDiscountedPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDiscountedPrice", Alias: (*Alias)(&obj)})
}

type ProductSetImageLabelAction struct {
	Sku       *string `json:"sku,omitempty"`
	VariantId *int    `json:"variantId,omitempty"`
	// The URL of the image.
	ImageUrl string `json:"imageUrl"`
	// The new image label.
	// If left blank or set to null, the label is removed.
	Label  *string `json:"label,omitempty"`
	Staged *bool   `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetImageLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetImageLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setImageLabel", Alias: (*Alias)(&obj)})
}

type ProductSetKeyAction struct {
	// User-specific unique identifier for the product.
	// If left blank or set to `null`, the product key is unset/removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ProductSetMetaDescriptionAction struct {
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	Staged          *bool            `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetMetaDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaDescription", Alias: (*Alias)(&obj)})
}

type ProductSetMetaKeywordsAction struct {
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	Staged       *bool            `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetMetaKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaKeywords", Alias: (*Alias)(&obj)})
}

type ProductSetMetaTitleAction struct {
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	Staged    *bool            `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}

type ProductSetPricesAction struct {
	VariantId *int         `json:"variantId,omitempty"`
	Sku       *string      `json:"sku,omitempty"`
	Prices    []PriceDraft `json:"prices"`
	Staged    *bool        `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetPricesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetPricesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPrices", Alias: (*Alias)(&obj)})
}

type ProductSetProductPriceCustomFieldAction struct {
	PriceId string      `json:"priceId"`
	Staged  *bool       `json:"staged,omitempty"`
	Name    string      `json:"name"`
	Value   interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetProductPriceCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductPriceCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductPriceCustomField", Alias: (*Alias)(&obj)})
}

type ProductSetProductPriceCustomTypeAction struct {
	PriceId string                  `json:"priceId"`
	Staged  *bool                   `json:"staged,omitempty"`
	Type    *TypeResourceIdentifier `json:"type,omitempty"`
	Fields  *FieldContainer         `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetProductPriceCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductPriceCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductPriceCustomType", Alias: (*Alias)(&obj)})
}

type ProductSetProductVariantKeyAction struct {
	VariantId *int    `json:"variantId,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	// If left blank or set to `null`, the key is unset/removed.
	Key    *string `json:"key,omitempty"`
	Staged *bool   `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetProductVariantKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductVariantKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductVariantKey", Alias: (*Alias)(&obj)})
}

type ProductSetSearchKeywordsAction struct {
	SearchKeywords SearchKeywords `json:"searchKeywords"`
	Staged         *bool          `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetSearchKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetSearchKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSearchKeywords", Alias: (*Alias)(&obj)})
}

type ProductSetSkuAction struct {
	VariantId int `json:"variantId"`
	// SKU must be unique.
	// If left blank or set to `null`, the sku is unset/removed.
	Sku    *string `json:"sku,omitempty"`
	Staged *bool   `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetSkuAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetSkuAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSku", Alias: (*Alias)(&obj)})
}

type ProductSetTaxCategoryAction struct {
	// If left blank or set to `null`, the tax category is unset/removed.
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductSetTaxCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetTaxCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTaxCategory", Alias: (*Alias)(&obj)})
}

type ProductTransitionStateAction struct {
	State *StateResourceIdentifier `json:"state,omitempty"`
	Force *bool                    `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

type ProductUnpublishAction struct {
}

// MarshalJSON override to set the discriminator value
func (obj ProductUnpublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "unpublish", Alias: (*Alias)(&obj)})
}