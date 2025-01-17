package iniconv

import (
	"strings"
)

// ParseStrings helps convery an INI value that is (supposed to be) a list into a []string.
func ParseStrings(text string, separator string) []string {
	if "" == text {
		return []string{}
	}

	var ss []string = strings.Split(text, separator)

	{
		var length int = len(ss)

		for index:=0; index<length; index++ {
			ss[index] = strings.TrimSpace(ss[index])
		}
	}

	return ss
}
