package xendit

import (
	"errors"
	"net/url"
	"net/http"
	"io"
	"bytes"
	"encoding/json"
	"fmt"
	"context"
)

const (
	libraryVersion = "0.1"
	userAgent      = "xendit-client/" + libraryVersion
	xenditURL      = "https://api.xendit.co/"
)

var (
	// ErrUnauthorized can be returned on any call on response status code 401.
	ErrUnauthorized = errors.New("xendit-client: unauthorized")
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	doer      Doer
	BaseURL   *url.URL
	UserAgent string
	EnableLog bool
	PublicKey string
	SecretKey string

	Balance        *Balance
	CreditCard     *CreditCard
	Disbursement   *Disbursement
	Forex          *Forex
	Invoice        *Invoice
	NameValidator  *NameValidator
	NPWPValidator  *NPWPValidator
	RetailOutlet   *RetailOutlet
	VirtualAccount *VirtualAccount

	httpClient *http.Client
}

type errorResponse struct {
	Error string `json:"error,omitempty"`
}

type DoerFunc func(req *http.Request) (resp *http.Response, err error)

// NewClient created new moc client with doer.
// If doer is nil then http.DefaultClient used instead.
func NewClient() *Client {

	baseUrl := xenditURL

	baseURL, _ := url.Parse(baseUrl)
	client := &Client{
		doer:      http.DefaultClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}

	client.Balance = &Balance{client}
	client.CreditCard = &CreditCard{client}
	client.Disbursement = &Disbursement{client}
	client.Forex = &Forex{client}
	client.Invoice = &Invoice{client}
	client.NameValidator = &NameValidator{client}
	client.NPWPValidator = &NPWPValidator{client}
	client.RetailOutlet = &RetailOutlet{client}
	client.VirtualAccount = &VirtualAccount{client}

	return client
}

func (c *Client) Request(method string, path string, data interface{}, v interface{}) error {

	urlStr := path

	rel, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	u := c.BaseURL.ResolveReference(rel)
	var body io.Reader

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)

		if c.EnableLog {
			fmt.Printf("Request %s to %s with data: %s \n", method, u.String(), string(b))
		}
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.SetBasicAuth(c.SecretKey, "")

	resp, err := c.doer.Do(req.WithContext(context.Background()))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	// Return error from xendit API
	if resp.StatusCode != http.StatusOK {
		rb := new(errorResponse)

		err = json.NewDecoder(resp.Body).Decode(rb)

		if rb.Error == "" {
			return errors.New("general error")
		}

		return errors.New(rb.Error)
	}

	// Decode to interface
	res := v
	err = json.NewDecoder(resp.Body).Decode(res)

	by, _ := json.Marshal(res)
	if c.EnableLog {
		fmt.Printf("Response %s from %s : %s \n", method, u.String(), string(by))
	}

	return err
}
