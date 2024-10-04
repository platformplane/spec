package cloudevents

import "time"

const (
	ReleaseCandidateCreatedType = "io.platformplane.releasecandidate.created"
	ReleaseCreatedType          = "io.platformplane.release.created"

	ReleaseApprovedType = "io.platformplane.release.approved"
	ReleaseRejectedType = "io.platformplane.release.rejected"

	ReleaseSyncedType   = "io.platformplane.release.synced"
	DeploymentReadyType = "io.platformplane.deployment.ready"
)

type Release struct {
	ID string `json:"id"`

	Name      string `json:"name"`
	Namespace string `json:"namespace"`

	Version string `json:"version"`

	DeploymentUnits []DeploymentUnit `json:"deployment_units"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ReleaseCandidateCreated struct {
	Release

	ViolationSummary     Summary `json:"violation_summary"`
	VulnerabilitySummary Summary `json:"vulnerability_summary"`

	Refs []DataRef `json:"data_ref"`
}

type ReleaseCreated struct {
	Release
}

type ReleaseApproved struct {
	Release
}

type ReleaseRejected struct {
	Release
}

type ReleaseSynced struct {
	Release
}

type DeploymentReady struct {
	Release

	Environment *Environment `json:"environment"`
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

type Summary struct {
	Low      int32 `json:"low"`
	Medium   int32 `json:"medium"`
	High     int32 `json:"high"`
	Critical int32 `json:"critical"`
}

type Environment struct {
	ID string `json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	Target string `json:"target"`
}
