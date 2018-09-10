package xendit

import "time"

type NameValidator struct {
	client *Client
}

type NameValidatorRequest struct {
	BankAccountNumber string `json:"bank_account_number,omitempty"`
	BankCode          string `json:"bank_code,omitempty"`
}

type NameValidatorResponse struct {
	BankAccountNumber     string    `json:"bank_account_number"`
	BankAccountHolderName string    `json:"bank_account_holder_name"`
	BankCode              string    `json:"bank_code"`
	Reference             string    `json:"reference"`
	Status                string    `json:"status"`
	Updated               time.Time `json:"updated"`
	Id                    string    `json:"id"`
}

func (c *NameValidator) Validate(body NameValidatorRequest) (NameValidatorResponse, error) {
	res := new(NameValidatorResponse)

	endpoint := "bank_account_data_requests"

	err := c.client.Request("POST", endpoint, nil, res)
	return *res, err
}
