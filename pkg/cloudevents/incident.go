package cloudevents

const (
	IncidentCreatedType = "platform.incident.created"
)

type IncidentCreated struct {
	Severity string `json:"severity"`

	Service     string `json:"service"`
	Environment string `json:"environment"`

	Description string `json:"description"`

	CustomDataContent     []byte `json:"custom_data_content"`
	CustomDataContentType string `json:"custom_data_content_type"`
}
