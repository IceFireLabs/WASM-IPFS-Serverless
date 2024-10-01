package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// sayHelloFiber handles the HTTP POST request to execute a WASM plugin function.
func sayHelloFiber(c *fiber.Ctx) error {
	// Extract the request body parameters.
	params := c.Body()

	// Retrieve the WASM plugin instance by name.
	pluginInst, err := WIS.WISPlugin.GetPluginByName("IPFS-hello")
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return c.SendString(err.Error())
	}

	// Lock the plugin instance to ensure thread-safe access.
	pluginInst.Mutex.Lock()
	defer pluginInst.Mutex.Unlock() // Ensure the lock is released after function execution.

	// Call the WASM plugin function "say_hello" with the provided parameters.
	_, out, err := pluginInst.Plugin.Call("say_hello", params)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusConflict)
		return c.SendString(err.Error())
	} else {
		c.Status(http.StatusOK)
		return c.SendString(string(out))
	}
}
