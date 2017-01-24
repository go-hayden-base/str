package str

import (
	"regexp"
	"strings"
)

// IsEmpty return the string argument is empty or not
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

const intReg = `^\s*(\+|-){0,1}\s*[0-9]+\s*$`

// IsInt reutrn the string argument is int number or not
func IsInt(s string) bool {
	reg := regexp.MustCompile(intReg)
	return reg.MatchString(s)
}

const floatReg = `^\s*(\+|-){0,1}\s*[0-9]+\.[0-9]+\s*$`

// IsFloat return the string argument is float or not
func IsFloat(s string) bool {
	reg := regexp.MustCompile(floatReg)
	return reg.MatchString(s)
}

// IsNumber return the string argument is number or not
func IsNumber(s string) bool {
	return IsInt(s) || IsFloat(s)
}

const httpURLReg = `^(?i:(http)|(https))://\S+$`

// IsHttpURL return the string argument is http or https url or not
func IsHttpURL(s string) bool {
	reg := regexp.MustCompile(httpURLReg)
	return reg.MatchString(s)
}
