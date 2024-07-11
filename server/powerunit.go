package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (pu *Powerunit) InitPUEndpoint(path string) {
	// Make proxy requests while following redirects
	pu.APP.Get(path, func(c *fiber.Ctx) error {
		fmt.Println("pu endpoint called")
		return nil
	})
}
