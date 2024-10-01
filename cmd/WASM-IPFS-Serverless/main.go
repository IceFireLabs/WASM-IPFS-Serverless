package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BlockCraftsman/WASM-IPFS-Serverless/core/wif"
	"github.com/BlockCraftsman/WASM-IPFS-Serverless/pkg/confer"
	"github.com/BlockCraftsman/WASM-IPFS-Serverless/pkg/ipfs"
	utils "github.com/BlockCraftsman/WASM-IPFS-Serverless/utils"
	extism "github.com/extism/go-sdk"
	"github.com/gofiber/fiber/v2"
	"github.com/tetratelabs/wazero"
	"github.com/urfave/cli"
)

// WIS is a global variable to hold the Wasm IPFS Serverless instance.
var WIS *wif.WIS

// init sets the log flags to include file name and line number.
func init() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
}

// main is the entry point of the application.
func main() {
	// Create a new CLI application.
	app := cli.NewApp()

	// Set the application name.
	app.Name = APP_NAME_WASM_WORKER

	// Define CLI flags.
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "config, c",
			Usage:    "init config file",
			EnvVar:   APP_CONFIG_ENV_NAME,
			Required: false,
		},
	}

	// Before function to set log flags and print banner.
	app.Before = func(ctx *cli.Context) error {
		log.SetFlags(log.Ldate | log.Ltime | log.Llongfile | log.LstdFlags)
		printBanner()
		return nil
	}

	// Action function to handle the main logic of the application.
	app.Action = func(ctx *cli.Context) error {
		// Remove all arguments except the first one (the program name).
		os.Args = os.Args[:1]

		// Recover from any panics and log the error.
		defer func() {
			if e := recover(); e != nil {
				log.Printf("app.Action recover %v \n", e)
			}
		}()

		// Get the config file path from the CLI context.
		configFile := ctx.String("config")
		if len(configFile) != 0 && !utils.IsFileExist(configFile) {
			return fmt.Errorf("%s does not exist", configFile)
		}

		// Set the default config file based on the application name.
		if len(configFile) == 0 {
			switch app.Name {
			case APP_NAME_WASM_WORKER:
				configFile = "config/wis_worker.yaml"
			default:
				configFile = "config/config.yaml"
			}
		}

		// Initialize the configuration.
		conf, err := confer.InitConfig(app.Name, configFile)
		if err != nil {
			log.Panicln(err)
		}

		// Start Go pprof debugging if enabled in the configuration.
		if conf.Opts.DebugConf.Enable {
			utils.GoWithRecover(func() {
				startDebug(conf.Opts.DebugConf.PprofBindAddr)
			}, nil)
		}

		// Create a context with cancel for the application.
		c, cancel := context.WithCancel(context.Background())
		offlineCtx, offlineCancel := context.WithCancel(c)
		WIS = wif.NewWIS(c, offlineCtx, conf)

		// Define initialization functions.
		initFunc := []func() error{
			WIS.LoadBasics,
		}

		// Add additional initialization functions based on the application type.
		switch conf.Opts.ApptypeConf {
		case confer.APP_TYPE_WASM_WORKER:
			//initFunc = append(initFunc, nil)
		default:
			panic("run APP type error:" + conf.Opts.ApptypeConf)
		}

		// Execute all initialization functions.
		for k := range initFunc {
			err := initFunc[k]()
			if err != nil {
				log.Println("init function run fall:", k, err)
				panic(err)
			}
		}

		// Create a context for the Wasm engine.
		ctxWasm := context.Background()

		// Configure the Wasm plugin.
		config := extism.PluginConfig{
			ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
			EnableWasi:   true,
		}

		// Create a manifest for the Wasm modules.
		manifest := extism.Manifest{Wasm: []extism.Wasm{}}

		// Load Wasm modules from files if enabled in the configuration.
		if conf.Opts.WASMModulesFiles.Enable {
			for _, wasmPath := range conf.Opts.WASMModulesFiles.WASMFilePaths {
				if _, err := os.Stat(wasmPath); err != nil {
					log.Panicln(err)
				}
				manifest.Wasm = append(manifest.Wasm, extism.WasmFile{Path: wasmPath})
			}
		}

		// Load Wasm modules from IPFS if enabled in the configuration.
		if conf.Opts.WASMModulesIPFS.Enable {
			ipfsC := ipfs.NewIPFSClient(conf.Opts.WASMModulesIPFS.LassieNet.Scheme,
				conf.Opts.WASMModulesIPFS.LassieNet.Host,
				conf.Opts.WASMModulesIPFS.LassieNet.Port)

			for _, wasmCID := range conf.Opts.WASMModulesIPFS.CIDS {
				D, err := ipfs.GetDATAFromIPFSCID(ipfsC, wasmCID)
				if err != nil {
					log.Panicln(err)
				}

				for _, wasmData := range D {
					manifest.Wasm = append(manifest.Wasm, extism.WasmData{Data: wasmData})
				}
			}
		}

		// Create an instance of the Wasm plugin.
		pluginInst, err := WIS.WISPlugin.NewWISPlugin(ctxWasm, manifest, config, nil)
		if err != nil {
			log.Panicln(err)
		}

		// Add the Wasm plugin to the WIS instance.
		err = WIS.WISPlugin.AddPlugin("IPFS-hello", *pluginInst, true)
		if err != nil {
			log.Panicln(err)
		}

		// Create a new HTTP server using Fiber.
		FASSFrame := WIS.NewHTTPServer(fiber.Config{DisableStartupMessage: true})

		// Define a route for the HTTP server.
		FASSFrame.Post("/", sayHelloFiber)

		// Start the HTTP server.
		fmt.Println("http server is listening on:", conf.Opts.NetWork.BindAddress)
		FASSFrame.Listen(conf.Opts.NetWork.BindAddress)

		// Cancel the context and other cleanup tasks.
		_ = cancel
		_ = offlineCtx
		_ = offlineCancel
		_ = initFunc

		return nil
	}

	// Run the CLI application.
	err := app.Run(os.Args)
	if err != nil {
		log.Printf("app.Run error: %v\n", err)
	}
}
