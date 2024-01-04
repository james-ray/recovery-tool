package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/james-ray/recovery-tool/common"
	"syscall/js"
)

func main() {
	js.Global().Set("generateChildExtendedPrivateKey", js.FuncOf(generateChildExtendedPrivateKey))
	done := make(chan struct{}, 0)
	<-done
}

func Hello() string {
	return "Greetings"
}

func generateChildExtendedPrivateKey(this js.Value, args []js.Value) interface{} {
	metadataFileStr := args[0].String()
	walletType := args[1].Float()
	vaultIndex := args[2].Float()
	chainInt := args[3].Float()
	subIndex := args[4].Float()
	fmt.Printf("params: %+v", args)
	// 实现生成子扩展私钥的逻辑
	hdPath := fmt.Sprintf("81/%d/%d/%d/%d", int(walletType), int(vaultIndex), int(chainInt), int(subIndex))
	var metadataMap map[string]string
	err := json.Unmarshal([]byte(metadataFileStr), &metadataMap)
	if err != nil {
		panic(err)
	}
	privBytes, err := common.DeriveChildPrivateKey(metadataMap, hdPath)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(privBytes)
}
