// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Developer",
            "email": "kongwoojin03@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/arch/{board}": {
            "get": {
                "description": "Get arch article list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "arch"
                ],
                "summary": "Get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/article/arch": {
            "get": {
                "description": "Get arch article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "arch"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/article/cse": {
            "get": {
                "description": "Get cse article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cse"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/article/dorm": {
            "get": {
                "description": "Get dorm article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dorm"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/article/emc": {
            "get": {
                "description": "Get emc article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emc"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/article/ide": {
            "get": {
                "description": "Get ide article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ide"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/article/ite": {
            "get": {
                "description": "Get ite article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ite"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/article/mechanical": {
            "get": {
                "description": "Get mechanical article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mechanical"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/article/mechatronics": {
            "get": {
                "description": "Get mechatronics article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mechatronics"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/article/school": {
            "get": {
                "description": "Get school article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "school"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/article/sim": {
            "get": {
                "description": "Get sim article by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sim"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid of article",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/cse/{board}": {
            "get": {
                "description": "Get cse article list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cse"
                ],
                "summary": "Get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/dorm/{board}": {
            "get": {
                "description": "Get dorm article list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dorm"
                ],
                "summary": "Get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/emc/{board}": {
            "get": {
                "description": "Get emc article list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "emc"
                ],
                "summary": "Get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/ide/{board}": {
            "get": {
                "description": "Get ide article list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ide"
                ],
                "summary": "Get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/ite/{board}": {
            "get": {
                "description": "Get ite article list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ite"
                ],
                "summary": "Get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/mechanical/{board}": {
            "get": {
                "description": "Get mechanical article list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mechanical"
                ],
                "summary": "Get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/mechatronics/{board}": {
            "get": {
                "description": "Get mechatronics article list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mechatronics"
                ],
                "summary": "Get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/school/{board}": {
            "get": {
                "description": "Get school article list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "school"
                ],
                "summary": "Get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/sim/{board}": {
            "get": {
                "description": "Get sim article list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sim"
                ],
                "summary": "Get article list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the board",
                        "name": "board",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.APIData"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.APIData": {
            "type": "object",
            "properties": {
                "last_page": {
                    "description": "Last page of requested board",
                    "type": "integer"
                },
                "posts": {
                    "description": "Data of board",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Board"
                    }
                }
            }
        },
        "model.Article": {
            "type": "object",
            "properties": {
                "article_url": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Files"
                    }
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "write_date": {
                    "type": "string"
                },
                "writer": {
                    "type": "string"
                }
            }
        },
        "model.Board": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "num": {
                    "type": "string"
                },
                "read_count": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "write_date": {
                    "type": "string"
                },
                "writer": {
                    "type": "string"
                }
            }
        },
        "model.Files": {
            "type": "object",
            "properties": {
                "file_name": {
                    "type": "string"
                },
                "file_url": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v3",
	Schemes:          []string{},
	Title:            "KOREATECH board REST API",
	Description:      "This is unofficial version of KOREATECH board REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
