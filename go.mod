module github.com/james-ray/recovery-tool

go 1.20

//replace github.com/btcsuite/btcutil v0.0.0-20191219182022-e17c9730c422 => github.com/btcsuite/btcd/btcutil v1.1.3

replace github.com/btcsuite/btcutil/hdkeychain v0.0.0-20191219182022-e17c9730c422 => github.com/btcsuite/btcd/btcuti/hdkeychain v1.1.3

//replace github.com/btcsuite/btcd/btcec v0.0.0-20191219182022-e17c9730c422 => github.com/btcsuite/btcd/btcec/v2 v2.2.1
replace github.com/btcsuite/btcd/btcec => ./package/github.com/btcsuite/btcd/btcec/v1

replace github.com/btcsuite/btcd/btcec/v2 => ./package/github.com/btcsuite/btcd/btcec/v2

require (
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.3.1
	golang.org/x/crypto v0.14.0
	github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect
    github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
)
