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
        "/CreateList/{id}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Sign-In with user credentials, for generated access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/Todo.ResponseJson"
                        }
                    }
                }
            }
        },
        "/CreateTask/{listid}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Sign-In with user credentials, for generated access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create Task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "List ID",
                        "name": "listid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Add Task",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Todo.SaveTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/Todo.ResponseJson"
                        }
                    }
                }
            }
        },
        "/DeleteList/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Sign-In with user credentials, for generated access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/Todo.ResponseJson"
                        }
                    }
                }
            }
        },
        "/DeleteTask/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Sign-In with user credentials, for generated access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/Todo.ResponseJson"
                        }
                    }
                }
            }
        },
        "/GetList/{userid}": {
            "get": {
                "description": "Sign-In with user credentials, for generated access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/Todo.ResponseJson"
                        }
                    }
                }
            }
        },
        "/GetUser/{id}": {
            "get": {
                "description": "Fetch User Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetUserById",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/authentication.ResponseJson"
                        }
                    }
                }
            }
        },
        "/Login": {
            "post": {
                "description": "Sign-In with user credentials, for generated access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authentication.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful",
                        "schema": {
                            "$ref": "#/definitions/authentication.ResponseJson"
                        }
                    }
                }
            }
        },
        "/Register": {
            "post": {
                "description": "Create User Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authentication.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/authentication.ResponseJson"
                        }
                    }
                }
            }
        },
        "/TaskCompleted/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    },
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Sign-In with user credentials, for generated access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Change Status Task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Change Status",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Todo.SetStatus"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/Todo.ResponseJson"
                        }
                    }
                }
            }
        },
        "/UpdateTask/{id}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Sign-In with user credentials, for generated access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update Task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Task",
                        "name": "Request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Todo.SaveTask"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/Todo.ResponseJson"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Todo.ResponseJson": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Success"
                }
            }
        },
        "Todo.SaveTask": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "Todo.SetStatus": {
            "type": "object",
            "required": [
                "isCompleted"
            ],
            "properties": {
                "isCompleted": {
                    "type": "boolean"
                }
            }
        },
        "authentication.ResponseJson": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Success"
                }
            }
        },
        "authentication.User": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Todo Web API",
	Description:      "Todo Web API with JWT Auth",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}