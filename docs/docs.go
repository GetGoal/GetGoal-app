// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/labels": {
            "get": {
                "description": "Find All Label",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Label"
                ],
                "summary": "Find All Label",
                "operationId": "FindAllLabel",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new label",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Label"
                ],
                "summary": "Save Label",
                "operationId": "SaveLabel",
                "parameters": [
                    {
                        "description": "Label Request",
                        "name": "label",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LabelRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/labels/:id": {
            "get": {
                "description": "Search label by provided ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Label"
                ],
                "summary": "Find Label By ID",
                "operationId": "FindLabelByID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Label ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update existing label",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Label"
                ],
                "summary": "Update Label",
                "operationId": "UpdateLabel",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Label ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Label Request",
                        "name": "label",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LabelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete existing label",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Label"
                ],
                "summary": "Delete Label",
                "operationId": "DeleteLabel",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Label ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/labels/search": {
            "get": {
                "description": "Get label with random order and limit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Label"
                ],
                "summary": "Search Label Filter",
                "operationId": "GetSearchLabel",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/programs": {
            "get": {
                "description": "Find All program",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Program"
                ],
                "summary": "Find All program",
                "operationId": "FindAllProgram",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Save new program",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Program"
                ],
                "summary": "Save new program",
                "operationId": "SaveProgram",
                "parameters": [
                    {
                        "description": "Program Create or Update",
                        "name": "program",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ProgramCreateOrUpdate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/programs/:id": {
            "get": {
                "description": "Find program by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Program"
                ],
                "summary": "Find Program by ID",
                "operationId": "FindProgramById",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Program ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/programs/filter": {
            "post": {
                "description": "Find program by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Program"
                ],
                "summary": "Find Program by ID",
                "operationId": "FindProgramByLabel",
                "parameters": [
                    {
                        "description": "label name",
                        "name": "labels",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/search": {
            "post": {
                "description": "Find program by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Program"
                ],
                "summary": "Find Program by ID",
                "operationId": "FindProgramByText",
                "parameters": [
                    {
                        "description": "Search Text",
                        "name": "text",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Search"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks": {
            "get": {
                "description": "Find All Tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Find All Tasks",
                "operationId": "FindAllTask",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Save Task",
                "operationId": "SaevTask",
                "parameters": [
                    {
                        "description": "Task Data",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TaskCreateOrUpdate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/:id": {
            "get": {
                "description": "Find a task by passing ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Find task by ID",
                "operationId": "FindTaskById",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a task by passing ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Update Task",
                "operationId": "UpdateTask",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task Data",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TaskCreateOrUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Existing Task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Delete Task",
                "operationId": "DeleteTask",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/done/:id": {
            "put": {
                "description": "update status to 1 (done)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "update status to done",
                "operationId": "UpdateDone",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task  ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/join-program/:program_id": {
            "post": {
                "description": "Create tasks from program",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "join program",
                "operationId": "JoinProgram",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Program ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Modifications",
                        "name": "modifications",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.JoinProgramModifications"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/plan/:program_id": {
            "get": {
                "description": "List task related to that program_id for planning",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "list task to plan program",
                "operationId": "PlanTask",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Program ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/to-do": {
            "get": {
                "description": "Find Task by Email and Date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Find Task by Email and Date",
                "operationId": "FindTaskByEmailDate",
                "parameters": [
                    {
                        "description": "Data",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ToDoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tasks/un-done/:id": {
            "put": {
                "description": "update status to 0 (todo)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "update status to todo",
                "operationId": "UpdateTodo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task  ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.GeneralResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.GeneralResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "count": {
                    "type": "integer"
                },
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "model.JoinProgramModifications": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "modifications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Modification"
                    }
                }
            }
        },
        "model.LabelRequest": {
            "type": "object",
            "required": [
                "label_name"
            ],
            "properties": {
                "label_name": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 4
                }
            }
        },
        "model.Modification": {
            "type": "object",
            "required": [
                "start_time"
            ],
            "properties": {
                "is_set_noti": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "string"
                },
                "time_before_notify": {
                    "type": "integer"
                }
            }
        },
        "model.ProgramCreateOrUpdate": {
            "type": "object",
            "required": [
                "expected_time",
                "media_url",
                "program_desc",
                "program_name",
                "tasks"
            ],
            "properties": {
                "expected_time": {
                    "type": "integer"
                },
                "labels": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.LabelRequest"
                    }
                },
                "media_url": {
                    "type": "string"
                },
                "program_desc": {
                    "type": "string"
                },
                "program_name": {
                    "type": "string"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TaskCreateOrUpdate"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Search": {
            "type": "object",
            "required": [
                "search_text"
            ],
            "properties": {
                "search_text": {
                    "type": "string"
                }
            }
        },
        "model.TaskCreateOrUpdate": {
            "type": "object",
            "required": [
                "owner",
                "task_name"
            ],
            "properties": {
                "category": {
                    "type": "string"
                },
                "is_set_noti": {
                    "type": "integer",
                    "default": 0
                },
                "link": {
                    "type": "string"
                },
                "media_url": {
                    "type": "string"
                },
                "owner": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "string"
                },
                "task_description": {
                    "type": "string"
                },
                "task_name": {
                    "type": "string",
                    "maxLength": 250
                },
                "time_before_notify": {
                    "type": "integer"
                }
            }
        },
        "model.ToDoRequest": {
            "type": "object",
            "required": [
                "date",
                "email"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
