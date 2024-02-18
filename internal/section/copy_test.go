package inisection

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
			Value:            "[the section]",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},



		{
			Value:            "[the section]\n",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\r",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\r\n",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\v",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\u0085",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\u2028",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\u2029",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},



		{
			Value:            "[the section]\napple=one",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\rapple=one",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\r\napple=one",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\vapple=one",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\u0085apple=one",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\u2028apple=one",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\u2029apple=one",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},



		{
			Value:            "[the section]\napple=one\n",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\rapple=one\r",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\r\napple=one\r\n",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\vapple=one\v",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\u0085apple=one\u0085",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\u2028apple=one\u2028",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
		{
			Value:            "[the section]\u2029apple=one\u2029",
			ExpectedValue:    "[the section]",
			ExpectedSize: len("[the section]"),
		},
	}


	for testNumber, test := range tests {

		runeScanner := utf8s.NewRuneScanner( strings.NewReader(test.Value) )

		token, actualSize, err := Read(runeScanner)
		if nil != err && io.EOF != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q.", testNumber, err, err)
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
			Value: "section]",
		},



		{
			Value: "; this is a comment.",
		},
		{
			Value: "# this is a comment.",
		},
		{
			Value: "key=value",
		},
		{
			Value: "key:value",
		},
		{
			Value: "    \r\n",
		},



		{
			Value: "    [section]",
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
