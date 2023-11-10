# HBC Recovery Tool
This tool should run in an offline environment


## Getting started

```
1. install git, go
2. download the project
   git clone git@github.com:james-ray/recovery-tool.git
3. Compile
cd recovery-tool
go mod vendor
go build  -o recovery-tool  main.go

4. Get user passphrase, hbc passphrase, and the RSA private key from coincover, then run 
  user passphrase: The customer should remember his own recovery passphrase
  hbc passphrase: After customer set the recovery passphrase, the App sends the encrypted key share to Hbc, then Hbc will send him
                  his Hbc recovery passphrase by email. 
  privkeyFilePath: Coincover would give a rsa private key in hex format, copy the whole hex string into a file privkeyFile, save the file.

The command description:
recovery-tool parseZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [privkeyFilePath]
eg: recovery-tool parseZipFile './zipTest.zip' '123123' '456456' './privkeyFile'

Change the ./zipTest.zip to the name of the real backup file that coincover gives you.
Change 123123 to the recovery phrase set in the App.
Change 456456 to the Hbc password in the email.

This should print something like this:
./recovery-tool parseZipFile './zipTest.zip' '123123' '456456' './privkeyFile'
map[chaincodes:["123123","234234","456456"] hbc.encrypted.0:223344 hbc.encrypted.1:445566 pubkeys:["123123ab","234234ab","456456ab"] user:112233]

I'll update the tool to derive extended private key.
```
