/**
  @author: fanyanan
  @date: 2022/6/8
  @note: //time operation method
**/
package date

import "time"

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
