package iniscanner_key

import (
	"github.com/reiver/go-utf8s"

	"io"
	"strings"

	"testing"
)

func TestRead(t *testing.T) {

	tests := []struct{
		Value         string
		ExpectedValue string
		ExpectedSize  int
	}{
		{ // 0
			Value:            "key=",
			ExpectedValue:    "key",
			ExpectedSize: len("key"),
		},
		{ // 1
			Value:            "key:",
			ExpectedValue:    "key",
			ExpectedSize: len("key"),
		},



		{ // 2
			Value:            "key۵=",
			ExpectedValue:    "key۵",
			ExpectedSize: len("key۵"),
		},
		{ // 3
			Value:            "key۵:",
			ExpectedValue:    "key۵",
			ExpectedSize: len("key۵"),
		},



		{ // 4
			Value:            "key =",
			ExpectedValue:    "key",
			ExpectedSize: len("key"),
		},
		{ // 5
			Value:            "key :",
			ExpectedValue:    "key",
			ExpectedSize: len("key"),
		},



		{ // 6
			Value:            "key۵ =",
			ExpectedValue:    "key۵",
			ExpectedSize: len("key۵"),
		},
		{ // 7
			Value:            "key۵ :",
			ExpectedValue:    "key۵",
			ExpectedSize: len("key۵"),
		},



		{ // 8
			Value:            `apple\=banana=cherry`,
			ExpectedValue:    "apple=banana",
			ExpectedSize: len(`apple\=banana`),
		},
		{ // 9
			Value:            `apple\=banana:cherry`,
			ExpectedValue:    "apple=banana",
			ExpectedSize: len(`apple\=banana`),
		},



		{ // 10
			Value:            `apple\:banana=cherry`,
			ExpectedValue:    "apple:banana",
			ExpectedSize: len(`apple\:banana`),
		},
		{ // 11
			Value:            `apple\:banana:cherry`,
			ExpectedValue:    "apple:banana",
			ExpectedSize: len(`apple\:banana`),
		},



		{ // 12
			Value:            `apple۵\=banana=cherry`,
			ExpectedValue:    "apple۵=banana",
			ExpectedSize: len(`apple۵\=banana`),
		},
		{ // 13
			Value:            `apple۵\=banana:cherry`,
			ExpectedValue:    "apple۵=banana",
			ExpectedSize: len(`apple۵\=banana`),
		},



		{ // 14
			Value:            `apple۵\:banana=cherry`,
			ExpectedValue:    "apple۵:banana",
			ExpectedSize: len(`apple۵\:banana`),
		},
		{ // 15
			Value:            `apple۵\:banana:cherry`,
			ExpectedValue:    "apple۵:banana",
			ExpectedSize: len(`apple۵\:banana`),
		},



		{ // 16
			Value:            `apple=banana\=cherry`,
			ExpectedValue:    "apple",
			ExpectedSize: len("apple"),
		},
		{ // 17
			Value:            `apple:banana\=cherry`,
			ExpectedValue:    "apple",
			ExpectedSize: len("apple"),
		},



		{ // 18
			Value:            `apple=banana\:cherry`,
			ExpectedValue:    "apple",
			ExpectedSize: len("apple"),
		},
		{ // 19
			Value:            `apple:banana\:cherry`,
			ExpectedValue:    "apple",
			ExpectedSize: len("apple"),
		},



		{ // 20
			Value:            `apple۵=banana\=cherry`,
			ExpectedValue:    "apple۵",
			ExpectedSize: len("apple۵"),
		},
		{ // 21
			Value:            `apple۵:banana\=cherry`,
			ExpectedValue:    "apple۵",
			ExpectedSize: len("apple۵"),
		},



		{ // 22
			Value:            `apple۵=banana\:cherry`,
			ExpectedValue:    "apple۵",
			ExpectedSize: len("apple۵"),
		},
		{ // 23
			Value:            `apple۵:banana\:cherry`,
			ExpectedValue:    "apple۵",
			ExpectedSize: len("apple۵"),
		},



		{ // 24
			Value:            `\\=`,
			ExpectedValue:     `\`,
			ExpectedSize: len(`\\`),
		},



		{ // 25
			Value:            `\0=`,
			ExpectedValue:  "\x00",
			ExpectedSize: len(`\0`),
		},



		{ // 26
			Value:            `\a=`,
			ExpectedValue:    "\a",
			ExpectedSize: len(`\a`),
		},



		{ // 27
			Value:            `\b=`,
			ExpectedValue:    "\b",
			ExpectedSize: len(`\b`),
		},



		{ // 28
			Value:            `\t=`,
			ExpectedValue:    "\t",
			ExpectedSize: len(`\t`),
		},



		{ // 29
			Value:            `\r=`,
			ExpectedValue:    "\r",
			ExpectedSize: len(`\r`),
		},



		{ // 30
			Value:            `\n=`,
			ExpectedValue:    "\n",
			ExpectedSize: len(`\n`),
		},



		{ // 31
			Value:            `\;=`,
			ExpectedValue:     ";",
			ExpectedSize: len(`\;`),
		},



		{ // 32
			Value:            `\#=`,
			ExpectedValue:     "#",
			ExpectedSize: len(`\#`),
		},



		{ // 33
			Value:            `\==`,
			ExpectedValue:     "=",
			ExpectedSize: len(`\=`),
		},



		{ // 34
			Value:            `\:=`,
			ExpectedValue:     ":",
			ExpectedSize: len(`\:`),
		},



		{ // 34
			Value:           `\\\0\a\b\t\r\n\;\#\=\:apple.banana.cherry=one two three`,
			ExpectedValue:   "\\\x00\a\b\t\r\n;#=:apple.banana.cherry",
			ExpectedSize: len(`\\\0\a\b\t\r\n\;\#\=\:apple.banana.cherry`),
		},



		{ // 35
			Value:            "key # this is the key\n= # this is the equals\nthis is the value\n",
			ExpectedValue:    "key",
			ExpectedSize: len(`key`),
		},



		{ // 36
			Value:            "key; this is the key\n=; this is the equals\nthis is the value\n",
			ExpectedValue:    "key",
			ExpectedSize: len(`key`),
		},



		{ // 36
			Value:            "key# this is the key\n=# this is the equals\nthis is the value\n",
			ExpectedValue:    "key",
			ExpectedSize: len(`key`),
		},

	}


	for testNumber, test := range tests {

		runeScanner := utf8s.NewRuneScanner( strings.NewReader(test.Value) )

		token, actualSize, err := Read(runeScanner)
		if nil != err && io.EOF != err {
			t.Errorf("For test #%d, for %q did not expect an error, but actually got one: (%T) %q.", testNumber, test.Value, err, err)
			continue
		}

		if expected, actual := test.ExpectedValue, token.Unwrap(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
		if expected, actual := test.ExpectedSize, actualSize; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		if io.EOF == err {
			continue
		}
	}
}

func TestReadError(t *testing.T) {

	tests := []struct{
		Value string
	}{
		{
			Value: "=value",
		},
		{
			Value: ":value",
		},



		{
			Value: "=value=wrong",
		},
		{
			Value: ":value=wrong",
		},



		{
			Value: "=value:wrong",
		},
		{
			Value: ":value:wrong",
		},



		{
			Value: "; comment",
		},
		{
			Value: "# comment",
		},



		{
			Value: "; comment=wrong",
		},
		{
			Value: "# comment=wrong",
		},



		{
			Value: "; comment:wrong",
		},
		{
			Value: "# comment:wrong",
		},
	}


	for testNumber, test := range tests {

		runeScanner := utf8s.NewRuneScanner( strings.NewReader(test.Value) )

		_, _, err := Read(runeScanner)
		if nil == err || io.EOF == err {
			t.Errorf("For test #%d, for %q expected an error, but did not actually get one: (%T) %q.", testNumber, test.Value, err, err)
			continue
		}

	}
}
