package server

import (
	"github.com/gofiber/fiber/v2"
)

type Powerunit struct { // rename to Server
	APP                   *fiber.App
	STARTUP_ERROR         bool
	STARTUP_ERROR_MESSAGE string
}

type PU_Response struct {
	Keys []struct {
		Kid     string   `json:"kid"`
		Kty     string   `json:"kty"`
		Alg     string   `json:"alg"`
		Use     string   `json:"use"`
		N       string   `json:"n"`
		E       string   `json:"e"`
		X5C     []string `json:"x5c"`
		X5T     string   `json:"x5t"`
		X5TS256 string   `json:"x5t#S256"`
	} `json:"keys"`
}
