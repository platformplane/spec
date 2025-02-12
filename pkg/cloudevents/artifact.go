package cloudevents

const (
	ArtifactQueuingFailedType = "io.platformplane.artifact.queuing.failed"
	ArtifactSyncFailedType    = "io.platformplane.artifact.sync.failed"
)

type ArtifactQueuingFailed struct {
	Purl string `json:"purl"`

	ErrorMessage string `json:"error_message"`

	LastAttempt bool `json:"last_attempt"`

	CorrelationID *string `json:"correlation_id"`
}

type ArtifactSyncFailed struct {
	Purl string `json:"purl"`

	ErrorMessage string `json:"error_message"`

	LastAttempt bool `json:"last_attempt"`

	CorrelationID *string `json:"correlation_id"`
}
