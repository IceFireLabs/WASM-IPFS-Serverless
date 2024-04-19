package confer

import (
	"sync"
)

type Confer struct {
	Mutex sync.RWMutex
	Opts  confS
}

type confS struct {
	ApptypeConf      string           `yaml:"app-type"`           // AppType: WASM_WORKER
	NetModelConf     string           `yaml:"net-model"`          // The network model of HTTP handle,NetPoll(gin) RAWEPOLL(fiber)
	NetWork          NetWorkS         `yaml:"NetWork"`            // TrafficInflow process HTTP traffic network entry configuration
	DebugConf        DebugConfS       `yaml:"debug"`              // Runtime memory debug configuration
	WASMModulesFiles WASMModulesFileS `yaml:"wasm-modules-files"` // Runtime WASM Modules files from paths
	WASMModulesIPFS  WASMModulesIPFSS `yaml:"wasm-modules-ipfs"`  // Runtime WASM Modules files from IPFS
}

type WASMModulesFileS struct {
	Enable        bool     `yaml:"enable"`
	WASMFilePaths []string `yaml:"path"`
}

type LassieNetS struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
}

type WASMModulesIPFSS struct {
	Enable    bool       `yaml:"enable"`
	LassieNet LassieNetS `yaml:"lassie-net"`
	CIDS      []string   `yaml:"cids"`
}

// TrafficInFlowS is used to handle incoming traffic network configuration
type NetWorkS struct {
	BindNetWork  string `yaml:"bind-network"`  //Network transport layer type: TCP | UDP
	ProtocolType string `yaml:"protocol-type"` //Application layer network protocol：HTTP | RESP | QUIC
	BindAddress  string `yaml:"bind-address"`  //Network listening address,indicating where the application will listen for incoming network traffic.
}

// DebugConfS id debug configure options
type DebugConfS struct {
	Enable        bool   `yaml:"enable"`
	PprofBindAddr string `yaml:"pprof-bind-addr"` // address of performance analysis network binding
}
