// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/accounts/{id}": {
            "get": {
                "description": "Get one exists account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "get one exists account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/api/transactions": {
            "post": {
                "description": "handle new transaction.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "handle new transaction",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GeneralResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.GeneralResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "models.TransactionRequest": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer"
                },
                "amount": {
                    "type": "number"
                },
                "messageType": {
                    "description": "PAYMENT | ADJUSTMENT",
                    "type": "string"
                },
                "origin": {
                    "description": "VISA | MASTER",
                    "type": "string"
                },
                "transactionId": {
                    "description": "GUID",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:9999",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Inspiration Tech Case",
	Description:      "Web API for an issuer bank",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
