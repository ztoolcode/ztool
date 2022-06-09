/**
  @author: fanyanan
  @date: 2022/6/9
  @note: //Common Algorithms for People's Republic of China ID Cards
**/
package idcard

import (
	"errors"
	"github.com/druidcaesa/ztool/date"
	"strconv"
)

// GetProvinceByIdCard get province
func (card *Idcard) GetProvinceByIdCard(address string) string {
	address = address[:6]
	provincialCode := address[:3] + "000"
	cityCode := address[:4] + "00"
	return area[provincialCode] + area[cityCode] + area[address]
}

// IsValidCard Verify that the ID is legal
func (card *Idcard) IsValidCard(id string) (bool, error) {
	if len(id) != 15 && len(id) != 18 {
		return false, errors.New("身份证长度不对")
	}
	return checkValidNo18(id)
}

// Convert15To18 15-digit ID card to 18-digit
func (card *Idcard) Convert15To18(id string) string {
	nLen := len(id)
	if nLen != 15 {
		return "身份证不是15位！"
	}
	id18 := make([]byte, 0)
	id18 = append(id18, id[:6]...)
	id18 = append(id18, '1', '9')
	id18 = append(id18, id[6:]...)

	sum := 0
	for i, v := range id18 {
		n, _ := strconv.Atoi(string(v))
		sum += n * weight[i]
	}
	mod := sum % 11
	id18 = append(id18, validValue[mod])
	return string(id18)
}

// GetBirthByIdCard get birthday
func (card *Idcard) GetBirthByIdCard(id string) (string, error) {
	birth := id[6:14]
	_, err := checkBirthdayCode(birth)
	if err != nil {
		return "", err
	}
	return birth[:4] + "-" + birth[4:6] + "-" + birth[6:], nil
}

// GetAgeByIdCard Get age based on ID
func (card *Idcard) GetAgeByIdCard(id string) int {
	d := new(date.DateTime)
	now := d.Now().Time()
	birth, _ := card.GetBirthByIdCard(id)
	startTime, err := d.Parse(birth)
	if err != nil {
		return 0
	}
	if startTime.Time().Before(now) {
		diff := now.Unix() - startTime.Time().Unix()
		Age := diff / (3600 * 365 * 24)
		return int(Age)
	}
	return 0
}

// GetYearByIdCard Year based on ID
func (card *Idcard) GetYearByIdCard(id string) string {
	return id[6:14][:4]
}

// GetMonthByIdCard Get the month of your birthday based on your ID
func (card *Idcard) GetMonthByIdCard(id string) string {
	return id[6:14][4:6]
}

// GetDayByIdCard Get birthday based on ID
func (card *Idcard) GetDayByIdCard(id string) string {
	return id[6:14][6:]
}

// GetSex Get gender based on ID
// Description 0 Female 1 Male 3 Identification ID card gender identification error
func (card *Idcard) GetSex(id string) uint {
	var unknown uint = 3
	sexStr := id[16:17]
	if sexStr == "" {
		return unknown
	}
	i, err := strconv.Atoi(sexStr)
	if err != nil {
		return unknown
	}
	if i%2 != 0 {
		return 1
	}
	return 0
}
