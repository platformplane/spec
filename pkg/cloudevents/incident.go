package cloudevents

const (
	IncidentCreatedType = "platform.incident.created"
)

type IncidentCreated struct {
	Severity string `json:"severity"`

	Service     string `json:"service"`
	Environment string `json:"environment"`

	Description string `json:"description"`

	CustomDataContent     []byte `json:"customDataContent"`
	CustomDataContentType string `json:"customDataContentType"`
}
