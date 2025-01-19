package ininamevalue

import (
	gobytes "bytes"
	"unicode/utf8"

	"github.com/reiver/go-opt"

	"github.com/reiver/go-ini/internal/eol"
	"github.com/reiver/go-ini/internal/name"
	"github.com/reiver/go-ini/internal/spacing"
	"github.com/reiver/go-ini/internal/value"
)

func ParseBytes(bytes []byte) (name string, value string, size int, err error) {
	const nada string = ""

	if len(bytes) <= 0 {
		return nada, nada, 0, nil
	}

	var p []byte = bytes

	{
		var spacingSize int

		spacingSize, err := inispacing.ParseBytes(p)
		if nil != err {
			return nada, nada, 0, err
		}

		size += spacingSize
		p = p[spacingSize:]
	}

	{
		var nameSize int

		name, nameSize, err = ininame.ParseBytes(p)
		if nil != err {
			return nada, nada, 0, err
		}

		size += nameSize
		p = p[nameSize:]
	}

	{
		var spacingSize int

		spacingSize, err := inispacing.ParseBytes(p)
		if nil != err {
			return nada, nada, 0, err
		}

		size += spacingSize
		p = p[spacingSize:]
	}

	var delimiter opt.Optional[rune]
	{
		r, runeSize := utf8.DecodeRune(p)
		if utf8.RuneError == r {
			switch size {
			case 1:
				return nada, nada, 0, errRuneError
			default:
				return nada, nada, 0, errInternalError
			}
		}

		switch r {
		case ':', '=', '&':
			size += runeSize
			p = p[runeSize:]

			delimiter = opt.Something(r)
		}
	}

	{
		var spacingSize int

		spacingSize, err := inispacing.ParseBytes(p)
		if nil != err {
			return nada, nada, 0, err
		}

		size += spacingSize
		p = p[spacingSize:]
	}

	{
		var valueSize int

		value, valueSize, err = inivalue.ParseBytes(p)
		if nil != err {
			return nada, nada, 0, err
		}

		size += valueSize
		p = p[valueSize:]
	}

	{
		var eolSize int

		eolSize, err := inieol.ParseBytes(p)
		if nil != err {
			return nada, nada, 0, err
		}

		size += eolSize
		p = p[eolSize:]
	}

	if opt.Something('&') == delimiter {

		var endOfString string = value
		value = ""

		var buffer [256]byte
		var end []byte = buffer[0:0]
		end = append(end, endOfString...)

		var index int = gobytes.Index(p, end)

		switch {
		case index < 0:
			value = string(p)
			valueSize := len(value)
			size += valueSize
			p = p[valueSize:]
		default:
			value = string(p[:index])
			valueSize := len(value)
			size += valueSize
			endSize := len(end)
			size += endSize
			p = p[valueSize+endSize:]
		}

		{
			var eolSize int

			eolSize, err := inieol.ParseBytes(p)
			if nil != err {
				return nada, nada, 0, err
			}

			size += eolSize
			p = p[eolSize:]
		}
	}

	return name, value, size, nil
}
