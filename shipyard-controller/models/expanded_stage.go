// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ExpandedStage expanded stage
//
// swagger:model ExpandedStage
type ExpandedStage struct {

	// last event context
	LastEventContext *EventContext `json:"lastEventContext,omitempty"`

	// services
	Services []*ExpandedService `json:"services"`

	// Stage name
	StageName string `json:"stageName,omitempty"`
}
