package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func sayHelloFiber(c *fiber.Ctx) error {
	params := c.Body()

	pluginInst, err := WIS.WISPlugin.GetPluginByName("IPFS-hello")
	pluginInst.Mutex.Lock()

	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return c.SendString(err.Error())
	}

	_, out, err := pluginInst.Plugin.Call("say_hello", params) // hello function

	pluginInst.Mutex.Unlock() //unlock wasm runtime resource

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusConflict)
		return c.SendString(err.Error())
	} else {
		c.Status(http.StatusOK)
		return c.SendString(string(out))
	}

}
