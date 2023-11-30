package common

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"github.com/pkg/errors"
)

func RSASign(data string, priKey string) (string, error) {
	hashMd5 := md5.Sum([]byte(data))
	hashed := hashMd5[:]

	// rsa
	block, _ := pem.Decode([]byte(priKey))
	if block == nil {
		return "", errors.New("private key error")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.MD5, hashed)
	//fmt.Println("消息的签名信息： ", base64.StdEncoding.EncodeToString(signature))
	if err != nil {
		return "", err
	}
	ciphertext := base64.StdEncoding.EncodeToString(signature)
	return ciphertext, nil
}

func RSAVerifySign(data string, pubKey string, signed string) error {
	hashMd5 := md5.Sum([]byte(data))
	hashed := hashMd5[:]

	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		return errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)

	decodeString, err := base64.StdEncoding.DecodeString(signed)
	if err != nil {
		return err
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.MD5, hashed, decodeString)
	if err != nil {
		return err
	}

	return nil
}

func RSAEncryptFromHexPubkey(data []byte, pubkeyStr string) ([]byte, error) {
	pub, err := parsePublicKey(pubkeyStr)
	if err != nil {
		return nil, err
	}
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		return nil, err
	}
	return encryptedBytes, nil
}

func RSADecryptFromHexPrivkey(encryptedBytes []byte, privkeyStr string) ([]byte, error) {
	priv, err := parsePrivateKey(privkeyStr)
	if err != nil {
		return nil, err
	}
	plainBytes, err := rsa.DecryptPKCS1v15(nil, priv, encryptedBytes)
	if err != nil {
		return nil, err
	}
	return plainBytes, nil
}

func parsePublicKey(pemData string) (*rsa.PublicKey, error) {
	/*pemData, err := hexToDER(hexStr)
	if err != nil {
		fmt.Println("Failed to convert hex to PEM:", err)
		return nil, err
	}*/
	p, _ := pem.Decode([]byte(pemData))
	// The key is expected to be in ASN.1 DER format.
	pub, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPub := pub.(*rsa.PublicKey)
	return rsaPub, nil
}

func parsePrivateKey(pemData string) (*rsa.PrivateKey, error) {
	// Convert the hex private key to PEM format
	/*pemData, err := hexToDER(hexStr)
	if err != nil {
		fmt.Println("Failed to convert hex to PEM:", err)
		return nil, err
	}*/
	//fmt.Printf("priv der: %s", string(pemData))
	// The key is expected to be in ASN.1 DER format.
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, errors.New("Failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		priv1, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		priv = priv1.(*rsa.PrivateKey)
	}

	return priv, nil
}

// Convert hex-encoded data to DER format (binary)
func hexToDER(hexData string) ([]byte, error) {
	// Decode the hex data
	decoded, err := hex.DecodeString(hexData)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
