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
                        "CqhttpSelfID": []
                    },
                    {
                        "CqhttpSignature": []
                    }
                ],
                "tags": [
                    "CQHttp API"
                ],
                "summary": "接收cqhttp的事件",
                "parameters": [
                    {
                        "description": "cqhttp event",
                        "name": "event",
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
        }
    },
    "securityDefinitions": {
        "AppToken": {
            "type": "apiKey",
            "name": "app_token",
            "in": "query"
        },
        "CqhttpSelfID": {
            "type": "apiKey",
            "name": "X-Self-ID",
            "in": "header"
        },
        "CqhttpSignature": {
            "type": "apiKey",
            "name": "X-Signature",
            "in": "header"
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
