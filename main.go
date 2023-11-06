package main

import (
	"fmt"
	"github.com/james-ray/recovery-tool/coincover"
	"os"
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
		_, err = coincover.MakeZipFile(os.Args[3], os.Args[4], string(pubkeyBytes), os.Args[6], os.Args[8:], os.Args[2])
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
		d, err := coincover.ParseFile(os.Args[2], string(privkeyBytes), os.Args[3], os.Args[4])
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
	fmt.Println("1. recovery-tool makeZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [pubkeyFilePath] [userMnemonic] [hbcMnemonics...]")
	fmt.Println("eg: recovery-tool makeZipFile './zipTest.zip' '123123' '456456' './pubkeyFile' 'user word1 word2' 'hbcnode1 word1 word2' 'hbcnode2 word1 word2'")
	fmt.Println("2. recovery-tool parseZipFile [zipFilePath] [userPassphrase] [hbcPassphrase] [privkeyFilePath]")
	fmt.Println("eg: recovery-tool parseZipFile './zipTest.zip' '123123' '456456' './privkeyFile'")
}
