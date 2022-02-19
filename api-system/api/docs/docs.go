// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

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
        "/v1/captcha": {
            "get": {
                "tags": [
                    "Captcha"
                ],
                "summary": "请求生成验证码",
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
        "/v1/captcha/{file}": {
            "get": {
                "consumes": [
                    "image/png"
                ],
                "tags": [
                    "Captcha"
                ],
                "summary": "获取验证码图片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "image file name",
                        "name": "file",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "lang",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "name": "reload",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/system/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "System"
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
                    "User"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "register form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.LoginForm"
                        }
                    },
                    {
                        "type": "string",
                        "name": "captcha_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "captcha_val",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.ErrJson"
                        }
                    }
                }
            }
        },
        "/v1/user/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "User"
                ],
                "summary": "User logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.ErrJson"
                        }
                    }
                }
            }
        },
        "/v1/user/register": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.RegisterForm"
                        }
                    },
                    {
                        "type": "string",
                        "name": "captcha_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "captcha_val",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.ErrJson"
                        }
                    }
                }
            }
        },
        "/v1/user/value/exist": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "判断主要的用户信息的值是否存在",
                "parameters": [
                    {
                        "description": "form",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.KeyFieldExistForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.ApiJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.ErrJson"
                        }
                    }
                }
            }
        },
        "/v1/visits": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Visit"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "query userinfo",
                        "name": "form",
                        "in": "query",
                        "required": true
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
        "/v1/visits/count": {
            "get": {
                "tags": [
                    "Visit"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "name": "timestamp",
                        "in": "query"
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
        "/v1/visits/visit": {
            "post": {
                "tags": [
                    "Visit"
                ],
                "summary": "登记访问信息",
                "parameters": [
                    {
                        "description": "query userinfo",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.PostVisitForm"
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
        "/v1/war_thunder/userinfo": {
            "get": {
                "tags": [
                    "WarThunder"
                ],
                "summary": "获取异步查询结果",
                "parameters": [
                    {
                        "type": "string",
                        "name": "query_id",
                        "in": "query",
                        "required": true
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
        "/v1/war_thunder/userinfo/queries": {
            "get": {
                "tags": [
                    "WarThunder"
                ],
                "summary": "查询游戏昵称的所有query_id",
                "parameters": [
                    {
                        "maxLength": 20,
                        "type": "string",
                        "description": "游戏的昵称",
                        "name": "nickname",
                        "in": "query",
                        "required": true
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
        "/v1/war_thunder/userinfo/query/count": {
            "get": {
                "tags": [
                    "WarThunder"
                ],
                "summary": "查询query的数量",
                "parameters": [
                    {
                        "type": "string",
                        "name": "timestamp",
                        "in": "query"
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
        "/v1/war_thunder/userinfo/refresh": {
            "post": {
                "tags": [
                    "WarThunder"
                ],
                "summary": "刷新一个游戏用户数据的最新数据",
                "parameters": [
                    {
                        "maxLength": 20,
                        "type": "string",
                        "description": "游戏的昵称",
                        "name": "nickname",
                        "in": "query",
                        "required": true
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
        "app.ErrJson": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "v1.KeyFieldExistForm": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "key": {
                    "type": "string",
                    "enum": [
                        "username",
                        "email"
                    ]
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "v1.LoginForm": {
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
        },
        "v1.PostVisitForm": {
            "type": "object",
            "properties": {
                "client_id": {
                    "description": "客户端生成id",
                    "type": "string"
                },
                "page": {
                    "description": "访问页面",
                    "type": "string"
                },
                "user_id": {
                    "description": "用户id",
                    "type": "integer"
                }
            }
        },
        "v1.RegisterForm": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "avatar_url": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "invited_code": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
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
	Title:       "Golang Gin Template API",
	Description: "An example of gin",
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
