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

// ReadDoc builds an OpenAPI (Swagger 2.0) document with auth, product and order endpoints.
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
          {"in": "body", "name": "body", "required": true, "schema": {"$ref": "#/definitions/UserCreate"}}
        ],
        "responses": {
          "200": {"description": "User created", "schema": {"$ref": "#/definitions/User"}}
        }
      }
    },
    "/login": {
      "post": {
        "summary": "Login user",
        "tags": ["auth"],
        "parameters": [
          {"in": "body", "name": "body", "required": true, "schema": {"$ref": "#/definitions/LoginRequest"}}
        ],
        "responses": {
          "200": {"description": "Login response", "schema": {"$ref": "#/definitions/AuthResponse"}}
        }
      }
    },
    "/auth/google/login": {
      "get": {
        "summary": "Start Google OAuth2 login",
        "tags": ["auth"],
        "responses": {
          "302": {"description": "Redirects to Google OAuth2 consent screen"}
        }
      }
    },
    "/auth/google/callback": {
      "get": {
        "summary": "Google OAuth2 callback",
        "tags": ["auth"],
        "parameters": [
          {"name": "state", "in": "query", "required": true, "type": "string"},
          {"name": "code", "in": "query", "required": true, "type": "string"}
        ],
        "responses": {
          "200": {"description": "Login response", "schema": {"$ref": "#/definitions/AuthResponse"}},
          "401": {"description": "Invalid state parameter"}
        }
      }
    },

    "/user/{id}": {
      "get": {
        "summary": "Get user by ID",
        "tags": ["user"],
        "parameters": [
          {"name": "id", "in": "path", "required": true, "type": "integer", "format": "int64"}
        ],
        "security": [{"BearerAuth": []}],
        "responses": {
          "200": {"description": "User", "schema": {"$ref": "#/definitions/User"}},
          "401": {"description": "Unauthorized"}
        }
      },
      "put": {
        "summary": "Update user by ID",
        "tags": ["user"],
        "parameters": [
          {"name": "id", "in": "path", "required": true, "type": "integer", "format": "int64"},
          {"in": "body", "name": "body", "required": true, "schema": {"$ref": "#/definitions/UserUpdate"}}
        ],
        "security": [{"BearerAuth": []}],
        "responses": {
          "200": {"description": "Updated user", "schema": {"$ref": "#/definitions/User"}},
          "401": {"description": "Unauthorized"}
        }
      },
      "delete": {
        "summary": "Delete user by ID",
        "tags": ["user"],
        "parameters": [
          {"name": "id", "in": "path", "required": true, "type": "integer", "format": "int64"}
        ],
        "security": [{"BearerAuth": []}],
        "responses": {
          "200": {"description": "Deleted"},
          "401": {"description": "Unauthorized"}
        }
      }
    },

    "/product": {
      "post": {
        "summary": "Create product (admin)",
        "tags": ["product"],
        "security": [{"BearerAuth": []}],
        "parameters": [
          {"in": "body", "name": "body", "required": true, "schema": {"$ref": "#/definitions/ProductCreate"}}
        ],
        "responses": {
          "200": {"description": "Product created", "schema": {"$ref": "#/definitions/Product"}},
          "401": {"description": "Unauthorized"}
        }
      },
      "get": {
        "summary": "List products",
        "tags": ["product"],
        "parameters": [
          {"name": "new", "in": "query", "required": false, "type": "string"},
          {"name": "categories", "in": "query", "required": false, "type": "string"},
          {"name": "color", "in": "query", "required": false, "type": "string"},
          {"name": "size", "in": "query", "required": false, "type": "string"}
        ],
        "responses": {
          "200": {"description": "Products", "schema": {"type": "array", "items": {"$ref": "#/definitions/Product"}}}
        }
      }
    },
    "/product/{id}": {
      "get": {
        "summary": "Get product by ID",
        "tags": ["product"],
        "parameters": [{"name": "id", "in": "path", "required": true, "type": "integer", "format": "int64"}],
        "responses": {"200": {"description": "Product", "schema": {"$ref": "#/definitions/Product"}}}
      },
      "put": {
        "summary": "Update product by ID (admin)",
        "tags": ["product"],
        "security": [{"BearerAuth": []}],
        "parameters": [
          {"name": "id", "in": "path", "required": true, "type": "integer", "format": "int64"},
          {"in": "body", "name": "body", "required": true, "schema": {"$ref": "#/definitions/ProductUpdate"}}
        ],
        "responses": {
          "200": {"description": "Updated product", "schema": {"$ref": "#/definitions/Product"}},
          "401": {"description": "Unauthorized"}
        }
      },
      "delete": {
        "summary": "Delete product by ID (admin)",
        "tags": ["product"],
        "security": [{"BearerAuth": []}],
        "parameters": [{"name": "id", "in": "path", "required": true, "type": "integer", "format": "int64"}],
        "responses": {
          "200": {"description": "Deleted"},
          "401": {"description": "Unauthorized"}
        }
      }
    },

    "/order": {
      "post": {
        "summary": "Create order",
        "tags": ["order"],
        "security": [{"BearerAuth": []}],
        "parameters": [
          {"in": "body", "name": "body", "required": true, "schema": {"$ref": "#/definitions/OrderCreate"}}
        ],
        "responses": {"200": {"description": "Order created", "schema": {"$ref": "#/definitions/Order"}}}
      },
      "get": {
        "summary": "Get orders for current user",
        "tags": ["order"],
        "security": [{"BearerAuth": []}],
        "responses": {
          "200": {"description": "Orders", "schema": {"type": "array", "items": {"$ref": "#/definitions/Order"}}}
        }
      }
    },
    "/getallorder": {
      "get": {
        "summary": "Admin: list orders",
        "tags": ["order"],
        "security": [{"BearerAuth": []}],
        "parameters": [{"name": "new", "in": "query", "required": false, "type": "string"}],
        "responses": {"200": {"description": "Orders", "schema": {"type": "array", "items": {"$ref": "#/definitions/Order"}}}}
      }
    },
    "/order/{id}": {
      "put": {
        "summary": "Update order by ID",
        "tags": ["order"],
        "security": [{"BearerAuth": []}],
        "parameters": [
          {"name": "id", "in": "path", "required": true, "type": "integer", "format": "int64"},
          {"in": "body", "name": "body", "required": true, "schema": {"$ref": "#/definitions/OrderUpdate"}}
        ],
        "responses": {"200": {"description": "Updated", "schema": {"$ref": "#/definitions/Order"}}}
      },
      "delete": {
        "summary": "Delete order by ID",
        "tags": ["order"],
        "security": [{"BearerAuth": []}],
        "parameters": [{"name": "id", "in": "path", "required": true, "type": "integer", "format": "int64"}],
        "responses": {"200": {"description": "Deleted"}}
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "id": {"type": "integer", "format": "int64"},
        "user_name": {"type": "string"},
        "email": {"type": "string"},
        "is_admin": {"type": "boolean"},
        "created_at": {"type": "string"},
        "updated_at": {"type": "string"}
      }
    },
    "UserCreate": {
      "type": "object",
      "required": ["user_name", "email", "password"],
      "properties": {
        "user_name": {"type": "string"},
        "email": {"type": "string"},
        "password": {"type": "string"},
        "is_admin": {"type": "boolean"}
      }
    },
    "UserUpdate": {
      "type": "object",
      "properties": {
        "user_name": {"type": "string"},
        "email": {"type": "string"},
        "password": {"type": "string"},
        "is_admin": {"type": "boolean"}
      }
    },

    "Category": {"type": "object", "properties": {"id": {"type": "integer", "format": "int64"}, "name": {"type": "string"}, "product_id": {"type": "integer", "format": "int64"}}},
    "Color": {"type": "object", "properties": {"id": {"type": "integer", "format": "int64"}, "name": {"type": "string"}, "product_id": {"type": "integer", "format": "int64"}}},
    "Size": {"type": "object", "properties": {"id": {"type": "integer", "format": "int64"}, "name": {"type": "string"}, "product_id": {"type": "integer", "format": "int64"}}},

    "Product": {
      "type": "object",
      "properties": {
        "id": {"type": "integer", "format": "int64"},
        "title": {"type": "string"},
        "description": {"type": "string"},
        "price": {"type": "integer", "format": "int64"},
        "categories": {"type": "array", "items": {"$ref": "#/definitions/Category"}},
        "color": {"type": "array", "items": {"$ref": "#/definitions/Color"}},
        "image_url": {"type": "string"},
        "size": {"type": "array", "items": {"$ref": "#/definitions/Size"}},
        "in_stock": {"type": "integer", "format": "int64"},
        "created_at": {"type": "string"},
        "updated_at": {"type": "string"}
      }
    },
    "ProductCreate": {
      "type": "object",
      "required": ["title", "description", "price", "image_url"],
      "properties": {
        "title": {"type": "string"},
        "description": {"type": "string"},
        "price": {"type": "integer", "format": "int64"},
        "categories": {"type": "array", "items": {"$ref": "#/definitions/Category"}},
        "color": {"type": "array", "items": {"$ref": "#/definitions/Color"}},
        "image_url": {"type": "string"},
        "size": {"type": "array", "items": {"$ref": "#/definitions/Size"}},
        "in_stock": {"type": "integer", "format": "int64"}
      }
    },
    "ProductUpdate": {
      "type": "object",
      "properties": {
        "title": {"type": "string"},
        "description": {"type": "string"},
        "price": {"type": "integer", "format": "int64"},
        "categories": {"type": "array", "items": {"$ref": "#/definitions/Category"}},
        "color": {"type": "array", "items": {"$ref": "#/definitions/Color"}},
        "image_url": {"type": "string"},
        "size": {"type": "array", "items": {"$ref": "#/definitions/Size"}},
        "in_stock": {"type": "integer", "format": "int64"}
      }
    },

    "OrderQty": {"type": "object", "properties": {"id": {"type": "integer", "format": "int64"}, "cart_id": {"type": "integer", "format": "int64"}, "quantity": {"type": "integer", "format": "int64"}, "order_id": {"type": "integer", "format": "int64"}}},
    "Address": {"type": "object", "properties": {"id": {"type": "integer", "format": "int64"}, "addresses": {"type": "string"}, "order_id": {"type": "integer", "format": "int64"}}},

    "Order": {
      "type": "object",
      "properties": {
        "id": {"type": "integer", "format": "int64"},
        "user_id": {"type": "integer", "format": "int64"},
        "order_quantity": {"type": "array", "items": {"$ref": "#/definitions/OrderQty"}},
        "Amount": {"type": "integer", "format": "int64"},
        "address": {"$ref": "#/definitions/Address"},
        "status": {"type": "string"},
        "created_at": {"type": "string"},
        "updated_at": {"type": "string"}
      }
    },
    "OrderCreate": {
      "type": "object",
      "required": ["Amount"],
      "properties": {
        "Amount": {"type": "integer", "format": "int64"},
        "order_quantity": {"type": "array", "items": {"$ref": "#/definitions/OrderQty"}},
        "address": {"$ref": "#/definitions/Address"}
      }
    },
    "OrderUpdate": {
      "type": "object",
      "properties": {
        "Amount": {"type": "integer", "format": "int64"},
        "order_quantity": {"type": "array", "items": {"$ref": "#/definitions/OrderQty"}},
        "status": {"type": "string"},
        "address": {"$ref": "#/definitions/Address"}
      }
    },

    "LoginRequest": {
      "type": "object",
      "required": ["email", "password"],
      "properties": {"email": {"type": "string"}, "password": {"type": "string"}}
    },
    "AuthResponse": {
      "type": "object",
      "properties": {
        "status": {"type": "boolean"},
        "message": {"type": "string"},
        "token": {"type": "string"},
        "user": {"$ref": "#/definitions/User"}
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
