package types

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

type BaseLevelProofData struct {
	*BatchedProofData
	BlockHeaders []ChainBlockHeader
}

type BatchedProofData struct {
	EndBlockHeight    uint32
	IncomingBlockHash chainhash.Hash
	OutgoingBlockHash chainhash.Hash
	FirstBlockHash    chainhash.Hash
	Bits              ChainBits
	FirstBlockTime    uint32
	LastBlockTime     uint32
	BlockCount        uint32
}
