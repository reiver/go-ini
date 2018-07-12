package iniscanner_value

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
		{
			Value:            "",
			ExpectedValue:    "",
			ExpectedSize: len(""),
		},



		{
			Value:            "value",
			ExpectedValue:    "value",
			ExpectedSize: len("value"),
		},



		{
			Value:            "value۵",
			ExpectedValue:    "value۵",
			ExpectedSize: len("value۵"),
		},



		{
			Value:            "value۵ 123",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},



		{
			Value:            "value۵ 123\n",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\r",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\r\n",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\v",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\u0085",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\u2028",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\u2029",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},



		{
			Value:            "value۵ 123\nsome=thing\n",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\rsome=thing\r",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\r\nsome=thing\r\n",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\vsome=thing\v",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\u0085some=thing\u0085",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\u2028some=thing\u2028",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\u2029some=thing\u2028",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},



		{
			Value:            "value۵ 123\nsome:thing\n",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\rsome:thing\r",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\r\nsome:thing\r\n",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\vsome:thing\v",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\u0085some:thing\u0085",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\u2028some:thing\u2028",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},
		{
			Value:            "value۵ 123\u2029some:thing\u2028",
			ExpectedValue:    "value۵ 123",
			ExpectedSize: len("value۵ 123"),
		},



		{
			Value:            `\0`,
			ExpectedValue:  "\x00",
			ExpectedSize: len(`\0`),
		},
		{
			Value:            `\a`,
			ExpectedValue:    "\a",
			ExpectedSize: len(`\a`),
		},
		{
			Value:            `\b`,
			ExpectedValue:    "\b",
			ExpectedSize: len(`\b`),
		},
		{
			Value:            `\t`,
			ExpectedValue:    "\t",
			ExpectedSize: len(`\t`),
		},
		{
			Value:            `\r`,
			ExpectedValue:    "\r",
			ExpectedSize: len(`\r`),
		},
		{
			Value:            `\n`,
			ExpectedValue:    "\n",
			ExpectedSize: len(`\n`),
		},
		{
			Value:            `\;`,
			ExpectedValue:     ";",
			ExpectedSize: len(`\;`),
		},
		{
			Value:            `\#`,
			ExpectedValue:     "#",
			ExpectedSize: len(`\#`),
		},
		{
			Value:            `\=`,
			ExpectedValue:     "=",
			ExpectedSize: len(`\=`),
		},
		{
			Value:            `\:`,
			ExpectedValue:     ":",
			ExpectedSize: len(`\:`),
		},



		{
			Value:            `\0`+"\u000A",
			ExpectedValue:  "\x00",
			ExpectedSize: len(`\0`),
		},
		{
			Value:            `\a`+"\u000A",
			ExpectedValue:    "\a",
			ExpectedSize: len(`\a`),
		},
		{
			Value:            `\b`+"\u000A",
			ExpectedValue:    "\b",
			ExpectedSize: len(`\b`),
		},
		{
			Value:            `\t`+"\u000A",
			ExpectedValue:    "\t",
			ExpectedSize: len(`\t`),
		},
		{
			Value:            `\r`+"\u000A",
			ExpectedValue:    "\r",
			ExpectedSize: len(`\r`),
		},
		{
			Value:            `\n`+"\u000A",
			ExpectedValue:    "\n",
			ExpectedSize: len(`\n`),
		},
		{
			Value:            `\;`+"\u000A",
			ExpectedValue:     ";",
			ExpectedSize: len(`\;`),
		},
		{
			Value:            `\#`+"\u000A",
			ExpectedValue:     "#",
			ExpectedSize: len(`\#`),
		},
		{
			Value:            `\=`+"\u000A",
			ExpectedValue:     "=",
			ExpectedSize: len(`\=`),
		},
		{
			Value:            `\:`+"\u000A",
			ExpectedValue:     ":",
			ExpectedSize: len(`\:`),
		},



		{
			Value:            `\0`+"\u000B",
			ExpectedValue:  "\x00",
			ExpectedSize: len(`\0`),
		},
		{
			Value:            `\a`+"\u000B",
			ExpectedValue:    "\a",
			ExpectedSize: len(`\a`),
		},
		{
			Value:            `\b`+"\u000B",
			ExpectedValue:    "\b",
			ExpectedSize: len(`\b`),
		},
		{
			Value:            `\t`+"\u000B",
			ExpectedValue:    "\t",
			ExpectedSize: len(`\t`),
		},
		{
			Value:            `\r`+"\u000B",
			ExpectedValue:    "\r",
			ExpectedSize: len(`\r`),
		},
		{
			Value:            `\n`+"\u000B",
			ExpectedValue:    "\n",
			ExpectedSize: len(`\n`),
		},
		{
			Value:            `\;`+"\u000B",
			ExpectedValue:     ";",
			ExpectedSize: len(`\;`),
		},
		{
			Value:            `\#`+"\u000B",
			ExpectedValue:     "#",
			ExpectedSize: len(`\#`),
		},
		{
			Value:            `\=`+"\u000B",
			ExpectedValue:     "=",
			ExpectedSize: len(`\=`),
		},
		{
			Value:            `\:`+"\u000B",
			ExpectedValue:     ":",
			ExpectedSize: len(`\:`),
		},



		{
			Value:            `\0`+"\u000C",
			ExpectedValue:  "\x00",
			ExpectedSize: len(`\0`),
		},
		{
			Value:            `\a`+"\u000C",
			ExpectedValue:    "\a",
			ExpectedSize: len(`\a`),
		},
		{
			Value:            `\b`+"\u000C",
			ExpectedValue:    "\b",
			ExpectedSize: len(`\b`),
		},
		{
			Value:            `\t`+"\u000C",
			ExpectedValue:    "\t",
			ExpectedSize: len(`\t`),
		},
		{
			Value:            `\r`+"\u000C",
			ExpectedValue:    "\r",
			ExpectedSize: len(`\r`),
		},
		{
			Value:            `\n`+"\u000C",
			ExpectedValue:    "\n",
			ExpectedSize: len(`\n`),
		},
		{
			Value:            `\;`+"\u000C",
			ExpectedValue:     ";",
			ExpectedSize: len(`\;`),
		},
		{
			Value:            `\#`+"\u000C",
			ExpectedValue:     "#",
			ExpectedSize: len(`\#`),
		},
		{
			Value:            `\=`+"\u000C",
			ExpectedValue:     "=",
			ExpectedSize: len(`\=`),
		},
		{
			Value:            `\:`+"\u000C",
			ExpectedValue:     ":",
			ExpectedSize: len(`\:`),
		},



		{
			Value:            `\0`+"\u000D",
			ExpectedValue:  "\x00",
			ExpectedSize: len(`\0`),
		},
		{
			Value:            `\a`+"\u000D",
			ExpectedValue:    "\a",
			ExpectedSize: len(`\a`),
		},
		{
			Value:            `\b`+"\u000D",
			ExpectedValue:    "\b",
			ExpectedSize: len(`\b`),
		},
		{
			Value:            `\t`+"\u000D",
			ExpectedValue:    "\t",
			ExpectedSize: len(`\t`),
		},
		{
			Value:            `\r`+"\u000D",
			ExpectedValue:    "\r",
			ExpectedSize: len(`\r`),
		},
		{
			Value:            `\n`+"\u000D",
			ExpectedValue:    "\n",
			ExpectedSize: len(`\n`),
		},
		{
			Value:            `\;`+"\u000D",
			ExpectedValue:     ";",
			ExpectedSize: len(`\;`),
		},
		{
			Value:            `\#`+"\u000D",
			ExpectedValue:     "#",
			ExpectedSize: len(`\#`),
		},
		{
			Value:            `\=`+"\u000D",
			ExpectedValue:     "=",
			ExpectedSize: len(`\=`),
		},
		{
			Value:            `\:`+"\u000D",
			ExpectedValue:     ":",
			ExpectedSize: len(`\:`),
		},



		{
			Value:            `\0`+"\u0085",
			ExpectedValue:  "\x00",
			ExpectedSize: len(`\0`),
		},
		{
			Value:            `\a`+"\u0085",
			ExpectedValue:    "\a",
			ExpectedSize: len(`\a`),
		},
		{
			Value:            `\b`+"\u0085",
			ExpectedValue:    "\b",
			ExpectedSize: len(`\b`),
		},
		{
			Value:            `\t`+"\u0085",
			ExpectedValue:    "\t",
			ExpectedSize: len(`\t`),
		},
		{
			Value:            `\r`+"\u0085",
			ExpectedValue:    "\r",
			ExpectedSize: len(`\r`),
		},
		{
			Value:            `\n`+"\u0085",
			ExpectedValue:    "\n",
			ExpectedSize: len(`\n`),
		},
		{
			Value:            `\;`+"\u0085",
			ExpectedValue:     ";",
			ExpectedSize: len(`\;`),
		},
		{
			Value:            `\#`+"\u0085",
			ExpectedValue:     "#",
			ExpectedSize: len(`\#`),
		},
		{
			Value:            `\=`+"\u0085",
			ExpectedValue:     "=",
			ExpectedSize: len(`\=`),
		},
		{
			Value:            `\:`+"\u0085",
			ExpectedValue:     ":",
			ExpectedSize: len(`\:`),
		},



		{
			Value:            `\0`+"\u2028",
			ExpectedValue:  "\x00",
			ExpectedSize: len(`\0`),
		},
		{
			Value:            `\a`+"\u2028",
			ExpectedValue:    "\a",
			ExpectedSize: len(`\a`),
		},
		{
			Value:            `\b`+"\u2028",
			ExpectedValue:    "\b",
			ExpectedSize: len(`\b`),
		},
		{
			Value:            `\t`+"\u2028",
			ExpectedValue:    "\t",
			ExpectedSize: len(`\t`),
		},
		{
			Value:            `\r`+"\u2028",
			ExpectedValue:    "\r",
			ExpectedSize: len(`\r`),
		},
		{
			Value:            `\n`+"\u2028",
			ExpectedValue:    "\n",
			ExpectedSize: len(`\n`),
		},
		{
			Value:            `\;`+"\u2028",
			ExpectedValue:     ";",
			ExpectedSize: len(`\;`),
		},
		{
			Value:            `\#`+"\u2028",
			ExpectedValue:     "#",
			ExpectedSize: len(`\#`),
		},
		{
			Value:            `\=`+"\u2028",
			ExpectedValue:     "=",
			ExpectedSize: len(`\=`),
		},
		{
			Value:            `\:`+"\u2028",
			ExpectedValue:     ":",
			ExpectedSize: len(`\:`),
		},



		{
			Value:            `\0`+"\u2029",
			ExpectedValue:  "\x00",
			ExpectedSize: len(`\0`),
		},
		{
			Value:            `\a`+"\u2029",
			ExpectedValue:    "\a",
			ExpectedSize: len(`\a`),
		},
		{
			Value:            `\b`+"\u2029",
			ExpectedValue:    "\b",
			ExpectedSize: len(`\b`),
		},
		{
			Value:            `\t`+"\u2029",
			ExpectedValue:    "\t",
			ExpectedSize: len(`\t`),
		},
		{
			Value:            `\r`+"\u2029",
			ExpectedValue:    "\r",
			ExpectedSize: len(`\r`),
		},
		{
			Value:            `\n`+"\u2029",
			ExpectedValue:    "\n",
			ExpectedSize: len(`\n`),
		},
		{
			Value:            `\;`+"\u2029",
			ExpectedValue:     ";",
			ExpectedSize: len(`\;`),
		},
		{
			Value:            `\#`+"\u2029",
			ExpectedValue:     "#",
			ExpectedSize: len(`\#`),
		},
		{
			Value:            `\=`+"\u2029",
			ExpectedValue:     "=",
			ExpectedSize: len(`\=`),
		},
		{
			Value:            `\:`+"\u2029",
			ExpectedValue:     ":",
			ExpectedSize: len(`\:`),
		},



		{
			Value:           `\\\0\a\b\t\r\n\;\#\=\:apple.banana.cherry=one two three`,
			ExpectedValue:   "\\\x00\a\b\t\r\n;#=:apple.banana.cherry=one two three",
			ExpectedSize: len(`\\\0\a\b\t\r\n\;\#\=\:apple.banana.cherry=one two three`),
		},



		{
			Value:            "value۵ line 1\\\nline 2",
			ExpectedValue:    "value۵ line 1\nline 2",
			ExpectedSize: len("value۵ line 1\\\nline 2"),
		},
		{
			Value:            "value۵ line 1\\\rline 2",
			ExpectedValue:    "value۵ line 1\rline 2",
			ExpectedSize: len("value۵ line 1\\\rline 2"),
		},
		{
			Value:            "value۵ line 1\\\r\nline 2",
			ExpectedValue:    "value۵ line 1\r\nline 2",
			ExpectedSize: len("value۵ line 1\\\r\nline 2"),
		},
		{
			Value:            "value۵ line 1\\\vline 2",
			ExpectedValue:    "value۵ line 1\vline 2",
			ExpectedSize: len("value۵ line 1\\\vline 2"),
		},
		{
			Value:            "value۵ line 1\\\u0085line 2",
			ExpectedValue:    "value۵ line 1\u0085line 2",
			ExpectedSize: len("value۵ line 1\\\u0085line 2"),
		},
		{
			Value:            "value۵ line 1\\\u2028line 2",
			ExpectedValue:    "value۵ line 1\u2028line 2",
			ExpectedSize: len("value۵ line 1\\\u2028line 2"),
		},
		{
			Value:            "value۵ line 1\\\u2029line 2",
			ExpectedValue:    "value۵ line 1\u2029line 2",
			ExpectedSize: len("value۵ line 1\\\u2029line 2"),
		},



		{
			Value:            "value۵ line 1\\\nline 2\n",
			ExpectedValue:    "value۵ line 1\nline 2",
			ExpectedSize: len("value۵ line 1\\\nline 2"),
		},
		{
			Value:            "value۵ line 1\\\rline 2\r",
			ExpectedValue:    "value۵ line 1\rline 2",
			ExpectedSize: len("value۵ line 1\\\rline 2"),
		},
		{
			Value:            "value۵ line 1\\\r\nline 2\r\n",
			ExpectedValue:    "value۵ line 1\r\nline 2",
			ExpectedSize: len("value۵ line 1\\\r\nline 2"),
		},
		{
			Value:            "value۵ line 1\\\vline 2\v",
			ExpectedValue:    "value۵ line 1\vline 2",
			ExpectedSize: len("value۵ line 1\\\vline 2"),
		},
		{
			Value:            "value۵ line 1\\\u0085line 2\u0085",
			ExpectedValue:    "value۵ line 1\u0085line 2",
			ExpectedSize: len("value۵ line 1\\\u0085line 2"),
		},
		{
			Value:            "value۵ line 1\\\u2028line 2\u2028",
			ExpectedValue:    "value۵ line 1\u2028line 2",
			ExpectedSize: len("value۵ line 1\\\u2028line 2"),
		},
		{
			Value:            "value۵ line 1\\\u2029line 2\u2029",
			ExpectedValue:    "value۵ line 1\u2029line 2",
			ExpectedSize: len("value۵ line 1\\\u2029line 2"),
		},
	}


	for testNumber, test := range tests {

		runeScanner := utf8s.NewRuneScanner( strings.NewReader(test.Value) )

		token, actualSize, err := Read(runeScanner)
		if nil != err && io.EOF != err {
			t.Errorf("For test #%d, for %q did not expect an error, but actually got one: (%T) %q.", testNumber, test.Value, err, err)
			continue
		}

		if expected, actual := test.ExpectedValue, token.String(); expected != actual {
			t.Errorf("For test #%d, for %q expected %q, but actually got %q.", testNumber, test.Value, expected, actual)
			continue
		}
		if expected, actual := test.ExpectedSize, actualSize; expected != actual {
			t.Errorf("For test #%d, for %q expected %d, but actually got %d.", testNumber, test.Value, expected, actual)
			t.Errorf("EXPECTED value: %q", test.ExpectedValue)
			t.Errorf("ACTUAL   value  %q", token.String())
			t.Errorf("")
			continue
		}

		if io.EOF == err {
			continue
		}
	}
}
