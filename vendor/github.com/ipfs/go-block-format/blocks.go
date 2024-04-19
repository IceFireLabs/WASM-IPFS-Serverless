// Deprecated: Use github.com/ipfs/go-libipfs/blocks instead.
package blocks

import (
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-libipfs/blocks"
)

// Deprecated: Use blocks.ErrWrongHash instead.
var ErrWrongHash = blocks.ErrWrongHash

// Deprecated: Use blocks.Block instead.
type Block = blocks.Block

// Deprecated: Use blocks.BasicBlock instead.
type BasicBlock = blocks.BasicBlock

// Deprecated: Use blocks.NewBlock instead.
func NewBlock(data []byte) *blocks.BasicBlock {
	return blocks.NewBlock(data)
}

// Deprecated: Use blocks.NewBlockWithCid instead.
func NewBlockWithCid(data []byte, c cid.Cid) (*blocks.BasicBlock, error) {
	return blocks.NewBlockWithCid(data, c)
}
