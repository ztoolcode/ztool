/**
  @author: fanyanan
  @date: 2022/6/8
  @note: //Time Manipulation Tool Setup Basics
**/
package date

import (
	"strings"
	"time"
)

type DateTime struct {
	t time.Time
}

const (
	year            = "2006"
	month           = "01"
	day             = "02"
	hour            = "15"
	minute          = "04"
	second          = "05"
	complete        = "2006-01-02 15:05:05"
	stringToTimeOne = "2006-01-02 15:04:05"
	stringToTimeTow = "2006-01-02"
)

//time format string
func replace(s string) (string, bool) {
	flag := false
	if strings.Contains(s, "YYYY") {
		s = strings.Replace(s, "YYYY", year, 1)
		flag = true
	}
	if strings.Contains(s, "MM") {
		s = strings.Replace(s, "MM", month, 1)
		flag = true
	}
	if strings.Contains(s, "DD") {
		s = strings.Replace(s, "DD", day, 1)
		flag = true
	}
	if strings.Contains(s, "hh") {
		s = strings.Replace(s, "hh", hour, 1)
		flag = true
	}
	if strings.Contains(s, "mm") {
		s = strings.Replace(s, "mm", minute, 1)
		flag = true
	}
	if strings.Contains(s, "ss") {
		s = strings.Replace(s, "ss", second, 1)
		flag = true
	}
	return s, flag
}
