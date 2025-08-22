package types

import (
	"encoding/binary"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

const (
	BlockHeaderLen  = 80
	PrevBlockOffset = 4
	HashLen         = 32
	BitsOffset      = 72
	BitsLen         = 4
	TimestampLen    = 4

	BlockHeaderStrSize = BlockHeaderLen * 2
	PrevBlockStrOffset = 4 * 2
	HashStrSize        = chainhash.HashSize * 2 // hexSize
)

type ChainBlockHeader [BlockHeaderLen]byte // encoding order
type ChainBits [BitsLen]byte               // encoding order
type ChainTimestamp [TimestampLen]byte     // encoding order

func ChainBitsToUint32(chainBits ChainBits) uint32 {
	return binary.LittleEndian.Uint32(chainBits[:])
}

func ChainTimeToUint32(chainTime ChainTimestamp) uint32 {
	return binary.LittleEndian.Uint32(chainTime[:])
}
