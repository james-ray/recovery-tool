package main

import (
	"encoding/hex"
	"fmt"
	"github.com/james-ray/recovery-tool/coincover"
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
		_, err = coincover.MakeZipFile([]byte(os.Args[3]), hbcPasswdBytes, string(pubkeyBytes), os.Args[6], strings.Split(os.Args[7], "|"), strings.Split(os.Args[8], "|"), strings.Split(os.Args[9], "|"), os.Args[2])
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
		d, err := coincover.ParseFile(os.Args[2], string(privkeyBytes), []byte(os.Args[3]), hbcPasswdBytes)
		if err != nil {
			panic(err)
		}
		fmt.Println(d)
	} else {
		printUsage()
	}
	os.Exit(0)
}

func printUsage() {
	fmt.Println("USAGE:")
	fmt.Println("1. recovery-tool makeZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [pubkeyFilePath] [userPrivatekeySlice|hbcPrivatekeySlice1|hbcPrivatekeySlice2] [chaincode1|chaincode2|chaincode3] [pubkeySlice1|pubkeySlice2|pubkeySlice3]")
	fmt.Println("eg: recovery-tool makeZipFile './zipTest.zip' '123123' '456456' './pubkeyFile' '112233' '223344|445566' '123123|234234|456456', '123123ab|234234ab|456456ab'")
	fmt.Println("2. recovery-tool parseZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [privkeyFilePath]")
	fmt.Println("eg: recovery-tool parseZipFile './zipTest.zip' '123123' '456456' './privkeyFile'")
}
