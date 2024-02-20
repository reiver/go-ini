package ininamevalue

import (
	"unicode/utf8"

	"github.com/reiver/go-ini/internal/eol"
	"github.com/reiver/go-ini/internal/name"
	"github.com/reiver/go-ini/internal/spacing"
	"github.com/reiver/go-ini/internal/value"
)

func ParseBytes(bytes []byte) (name string, value string, size int, err error) {
	if len(bytes) <= 0 {
		return "", "", 0, nil
	}

	var p []byte = bytes

	{
		var spacingSize int

		spacingSize, err := inispacing.ParseBytes(p)
		if nil != err {
			return "", "", 0, err
		}

		size += spacingSize
		p = p[spacingSize:]
	}

	{
		var nameSize int

		name, nameSize, err = ininame.ParseBytes(p)
		if nil != err {
			return "", "", 0, err
		}

		size += nameSize
		p = p[nameSize:]
	}

	{
		var spacingSize int

		spacingSize, err := inispacing.ParseBytes(p)
		if nil != err {
			return "", "", 0, err
		}

		size += spacingSize
		p = p[spacingSize:]
	}

	{
		r, runeSize := utf8.DecodeRune(p)
		if utf8.RuneError == r {
			switch size {
			case 1:
				return "", "", 0, errRuneError
			default:
				return "", "", 0, errInternalError
			}
		}

		switch r {
		case ':', '=':
			size += runeSize
			p = p[runeSize:]
		}
	}

	{
		var spacingSize int

		spacingSize, err := inispacing.ParseBytes(p)
		if nil != err {
			return "", "", 0, err
		}

		size += spacingSize
		p = p[spacingSize:]
	}

	{
		var valueSize int

		value, valueSize, err = inivalue.ParseBytes(p)
		if nil != err {
			return "", "", 0, err
		}

		size += valueSize
		p = p[valueSize:]
	}

	{
		var eolSize int

		eolSize, err := inieol.ParseBytes(p)
		if nil != err {
			return "", "", 0, err
		}

		size += eolSize
		p = p[eolSize:]
	}

	return name, value, size, nil
}
