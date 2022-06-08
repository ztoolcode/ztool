/**
  @author: fanyanan
  @date: 2022/6/8
  @note: //time operation method
**/
package date

import (
	"github.com/druidcaesa/ztool/common"
	"time"
)

// Time return time structure
func (dateTime *DateTime) Time() time.Time {
	return dateTime.t
}

// Now get standard time
// If time Now is empty, the single-sign time is obtained, otherwise the set time
func (dateTime *DateTime) Now() *DateTime {
	dateTime.t = time.Now()
	return dateTime
}

// Format format time
func (dateTime *DateTime) Format(s ...string) string {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	if len(s) > 0 {
		replace, b := replace(s[0])
		if !b {
			return dateTime.t.Format(complete)
		} else {
			return dateTime.t.Format(replace)
		}
	}
	return dateTime.t.Format(complete)
}

// SetTime set time
func (dateTime *DateTime) SetTime(t ...time.Time) *DateTime {
	dateTime.t = time.Now()
	if len(t) > 0 {
		dateTime.t = t[0]
	}
	return dateTime
}

// OperationDay Add and subtract time by day
// Add when num is positive, subtract when negative
// This method is mostly used for user or customer validity period operation,
// which is convenient for time operation.
func (dateTime *DateTime) OperationDay(num int) *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	dateTime.t = dateTime.t.AddDate(0, 0, num)
	return dateTime
}

// OperationMoon Add and subtract time by Moon
// Add when num is positive, subtract when negative
// This method is mostly used for user or customer validity period operation,
// which is convenient for time operation.
func (dateTime *DateTime) OperationMoon(num int) *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	dateTime.t = dateTime.t.AddDate(0, num, 0)
	return dateTime
}

// OperationYears Add and subtract time by Years
// Add when num is positive, subtract when negative
// This method is mostly used for user or customer validity period operation,
// which is convenient for time operation.
func (dateTime *DateTime) OperationYears(num int) *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	dateTime.t = dateTime.t.AddDate(num, 0, 0)
	return dateTime
}

// StartTimeOfDay start time of day
// This method can get the start time of the day, which is year:month:day 00:00:00
func (dateTime *DateTime) StartTimeOfDay() *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	y, m, d := dateTime.t.Date()
	dateTime.t = time.Date(y, m, d, 0, 0, 0, 0, dateTime.t.Location())
	return dateTime
}

// StartOfWeek Weekly start time
func (dateTime *DateTime) StartOfWeek() *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
		dateTime.weekStartDay = weekStartDay
	}
	t := dateTime.StartTimeOfDay()
	weekday := int(t.Time().Weekday())
	if dateTime.weekStartDay != common.Sunday {
		weekStartDayInt := int(dateTime.weekStartDay)
		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	dateTime.t = t.t.AddDate(0, 0, -weekday)
	return dateTime
}

// StartOfMonth beginning of month
// This method user gets the start time of each month
func (dateTime *DateTime) StartOfMonth() *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	y, m, _ := dateTime.t.Date()
	dateTime.t = time.Date(y, m, 1, 0, 0, 0, 0, dateTime.t.Location())
	return dateTime
}

// SetWeekStartDay Set the start time of the week, the default is Sunday. China generally uses a week to start on Monday
//Use this method to quickly set up The value range is 0-6, please use the enumeration preset by the system
func (dateTime *DateTime) SetWeekStartDay(weekday common.Weekday) *DateTime {
	if dateTime.t.IsZero() {
		dateTime.t = time.Now()
	}
	if weekday < 0 || weekday > 6 {
		weekday = common.Sunday
	}
	dateTime.weekStartDay = weekday
	return dateTime
}
