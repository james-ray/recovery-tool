package coincover

import (
	"archive/zip"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/james-ray/recovery-tool/common"
	"github.com/james-ray/recovery-tool/utils"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func MakeZipFile(userPassphrase, hbcPassphrase []byte, pubkeyHex, userPrivateSlice string, hbcPrivateSlices []string, chaincodes []string, ownerPubkeySlices []string, saveFilePath string) (*os.File, error) {
	archive, err := os.Create(saveFilePath)
	if err != nil {
		return nil, err
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()
	w1, err := zipWriter.Create("user_share.encrypted")
	if err != nil {
		return nil, err
	}
	encryptedPrivkeySlice, err := utils.AesGcmEncrypt(userPassphrase, []byte(userPrivateSlice))
	if err != nil {
		return nil, err
	}
	encryptedPrivkeySlice, err = common.RSAEncryptFromHexPubkey(encryptedPrivkeySlice, pubkeyHex)
	if err != nil {
		return nil, err
	}
	bs := bytes.NewBufferString(hex.EncodeToString(encryptedPrivkeySlice))
	if _, err = io.Copy(w1, bs); err != nil {
		return nil, err
	}

	fileName := "chaincodes"
	if len(hbcPassphrase) == 0 {
		fileName += "_hbc"
	}
	w1, err = zipWriter.Create(fileName)
	if err != nil {
		return nil, err
	}
	chaincodesBytes, err := json.Marshal(chaincodes)
	if err != nil {
		return nil, err
	}
	if len(hbcPassphrase) > 0 {
		encryptedPrivkeySlice, err = utils.AesGcmEncrypt(hbcPassphrase, chaincodesBytes)
		if err != nil {
			return nil, err
		}
	}

	encryptedPrivkeySlice, err = common.RSAEncryptFromHexPubkey(encryptedPrivkeySlice, pubkeyHex)
	if err != nil {
		return nil, err
	}
	bs = bytes.NewBufferString(hex.EncodeToString(encryptedPrivkeySlice))
	if _, err = io.Copy(w1, bs); err != nil {
		return nil, err
	}

	fileName = "pubkeys"
	if len(hbcPassphrase) == 0 {
		fileName += "_hbc"
	}
	w1, err = zipWriter.Create(fileName)
	if err != nil {
		return nil, err
	}
	pubkeysBytes, err := json.Marshal(ownerPubkeySlices)
	if err != nil {
		return nil, err
	}
	if len(hbcPassphrase) > 0 {
		encryptedPrivkeySlice, err = utils.AesGcmEncrypt(hbcPassphrase, pubkeysBytes)
		if err != nil {
			return nil, err
		}
	}
	encryptedPrivkeySlice, err = common.RSAEncryptFromHexPubkey(encryptedPrivkeySlice, pubkeyHex)
	if err != nil {
		return nil, err
	}
	bs = bytes.NewBufferString(hex.EncodeToString(encryptedPrivkeySlice))
	if _, err = io.Copy(w1, bs); err != nil {
		return nil, err
	}

	for i := 0; i < len(hbcPrivateSlices); i++ {
		fileName = fmt.Sprintf("hbc.encrypted.%d", i)
		if len(hbcPassphrase) == 0 {
			fileName += "_hbc"
		}
		w1, err = zipWriter.Create(fileName)
		if err != nil {
			return nil, err
		}
		if len(hbcPassphrase) > 0 {
			encryptedPrivkeySlice, err = utils.AesGcmEncrypt(hbcPassphrase, []byte(hbcPrivateSlices[i]))
			if err != nil {
				return nil, err
			}
		}
		encryptedPrivkeySlice, err = common.RSAEncryptFromHexPubkey(encryptedPrivkeySlice, pubkeyHex)
		if err != nil {
			return nil, err
		}
		bs = bytes.NewBufferString(hex.EncodeToString(encryptedPrivkeySlice))
		if _, err = io.Copy(w1, bs); err != nil {
			return nil, err
		}
	}
	return archive, nil
}

func readAll(file *zip.File) ([]byte, error) {
	fc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fc.Close()

	content, err := ioutil.ReadAll(fc)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func ParseFile(zipFilePath string, privKeyHex string, userPassphrase, hbcPassphrase []byte) (map[string]string, error) {
	zf, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return nil, err
	}
	defer zf.Close()

	dataMap := make(map[string]string)
	for _, file := range zf.File {
		fileBytes, err := readAll(file)
		if err != nil {
			return nil, err
		}
		fileContent := string(fileBytes)
		textBytes, err := hex.DecodeString(fileContent)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		//fmt.Printf("=%s\n", file.Name)
		//fmt.Printf("%x\n\n", fileBytes) // file content
		encryptedBytes, err := common.RSADecryptFromHexPrivkey(textBytes, privKeyHex)
		if err != nil {
			return nil, err
		}
		if file.Name == "user_share.encrypted" {
			plainBytes, err := utils.AesGcmDecrypt(userPassphrase, encryptedBytes)
			if err != nil {
				return nil, err
			}
			dataMap["user"] = string(plainBytes)
		} else if strings.Contains(file.Name, "chaincodes") {
			if len(hbcPassphrase) > 0 {
				plainBytes, err := utils.AesGcmDecrypt(hbcPassphrase, encryptedBytes)
				if err != nil {
					return nil, err
				}
				dataMap["user"] = string(plainBytes)
			} else {

			}

		} else {
			if len(hbcPassphrase) > 0 {
				plainBytes, err := utils.AesGcmDecrypt(hbcPassphrase, encryptedBytes)
				if err != nil {
					return nil, err
				}
				dataMap[file.Name] = string(plainBytes)
			} else {
				dataMap[file.Name] = string(encryptedBytes)
			}
		}

	}
	return dataMap, nil
}
