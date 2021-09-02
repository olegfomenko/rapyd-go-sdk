package resources

type WebhookType string

const (
	PaymentCompletedWebhook WebhookType = "PAYMENT_COMPLETED"
	PaymentFailedWebhook    WebhookType = "PAYMENT_FAILED"
)

type Webhook struct {
	Id   string      `json:"id"`
	Type WebhookType `json:"type"`
	Data Data        `json:"data"`
}
