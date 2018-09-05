package xendit

type VirtualAccount struct {
	client *Client
}

type AvailableVirtualAccountResponse struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func (c *VirtualAccount) GetAvailableVirtualAccount() ([]AvailableVirtualAccountResponse, error) {
	res := new([]AvailableVirtualAccountResponse)

	endpoint := "available_virtual_account_banks"

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}
