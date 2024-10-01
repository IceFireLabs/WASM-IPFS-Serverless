package wif

import (
	"github.com/gofiber/fiber/v2"
)

// NewHTTPServer creates a new HTTP server instance using the provided fiber.Config.
// It initializes the Fiber app with the given configuration and assigns it to the WIS instance.
func (WIS *WIS) NewHTTPServer(fc fiber.Config) *fiber.App {
	WIS.HTTPServer = fiber.New(fc)
	return WIS.HTTPServer
}

// HTTPServerListen starts listening for incoming HTTP requests on the specified address.
// It returns an error if the server fails to start or encounters any issues during the listening process.
func (WIS *WIS) HTTPServerListen(addr string) error {
	return WIS.HTTPServer.Listen(addr)
}
