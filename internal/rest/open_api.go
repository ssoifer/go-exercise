package rest

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/gorilla/mux"
)

//go:generate go run ../../cmd/openapi-gen/main.go -path .
//go:generate oapi-codegen -package openapi3 -generate types  -o ../../pkg/openapi3/task_types.gen.go openapi3.yaml
//go:generate oapi-codegen -package openapi3 -generate chi-server -o ../../pkg/openapi3/openapi_server.gen.go openapi3.yaml

// NewOpenAPI3 instantiates the OpenAPI specification for this service.
func NewOpenAPI3() openapi3.T {
	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Edge BE onboarding exercise - API",
			Description: "REST APIs used for interacting with the Edge BE onboarding Service",
			Version:     "0.0.1",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				URL: "shay_soifer@dell.com",
			},
		},
		Tags: openapi3.Tags{
			&openapi3.Tag{
				Description: "Task management services",
				Name:        "Task management",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://0.0.0.0:9234",
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"Task": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewUUIDSchema()).
				WithProperty("title", openapi3.NewStringSchema()).
				WithProperty("content", openapi3.NewStringSchema()).
				WithProperty("views", openapi3.NewInt64Schema()).
				WithProperty("timestamp", openapi3.NewInt64Schema())),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CreateTasksRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for creating a task.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("task", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Task",
					})),
		},
		//"UpdateTasksRequest": &openapi3.RequestBodyRef{
		//	Value: openapi3.NewRequestBody().
		//		WithDescription("Request used for updating a task.").
		//		WithRequired(true).
		//		WithJSONSchema(openapi3.NewSchema().
		//			WithProperty("description", openapi3.NewStringSchema().
		//				WithMinLength(1)).
		//			WithProperty("is_done", openapi3.NewBoolSchema().
		//				WithDefault(false)).
		//			WithPropertyRef("priority", &openapi3.SchemaRef{
		//				Ref: "#/components/schemas/Priority",
		//			}).
		//			WithPropertyRef("dates", &openapi3.SchemaRef{
		//				Ref: "#/components/schemas/Dates",
		//			})),
		//},
	}

	swagger.Components.Responses = openapi3.Responses{
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when errors happen.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("error", openapi3.NewStringSchema()))),
		},
		"CreateTasksResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after creating tasks.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("success", openapi3.NewBoolSchema()))),
		},
		"ReadTaskByIdResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after searching one task.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("task", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Task",
					}))),
		},
		"ReadAllTasksResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after searching one task.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewArraySchema().
					WithItems(openapi3.NewObjectSchema().WithPropertyRef("task", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Task",
					})),
				)),
		},
	}

	swagger.Paths = openapi3.Paths{
		"/tasks": &openapi3.PathItem{
			Post: &openapi3.Operation{
				Tags:        []string{"Task management"},
				OperationID: "CreateTask",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CreateTasksRequest",
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/CreateTasksResponse",
					},
				},
			},
			Get: &openapi3.Operation{
				Tags:        []string{"Task management"},
				OperationID: "GetAllTasks",
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/ReadAllTasksResponse",
					},
				},
			},
		},
		"/tasks/{taskId}": &openapi3.PathItem{
			Get: &openapi3.Operation{
				Tags:        []string{"Task management"},
				OperationID: "ReadTask",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("taskId").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				Responses: openapi3.Responses{
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/ReadTaskByIdResponse",
					},
				},
			},
			//Put: &openapi3.Operation{
			//	OperationID: "UpdateTask",
			//	Parameters: []*openapi3.ParameterRef{
			//		{
			//			Value: openapi3.NewPathParameter("taskId").
			//				WithSchema(openapi3.NewUUIDSchema()),
			//		},
			//	},
			//	RequestBody: &openapi3.RequestBodyRef{
			//		Ref: "#/components/requestBodies/UpdateTasksRequest",
			//	},
			//	Responses: openapi3.Responses{
			//		"400": &openapi3.ResponseRef{
			//			Ref: "#/components/responses/ErrorResponse",
			//		},
			//		"500": &openapi3.ResponseRef{
			//			Ref: "#/components/responses/ErrorResponse",
			//		},
			//		"200": &openapi3.ResponseRef{
			//			Value: openapi3.NewResponse().WithDescription("Task was updated"),
			//		},
			//	},
			//},
		},
	}

	return swagger
}

func RegisterOpenAPI(r *mux.Router) {
	swagger := NewOpenAPI3()

	r.HandleFunc("/openapi3.json", func(w http.ResponseWriter, r *http.Request) {
		renderResponse(w, &swagger, http.StatusOK)
	}).Methods(http.MethodGet)

	r.HandleFunc("/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")

		data, _ := yaml.Marshal(&swagger)

		_, _ = w.Write(data)

		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
}
