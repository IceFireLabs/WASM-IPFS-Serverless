package main

import (
	"fmt"
	"log"
	"net/http"

	_ "go.uber.org/automaxprocs"
)

// BuildDate: Binary file compilation time
// BuildVersion: Binary compiled GIT version
var (
	BuildDate    string
	BuildVersion string
)

const (
	// APP_NAME_WASM_WORKER: Name of the WASM worker application
	APP_NAME_WASM_WORKER = "WASM-IPFS-Serverless-WORKER"
	// APP_CONFIG_ENV_NAME: Environment variable name for the worker configuration
	APP_CONFIG_ENV_NAME = "WASM-IPFS-Serverless_WORKER_CONFIG"
)

// printBanner prints the application banner along with build version and date.
func printBanner() {
	bannerData := `╦ ╦╔═╗╔═╗╔╦╗   ╦╔═╗╔═╗╔═╗   ╔═╗┌─┐┬  ┬┌─┐┬─┐┬  ┌─┐┌─┐┌─┐
║║║╠═╣╚═╗║║║───║╠═╝╠╣ ╚═╗───╚═╗├┤ └┐┌┘├┤ ├┬┘│  ├┤ └─┐└─┐
╚╩╝╩ ╩╚═╝╩ ╩   ╩╩  ╚  ╚═╝   ╚═╝└─┘ └┘ └─┘┴└─┴─┘└─┘└─┘└─┘`
	fmt.Println(bannerData)
	fmt.Println("Build Version: ", BuildVersion, "  Date: ", BuildDate)
}

// startDebug starts the debug server on the specified address.
func startDebug(pprofBind string) {
	log.Printf("%s pprof listen on: %s\n", APP_NAME_WASM_WORKER, pprofBind)
	// Start the HTTP server for pprof
	err := http.ListenAndServe(pprofBind, nil)
	if err != nil {
		// Log the error and return if the server fails to start
		log.Printf("Failed to start pprof server: %v", err)
		return
	}
}
