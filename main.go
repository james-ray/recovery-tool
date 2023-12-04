package main

import (
	"encoding/hex"
	"fmt"
	"github.com/james-ray/recovery-tool/common"
	"github.com/james-ray/recovery-tool/tss/crypto"
	"github.com/james-ray/recovery-tool/tss/tss"
	"math/big"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	if os.Args[1] == "makeZipFile" {
		if len(os.Args) < 8 {
			printUsage()
			os.Exit(1)
		}
		pubkeyBytes, err := os.ReadFile(os.Args[5])
		if err != nil {
			panic(err)
		}
		pubkeyBytes = pubkeyBytes[:len(pubkeyBytes)-1]
		/*fmt.Println(string(pubkeyBytes))
		fmt.Println(len(string(pubkeyBytes)))
		fmt.Println(len(pubkeyBytes))
		fmt.Printf("%x \n", pubkeyBytes)*/
		hbcPasswdBytes, err := hex.DecodeString(os.Args[4])
		if err != nil {
			panic(err)
		}
		_, err = common.MakeZipFile([]byte(os.Args[3]), hbcPasswdBytes, string(pubkeyBytes), os.Args[6], strings.Split(os.Args[7], "|"), strings.Split(os.Args[8], "|"), strings.Split(os.Args[9], "|"), os.Args[2])
		if err != nil {
			panic(err)
		}
	} else if os.Args[1] == "parseZipFile" {
		if len(os.Args) < 6 {
			printUsage()
			os.Exit(1)
		}
		privkeyBytes, err := os.ReadFile(os.Args[5])
		if err != nil {
			panic(err)
		}
		privkeyBytes = privkeyBytes[:len(privkeyBytes)-1]
		/*fmt.Println(string(privkeyBytes))
		fmt.Println(len(string(privkeyBytes)))
		fmt.Println(len(privkeyBytes))
		fmt.Printf("%x \n", privkeyBytes)*/
		hbcPasswdBytes, err := hex.DecodeString(os.Args[4])
		if err != nil {
			panic(err)
		}
		d, err := common.ParseFile(os.Args[2], string(privkeyBytes), []byte(os.Args[3]), hbcPasswdBytes)
		if err != nil {
			panic(err)
		}
		fmt.Println(d)
	} else if os.Args[1] == "deriveChildPrivateKey" {
		if len(os.Args) < 7 {
			printUsage()
			os.Exit(1)
		}
		privkeyBytes, err := os.ReadFile(os.Args[5])
		if err != nil {
			panic(err)
		}
		privkeyBytes = privkeyBytes[:len(privkeyBytes)-1]
		/*fmt.Println(string(privkeyBytes))
		fmt.Println(len(string(privkeyBytes)))
		fmt.Println(len(privkeyBytes))
		fmt.Printf("%x \n", privkeyBytes)*/
		hbcPasswdBytes, err := hex.DecodeString(os.Args[4])
		if err != nil {
			panic(err)
		}
		d, err := common.ParseFile(os.Args[2], string(privkeyBytes), []byte(os.Args[3]), hbcPasswdBytes)
		if err != nil {
			panic(err)
		}
		fmt.Println(d)
		p, err := common.DeriveChildPrivateKey(d, os.Args[6])
		if err != nil {
			panic(err)
		}
		pubECPoint := crypto.ScalarBaseMult(tss.S256(), big.NewInt(0).SetBytes(p))
		fmt.Printf("priv key is: %x  pubkey x: %x y: %x \n", p, pubECPoint.X().Bytes(), pubECPoint.Y().Bytes())
	} else if os.Args[1] == "generateKeyPair" {
		err := common.GenerateRSAKeyPair()
		if err != nil {
			panic(err)
		}
	} else {
		printUsage()
	}
	os.Exit(0)
}

func printUsage() {
	fmt.Println("USAGE:")
	fmt.Println("1. recovery-tool makeZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [pubkeyFilePath] [userPrivatekeySlice] [hbcPrivatekeySlice1|hbcPrivatekeySlice2] [chaincode1|chaincode2|chaincode3] [pubkeySlice1|pubkeySlice2|pubkeySlice3]")
	fmt.Println("eg: recovery-tool makeZipFile './zipTest.zip' '123123' '456456' './pubkeyFile' '5ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442c9' '7ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442c9|8ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442c9' '4ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442c9|4ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442ca|4ecd00a8164031b61c7c61578137b83d5c0b57d6dbd8617ece480ec9078442cb' '033669d206967b384d588b366b6400733987befc6604fec764f9fc2d42a3bf7a86|021b491468a9c042e6d4e994c3979df14454cd99e4fc207161a929e719f644540b|02f75cebd23a9ac7e1364d0462be378f09aaf26474eb46cc43bdef5de2817932e5'")
	fmt.Println("2. recovery-tool parseZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [privkeyFilePath]")
	fmt.Println("eg: recovery-tool parseZipFile './zipTest.zip' '123123' '456456' './private_key.pem'")
	fmt.Println("3. recovery-tool deriveChildPrivateKey [zipFilePath] [userPassphrase] [hbcPassphrase] [privkeyFilePath] [derivePath]")
	fmt.Println("eg: recovery-tool deriveChildPrivateKey './zipTest.zip' '123123' '456456' './private_key.pem' '81/0/46/0/0'")
}
