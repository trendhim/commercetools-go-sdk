// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Payment struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// A reference to the customer this payment belongs to.
	Customer *CustomerReference `json:"customer,omitempty"`
	// Identifies payments belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId *string `json:"anonymousId,omitempty"`
	ExternalId  *string `json:"externalId,omitempty"`
	// The identifier that is used by the interface that manages the payment (usually the PSP).
	// Cannot be changed once it has been set.
	// The combination of this ID and the PaymentMethodInfo `paymentInterface` must be unique.
	InterfaceId *string `json:"interfaceId,omitempty"`
	// How much money this payment intends to receive from the customer.
	// The value usually matches the cart or order gross total.
	AmountPlanned     TypedMoney        `json:"amountPlanned"`
	AmountAuthorized  TypedMoney        `json:"amountAuthorized,omitempty"`
	AuthorizedUntil   *string           `json:"authorizedUntil,omitempty"`
	AmountPaid        TypedMoney        `json:"amountPaid,omitempty"`
	AmountRefunded    TypedMoney        `json:"amountRefunded,omitempty"`
	PaymentMethodInfo PaymentMethodInfo `json:"paymentMethodInfo"`
	PaymentStatus     PaymentStatus     `json:"paymentStatus"`
	// A list of financial transactions of different TransactionTypes with different TransactionStates.
	Transactions []Transaction `json:"transactions"`
	// Interface interactions can be requests sent to the PSP, responses received from the PSP or notifications received from the PSP.
	// Some interactions may result in a transaction.
	// If so, the `interactionId` in the Transaction should be set to match the ID of the PSP for the interaction.
	// Interactions are managed by the PSP integration and are usually neither written nor read by the user facing frontends or other services.
	InterfaceInteractions []CustomFields `json:"interfaceInteractions"`
	Custom                *CustomFields  `json:"custom,omitempty"`
	// User-specific unique identifier for the payment (max.
	// 256 characters).
	Key *string `json:"key,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Payment) UnmarshalJSON(data []byte) error {
	type Alias Payment
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.AmountPlanned != nil {
		var err error
		obj.AmountPlanned, err = mapDiscriminatorTypedMoney(obj.AmountPlanned)
		if err != nil {
			return err
		}
	}
	if obj.AmountAuthorized != nil {
		var err error
		obj.AmountAuthorized, err = mapDiscriminatorTypedMoney(obj.AmountAuthorized)
		if err != nil {
			return err
		}
	}
	if obj.AmountPaid != nil {
		var err error
		obj.AmountPaid, err = mapDiscriminatorTypedMoney(obj.AmountPaid)
		if err != nil {
			return err
		}
	}
	if obj.AmountRefunded != nil {
		var err error
		obj.AmountRefunded, err = mapDiscriminatorTypedMoney(obj.AmountRefunded)
		if err != nil {
			return err
		}
	}
	return nil
}

type PaymentDraft struct {
	// A reference to the customer this payment belongs to.
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
	// Identifies payments belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId *string `json:"anonymousId,omitempty"`
	ExternalId  *string `json:"externalId,omitempty"`
	// The identifier that is used by the interface that manages the payment (usually the PSP).
	// Cannot be changed once it has been set.
	// The combination of this ID and the PaymentMethodInfo `paymentInterface` must be unique.
	InterfaceId *string `json:"interfaceId,omitempty"`
	// How much money this payment intends to receive from the customer.
	// The value usually matches the cart or order gross total.
	AmountPlanned     Money               `json:"amountPlanned"`
	AmountAuthorized  *Money              `json:"amountAuthorized,omitempty"`
	AuthorizedUntil   *string             `json:"authorizedUntil,omitempty"`
	AmountPaid        *Money              `json:"amountPaid,omitempty"`
	AmountRefunded    *Money              `json:"amountRefunded,omitempty"`
	PaymentMethodInfo *PaymentMethodInfo  `json:"paymentMethodInfo,omitempty"`
	PaymentStatus     *PaymentStatusDraft `json:"paymentStatus,omitempty"`
	// A list of financial transactions of different TransactionTypes with different TransactionStates.
	Transactions []TransactionDraft `json:"transactions,omitempty"`
	// Interface interactions can be requests send to the PSP, responses received from the PSP or notifications received from the PSP.
	// Some interactions may result in a transaction.
	// If so, the `interactionId` in the Transaction should be set to match the ID of the PSP for the interaction.
	// Interactions are managed by the PSP integration and are usually neither written nor read by the user facing frontends or other services.
	InterfaceInteractions []CustomFieldsDraft `json:"interfaceInteractions,omitempty"`
	Custom                *CustomFieldsDraft  `json:"custom,omitempty"`
	// User-specific unique identifier for the payment (max.
	// 256 characters).
	Key *string `json:"key,omitempty"`
}

type PaymentMethodInfo struct {
	// The interface that handles the payment (usually a PSP).
	// Cannot be changed once it has been set.
	// The combination of Payment`interfaceId` and this field must be unique.
	PaymentInterface *string `json:"paymentInterface,omitempty"`
	// The payment method that is used, e.g.
	// e.g.
	// a conventional string representing Credit Card, Cash Advance etc.
	Method *string `json:"method,omitempty"`
	// A human-readable, localized name for the payment method, e.g.
	// 'Credit Card'.
	Name *LocalizedString `json:"name,omitempty"`
}

type PaymentPagedQueryResponse struct {
	Limit   int       `json:"limit"`
	Count   int       `json:"count"`
	Total   *int      `json:"total,omitempty"`
	Offset  int       `json:"offset"`
	Results []Payment `json:"results"`
}

type PaymentReference struct {
	// Unique ID of the referenced resource.
	ID  string   `json:"id"`
	Obj *Payment `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}

type PaymentResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias PaymentResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}

type PaymentStatus struct {
	// A code describing the current status returned by the interface that processes the payment.
	InterfaceCode *string `json:"interfaceCode,omitempty"`
	// A text describing the current status returned by the interface that processes the payment.
	InterfaceText *string         `json:"interfaceText,omitempty"`
	State         *StateReference `json:"state,omitempty"`
}

type PaymentStatusDraft struct {
	InterfaceCode *string                  `json:"interfaceCode,omitempty"`
	InterfaceText *string                  `json:"interfaceText,omitempty"`
	State         *StateResourceIdentifier `json:"state,omitempty"`
}

type PaymentUpdate struct {
	Version int                   `json:"version"`
	Actions []PaymentUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentUpdate) UnmarshalJSON(data []byte) error {
	type Alias PaymentUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type PaymentUpdateAction interface{}

func mapDiscriminatorPaymentUpdateAction(input interface{}) (PaymentUpdateAction, error) {

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
	case "addInterfaceInteraction":
		obj := PaymentAddInterfaceInteractionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addTransaction":
		obj := PaymentAddTransactionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAmountPlanned":
		obj := PaymentChangeAmountPlannedAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTransactionInteractionId":
		obj := PaymentChangeTransactionInteractionIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTransactionState":
		obj := PaymentChangeTransactionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTransactionTimestamp":
		obj := PaymentChangeTransactionTimestampAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAmountPaid":
		obj := PaymentSetAmountPaidAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAmountRefunded":
		obj := PaymentSetAmountRefundedAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAnonymousId":
		obj := PaymentSetAnonymousIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAuthorization":
		obj := PaymentSetAuthorizationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := PaymentSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := PaymentSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomer":
		obj := PaymentSetCustomerAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setExternalId":
		obj := PaymentSetExternalIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setInterfaceId":
		obj := PaymentSetInterfaceIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := PaymentSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMethodInfoInterface":
		obj := PaymentSetMethodInfoInterfaceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMethodInfoMethod":
		obj := PaymentSetMethodInfoMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMethodInfoName":
		obj := PaymentSetMethodInfoNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStatusInterfaceCode":
		obj := PaymentSetStatusInterfaceCodeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStatusInterfaceText":
		obj := PaymentSetStatusInterfaceTextAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionState":
		obj := PaymentTransitionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type Transaction struct {
	// The unique ID of this object.
	ID string `json:"id"`
	// The time at which the transaction took place.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The type of this transaction.
	Type   TransactionType `json:"type"`
	Amount TypedMoney      `json:"amount"`
	// The identifier that is used by the interface that managed the transaction (usually the PSP).
	// If a matching interaction was logged in the `interfaceInteractions` array, the corresponding interaction should be findable with this ID.
	InteractionId *string `json:"interactionId,omitempty"`
	// The state of this transaction.
	State *TransactionState `json:"state,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Amount != nil {
		var err error
		obj.Amount, err = mapDiscriminatorTypedMoney(obj.Amount)
		if err != nil {
			return err
		}
	}
	return nil
}

type TransactionDraft struct {
	// The time at which the transaction took place.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The type of this transaction.
	Type   TransactionType `json:"type"`
	Amount Money           `json:"amount"`
	// The identifier that is used by the interface that managed the transaction (usually the PSP).
	// If a matching interaction was logged in the `interfaceInteractions` array, the corresponding interaction should be findable with this ID.
	InteractionId *string `json:"interactionId,omitempty"`
	// The state of this transaction.
	// If not set, defaults to `Initial`.
	State *TransactionState `json:"state,omitempty"`
}

type TransactionState string

const (
	TransactionStateInitial TransactionState = "Initial"
	TransactionStatePending TransactionState = "Pending"
	TransactionStateSuccess TransactionState = "Success"
	TransactionStateFailure TransactionState = "Failure"
)

type TransactionType string

const (
	TransactionTypeAuthorization       TransactionType = "Authorization"
	TransactionTypeCancelAuthorization TransactionType = "CancelAuthorization"
	TransactionTypeCharge              TransactionType = "Charge"
	TransactionTypeRefund              TransactionType = "Refund"
	TransactionTypeChargeback          TransactionType = "Chargeback"
)

type PaymentAddInterfaceInteractionAction struct {
	Type   TypeResourceIdentifier `json:"type"`
	Fields *FieldContainer        `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentAddInterfaceInteractionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddInterfaceInteractionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addInterfaceInteraction", Alias: (*Alias)(&obj)})
}

type PaymentAddTransactionAction struct {
	Transaction TransactionDraft `json:"transaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentAddTransactionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddTransactionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTransaction", Alias: (*Alias)(&obj)})
}

type PaymentChangeAmountPlannedAction struct {
	Amount Money `json:"amount"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeAmountPlannedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeAmountPlannedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAmountPlanned", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionInteractionIdAction struct {
	TransactionId string `json:"transactionId"`
	InteractionId string `json:"interactionId"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeTransactionInteractionIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionInteractionIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionInteractionId", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionStateAction struct {
	TransactionId string           `json:"transactionId"`
	State         TransactionState `json:"state"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeTransactionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionState", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionTimestampAction struct {
	TransactionId string    `json:"transactionId"`
	Timestamp     time.Time `json:"timestamp"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeTransactionTimestampAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionTimestampAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionTimestamp", Alias: (*Alias)(&obj)})
}

type PaymentSetAmountPaidAction struct {
	Amount *Money `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAmountPaidAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountPaidAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountPaid", Alias: (*Alias)(&obj)})
}

type PaymentSetAmountRefundedAction struct {
	Amount *Money `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAmountRefundedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountRefundedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountRefunded", Alias: (*Alias)(&obj)})
}

type PaymentSetAnonymousIdAction struct {
	// Anonymous ID of the anonymous customer that this payment belongs to.
	// If this field is not set any existing `anonymousId` is removed.
	AnonymousId *string `json:"anonymousId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAnonymousIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAnonymousIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

type PaymentSetAuthorizationAction struct {
	Amount *Money     `json:"amount,omitempty"`
	Until  *time.Time `json:"until,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAuthorizationAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAuthorizationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthorization", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomTypeAction struct {
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the custom fields to this value.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomerAction struct {
	// A reference to the customer this payment belongs to.
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

type PaymentSetExternalIdAction struct {
	ExternalId *string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetExternalIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetExternalIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

type PaymentSetInterfaceIdAction struct {
	InterfaceId string `json:"interfaceId"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetInterfaceIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetInterfaceIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setInterfaceId", Alias: (*Alias)(&obj)})
}

type PaymentSetKeyAction struct {
	// User-specific unique identifier for the payment (max.
	// 256 characters).
	// If not provided an existing key will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoInterfaceAction struct {
	Interface string `json:"interface"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoInterfaceAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoInterfaceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoInterface", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoMethodAction struct {
	// If not provided, the method is unset.
	Method *string `json:"method,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoMethodAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoMethod", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoNameAction struct {
	// If not provided, the name is unset.
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoNameAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoName", Alias: (*Alias)(&obj)})
}

type PaymentSetStatusInterfaceCodeAction struct {
	InterfaceCode *string `json:"interfaceCode,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetStatusInterfaceCodeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetStatusInterfaceCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatusInterfaceCode", Alias: (*Alias)(&obj)})
}

type PaymentSetStatusInterfaceTextAction struct {
	InterfaceText string `json:"interfaceText"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetStatusInterfaceTextAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetStatusInterfaceTextAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatusInterfaceText", Alias: (*Alias)(&obj)})
}

type PaymentTransitionStateAction struct {
	State StateResourceIdentifier `json:"state"`
	Force *bool                   `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}