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
  user passphrase: The customer should remember his own recovery passphrase which he set in the App.
  hbc passphrase: After customer set the recovery passphrase, the App sends the encrypted key share to Hbc, then Hbc will send him
                  his Hbc recovery passphrase by email. 
  privkeyFilePath: Coincover would give a rsa private key in hex format, copy the whole hex string into a file privkeyFile, save the file.

The command description:
a. parseZipFile
recovery-tool parseZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [privkeyFilePath]
eg: recovery-tool parseZipFile './zipTest.zip' '123123' '456456' './privkeyFile'

Change the ./zipTest.zip to the name of the real backup file that coincover gives you.
Change 123123 to the recovery phrase set in the App.
Change 456456 to the Hbc password in the email.

This should print something like this:
map[chaincodes:["123123","234234","456456"] hbc.encrypted.0:223344 hbc.encrypted.1:445566 pubkeys:["123123ab","234234ab","456456ab"] user:112233]

This outputs three arrays:
hbc.encrypted.0/1 and user are private key slices. 
pubkeys is an array of the public key slices.
chaincodes is also an array.

After you get the metadata plaintext, you can use the following three ways to recover the extended child private key(s).

b.deriveChildPrivateKey
recovery-tool deriveChildPrivateKey [metadataFilePath] [derivePath]
eg: recovery-tool deriveChildPrivateKey './metadata.json' '81/0/1/60/2'

the derivePath is used for derive the child private key, the path is like '81/0/46/0/0'

c.recovery-tool deriveCsvFile [metadataFilePath] [csvFilePath]
This command is used for batch recovering. You can download the csv file from hbc backend.

d.Use the UI
Follow the three steps to generate the derive.html, it is a UI for extended child private key.
1. GOOS=js GOARCH=wasm go build -tags=osusergo -o recovery-tool.wasm helpers/helper.go
2. If you are Ubuntu/Linux:  base64 -w 0 recovery-tool.wasm  >  wasmstr.txt
   If you are Mac: base64 -b -i recovery-tool.wasm -o wasmstr.txt
3. awk 'NR==FNR{a[i++]=$0;next} /var base64String = ".*";/{sub(/var base64String = ".*";/, "var base64String = \""a[0]"\";")}1' wasmstr.txt derive_template.html > tmp && mv tmp derive.html
You should install awk before execute this step. 

Explain the three steps:
step1 generates recovery-tool.wasm
step2 generates the base64 string of recovery-tool.wasm
step3 replaces the content in template html to the real base64 string

Now you can double click the derive.html.

The text boxes or drop down box in the UI:
Metadata: Paste the metadata json string in this text box.
WalletType: Fund Wallet or Api Wallet.
VaultIndex: If it is Fund Wallet, it starts from 1. Else if it is Api Wallet, this is fixed 0.
Chain: Choose the destination chain from the list.
SubIndex: If it is Fund Wallet, this is fixed 0. Else if it is Api Wallet, it starts from 1.

```
