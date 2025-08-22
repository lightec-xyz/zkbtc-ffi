export btcUrl="http://127.0.0.1:12000/"
export btcUser=""
export btcPwd=""
export btcSoPath="./lib/btcprovers.so"
export btcSetup="./circuits/btc_provers.so"
export ethSoPath="./lib/btcprovers.so"
export ethSetup="./circuits/btc_provers.so"

go test -tags="testnet4" -run TestCircuit_BtcBaseProve

