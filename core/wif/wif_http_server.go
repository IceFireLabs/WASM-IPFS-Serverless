package wif

import "github.com/gofiber/fiber/v2"

func (WIS *WIS) NewHTTPServer(fc fiber.Config) *fiber.App {
	// NewHTTPServer creates a new HTTP server instance using the provided fiber.Config.
	WIS.HTTPServer = fiber.New(fc)
	return WIS.HTTPServer
}

func (WIS *WIS) HTTPServerListen(addr string) error {
	// HTTPServerListen starts listening for incoming HTTP requests on the specified address.
	// It returns an error if the server fails to start or encounters any issues during the listening process.
	return WIS.HTTPServer.Listen(addr)
}
