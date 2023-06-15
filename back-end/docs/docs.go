// Code generated by swaggo/swag. DO NOT EDIT.

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
            "name": "API Support",
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
        "/admin/": {
            "post": {
                "description": "create admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "create admin",
                "parameters": [
                    {
                        "description": "admin 's username and password",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Administrator"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/admin/login/": {
            "post": {
                "description": "admin login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Admin Login",
                "parameters": [
                    {
                        "description": "Admin login with username and password",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Administrator"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/booking/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "book seat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "book seat",
                "parameters": [
                    {
                        "description": "Book a seat by giving seat_id and user_id",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/booking/bookingHistory/{userID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get booking history",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "get booking history",
                "parameters": [
                    {
                        "type": "string",
                        "description": "booking history",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/booking/deleteBooking": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete booking by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "delete booking",
                "parameters": [
                    {
                        "description": "booking id",
                        "name": "booking_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/booking/getBookingByID/{booking_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get booking by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "get booking by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "booking id",
                        "name": "booking_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/booking/getBookingByUserID/{user_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get booking by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "get booking by user ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/booking/updateBooking": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update booking , change isSigned",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Booking"
                ],
                "summary": "update booking",
                "parameters": [
                    {
                        "description": "booking information",
                        "name": "booking",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Booking"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/room/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all room information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "get all room",
                "responses": {}
            }
        },
        "/room/auth/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get room by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "get room",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Create Room by giving room information",
                        "name": "admin",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/room/auth/createRoom": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "create room",
                "parameters": [
                    {
                        "description": "Create Room by giving room information",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Room"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/room/auth/deleteRoom": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete room by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "delete room",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delete Room by giving room id",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/room/auth/updateRoom": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update room",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room"
                ],
                "summary": "update room",
                "parameters": [
                    {
                        "description": "Update Room by giving room id and room details",
                        "name": "room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Room"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/seat/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all seats",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seat"
                ],
                "summary": "get seat",
                "responses": {}
            }
        },
        "/seat/auth/createSeat": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create seat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seat"
                ],
                "summary": "create seat",
                "parameters": [
                    {
                        "description": "Create Seat by giving seat information",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Seat"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/seat/auth/deleteSeat": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete seat by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seat"
                ],
                "summary": "delete seat",
                "parameters": [
                    {
                        "description": "Delete Seat by giving seat id",
                        "name": "seat_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/seat/auth/updateSeat": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update seat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seat"
                ],
                "summary": "update seat",
                "parameters": [
                    {
                        "description": "Update Seat by giving seat information",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Seat"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/seat/getSeatByRoomID/{room_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get seat by room id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seat"
                ],
                "summary": "get seat by room id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get Seats by giving room id",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/seat/{seat_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get seat by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seat"
                ],
                "summary": "get seat by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get Seat by giving seat id",
                        "name": "seat_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/": {
            "post": {
                "description": "create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "create user",
                "parameters": [
                    {
                        "description": "user 's username and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/auth/deleteUser": {
            "post": {
                "description": "delete user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "delete user",
                "parameters": [
                    {
                        "description": "user id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/auth/getUserByID/{user_id}": {
            "get": {
                "description": "get user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/auth/getUserByUsername/{username}": {
            "get": {
                "description": "get user by username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get user by username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/auth/password": {
            "post": {
                "description": "update password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "update password",
                "parameters": [
                    {
                        "description": "userID and password",
                        "name": "userinfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/auth/updateUser": {
            "post": {
                "description": "update user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "update user",
                "parameters": [
                    {
                        "description": "user information",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/login": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "user login",
                "parameters": [
                    {
                        "description": "user 's username and password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.Administrator": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Booking": {
            "type": "object",
            "properties": {
                "bookingTime": {
                    "description": "booking time, after it 15 min will auto cancel booking and free seat and notify student and record one default",
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "duration": {
                    "description": "max 4h, unit hour",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "isSigned": {
                    "description": "0 represent waiting, 1 represent attend, 2 represent delay",
                    "type": "integer"
                },
                "roomID": {
                    "type": "integer"
                },
                "seatID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.Room": {
            "type": "object"
        },
        "models.Seat": {
            "type": "object"
        },
        "models.User": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "updatedAt": {
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "10.177.88.190:8800",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "iBooking",
	Description:      "iBooking back-end api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}