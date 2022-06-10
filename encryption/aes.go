//MIT License
//
//Copyright (c) 2022 druidcaesa
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

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
