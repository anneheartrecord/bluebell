// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "这里写联系人信息",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/posts2": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "可按社区按时间或分数排序查询帖子列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "升级版帖子列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostList"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller._ResponsePostList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ApiPostDetail"
                    }
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "models.ApiPostDetail": {
            "type": "object",
            "required": [
                "community_id",
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "type": "string",
                    "example": "0"
                },
                "author_name": {
                    "type": "string"
                },
                "community_id": {
                    "type": "string",
                    "example": "0"
                },
                "content": {
                    "type": "string"
                },
                "creat_time": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "0"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "vote_num": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "这里写接口服务的host",
	BasePath:         "这里写base path",
	Schemes:          []string{},
	Title:            "这里写标题",
	Description:      "这里写描述信息",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
