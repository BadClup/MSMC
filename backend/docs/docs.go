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
        "/get-jwt-payload": {
            "post": {
                "description": "This is a handler for auth-service only.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Get JWT payload",
                "parameters": [
                    {
                        "description": "Token",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal.GetJwtPayloadDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "payload": {
                                    "$ref": "#/definitions/internal.JwtPayload"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/login-remote": {
            "post": {
                "description": "This is a handler for auth-service only. It returns a token, which should be forwarded to the client.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login handler for auth-service only",
                "parameters": [
                    {
                        "description": "User ID",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal.LoginRemoteDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "token": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "error": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                }
            }
        },
        "/server": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "Get all servers",
                "operationId": "get-all-servers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/shared.ServerInstanceStatus"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "Create a new server",
                "operationId": "create-server",
                "parameters": [
                    {
                        "description": "Server data",
                        "name": "server",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mcserver.createServerDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Server created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internal.GetJwtPayloadDto": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "internal.JwtPayload": {
            "type": "object",
            "properties": {
                "ExpiresAt": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "internal.LoginRemoteDto": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "mcserver.createServerDto": {
            "type": "object",
            "properties": {
                "engine": {
                    "default": "vanilla",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.MinecraftEngine"
                        }
                    ]
                },
                "name": {
                    "type": "string",
                    "default": "example-minecraft-server"
                },
                "port": {
                    "type": "integer",
                    "default": 25565
                },
                "seed": {
                    "type": "string",
                    "default": "example-seed"
                },
                "version": {
                    "type": "string",
                    "default": "1.16.5"
                }
            }
        },
        "shared.MinecraftEngine": {
            "type": "string",
            "enum": [
                "vanilla",
                "forge"
            ],
            "x-enum-varnames": [
                "VanillaEngine",
                "ForgeEngine"
            ]
        },
        "shared.ServerInstance": {
            "type": "object",
            "properties": {
                "docker_id": {
                    "type": "string",
                    "default": "af5bb532db04"
                },
                "engine": {
                    "default": "vanilla",
                    "allOf": [
                        {
                            "$ref": "#/definitions/shared.MinecraftEngine"
                        }
                    ]
                },
                "name": {
                    "type": "string",
                    "default": "My Server"
                },
                "port": {
                    "type": "integer",
                    "default": 25565
                },
                "seed": {
                    "type": "string",
                    "default": "example-seed"
                },
                "version": {
                    "type": "string",
                    "default": "1.16.5"
                }
            }
        },
        "shared.ServerInstanceStatus": {
            "type": "object",
            "properties": {
                "running": {
                    "type": "boolean",
                    "default": false
                },
                "server_instance": {
                    "$ref": "#/definitions/shared.ServerInstance"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MSMC API",
	Description:      "This is MSMC API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
