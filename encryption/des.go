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
  @note: //DES Encryption and decryption
**/
package encryption

import (
	"crypto/des"
	"encoding/hex"
	"errors"
)

//DesEncrypt des encrypted function
//salt The parameter is the salt value of encryption and decryption, and the maximum
//length is 8. If the length is exceeded, the corresponding exception will be thrown
func DesEncrypt(text string, salt ...string) (string, error) {
	key := []byte(defaultDesSalt)
	if len(salt) > 0 {
		if len(salt[0]) > 8 {
			return "", errors.New("DES The maximum length of the encrypted salt value is 8")
		}
		key = []byte(salt[0])
	}
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = zeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("What is required is an integer multiple of the size")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

//DesDecrypt des encrypted function
//salt The parameter is the salt value of encryption and decryption, and the maximum
//length is 8. If the length is exceeded, the corresponding exception will be thrown
func DesDecrypt(decrypted string, salt ...string) (string, error) {
	key := []byte(defaultDesSalt)
	if len(salt) > 0 {
		if len(salt[0]) > 8 {
			return "", errors.New("DES The maximum length of the encrypted salt value is 8")
		}
		key = []byte(salt[0])
	}
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = zeroUnPadding(out)
	return string(out), nil
}
