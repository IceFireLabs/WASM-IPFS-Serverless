package wif

import (
	"context"
	"fmt"
	"sync"

	extism "github.com/extism/go-sdk"
)

// WISPlugins represents a collection of wasm plugins.
type WISPlugins struct {
	RWISutex sync.RWMutex
	Plugins  map[string]*WISPlugin
}

// WISPlugin Struct: actor pool
type WISPlugin struct {
	Name   string
	Plugin extism.Plugin
	Mutex  sync.Mutex
}

func newWISPlugin(pluginName string, plugin extism.Plugin) *WISPlugin {
	return &WISPlugin{Name: pluginName, Plugin: plugin}
}

// NewWISPlugins creates a new instance of WISPlugins.
func NewWISPlugins() (wisPlugins *WISPlugins) {
	wisPlugins = &WISPlugins{
		Plugins: make(map[string]*WISPlugin),
	}

	return
}

// AddPlugin adds a new wasm plugin to the WISPlugins collection.
// If the plugin with the given name already exists, it will return an error unless the force flag is set to true.
// If force is true, the existing plugin will be replaced with the new plugin.
func (WISPlugins *WISPlugins) AddPlugin(pluginName string, plugin extism.Plugin, force bool) (err error) {
	WISPlugins.RWISutex.Lock()
	defer WISPlugins.RWISutex.Unlock()

	if _, ok := WISPlugins.Plugins[pluginName]; !ok {
		WISp := newWISPlugin(pluginName, plugin)
		WISPlugins.Plugins[pluginName] = WISp
		return nil
	}

	// If plugin is extism, judge if force replace is required
	if !force {
		// If force is not enabled, return an error
		return fmt.Errorf("failed to add the wasm plugin(%s), the plug-in already exists", pluginName)
	}

	// If force is true, then force replace the wasm plugin by name.
	WISp := newWISPlugin(pluginName, plugin)
	WISPlugins.Plugins[pluginName] = WISp
	return nil
}

// GetPluginByName returns the wasm plugin with the given name from the WISPlugins collection.
// It acquires a read lock on the RWISutex to ensure thread safety.
// If the plugin with the given name exists, it returns the plugin and a nil error.
// If the plugin with the given name does not exist, it returns an empty extism.Plugin and an error indicating that the plugin does not exist.
func (WISPlugins *WISPlugins) GetPluginByName(pluginName string) (*WISPlugin, error) {
	WISPlugins.RWISutex.RLock()
	defer WISPlugins.RWISutex.RUnlock()

	if plugin, ok := WISPlugins.Plugins[pluginName]; ok {
		return plugin, nil
	}

	return nil, fmt.Errorf("no wasm plugin named %s", pluginName)
}

// Create a new wasm plugin using the extism.NewPlugin function
func (WISPlugins *WISPlugins) NewWISPlugin(ctx context.Context,
	manifest extism.Manifest,
	config extism.PluginConfig,
	functions []extism.HostFunction) (*extism.Plugin, error) {

	return extism.NewPlugin(ctx, manifest, config, nil)
}
