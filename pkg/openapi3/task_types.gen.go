// Package openapi3 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.10.0 DO NOT EDIT.
package openapi3

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Task defines model for Task.
type Task struct {
	Content   *string             `json:"content,omitempty"`
	Id        *openapi_types.UUID `json:"id,omitempty"`
	Timestamp *int64              `json:"timestamp,omitempty"`
	Title     *string             `json:"title,omitempty"`
	Views     *int64              `json:"views,omitempty"`
}

// CreateTasksResponse defines model for CreateTasksResponse.
type CreateTasksResponse struct {
	Success *bool `json:"success,omitempty"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// ReadAllTasksResponse defines model for ReadAllTasksResponse.
type ReadAllTasksResponse []struct {
	Task *Task `json:"task,omitempty"`
}

// ReadTaskByIdResponse defines model for ReadTaskByIdResponse.
type ReadTaskByIdResponse struct {
	Task *Task `json:"task,omitempty"`
}

// CreateTasksRequest defines model for CreateTasksRequest.
type CreateTasksRequest struct {
	Task *Task `json:"task,omitempty"`
}

// CreateTaskJSONRequestBody defines body for CreateTask for application/json ContentType.
type CreateTaskJSONRequestBody CreateTasksRequest
