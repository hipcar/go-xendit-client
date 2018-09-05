package xendit

import "fmt"

type Disbursement struct {
	client *Client
}

type GetAvailableDisbursementBankResponse struct {
	Name            string `json:"name"`
	Code            string `json:"code"`
	CanDisburse     bool   `json:"can_disburse"`
	CanNameValidate bool   `json:"can_name_validate"`
}

type CreateDisbursementRequest struct {
	ExternalId        string  `json:"external_id"`
	BankCode          string  `json:"bank_code"`
	AccountHolderName string  `json:"account_holder_name"`
	AccountNumber     string  `json:"account_number"`
	Description       string  `json:"description"`
	Amount            float64 `json:"amount"`
}

type CreateDisbursementResponse struct {
	UserId                  string  `json:"user_id"`
	ExternalId              string  `json:"external_id"`
	Amount                  float64 `json:"amount"`
	BankCode                string  `json:"bank_code"`
	AccountHolderName       string  `json:"account_holder_name"`
	DisbursementDescription string  `json:"disbursement_description"`
	Status                  string  `json:"status"`
	Id                      string  `json:"id"`
}

type GetDisbursementResponse struct {
	UserId                  string  `json:"user_id"`
	ExternalId              string  `json:"external_id"`
	Amount                  float64 `json:"amount"`
	BankCode                string  `json:"bank_code"`
	AccountHolderName       string  `json:"account_holder_name"`
	DisbursementDescription string  `json:"disbursement_description"`
	Status                  string  `json:"status"`
	Id                      string  `json:"id"`
}

type DisbursementCallbackRequest struct {
	IsInstant               bool    `json:"is_instant"`
	UserId                  string  `json:"user_id"`
	ExternalId              string  `json:"external_id"`
	Amount                  float64 `json:"amount"`
	BankCode                string  `json:"bank_code"`
	AccountHolderName       string  `json:"account_holder_name"`
	DisbursementDescription string  `json:"disbursement_description"`
	Status                  string  `json:"status"`
	FailureCode             string  `json:"failure_code"`
	Id                      string  `json:"id"`
}

func (c *Disbursement) GetAvailableDisbursementBank() ([]GetAvailableDisbursementBankResponse, error) {
	res := new([]GetAvailableDisbursementBankResponse)

	endpoint := "available_disbursements_banks"

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}

func (c *Disbursement) CreateDisbursement(body CreateDisbursementRequest) (CreateDisbursementResponse, error) {
	res := new(CreateDisbursementResponse)

	endpoint := "disbursements"

	err := c.client.Request("POST", endpoint, body, res)
	return *res, err
}

func (c *Disbursement) GetDisbursementById(disbursementId string) (GetDisbursementResponse, error) {
	res := new(GetDisbursementResponse)

	endpoint := fmt.Sprintf("disbursements/%s", disbursementId)

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}

func (c *Disbursement) GetDisbursementByExternalId(externalId string) ([]GetDisbursementResponse, error) {
	res := new([]GetDisbursementResponse)

	endpoint := fmt.Sprintf("disbursements?external_id=%s", externalId)

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}
