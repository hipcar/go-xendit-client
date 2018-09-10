package xendit

type Balance struct {
	client *Client
}

const (
	AccountTypeCash    = "CASH"
	AccountTypeHolding = "HOLDING"
	AccountTypeTax     = "TAX"
)

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

func (c *Balance) GetBalance(accountType string) (BalanceResponse, error) {
	res := new(BalanceResponse)

	switch accountType {
	case AccountTypeCash:
	case AccountTypeHolding:
	case AccountTypeTax:
		break
	default:
		accountType = AccountTypeCash
		break
	}

	endpoint := "balance?account_type=" + accountType

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}
