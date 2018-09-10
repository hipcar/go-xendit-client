package xendit

import (
	"fmt"
	"time"
)

type Invoice struct {
	client *Client
}

type CreateInvoiceRequest struct {
	ExternalId               string  `json:"external_id,omitempty"`
	PayerEmail               string  `json:"payer_email,omitempty"`
	Description              string  `json:"description,omitempty"`
	Amount                   float64 `json:"amount,omitempty"`
	ShouldSendEmail          bool    `json:"should_send_email,omitempty"`
	CallbackVirtualAccountId string  `json:"callback_virtual_account_id,omitempty"`
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
	ExpiryDate                time.Time                              `json:"expiry_date"`
	ShouldExcludeCreditCard   bool                                   `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool                                   `json:"should_send_email"`
	Created                   time.Time                              `json:"created"`
	Updated                   time.Time                              `json:"updated"`
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
	Id                     string    `json:"id,omitempty"`
	UserId                 string    `json:"user_id,omitempty"`
	ExternalId             string    `json:"external_id,omitempty"`
	IsHigh                 bool      `json:"is_high,omitempty"`
	MerchantName           string    `json:"merchant_name,omitempty"`
	Amount                 float64   `json:"amount,omitempty"`
	FeesPaidAmount         float64   `json:"fees_paid_amount,omitempty"`
	Status                 string    `json:"status,omitempty"`
	PayerEmail             string    `json:"payer_email,omitempty"`
	Description            string    `json:"description,omitempty"`
	AdjustedReceivedAmount float64   `json:"adjusted_received_amount,omitempty"`
	PaymentMethod          string    `json:"payment_method,omitempty"`
	BankCode               string    `json:"bank_code,omitempty"`
	RetailOutletName       string    `json:"retail_outlet_name,omitempty"`
	PaidAmount             string    `json:"paid_amount,omitempty"`
	Updated                time.Time `json:"updated,omitempty"`
	Created                time.Time `json:"created,omitempty"`
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
