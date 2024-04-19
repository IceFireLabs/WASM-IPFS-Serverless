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
	APP_NAME_WASM_WORKER = "WASM-IPFS-Serverless-WORKER" //wasm worker
	APP_CONFIG_ENV_NAME  = "WASM-IPFS-Serverless_WORKER_CONFIG"
)

func printBanner() {
	bannerData := `╦ ╦╔═╗╔═╗╔╦╗   ╦╔═╗╔═╗╔═╗   ╔═╗┌─┐┬  ┬┌─┐┬─┐┬  ┌─┐┌─┐┌─┐
║║║╠═╣╚═╗║║║───║╠═╝╠╣ ╚═╗───╚═╗├┤ └┐┌┘├┤ ├┬┘│  ├┤ └─┐└─┐
╚╩╝╩ ╩╚═╝╩ ╩   ╩╩  ╚  ╚═╝   ╚═╝└─┘ └┘ └─┘┴└─┴─┘└─┘└─┘└─┘`
	fmt.Println(bannerData)
	fmt.Println("Build Version: ", BuildVersion, "  Date: ", BuildDate)
}

func startDebug(pprofBind string) {
	log.Printf("%s pprof listen on: %s\n", APP_NAME_WASM_WORKER, pprofBind)
	err := http.ListenAndServe(pprofBind, nil)
	if err != nil {
		return
	}
}
