package cloudevents

import "time"

const (
	ReleaseApprovedType = "io.platformplane.release.approved"

	ReleaseCreatedType          = "io.platformplane.release.created"
	ReleaseCandidateCreatedType = "io.platformplane.releasecandidate.created"

	ReleaseSyncedType = "io.platformplane.release.synced"
)

type ReleaseApproved struct {
	ID string `json:"id"`

	Name      string `json:"name"`
	Namespace string `json:"namespace"`

	Version string `json:"version"`

	DeploymentUnits []DeploymentUnit `json:"deployment_units"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ReleaseCreated struct {
	ID string `json:"id"`

	Name      string `json:"name"`
	Namespace string `json:"namespace"`

	Version string `json:"version"`

	DeploymentUnits []DeploymentUnit `json:"deployment_units"`
}

type ReleaseCandidateCreated struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Namespace *Namespace `json:"namespace"`

	Version string `json:"version"`

	DeploymentUnits []DeploymentUnit `json:"deployment_units"`

	ViolationSummary     Summary `json:"violation_summary"`
	VulnerabilitySummary Summary `json:"vulnerability_summary"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Namespace struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

type DeploymentUnit struct {
	ID string `json:"id"`

	Name      string `json:"name"`
	Namespace string `json:"namespace"`

	Components []Component `json:"components"`
}

type Component struct {
	ID string `json:"id"`

	Name      string     `json:"name"`
	Artifacts []Artifact `json:"artifacts"`
}

type Artifact struct {
	ID string `json:"id"`

	Purl   string `json:"purl"`
	Digest string `json:"digest"`
}

type ReleaseSynced struct {
	ID string `json:"id"`

	Name      string `json:"name"`
	Namespace string `json:"namespace"`

	Version string `json:"version"`

	DeploymentUnits []DeploymentUnit `json:"deployment_units"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Summary struct {
	Low      int32 `json:"low"`
	Medium   int32 `json:"medium"`
	High     int32 `json:"high"`
	Critical int32 `json:"critical"`
}
