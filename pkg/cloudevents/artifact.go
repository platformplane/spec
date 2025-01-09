package cloudevents

const (
	ArtifactSyncFailedType = "io.platformplane.artifact.sync.failed"
)

type ArtifactSyncFailed struct {
	Purl string `json:"purl"`

	ErrorMessage string `json:"error_message"`

	LastAttempt bool `json:"last_attempt"`

	CorrelationID *string `json:"correlation_id"`
}
