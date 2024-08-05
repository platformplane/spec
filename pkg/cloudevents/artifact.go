package cloudevents

const (
	ArtifactSyncFailedType = "io.platformplane.artifact.sync.failed"
)

type ArtifactSyncFailed struct {
	Purl string `json:"purl"`

	Message string `json:"error_message"`

	LastAttempt bool `json:"lastAttempt"`

	CorrelationID *string `json:"correlation_id"`
}
