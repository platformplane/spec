package custom

type ArtifactPublished struct {
	CorrelationID *string `json:"correlation_id"`

	PURL string `json:"purl"`
}
