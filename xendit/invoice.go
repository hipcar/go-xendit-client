package xendit

import "fmt"

type Invoice struct {
	client *Client
}

type CreateInvoiceRequest struct {
	ExternalId               string  `json:"external_id"`
	PayerEmail               string  `json:"payer_email"`
	Description              string  `json:"description"`
	Amount                   float64 `json:"amount"`
	ShouldSendEmail          bool    `json:"should_send_email"`
	CallbackVirtualAccountId string  `json:"callback_virtual_account_id"`
}

type InvoiceResponse struct {
	Id                        string                                 `json:"id"`
	UserId                    string                                 `json:"user_id"`
	ExternalId                string                                 `json:"external_id"`
	Status                    string                                 `json:"status"`
	MerchantName              string                                 `json:"merchant_name"`
	MerchantProfilePictureUrl string                                 `json:"merchant_profile_picture_url"`
	Amount                    float64                                `json:"amount"`
	PayerEmail                string                                 `json:"payer_email"`
	Description               string                                 `json:"description"`
	InvoiceUrl                string                                 `json:"invoice_url"`
	ExpiryDate                string                                 `json:"expiry_date"`
	ShouldExcludeCreditCard   bool                                   `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool                                   `json:"should_send_email"`
	Created                   string                                 `json:"created"`
	Updated                   string                                 `json:"updated"`
	AvailableBanks            []InvoiceAvailableBankResponse         `json:"available_banks"`
	AvailableRetailOutlets    []InvoiceAvailableRetailOutletResponse `json:"available_retail_outlets"`
}

type InvoiceAvailableBankResponse struct {
	BankCode          string  `json:"bank_code"`
	CollectionType    string  `json:"collection_type"`
	BankAccountNumber string  `json:"bank_account_number"`
	TransferAmount    float64 `json:"transfer_amount"`
	BankBranch        string  `json:"bank_branch"`
	AccountHolderName string  `json:"account_holder_name"`
	IdentityAmount    float64 `json:"identity_amount"`
}

type InvoiceAvailableRetailOutletResponse struct {
	RetailOutletName string  `json:"retail_outlet_name"`
	PaymentCode      string  `json:"payment_code"`
	TransferAmount   float64 `json:"transfer_amount"`
}

type InvoiceCallbackRequest struct {
	Id                     string  `json:"id"`
	UserId                 string  `json:"user_id"`
	ExternalId             string  `json:"external_id"`
	IsHigh                 bool    `json:"is_high"`
	MerchantName           string  `json:"merchant_name"`
	Amount                 float64 `json:"amount"`
	FeesPaidAmount         float64 `json:"fees_paid_amount"`
	Status                 string  `json:"status"`
	PayerEmail             string  `json:"payer_email"`
	Description            string  `json:"description"`
	AdjustedReceivedAmount float64 `json:"adjusted_received_amount"`
	PaymentMethod          string  `json:"payment_method"`
	BankCode               string  `json:"bank_code"`
	RetailOutletName       string  `json:"retail_outlet_name"`
	PaidAmount             string  `json:"paid_amount"`
	Updated                string  `json:"updated"`
	Created                string  `json:"created"`
}

func (c *Invoice) CreateInvoice(body CreateInvoiceRequest) (InvoiceResponse, error) {
	res := new(InvoiceResponse)

	endpoint := "v2/invoices"

	err := c.client.Request("POST", endpoint, nil, res)
	return *res, err
}

func (c *Invoice) GetInvoice(invoiceId string) (InvoiceResponse, error) {
	res := new(InvoiceResponse)

	endpoint := fmt.Sprintf("v2/invoices/%s", invoiceId)

	err := c.client.Request("GET", endpoint, nil, res)
	return *res, err
}

func (c *Invoice) ExpireInvoice(invoiceId string) (InvoiceResponse, error) {
	res := new(InvoiceResponse)

	endpoint := fmt.Sprintf("invoices/%s/expire!", invoiceId)

	err := c.client.Request("POST", endpoint, nil, res)
	return *res, err
}
