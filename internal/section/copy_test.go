package inisection

import (
	"sourcecode.social/reiver/go-utf8"

	"io"
	"strings"

	"testing"
)

func TestCopy(t *testing.T) {

	tests := []struct{
		Value         string
		ExpectedValue string
	}{
		{
			Value:            "[the section]",
			ExpectedValue:    "[the section]",
		},



		{
			Value:            "[the section]\n",
			ExpectedValue:    "[the section]\n",
		},
		{
			Value:            "[the section]\n\r",
			ExpectedValue:    "[the section]\n\r",
		},
		{
			Value:            "[the section]\r",
			ExpectedValue:    "[the section]\r",
		},
		{
			Value:            "[the section]\r\n",
			ExpectedValue:    "[the section]\r\n",
		},
		{
			Value:            "[the section]\u0085",
			ExpectedValue:    "[the section]\u0085",
		},
		{
			Value:            "[the section]\u2028",
			ExpectedValue:    "[the section]\u2028",
		},



		{
			Value:            "[the section]\napple=one",
			ExpectedValue:    "[the section]\n",
		},
		{
			Value:            "[the section]\n\rapple=one",
			ExpectedValue:    "[the section]\n\r",
		},
		{
			Value:            "[the section]\rapple=one",
			ExpectedValue:    "[the section]\r",
		},
		{
			Value:            "[the section]\r\napple=one",
			ExpectedValue:    "[the section]\r\n",
		},
		{
			Value:            "[the section]\u0085apple=one",
			ExpectedValue:    "[the section]\u0085",
		},
		{
			Value:            "[the section]\u2028apple=one",
			ExpectedValue:    "[the section]\u2028",
		},



		{
			Value:            "[the section]\napple=one\n",
			ExpectedValue:    "[the section]\n",
		},
		{
			Value:            "[the section]\n\rapple=one\n",
			ExpectedValue:    "[the section]\n\r",
		},
		{
			Value:            "[the section]\rapple=one\r",
			ExpectedValue:    "[the section]\r",
		},
		{
			Value:            "[the section]\r\napple=one\r\n",
			ExpectedValue:    "[the section]\r\n",
		},
		{
			Value:            "[the section]\u0085apple=one\u0085",
			ExpectedValue:    "[the section]\u0085",
		},
		{
			Value:            "[the section]\u2028apple=one\u2028",
			ExpectedValue:    "[the section]\u2028",
		},
	}


	for testNumber, test := range tests {

		runeScanner := utf8.NewRuneScanner( strings.NewReader(test.Value) )

		var result strings.Builder

		err := Copy(&result, runeScanner)
		if nil != err && io.EOF != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q.", testNumber, err, err)
			continue
		}

		if expected, actual := test.ExpectedValue, result.String(); expected != actual {
			t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		if io.EOF == err {
			continue
		}
	}
}

func TestCopyError(t *testing.T) {

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



/*
		{
			Value: "[section",
		},
		{
			Value: "[section\n",
		},
		{
			Value: "[section\n\r",
		},
		{
			Value: "[section\r",
		},
		{
			Value: "[section\r\n",
		},
		{
			Value: "[section\u0085",
		},
		{
			Value: "[section\u2028",
		},
*/
	}

	for testNumber, test := range tests {

		runeScanner := utf8.NewRuneScanner( strings.NewReader(test.Value) )

		var result strings.Builder

		err := Copy(&result, runeScanner)
		if nil == err || io.EOF == err {
			t.Errorf("For test #%d, expected an error, but did not actually got one: (%T) %q.", testNumber, err, err)
			continue
		}
	}
}
