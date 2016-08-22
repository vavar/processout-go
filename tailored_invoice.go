package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TailoredInvoices manages the TailoredInvoice struct
type TailoredInvoices struct {
	p *ProcessOut
}

type TailoredInvoice struct {
	// ID : ID of the tailored invoice
	ID string `json:"id"`
	// Name : Name of the tailored invoice
	Name string `json:"name"`
	// Price : Price of the tailored invoice
	Price string `json:"price"`
	// Currency : Currency of the tailored invoice
	Currency string `json:"currency"`
	// Taxes : Taxes applied on the tailored invoice (on top of the price)
	Taxes string `json:"taxes"`
	// Shipping : Shipping fees applied on the tailored invoice (on top of the price)
	Shipping string `json:"shipping"`
	// RequestEmail : Choose whether or not to request the email during the checkout process
	RequestEmail bool `json:"request_email"`
	// RequestShipping : Choose whether or not to request the shipping address during the checkout process
	RequestShipping bool `json:"request_shipping"`
	// ReturnURL : URL where the customer will be redirected upon payment
	ReturnURL string `json:"return_url"`
	// CancelURL : URL where the customer will be redirected if the paymen was canceled
	CancelURL string `json:"cancel_url"`
	// Custom : Custom variable passed along in the events/webhooks
	Custom string `json:"custom"`
	// CreatedAt : Date at which the tailored invoice was created
	CreatedAt time.Time `json:"created_at"`
}


// Invoice : Create a new invoice from the tailored invoice.
func (s TailoredInvoices) Invoice(tailoredInvoice *TailoredInvoice) (*Invoice, error) {

	type Response struct {
		Invoice `json:"invoice"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 _ , err := json.Marshal(map[string]interface{}{

	})
	if err != nil {
		return nil, err
	}

	path := "/tailored-invoices/"+url.QueryEscape(tailoredInvoice.ID)+"/invoices"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.Invoice, nil
}

// All : Get all the tailored invoices.
func (s TailoredInvoices) All() ([]*TailoredInvoice, error) {

	type Response struct {
		TailoredInvoices []*TailoredInvoice `json:"tailored_invoices"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 _ , err := json.Marshal(map[string]interface{}{

	})
	if err != nil {
		return nil, err
	}

	path := "/tailored-invoices"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return payload.TailoredInvoices, nil
}

// Create : Create a new tailored invoice.
func (s TailoredInvoices) Create(tailoredInvoice *TailoredInvoice) (*TailoredInvoice, error) {

	type Response struct {
		TailoredInvoice `json:"tailored_invoice"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 body , err := json.Marshal(map[string]interface{}{
		"name": tailoredInvoice.Name,
		"price": tailoredInvoice.Price,
		"taxes": tailoredInvoice.Taxes,
		"shipping": tailoredInvoice.Shipping,
		"currency": tailoredInvoice.Currency,
		"request_email": tailoredInvoice.RequestEmail,
		"request_shipping": tailoredInvoice.RequestShipping,
		"return_url": tailoredInvoice.ReturnURL,
		"cancel_url": tailoredInvoice.CancelURL,
		"custom": tailoredInvoice.Custom,

	})
	if err != nil {
		return nil, err
	}

	path := "/tailored-invoices"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.TailoredInvoice, nil
}

// Find : Find a tailored invoice by its ID.
func (s TailoredInvoices) Find(tailoredInvoiceID string) (*TailoredInvoice, error) {

	type Response struct {
		TailoredInvoice `json:"tailored_invoice"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 _ , err := json.Marshal(map[string]interface{}{

	})
	if err != nil {
		return nil, err
	}

	path := "/tailored-invoices/"+url.QueryEscape(tailoredInvoiceID)+""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.TailoredInvoice, nil
}

// Save : Save the updated tailored invoice attributes.
func (s TailoredInvoices) Save(tailoredInvoice *TailoredInvoice) (*TailoredInvoice, error) {

	type Response struct {
		TailoredInvoice `json:"tailored_invoice"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 _ , err := json.Marshal(map[string]interface{}{

	})
	if err != nil {
		return nil, err
	}

	path := "/tailored-invoices/"+url.QueryEscape(tailoredInvoice.ID)+""

	req, err := http.NewRequest(
		"PUT",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.TailoredInvoice, nil
}

// Delete : Delete the tailored invoice.
func (s TailoredInvoices) Delete(tailoredInvoice *TailoredInvoice) error {

	type Response struct {Success bool `json:"success"`
		Message string `json:"message"`
	}

	 _ , err := json.Marshal(map[string]interface{}{

	})
	if err != nil {
		return err
	}

	path := "/tailored-invoices/"+url.QueryEscape(tailoredInvoice.ID)+""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return err
	}

	if !payload.Success {
		return errors.New(payload.Message)
	}
	return nil
}


// dummyTailoredInvoice is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyTailoredInvoice() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
		f url.URL
	}
	errors.New("")
}