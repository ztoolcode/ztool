/**
  @author: fanyanan
  @date: 2022/6/8
  @note: //Tool registration calls API
**/
package ztool

import (
	"github.com/druidcaesa/ztool/date"
	"github.com/druidcaesa/ztool/encryption"
	"github.com/druidcaesa/ztool/file"
	"github.com/druidcaesa/ztool/id"
	"github.com/druidcaesa/ztool/idcard"
	"github.com/druidcaesa/ztool/net"
	"github.com/druidcaesa/ztool/str"
)

var (
	// HttpUtils http client tool
	HttpUtils net.Client
	// DateUtils time manipulation tool
	DateUtils date.DateTime
	// EncryptionUtils Encryption and decryption tools
	EncryptionUtils encryption.Encryption
	// FileUtils file manipulation tool
	FileUtils file.ZFile
	// IdUtils tool package
	IdUtils id.Id
	//StringUtils String tools
	StringUtils str.Str
	//People's Republic of China ID Card Verification Tool
	IDCardUtils idcard.Idcard
)
