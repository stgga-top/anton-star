// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "axiangcoding",
            "email": "axiangcoding@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/cqhttp/receive/event": {
            "post": {
                "security": [
                    {
                        "AppToken": []
                    }
                ],
                "tags": [
                    "CQHttp API"
                ],
                "summary": "接收cqhttp的事件",
                "parameters": [
                    {
                        "description": "getParam",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    }
                }
            }
        },
        "/v1/cqhttp/status": {
            "get": {
                "tags": [
                    "CQHttp API"
                ],
                "summary": "获取cqhttp的最新状态",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    }
                }
            }
        },
        "/v1/system/info": {
            "get": {
                "tags": [
                    "System API"
                ],
                "summary": "System Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    }
                }
            }
        },
        "/v1/user/login": {
            "post": {
                "tags": [
                    "User API"
                ],
                "parameters": [
                    {
                        "description": "login param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.LoginParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    }
                }
            }
        },
        "/v1/user/me": {
            "post": {
                "tags": [
                    "User API"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    }
                }
            }
        },
        "/v1/user/register": {
            "post": {
                "tags": [
                    "User API"
                ],
                "parameters": [
                    {
                        "description": "register param",
                        "name": "param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.RegisterParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.ApiJson": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "v1.LoginParam": {
            "type": "object",
            "required": [
                "loginName",
                "password"
            ],
            "properties": {
                "loginName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "v1.RegisterParam": {
            "type": "object",
            "required": [
                "displayName",
                "email",
                "password"
            ],
            "properties": {
                "displayName": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AppToken": {
            "type": "apiKey",
            "name": "app_token",
            "in": "query"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "axiangcoding/anton-star",
	Description: "api system build by ax-web",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
