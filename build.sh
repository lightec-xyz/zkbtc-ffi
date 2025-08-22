cd circuit && rm -f *.so *.h &&go build -tags="testnet4" -o libzkbtc.so -buildmode=c-shared api.go
\cp -f *.so *.h ..
\cp -f *.so *.h ../cgo
cd ..
go test -tags="testnet4" -run TestBtcBaseProve
