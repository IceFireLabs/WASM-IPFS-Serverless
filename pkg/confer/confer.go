package confer

import (
	"errors"
	"os"

	utils "github.com/BlockCraftsman/WASM-IPFS-Serverless/utils"
	"github.com/go-chassis/go-archaius"
)

const (
	NET_MODEL_NETPOLL    = "NETPOLL"
	NET_MODEL_RAWEPOLL   = "RAWEPOLL"
	APP_TYPE_WASM_WORKER = "WASM-IPFS-Serverless-WORKER"
)

func InitConfig(appName, confFilePath string) (*Confer, error) {
	// Initialize the archaius with the required configuration file
	if err := archaius.Init(
		archaius.WithRequiredFiles([]string{confFilePath}),
	); err != nil {
		return nil, err
	}

	// Create a new Confer instance
	conf := &Confer{}

	// Unmarshal the configuration into conf.Opts
	if err := archaius.UnmarshalConfig(&conf.Opts); err != nil {
		return nil, err
	}

	// Return the pointer to the conf and nil error
	return conf, nil
}

var _confer *Confer

// GetNewConfer is core function of Service Car : get a Confer entity.
func GetNewConfer(appType, confFileURI string) (confer *Confer, err error) {
	if _confer != nil {
		return _confer, nil
	}
	confer = &Confer{}

	//parse configure from file
	confer.Opts, err = parseYamlFromFile(confFileURI)

	if err != nil {
		return nil, err
	}

	confer.Opts.ApptypeConf = appType

	//NetModel options start
	confer.replaceByEnv(&confer.Opts.NetModelConf)

	if !utils.InArray(confer.Opts.NetModelConf, []string{NET_MODEL_NETPOLL, NET_MODEL_RAWEPOLL}) {
		confer.Opts.NetModelConf = NET_MODEL_NETPOLL
	}
	//NetModel options end

	// Check for Debug parameters: If Debug monitoring is enabled
	if confer.Opts.DebugConf.Enable {
		// If the listening address is an empty string in the scenario where Enable is enabled, it indicates that the configuration is incorrect
		if len(confer.Opts.DebugConf.PprofBindAddr) == 0 {
			err = errors.New("pprof network listening address cannot be empty")
			return
		}
	}

	_confer = confer
	return _confer, nil
}

// eg: the value of confName is "UserName",then get enviroment var by key "UserName",final replace config name by pointer.
func (*Confer) replaceByEnv(confName *string) {
	// Get the value of the environment variable specified by 'conf'
	if s := os.Getenv(*confName); len(s) > 0 {
		// Update the value of 'conf' with the value from the environment variable
		*confName = s
	}
}

// Get global configure object
func Global() *Confer {
	if _confer == nil {
		return &Confer{} // Return a new instance of Confer if _confer is nil
	}
	return _confer // Return the existing instance of Confer
}
