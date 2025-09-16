package docs

import (
	"fmt"
	"strings"

	"github.com/swaggo/swag"
)

// SwaggerInformation allows runtime overrides from main.go (optional)
type SwaggerInformation struct {
	Title       string
	Description string
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
}

// SwaggerInfo holds default values; main.go may override these.
var SwaggerInfo = SwaggerInformation{
	Title:       "E-Commerce Backend API",
	Description: "REST API for users, products, carts, and orders.",
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{"http"},
}

// s implements swag.Spec and returns the Swagger JSON.
type s struct{}

// ReadDoc builds an OpenAPI (Swagger 2.0) document with auth and user endpoints.
func (s *s) ReadDoc() string {
	// format schemes array
	schemes := ""
	for i, sc := range SwaggerInfo.Schemes {
		if i > 0 {
			schemes += ","
		}
		schemes += fmt.Sprintf("\"%s\"", sc)
	}

	doc := fmt.Sprintf(`{
  "swagger": "2.0",
  "info": {
    "description": %q,
    "title": %q,
    "version": %q
  },
  "host": %q,
  "basePath": %q,
  "schemes": [%s],
  "consumes": ["application/json"],
  "produces": ["application/json"],
  "securityDefinitions": {
    "BearerAuth": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header",
      "description": "JWT bearer token. Format: Bearer {token}"
    }
  },
  "paths": {
    "/register": {
      "post": {
        "summary": "Register a new user",
        "tags": ["auth"],
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/UserCreate" }
          }
        ],
        "responses": {
          "200": {
            "description": "User created",
            "schema": { "$ref": "#/definitions/User" }
          }
        }
      }
    },
    "/login": {
      "post": {
        "summary": "Login user",
        "tags": ["auth"],
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": { "$ref": "#/definitions/LoginRequest" }
          }
        ],
        "responses": {
          "200": {
            "description": "Login response",
            "schema": { "$ref": "#/definitions/AuthResponse" }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "summary": "Get user by ID",
        "tags": ["user"],
        "parameters": [
          { "name": "id", "in": "path", "required": true, "type": "integer", "format": "int64" }
        ],
        "security": [{ "BearerAuth": [] }],
        "responses": {
          "200": { "description": "User", "schema": { "$ref": "#/definitions/User" } },
          "401": { "description": "Unauthorized" }
        }
      },
      "put": {
        "summary": "Update user by ID",
        "tags": ["user"],
        "parameters": [
          { "name": "id", "in": "path", "required": true, "type": "integer", "format": "int64" },
          { "in": "body", "name": "body", "required": true, "schema": { "$ref": "#/definitions/UserUpdate" } }
        ],
        "security": [{ "BearerAuth": [] }],
        "responses": {
          "200": { "description": "Updated user", "schema": { "$ref": "#/definitions/User" } },
          "401": { "description": "Unauthorized" }
        }
      },
      "delete": {
        "summary": "Delete user by ID",
        "tags": ["user"],
        "parameters": [
          { "name": "id", "in": "path", "required": true, "type": "integer", "format": "int64" }
        ],
        "security": [{ "BearerAuth": [] }],
        "responses": {
          "200": { "description": "Deleted" },
          "401": { "description": "Unauthorized" }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "id": { "type": "integer", "format": "int64" },
        "user_name": { "type": "string" },
        "email": { "type": "string" },
        "is_admin": { "type": "boolean" },
        "created_at": { "type": "string" },
        "updated_at": { "type": "string" }
      }
    },
    "UserCreate": {
      "type": "object",
      "required": ["user_name", "email", "password"],
      "properties": {
        "user_name": { "type": "string" },
        "email": { "type": "string" },
        "password": { "type": "string" },
        "is_admin": { "type": "boolean" }
      }
    },
    "UserUpdate": {
      "type": "object",
      "properties": {
        "user_name": { "type": "string" },
        "email": { "type": "string" },
        "password": { "type": "string" },
        "is_admin": { "type": "boolean" }
      }
    },
    "LoginRequest": {
      "type": "object",
      "required": ["email", "password"],
      "properties": {
        "email": { "type": "string" },
        "password": { "type": "string" }
      }
    },
    "AuthResponse": {
      "type": "object",
      "properties": {
        "status": { "type": "boolean" },
        "message": { "type": "string" },
        "token": { "type": "string" },
        "user": { "$ref": "#/definitions/User" }
      }
    }
  }
}`,
		SwaggerInfo.Description,
		SwaggerInfo.Title,
		SwaggerInfo.Version,
		SwaggerInfo.Host,
		SwaggerInfo.BasePath,
		schemes,
	)

	return strings.TrimSpace(doc)
}

func init() {
	swag.Register(swag.Name, &s{})
}
