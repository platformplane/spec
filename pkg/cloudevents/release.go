package cloudevents

import "time"

const (
	ReleaseApprovedType = "io.platformplane.release.approved"
	ReleaseRejectedType = "io.platformplane.release.rejected"

	ReleaseCreatedType          = "io.platformplane.release.created"
	ReleaseCandidateCreatedType = "io.platformplane.releasecandidate.created"

	ReleaseSyncedType = "io.platformplane.release.synced"

	DeploymentReadyType = "io.platformplane.deployment.ready"
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

type ReleaseRejected struct {
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

	Refs []DataRef `json:"data_ref"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DataRef struct {
	Name string `json:"name"`
	Type string `json:"type"`

	Ref string `json:"ref"`
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

type DeploymentReady struct {
	ID string `json:"id"`

	Name      string     `json:"name"`
	Namespace *Namespace `json:"namespace"`

	Version string `json:"version"`

	DeploymentUnits []DeploymentUnit `json:"deployment_units"`

	Environment *Environment `json:"environment"`
}

type Environment struct {
	ID string `json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	Target string `json:"target"`
}
