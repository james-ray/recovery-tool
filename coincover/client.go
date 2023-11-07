package coincover

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/james-ray/recovery-tool/common"
	"github.com/james-ray/recovery-tool/utils"
	"io"
	"io/ioutil"
	"os"
)

func MakeZipFile(userPassphrase, hbcPassphrase []byte, pubkeyHex, userMnemonic string, hbcMnemonics []string, saveFilePath string) (*os.File, error) {
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
	encryptedMnenomic, err := utils.AesGcmEncrypt(userPassphrase, []byte(userMnemonic))
	if err != nil {
		return nil, err
	}
	encryptedMnenomic, err = common.RSAEncryptFromHexPubkey(encryptedMnenomic, pubkeyHex)
	if err != nil {
		return nil, err
	}
	bs := bytes.NewBuffer(encryptedMnenomic)
	if _, err = io.Copy(w1, bs); err != nil {
		return nil, err
	}

	for i := 0; i < len(hbcMnemonics); i++ {
		w1, err = zipWriter.Create(fmt.Sprintf("hbc.encrypted.%d", i))
		if err != nil {
			return nil, err
		}
		encryptedMnenomic, err = utils.AesGcmEncrypt(hbcPassphrase, []byte(hbcMnemonics[i]))
		if err != nil {
			return nil, err
		}
		encryptedMnenomic, err = common.RSAEncryptFromHexPubkey(encryptedMnenomic, pubkeyHex)
		if err != nil {
			return nil, err
		}
		bs = bytes.NewBuffer(encryptedMnenomic)
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
		//fmt.Printf("=%s\n", file.Name)
		//fmt.Printf("%x\n\n", fileBytes) // file content
		encryptedBytes, err := common.RSADecryptFromHexPrivkey(fileBytes, privKeyHex)
		if err != nil {
			return nil, err
		}
		if file.Name == "user_share.encrypted" {
			plainBytes, err := utils.AesGcmDecrypt(userPassphrase, encryptedBytes)
			if err != nil {
				return nil, err
			}
			dataMap["user"] = string(plainBytes)
		} else {
			plainBytes, err := utils.AesGcmDecrypt(hbcPassphrase, encryptedBytes)
			if err != nil {
				return nil, err
			}
			dataMap[file.Name] = string(plainBytes)
		}

	}
	return dataMap, nil
}
