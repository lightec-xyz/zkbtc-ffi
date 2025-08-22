module github.com/lightec-xyz/zkbtcCircuitApi

go 1.24.0

toolchain go1.24.3

replace (
	github.com/consensys/gnark => github.com/lightec-xyz/gnark v0.0.0-20250407032304-d7a29c9b22cb
	github.com/consensys/gnark-crypto => github.com/lightec-xyz/gnark-crypto v0.0.0-20250403072623-6d10578c6e27
	github.com/lightec-xyz/RLPark => ../../RLPark
	github.com/lightec-xyz/btc_provers => ../../btc_provers
	github.com/lightec-xyz/gMPTark => ../../gMPTark
	github.com/lightec-xyz/provers => ../../provers
)

require (
	github.com/btcsuite/btcd v0.24.0 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	github.com/consensys/gnark v0.12.0 // indirect
	github.com/consensys/gnark-crypto v0.17.1-0.20250326164229-5fd6610ac2a1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/lightec-xyz/btc_provers v0.0.0-00010101000000-000000000000
	github.com/lightec-xyz/chainark v0.5.7 // indirect
	github.com/lightec-xyz/common v0.2.9 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/ybbus/jsonrpc/v3 v3.1.5 // indirect
)

require (
	github.com/bits-and-blooms/bitset v1.20.0 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/btcsuite/btcd/btcutil v1.1.5 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/consensys/bavard v0.1.31-0.20250314194434-b30d4344e6d4 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.1 // indirect
	github.com/ethereum/go-ethereum v1.13.5 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/google/pprof v0.0.0-20240727154555-813a5fbdbec8 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/prysmaticlabs/fastssz v0.0.0-20241008181541-518c4ce73516 // indirect
	github.com/prysmaticlabs/go-bitfield v0.0.0-20240328144219-a1caa50c3a1e // indirect
	github.com/prysmaticlabs/gohashtree v0.0.4-beta.0.20240624100937-73632381301b // indirect
	github.com/prysmaticlabs/prysm/v5 v5.2.0 // indirect
	github.com/ronanh/intcomp v1.1.0 // indirect
	github.com/thomaso-mirodin/intmath v0.0.0-20160323211736-5dc6d854e46e // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/exp v0.0.0-20241217172543-b2144cdd0a67 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
