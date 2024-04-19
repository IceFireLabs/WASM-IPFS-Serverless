package wif

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/BlockCraftsman/WASM-IPFS-Serverless/pkg/confer"
	"github.com/gofiber/fiber/v2"
)

// WIS is the core data structure of WIS and is used to extend and integrate other structures and behaviors.
type WIS struct {
	Ctx        context.Context // Global context of WIS
	OfflineCtx context.Context // OfflineCtx is used to propagate WIS offline signals
	Confer     *confer.Confer  // Confer entity of WIS

	ShutdownWG     sync.WaitGroup // Wait for all cleanup operations to complete on WIS shutdown
	InFlowProtocol atomic.Value   // Inbound traffic protocol

	HTTPServer *fiber.App

	WISPlugin *WISPlugins

	sync.RWMutex
}

// NewWIS creates a new instance of WIS with the specified context, offline context, and confer.
// It ensures that only one instance of WIS is created and returns the existing instance if it already exists.
func NewWIS(ctx, offlineCtx context.Context, conf *confer.Confer) *WIS {

	_WIS := &WIS{Ctx: ctx, Confer: conf, OfflineCtx: offlineCtx}

	_WIS.WISPlugin = NewWISPlugins()
	return _WIS
}
