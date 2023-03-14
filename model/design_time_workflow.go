package model

import (
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type Workflow struct {
	Id             string                 `json:"id,omitempty"`
	Name           string                 `json:"name,omitempty"`
	CreatedAt      *timestamppb.Timestamp `json:"created_at,omitempty"`
	ModifiedAt     *timestamppb.Timestamp `json:"modified_at,omitempty"`
	ExportedOn     *timestamppb.Timestamp `json:"exported_on,omitempty"`
	CreatedBy      User                   `json:"created_by,omitempty"`
	ModifiedBy     User                   `json:"modified_by,omitempty"`
	OrganizationId string                 `json:"organization_id,omitempty"`
	ProjectId      string                 `json:"project_id,omitempty"`
	Config         Config                 `json:"config,omitempty"`
	Activities     []*Activity            `json:"activities,omitempty"`
	CurentState    RunState               `json:"current_state,omitempty"`
	CompletedState RunState               `json:"completed_state,omitempty"`
	Reason         string                 `json:"reason,omitempty"`
	Version        uint32                 `json:"version,omitempty"`
	RetryCount     uint32                 `json:"retry_count,omitempty"`
}

type Activity struct {
	Id             string                 `json:"id,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Children       []string               `json:"children,omitempty"`
	OrganizationId string                 `json:"organization_id,omitempty"`
	ProjectId      string                 `json:"project_id,omitempty"`
	Config         Config                 `json:"config,omitempty"`
	CurrentState   RunState               `json:"current_state,omitempty"`
	CompletedState RunState               `json:"completed_state,omitempty"`
	Reason         string                 `json:"reason,omitempty"`
	StartedAt      *timestamppb.Timestamp `json:"started_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp `json:"updated_at,omitempty"`
	RetryCount     uint32                 `json:"retry_count,omitempty"`
	Skip           bool                   `json:"skip,omitempty"`
	Files          []*ActivityRunFile     `json:"files,omitempty"`
}

type RunState int32

const (
	RUN_COMPLETED_STATE_UNSPECIFIED = 0
	RUN_COMPLETED_STATE_SUCCESS     = 1
	RUN_COMPLETED_STATE_FAILED      = 2
	RUN_COMPLETED_STATE_CANCELLED   = 3
	RUN_COMPLETED_STATE_SKIPPED     = 4
	RUN_COMPLETED_STATE_POLL_AGAIN  = 5
)

type User struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Admin bool   `json:"admin,omitempty"`
}

type Config struct {
	Name      string            `json:"name,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
}

type ActivityRunFile struct {
	Name     string           `json:"name,omitempty"`
	Location string           `json:"location,omitempty"`
	Output   *structpb.Struct `json:"output,omitempty"`
	Log      []byte           `json:"log,omitempty"`
}
