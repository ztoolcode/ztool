/**
  @author: fanyanan
  @date: 2022/6/8
  @note: //Tool registration calls API
**/
package ztool

import (
	"github.com/druidcaesa/ztool/date"
	"github.com/druidcaesa/ztool/net"
)

var (
	// HttpUtils http client tool
	HttpUtils net.Client
	// DateUtils time manipulation tool
	DateUtils date.DateTime
)
