# HBC Recovery Tool
This tool should run in an offline environment


## Getting started

```
1、Compile
go mod vendor
go build  -o recovery-tool  main.go

2、Get user passphrase, hbc passphrase, and the RSA private key from coincover, then run 
recovery-tool parseZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [privkeyFilePath]

This should print three mnemonics

I'll update the tool to derive extended private key.
```
