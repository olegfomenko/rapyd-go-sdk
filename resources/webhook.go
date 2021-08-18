package resources

import (
	"encoding/json"
	"net/http"
)

type WebhookType string

const (
	PaymentCompletedWebhook = "PAYMENT_COMPLETED"
)

type WebhookBody struct {
	Id   string      `json:"id"`
	Type WebhookType `json:"type"`
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}

type Webhook struct {
	Headers http.Header     `json:"headers"`
	Body    json.RawMessage `json:"body"`
}

func (w *Webhook) UnmarshallBody() (WebhookBody, error) {
	var body WebhookBody
	return body, json.Unmarshal(w.Body, &body)
}
