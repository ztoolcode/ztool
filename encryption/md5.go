/**
  @author: fanyanan
  @date: 2022/6/9
  @note: //md5 Encryption and decryption tools
**/
package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//Md5Check md5 Check method
func (e *Encryption) Md5Check(content, encrypted string) bool {
	return strings.EqualFold(e.Md5Encode(content), encrypted)
}

//Md5Encode md5 Signature function
func (e *Encryption) Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
