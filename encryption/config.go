/**
  @author: fanyanan
  @date: 2022/6/9
  @note: //Encrypt and decrypt global public resource
**/
package encryption

import "bytes"

const defaultAesSalt = "xEtOeGkBld5djgk&"
const defaultDesSalt = "e!FJSchz&"

//Encryption and decryption overall structure
type Encryption struct {
}

//complement method
func pKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//to code method
func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//zeroPadding zero padding function
func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

//zeroUnPadding zero unPadding function
func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
