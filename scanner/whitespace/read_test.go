package iniscanner_whitespace

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
			Value:            "\u0009",
			ExpectedValue:    "\u0009", // horizontal tab
			ExpectedSize: len("\u0009"),
		},



		{ // 1
			Value:            "\u000A",
			ExpectedValue:    "\u000A", // line feed
			ExpectedSize: len("\u000A"),
		},



		{ // 2
			Value:            "\u000B",
			ExpectedValue:    "\u000B", // vertical tab
			ExpectedSize: len("\u000B"),
		},



		{ // 3
			Value:            "\u000C",
			ExpectedValue:    "\u000C", // form feed
			ExpectedSize: len("\u000C"),
		},



		{ // 4
			Value:            "\u000D",
			ExpectedValue:    "\u000D", // carriage return
			ExpectedSize: len("\u000D"),
		},



		{ // 5
			Value:            "\u0020",
			ExpectedValue:    "\u0020", // space
			ExpectedSize: len("\u0020"),
		},



		{ // 6
			Value:            "\u0085",
			ExpectedValue:    "\u0085", // next line
			ExpectedSize: len("\u0085"),
		},



		{ // 7
			Value:            "\u00A0",
			ExpectedValue:    "\u00A0", // no-break space
			ExpectedSize: len("\u00A0"),
		},



		{ // 8
			Value:            "\u1680",
			ExpectedValue:    "\u1680", // ogham space mark
			ExpectedSize: len("\u1680"),
		},



		{ // 9
			Value:            "\u180E",
			ExpectedValue:    "\u180E", // mongolian vowel separator
			ExpectedSize: len("\u180E"),
		},



		{ // 10
			Value:            "\u2000",
			ExpectedValue:    "\u2000", // en quad
			ExpectedSize: len("\u2000"),
		},



		{ // 11
			Value:            "\u2001",
			ExpectedValue:    "\u2001", // em quad
			ExpectedSize: len("\u2001"),
		},



		{ // 12
			Value:            "\u2002",
			ExpectedValue:    "\u2002", // en space
			ExpectedSize: len("\u2002"),
		},



		{ // 13
			Value:            "\u2003",
			ExpectedValue:    "\u2003", // em space
			ExpectedSize: len("\u2003"),
		},



		{ // 14
			Value:            "\u2004",
			ExpectedValue:    "\u2004", // three-per-em space
			ExpectedSize: len("\u2004"),
		},



		{ // 15
			Value:            "\u2005",
			ExpectedValue:    "\u2005", // four-per-em space
			ExpectedSize: len("\u2005"),
		},



		{ // 16
			Value:            "\u2006",
			ExpectedValue:    "\u2006", // six-per-em space
			ExpectedSize: len("\u2006"),
		},



		{ // 17
			Value:            "\u2007",
			ExpectedValue:    "\u2007", // figure space
			ExpectedSize: len("\u2007"),
		},



		{ // 18
			Value:            "\u2008",
			ExpectedValue:    "\u2008", // punctuation space
			ExpectedSize: len("\u2008"),
		},



		{ // 19
			Value:            "\u2009",
			ExpectedValue:    "\u2009", // thin space
			ExpectedSize: len("\u2009"),
		},



		{ // 20
			Value:            "\u200A",
			ExpectedValue:    "\u200A", // hair space
			ExpectedSize: len("\u200A"),
		},



		{ // 21
			Value:            "\u2028",
			ExpectedValue:    "\u2028", // line separator
			ExpectedSize: len("\u2028"),
		},



		{ // 22
			Value:            "\u2029",
			ExpectedValue:    "\u2029", // paragraph separator
			ExpectedSize: len("\u2029"),
		},



		{ // 23
			Value:            "\u202F",
			ExpectedValue:    "\u202F", // narrow no-break space
			ExpectedSize: len("\u202F"),
		},



		{ // 24
			Value:            "\u205F",
			ExpectedValue:    "\u205F", // medium mathematical space
			ExpectedSize: len("\u205F"),
		},



		{ // 25
			Value:            "\u3000",
			ExpectedValue:    "\u3000", // ideographic space
			ExpectedSize: len("\u3000"),
		},









		{ // 26
			Value:            "\u0009.",
			ExpectedValue:    "\u0009", // horizontal tab
			ExpectedSize: len("\u0009"),
		},



		{ // 27
			Value:            "\u000A.",
			ExpectedValue:    "\u000A", // line feed
			ExpectedSize: len("\u000A"),
		},



		{ // 28
			Value:            "\u000B.",
			ExpectedValue:    "\u000B", // vertical tab
			ExpectedSize: len("\u000B"),
		},



		{ // 29
			Value:            "\u000C.",
			ExpectedValue:    "\u000C", // form feed
			ExpectedSize: len("\u000C"),
		},



		{ // 30
			Value:            "\u000D.",
			ExpectedValue:    "\u000D", // carriage return
			ExpectedSize: len("\u000D"),
		},



		{ // 31
			Value:            "\u0020.",
			ExpectedValue:    "\u0020", // space
			ExpectedSize: len("\u0020"),
		},



		{ // 32
			Value:            "\u0085.",
			ExpectedValue:    "\u0085", // next line
			ExpectedSize: len("\u0085"),
		},



		{ // 33
			Value:            "\u00A0.",
			ExpectedValue:    "\u00A0", // no-break space
			ExpectedSize: len("\u00A0"),
		},



		{ // 34
			Value:            "\u1680.",
			ExpectedValue:    "\u1680", // ogham space mark
			ExpectedSize: len("\u1680"),
		},



		{ // 35
			Value:            "\u180E.",
			ExpectedValue:    "\u180E", // mongolian vowel separator
			ExpectedSize: len("\u180E"),
		},



		{ // 36
			Value:            "\u2000.",
			ExpectedValue:    "\u2000", // en quad
			ExpectedSize: len("\u2000"),
		},



		{ // 37
			Value:            "\u2001.",
			ExpectedValue:    "\u2001", // em quad
			ExpectedSize: len("\u2001"),
		},



		{ // 38
			Value:            "\u2002.",
			ExpectedValue:    "\u2002", // en space
			ExpectedSize: len("\u2002"),
		},



		{ // 39
			Value:            "\u2003.",
			ExpectedValue:    "\u2003", // em space
			ExpectedSize: len("\u2003"),
		},



		{ // 40
			Value:            "\u2004.",
			ExpectedValue:    "\u2004", // three-per-em space
			ExpectedSize: len("\u2004"),
		},



		{ // 41
			Value:            "\u2005.",
			ExpectedValue:    "\u2005", // four-per-em space
			ExpectedSize: len("\u2005"),
		},



		{ // 42
			Value:            "\u2006.",
			ExpectedValue:    "\u2006", // six-per-em space
			ExpectedSize: len("\u2006"),
		},



		{ // 43
			Value:            "\u2007.",
			ExpectedValue:    "\u2007", // figure space
			ExpectedSize: len("\u2007"),
		},



		{ // 44
			Value:            "\u2008.",
			ExpectedValue:    "\u2008", // punctuation space
			ExpectedSize: len("\u2008"),
		},



		{ // 45
			Value:            "\u2009.",
			ExpectedValue:    "\u2009", // thin space
			ExpectedSize: len("\u2009"),
		},



		{ // 46
			Value:            "\u200A.",
			ExpectedValue:    "\u200A", // hair space
			ExpectedSize: len("\u200A"),
		},



		{ // 47
			Value:            "\u2028.",
			ExpectedValue:    "\u2028", // line separator
			ExpectedSize: len("\u2028"),
		},



		{ // 48
			Value:            "\u2029.",
			ExpectedValue:    "\u2029", // paragraph separator
			ExpectedSize: len("\u2029"),
		},



		{ // 49
			Value:            "\u202F.",
			ExpectedValue:    "\u202F", // narrow no-break space
			ExpectedSize: len("\u202F"),
		},



		{ // 50
			Value:            "\u205F.",
			ExpectedValue:    "\u205F", // medium mathematical space
			ExpectedSize: len("\u205F"),
		},



		{ // 51
			Value:            "\u3000.",
			ExpectedValue:    "\u3000", // ideographic space
			ExpectedSize: len("\u3000"),
		},









		{ // 52
			Value:            "\r\n\r\n    \r\n",
			ExpectedValue:    "\r\n\r\n    \r\n",
			ExpectedSize: len("\r\n\r\n    \r\n"),
		},
		{ // 53
			Value:            "\r\n\r\n    \r\n[the section]",
			ExpectedValue:    "\r\n\r\n    \r\n",
			ExpectedSize: len("\r\n\r\n    \r\n"),
		},
		{ // 53
			Value:            "\r\n\r\n    \r\n[the section]\r\n\r\nkey=value\r\n",
			ExpectedValue:    "\r\n\r\n    \r\n",
			ExpectedSize: len("\r\n\r\n    \r\n"),
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
			Value: "[section]",
		},
		{
			Value: "key=value",
		},
		{
			Value: "key:value",
		},
		{
			Value: "; comment",
		},
		{
			Value: "# comment",
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
