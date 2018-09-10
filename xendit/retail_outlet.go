package xendit

import (
	"fmt"
	"time"
)

type RetailOutlet struct {
	client *Client
}

type CreateFixedPaymentCodeRequest struct {
	ExternalId       string  `json:"external_id"`
	RetailOutletName string  `json:"retail_outlet_name"`
	Name             string  `json:"name"`
	ExpectedAmount   float64 `json:"expected_amount"`
	PaymentCode      string  `json:"payment_code"`
	ExpirationDate   string  `json:"expiration_date"`
	IsSingleUse      bool    `json:"is_single_use"`
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
	Name           string    `json:"name"`
	ExpectedAmount float64   `json:"expected_amount"`
	ExpirationDate time.Time `json:"expiration_date"`
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
	FixedPaymentCodePaymentId string    `json:"fixed_payment_code_payment_id"`
	OwnerId                   string    `json:"owner_id"`
	FixedPaymentCodeId        string    `json:"fixed_payment_code_id"`
	PaymentId                 string    `json:"payment_id"`
	ExternalId                string    `json:"external_id"`
	PaymentCode               string    `json:"payment_code"`
	Prefix                    string    `json:"prefix"`
	RetailOutletName          string    `json:"retail_outlet_name"`
	Amount                    float64   `json:"amount"`
	Name                      string    `json:"name"`
	TransactionTimestamp      time.Time `json:"transaction_timestamp"`
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
