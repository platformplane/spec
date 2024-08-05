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

type OCIArtifact struct {
	Registry   string
	Repository string
	Image      string

	Version string
	Digest  string
}

type HelmArtifact struct {
	RepoURL string
	Chart   string
	Version string
}

type GenericArtifact struct {
	Infos string
}
