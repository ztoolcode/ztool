package test

import (
	"fmt"
	"github.com/druidcaesa/ztool"
	"testing"
)

//驼峰转蛇形
func TestSnakeString(t *testing.T) {
	str := "userName"
	fmt.Printf("驼峰转蛇形------->%s", ztool.StrUtils.SnakeString(str))
}

//蛇形转驼峰
func TestCamelString(t *testing.T) {
	str := "user_name"
	fmt.Printf("蛇形转驼峰------->%s", ztool.StrUtils.CamelString(str))
}

//不可见字符串判空
func TestHasBlank(t *testing.T) {
	str := "    "
	str1 := ""
	str3 := "222222"
	fmt.Printf("不可见字符串------->%t\n", ztool.StrUtils.HasBlank(str))
	fmt.Printf("空字符串------->%t\n", ztool.StrUtils.HasBlank(str1))
	fmt.Printf("正常字符串------->%t\n", ztool.StrUtils.HasBlank(str3))
}

//
func TestHasEmpty(t *testing.T) {

}
