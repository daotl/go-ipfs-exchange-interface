// Package exchange defines the IPFS exchange interface
package exchange

import (
	"context"
	"io"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
)

// Channel represents a exchange channel within which blocks can be exchanged
// possibly with access control.
type Channel string

const (
	// PublicChannel represents the default public exchange channel without access control.
	PublicChannel Channel = ""
)

// Interface defines the functionality of the IPFS block exchange protocol.
type Interface interface { // type Exchanger interface
	Fetcher

	// TODO Should callers be concerned with whether the block was made
	// available on the network?
	HasBlock(blocks.Block) error

	HasBlockInChannel(Channel, blocks.Block) error

	IsOnline() bool

	io.Closer
}

// Fetcher is an object that can be used to retrieve blocks
type Fetcher interface {
	// GetBlock returns the public block associated with a given key.
	GetBlock(context.Context, cid.Cid) (blocks.Block, error)
	// GetBlockFromChannel returns the block associated with a given key from
	// the specified exchange channel.
	GetBlockFromChannel(context.Context, Channel, cid.Cid) (blocks.Block, error)
	// GetBlocks returns a channel where the caller may receive public blocks
	// that correspond to the provided keys.
	GetBlocks(context.Context, []cid.Cid) (<-chan blocks.Block, error)
	// GetBlocksFromChannel returns a channel where the caller may receive blocks
	// that correspond to the provided keys form the specified channel.
	GetBlocksFromChannel(context.Context, Channel, []cid.Cid) (<-chan blocks.Block, error)
}

// SessionExchange is an exchange.Interface which supports
// sessions.
type SessionExchange interface {
	Interface
	NewSession(context.Context) Fetcher
}
