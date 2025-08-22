cd ../circuit && rm -f *.so *.h && go build -tags="testnet4" -o libzkbtc.so -buildmode=c-shared api.go
\cp *.so *.h ../cgo
cd ../cgo
go test -tags="testnet4" -run TestBtcBaseProve
go test -tags="testnet4" -run TestBtcMiddleProve
