package confer

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	// Test case 1: _confer is already initialized
	// _confer = &Confer{}
	// conf, err := InitConfig("app1", "config.yaml")
	// assert.NoError(t, err)
	// assert.Equal(t, _confer, conf)

	// // Test case 2: Successful initialization
	// _confer = nil
	// conf, err = InitConfig("app2", "config.yaml")
	// assert.NoError(t, err)
	// assert.NotNil(t, conf)
	// assert.Equal(t, "app2", "app2")

	// // Test case 3: Error initializing archaius
	// _confer = nil
	// conf, err = InitConfig("app3", "nonexistent.yaml")
	// assert.Error(t, err)
	// assert.Nil(t, conf)

	// // Test case 4: Error unmarshaling configuration
	// _confer = nil
	// conf, err = InitConfig("app4", "invalid.yaml")
	// assert.Error(t, err)
	// assert.Nil(t, conf)
}
