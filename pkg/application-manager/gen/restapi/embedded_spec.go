///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch - Application Manager APIs\n",
    "title": "Application Manager",
    "contact": {
      "email": "dispatch@vmware.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1/application",
  "paths": {
    "/": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "application"
        ],
        "summary": "List all existing Applications",
        "operationId": "getApps",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter on Application tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/getAppsOKBody"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "default": {
            "description": "Unexpected Error",
            "schema": {
              "$ref": "#/definitions/Error"
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
          "application"
        ],
        "summary": "Add a new Application",
        "operationId": "addApp",
        "parameters": [
          {
            "description": "Application object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Application"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Application created",
            "schema": {
              "$ref": "#/definitions/Application"
            }
          },
          "400": {
            "description": "Invalid Input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "401": {
            "description": "Unauthorized Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/{application}": {
      "get": {
        "description": "get an Application by name",
        "produces": [
          "application/json"
        ],
        "tags": [
          "application"
        ],
        "summary": "Find Application by name",
        "operationId": "getApp",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Application"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Application not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "application"
        ],
        "summary": "Update an Application",
        "operationId": "updateApp",
        "parameters": [
          {
            "description": "Application object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Application"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful update",
            "schema": {
              "$ref": "#/definitions/Application"
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Application not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "application"
        ],
        "summary": "Deletes an Application",
        "operationId": "deleteApp",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Application"
            }
          },
          "400": {
            "description": "Invalid Name supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Application not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of Application to work on",
          "name": "application",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Application": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "createdTime": {
          "type": "integer",
          "readOnly": true
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "modifiedTime": {
          "type": "integer",
          "readOnly": true
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "status": {
          "$ref": "#/definitions/Status",
          "readOnly": true
        },
        "tags": {
          "$ref": "#/definitions/applicationTags"
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Status": {
      "type": "string",
      "enum": [
        "CREATING",
        "READY",
        "UPDATING",
        "DELETED",
        "ERROR"
      ]
    },
    "Tag": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "applicationTags": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Tag"
      },
      "x-go-gen-location": "models"
    },
    "getAppsOKBody": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Application"
      },
      "x-go-gen-location": "operations"
    }
  },
  "securityDefinitions": {
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookie": []
    }
  ],
  "tags": [
    {
      "description": "CRUD operations on Applications",
      "name": "application"
    }
  ]
}`))
}
