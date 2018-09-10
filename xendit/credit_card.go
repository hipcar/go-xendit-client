package xendit

import (
	"fmt"
	"time"
)

type CreditCard struct {
	client *Client
}

type TokenizationRequest struct {
	Amount             string `json:"amount"`
	CardNumber         string `json:"card_number"`
	CardExpMonth       string `json:"card_exp_month"`
	CardExpYear        string `json:"card_exp_year"`
	CardCvn            string `json:"card_cvn"`
	IsMultipleUse      bool   `json:"is_multiple_use"`
	ShouldAuthenticate bool   `json:"should_authenticate"`
}

type TokenizationResponse struct {
	Id                     string `json:"id"`
	AuthenticationId       string `json:"authentication_id"`
	MaskedCardNumber       string `json:"masked_card_number"`
	Status                 string `json:"status"`
	PayerAuthenticationUrl string `json:"payer_authentication_url"`
	FailureReason          string `json:"failure_reason"`
}

type ChargeRequest struct {
	TokenId          string  `json:"token_id"`
	ExternalId       string  `json:"external_id"`
	Amount           float64 `json:"amount"`
	AuthenticationId string  `json:"authentication_id"`
	CardCvn          string  `json:"card_cvn"`
	Capture          bool    `json:"capture"`
	Descriptor       string  `json:"descriptor"`
}

type ChargeResponse struct {
	Created               time.Time `json:"created"`
	Status                string    `json:"status"`
	BusinessId            string    `json:"business_id"`
	AuthorizedAmount      float64   `json:"authorized_amount"`
	ExternalId            string    `json:"external_id"`
	MerchantId            string    `json:"merchant_id"`
	MerchantReferenceCode string    `json:"merchant_reference_code"`
	CardType              string    `json:"card_type"`
	MaskedCardNumber      string    `json:"masked_card_number"`
	ChargeType            string    `json:"charge_type"`
	CardBrand             string    `json:"card_brand"`
	BankReconciliationId  string    `json:"bank_reconciliation_id"`
	ECI                   string    `json:"eci"`
	CaptureAmount         float64   `json:"capture_amount"`
	Id                    string    `json:"id"`
	Descriptor            string    `json:"descriptor"`
	FailureReason         string    `json:"failure_reason"`
}

type CaptureChargeRequest struct {
	Amount string `json:"amount"`
}

type CaptureChargeResponse struct {
	Created               time.Time `json:"created"`
	Status                string    `json:"status"`
	BusinessId            string    `json:"business_id"`
	AuthorizedAmount      float64   `json:"authorized_amount"`
	ExternalId            string    `json:"external_id"`
	MerchantId            string    `json:"merchant_id"`
	MerchantReferenceCode string    `json:"merchant_reference_code"`
	CardType              string    `json:"card_type"`
	MaskedCardNumber      string    `json:"masked_card_number"`
	ChargeType            string    `json:"charge_type"`
	CardBrand             string    `json:"card_brand"`
	BankReconciliationId  string    `json:"bank_reconciliation_id"`
	CaptureAmount         float64   `json:"capture_amount"`
	Id                    string    `json:"id"`
	Descriptor            string    `json:"descriptor"`
}

// TODO
func (c *CreditCard) CreateToken(tokenRequest TokenizationResponse) (TokenizationResponse, error) {
	res := new(TokenizationResponse)

	endpoint := "token"

	err := c.client.Request("GET", endpoint, tokenRequest, res)
	return *res, err
}

func (c *CreditCard) CreateCharge(chargeRequest ChargeRequest) (ChargeResponse, error) {
	res := new(ChargeResponse)

	endpoint := "credit_card_charges"

	err := c.client.Request("POST", endpoint, chargeRequest, res)
	return *res, err
}

func (c *CreditCard) CaptureCharge(creditCardChargeId string) (ChargeResponse, error) {
	res := new(ChargeResponse)

	endpoint := fmt.Sprintf("credit_card_charges/%s/capture", creditCardChargeId)

	err := c.client.Request("POST", endpoint, nil, res)
	return *res, err
}

func (c *CreditCard) GetCharge(creditCardChargeId string) (ChargeResponse, error) {
	res := new(ChargeResponse)

	endpoint := fmt.Sprintf("credit_card_charges/%s", creditCardChargeId)

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}
