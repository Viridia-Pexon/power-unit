package server

import (
	"github.com/gofiber/fiber/v2"
)

/*
exposes health enpoint which also reflects token refresh status
---
bietet einen health Endpunkt an der auch Refplekitert ob es aktuell ein Problem mit dem Aktualisieren des GIS-Tokens gibt
*/
func (pu *Powerunit) InitHealthEndpoint(path string) {
	// Make proxy requests while following redirects
	pu.APP.Get(path, func(c *fiber.Ctx) error {
		if pu.STARTUP_ERROR {
			return fiber.NewError(fiber.StatusInternalServerError, pu.STARTUP_ERROR_MESSAGE)
		} else {
			return c.SendString("OK")
		}

	})
}
