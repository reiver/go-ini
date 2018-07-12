package iniscanner_separator

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
			Value:            "=value",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value\n",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value\r",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value\r\n",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value\v",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value\u0085",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value\u2028",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value\u2029",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},



		{
			Value:            ":value",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value\n",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value\r",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value\r\n",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value\v",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value\u0085",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value\u2028",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value\u2029",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},



		{
			Value:            "=value ; comment",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value ; comment\n",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value ; comment\r",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value ; comment\r\n",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value ; comment\v",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value ; comment\u0085",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value ; comment\u2028",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value ; comment\u2029",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},



		{
			Value:            ":value ; comment",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value ; comment\n",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value ; comment\r",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value ; comment\r\n",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value ; comment\v",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value ; comment\u0085",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value ; comment\u2028",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value ; comment\u2029",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},



		{
			Value:            "=value # comment",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value # comment\n",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value # comment\r",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value # comment\r\n",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value # comment\v",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value # comment\u0085",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value # comment\u2028",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},
		{
			Value:            "=value # comment\u2029",
			ExpectedValue:    "=",
			ExpectedSize: len("="),
		},



		{
			Value:            ":value # comment",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value # comment\n",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value # comment\r",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value # comment\r\n",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value # comment\v",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value # comment\u0085",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value # comment\u2028",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
		{
			Value:            ":value # comment\u2029",
			ExpectedValue:    ":",
			ExpectedSize: len(":"),
		},
	}


	for testNumber, test := range tests {

		runeScanner := utf8s.NewRuneScanner( strings.NewReader(test.Value) )

		token, actualSize, err := Read(runeScanner)
		if nil != err && io.EOF != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q.", testNumber, err, err)
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
			Value: "key=value",
		},
		{
			Value: "[section]",
		},
		{
			Value: "    \r\n",
		},
		{
			Value: "\n\n[section]\napple=one\nbanana=two\ncherry=three\n",
		},



		{
			Value: " =value",
		},
		{
			Value: " :value",
		},
	}


	for testNumber, test := range tests {

		runeScanner := utf8s.NewRuneScanner( strings.NewReader(test.Value) )

		_, _, err := Read(runeScanner)
		if nil == err || io.EOF == err {
			t.Errorf("For test #%d, expected an error, but did not actually got one: (%T) %q.", testNumber, err, err)
			continue
		}
	}
}
