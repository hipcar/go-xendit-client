package xendit

import "time"

type NPWPValidator struct {
	client *Client
}

type NPWPValidatorRequest struct {
	AccountNumber string `json:"account_number"`
}

type NPWPValidatorResponse struct {
	Id            string    `json:"id"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
	AccountNumber string    `json:"account_number"`
	AccountName   string    `json:"account_name"`
	Status        string    `json:"status"`
	FailureReason string    `json:"failure_reason"`
}

func (c *NPWPValidator) Validate(body NPWPValidatorRequest) (NPWPValidatorResponse, error) {
	res := new(NPWPValidatorResponse)

	endpoint := "npwp_data_requests"

	err := c.client.Request("POST", endpoint, nil, res)
	return *res, err
}
