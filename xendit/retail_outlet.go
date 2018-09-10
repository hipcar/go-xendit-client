package xendit

import (
	"fmt"
	"time"
)

type RetailOutlet struct {
	client *Client
}

type CreateFixedPaymentCodeRequest struct {
	ExternalId       string  `json:"external_id,omitempty"`
	RetailOutletName string  `json:"retail_outlet_name,omitempty"`
	Name             string  `json:"name,omitempty"`
	ExpectedAmount   float64 `json:"expected_amount,omitempty"`
	PaymentCode      string  `json:"payment_code,omitempty"`
	ExpirationDate   string  `json:"expiration_date,omitempty"`
	IsSingleUse      bool    `json:"is_single_use,omitempty"`
}

type CreateFixedPaymentCodeResponse struct {
	OwnerId          string    `json:"owner_id"`
	ExternalId       string    `json:"external_id"`
	RetailOutletName string    `json:"retail_outlet_name"`
	Prefix           string    `json:"prefix"`
	Name             string    `json:"name"`
	PaymentCode      string    `json:"payment_code"`
	ExpectedAmount   float64   `json:"expected_amount"`
	IsSingleUse      bool      `json:"is_single_use"`
	ExpirationDate   time.Time `json:"expiration_date"`
	Id               string    `json:"id"`
}

type UpdateFixedPaymentCodeRequest struct {
	Name           string    `json:"name,omitempty"`
	ExpectedAmount float64   `json:"expected_amount,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
}

type UpdateFixedPaymentCodeResponse struct {
	OwnerId          string    `json:"owner_id"`
	ExternalId       string    `json:"external_id"`
	RetailOutletName string    `json:"retail_outlet_name"`
	Prefix           string    `json:"prefix"`
	Name             string    `json:"name"`
	PaymentCode      string    `json:"payment_code"`
	ExpectedAmount   float64   `json:"expected_amount"`
	IsSingleUse      bool      `json:"is_single_use"`
	ExpirationDate   time.Time `json:"expiration_date"`
	Id               string    `json:"id"`
}

type GetFixedPaymentCodeResponse struct {
	OwnerId          string    `json:"owner_id"`
	ExternalId       string    `json:"external_id"`
	RetailOutletName string    `json:"retail_outlet_name"`
	Prefix           string    `json:"prefix"`
	Name             string    `json:"name"`
	PaymentCode      string    `json:"payment_code"`
	ExpectedAmount   float64   `json:"expected_amount"`
	IsSingleUse      bool      `json:"is_single_use"`
	ExpirationDate   time.Time `json:"expiration_date"`
	Id               string    `json:"id"`
}

type FixedPaymentCodeCallbackRequest struct {
	FixedPaymentCodePaymentId string    `json:"fixed_payment_code_payment_id,omitempty"`
	OwnerId                   string    `json:"owner_id,omitempty"`
	FixedPaymentCodeId        string    `json:"fixed_payment_code_id,omitempty"`
	PaymentId                 string    `json:"payment_id,omitempty"`
	ExternalId                string    `json:"external_id,omitempty"`
	PaymentCode               string    `json:"payment_code,omitempty"`
	Prefix                    string    `json:"prefix,omitempty"`
	RetailOutletName          string    `json:"retail_outlet_name,omitempty"`
	Amount                    float64   `json:"amount,omitempty"`
	Name                      string    `json:"name,omitempty"`
	TransactionTimestamp      time.Time `json:"transaction_timestamp,omitempty"`
}

func (c *RetailOutlet) CreateFixedPaymentCode(body CreateFixedPaymentCodeRequest) (CreateFixedPaymentCodeResponse, error) {
	res := new(CreateFixedPaymentCodeResponse)

	endpoint := "fixed_payment_code"

	err := c.client.Request("POST", endpoint, body, res)
	return *res, err
}

func (c *RetailOutlet) UpdateFixedPaymentCode(fixedPaymentCodeId string, body UpdateFixedPaymentCodeRequest) (UpdateFixedPaymentCodeResponse, error) {
	res := new(UpdateFixedPaymentCodeResponse)

	endpoint := fmt.Sprintf("fixed_payment_code/%s", fixedPaymentCodeId)

	err := c.client.Request("PATCH", endpoint, body, res)
	return *res, err
}

func (c *RetailOutlet) GetFixedPaymentCode(fixedPaymentCodeId string) (GetFixedPaymentCodeResponse, error) {
	res := new(GetFixedPaymentCodeResponse)

	endpoint := fmt.Sprintf("fixed_payment_code/%s", fixedPaymentCodeId)

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}
