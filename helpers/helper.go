package main

import (
	"encoding/hex"
	"fmt"
	"github.com/james-ray/recovery-tool/common"
)

//export hello
func hello() string {
	return "Greetings"
}

//export generateChildExtendedPrivateKey
func generateChildExtendedPrivateKey(metadataFilePath string, walletType int, vaultIndex int, chainInt int, subIndex int) string {
	// 实现生成子扩展私钥的逻辑
	hdPath := fmt.Sprintf("81/%d/%d/%d/%d", walletType, vaultIndex, chainInt, subIndex)
	metadataMap, err := common.ReadMetadataFile(metadataFilePath)
	if err != nil {
		panic(err)
	}
	privBytes, err := common.DeriveChildPrivateKey(metadataMap, hdPath)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(privBytes)
}

func main() {

}
