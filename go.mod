module github.com/james-ray/recovery-tool

go 1.20

//replace github.com/btcsuite/btcutil v0.0.0-20191219182022-e17c9730c422 => github.com/btcsuite/btcd/btcutil v1.1.3

replace github.com/btcsuite/btcutil/hdkeychain v0.0.0-20191219182022-e17c9730c422 => github.com/btcsuite/btcd/btcuti/hdkeychain v1.1.3

//replace github.com/btcsuite/btcd/btcec v0.0.0-20191219182022-e17c9730c422 => github.com/btcsuite/btcd/btcec/v2 v2.2.1
replace github.com/btcsuite/btcd/btcec => ./package/github.com/btcsuite/btcd/btcec/v1

replace github.com/btcsuite/btcd/btcec/v2 v2.1.3 => ./package/github.com/btcsuite/btcd/btcec/v2

require (
		github.com/btcsuite/btcd v0.23.0
    	github.com/btcsuite/btcd/btcec v0.0.0-00010101000000-000000000000
    	github.com/btcsuite/btcd/btcec/v2 v2.1.3
    	github.com/btcsuite/btcd/btcutil v1.1.3
    	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
    		github.com/decred/dcrd/dcrec/edwards/v2 v2.0.1
	github.com/golang/protobuf v1.4.2
	github.com/james-ray/hcd v0.0.0-20230524063416-4917c422bd33
	github.com/otiai10/primes v0.0.0-20210501021515-f1b2be525a11
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.3.1
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.14.0
	google.golang.org/protobuf v1.23.0
)

require (
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
