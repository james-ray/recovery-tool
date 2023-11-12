module github.com/james-ray/recovery-tool

go 1.18

//replace github.com/btcsuite/btcutil v0.0.0-20191219182022-e17c9730c422 => github.com/btcsuite/btcd/btcutil v1.1.3

replace github.com/btcsuite/btcutil/hdkeychain v0.0.0-20191219182022-e17c9730c422 => github.com/btcsuite/btcd/btcuti/hdkeychain v1.1.3

//replace github.com/btcsuite/btcd/btcec v0.0.0-20191219182022-e17c9730c422 => github.com/btcsuite/btcd/btcec/v2 v2.2.1
replace github.com/btcsuite/btcd/btcec => ./package/github.com/btcsuite/btcd/btcec/v1

replace github.com/btcsuite/btcd/btcec/v2 => ./package/github.com/btcsuite/btcd/btcec/v2

require (
	github.com/btcsuite/btcd/btcec v0.0.0-00010101000000-000000000000
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/golang/protobuf v1.5.3
	github.com/james-ray/hcd v0.0.0-20230524063416-4917c422bd33
	github.com/otiai10/primes v0.0.0-20210501021515-f1b2be525a11
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.3.1
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.4
	golang.org/x/crypto v0.15.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
