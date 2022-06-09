/**
  @author: fanyanan
  @date: 2022/6/9
  @note: //Aes Encryption and decryption tools
**/
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//AesEncrypt function
//original The original password, salt is an optional variable.
//If the salt exists, the passed variable is used.
//If it does not exist, the system default salt value is used.
func (e *Encryption) AesEncrypt(original string, salt ...string) string {
	var saltValue = defaultAesSalt
	if len(salt) > 0 {
		saltValue = salt[0]
	}
	// Convert to byte array
	origData := []byte(original)
	k := []byte(saltValue)
	// group key
	block, _ := aes.NewCipher(k)
	// Get the length of the key block
	blockSize := block.BlockSize()
	// Completion code
	origData = pKCS7Padding(origData, blockSize)
	// encryption mode
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// create array
	crated := make([]byte, len(origData))
	// encryption
	blockMode.CryptBlocks(crated, origData)
	return base64.StdEncoding.EncodeToString(crated)

}

//AesDecrypt function
//original The original password, salt is an optional variable.
//If the salt exists, the passed variable is used.
//If it does not exist, the system default salt value is used.
func (e *Encryption) AesDecrypt(crated string, salt ...string) string {
	var saltValue = defaultAesSalt
	if len(salt) > 0 {
		saltValue = salt[0]
	}
	// Convert to byte array
	cratedByte, _ := base64.StdEncoding.DecodeString(crated)
	k := []byte(saltValue)
	// group key
	block, _ := aes.NewCipher(k)
	// Get the length of the key block
	blockSize := block.BlockSize()
	// encryption mode
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// create array
	orig := make([]byte, len(cratedByte))
	// decrypt
	blockMode.CryptBlocks(orig, cratedByte)
	// to complete the code
	orig = pKCS7UnPadding(orig)
	return string(orig)
}
