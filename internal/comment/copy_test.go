package inicomment

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
			Value:            ";this is a comment",
			ExpectedValue:    ";this is a comment",
		},
		{
			Value:            ";this is a comment\n",
			ExpectedValue:    ";this is a comment\n",
		},
		{
			Value:            ";this is a comment\n\r",
			ExpectedValue:    ";this is a comment\n\r",
		},
		{
			Value:            ";this is a comment\r",
			ExpectedValue:    ";this is a comment\r",
		},
		{
			Value:            ";this is a comment\r\n",
			ExpectedValue:    ";this is a comment\r\n",
		},
		{
			Value:            ";this is a comment\u0085",
			ExpectedValue:    ";this is a comment\u0085",
		},
		{
			Value:            ";this is a comment\u2028",
			ExpectedValue:    ";this is a comment\u2028",
		},



		{
			Value:            "#this is a comment",
			ExpectedValue:    "#this is a comment",
		},
		{
			Value:            "#this is a comment\n",
			ExpectedValue:    "#this is a comment\n",
		},
		{
			Value:            "#this is a comment\n\r",
			ExpectedValue:    "#this is a comment\n\r",
		},
		{
			Value:            "#this is a comment\r",
			ExpectedValue:    "#this is a comment\r",
		},
		{
			Value:            "#this is a comment\r\n",
			ExpectedValue:    "#this is a comment\r\n",
		},
		{
			Value:            "#this is a comment\u0085",
			ExpectedValue:    "#this is a comment\u0085",
		},
		{
			Value:            "#this is a comment\u2028",
			ExpectedValue:    "#this is a comment\u2028",
		},



		{
			Value:            ";this is a comment\n[the section]",
			ExpectedValue:    ";this is a comment\n",
		},
		{
			Value:            ";this is a comment\n\r[the section]",
			ExpectedValue:    ";this is a comment\n\r",
		},
		{
			Value:            ";this is a comment\r[the section]",
			ExpectedValue:    ";this is a comment\r",
		},
		{
			Value:            ";this is a comment\r\n[the section]",
			ExpectedValue:    ";this is a comment\r\n",
		},
		{
			Value:            ";this is a comment\u0085[the section]",
			ExpectedValue:    ";this is a comment\u0085",
		},
		{
			Value:            ";this is a comment\u2028[the section]",
			ExpectedValue:    ";this is a comment\u2028",
		},



		{
			Value:            "#this is a comment\n[the section]",
			ExpectedValue:    "#this is a comment\n",
		},
		{
			Value:            "#this is a comment\n\r[the section]",
			ExpectedValue:    "#this is a comment\n\r",
		},
		{
			Value:            "#this is a comment\r[the section]",
			ExpectedValue:    "#this is a comment\r",
		},
		{
			Value:            "#this is a comment\r\n[the section]",
			ExpectedValue:    "#this is a comment\r\n",
		},
		{
			Value:            "#this is a comment\u0085[the section]",
			ExpectedValue:    "#this is a comment\u0085",
		},
		{
			Value:            "#this is a comment\u2028[the section]",
			ExpectedValue:    "#this is a comment\u2028",
		},



		{
			Value:            ";this is a comment\n[the section]\n",
			ExpectedValue   : ";this is a comment\n",
		},
		{
			Value:            ";this is a comment\n\r[the section]\n",
			ExpectedValue   : ";this is a comment\n\r",
		},
		{
			Value:            ";this is a comment\r[the section]\r",
			ExpectedValue:    ";this is a comment\r",
		},
		{
			Value:            ";this is a comment\r\n[the section]\r\n",
			ExpectedValue:    ";this is a comment\r\n",
		},
		{
			Value:            ";this is a comment\u0085[the section]\u0085",
			ExpectedValue:    ";this is a comment\u0085",
		},
		{
			Value:            ";this is a comment\u2028[the section]\u2028",
			ExpectedValue:    ";this is a comment\u2028",
		},



		{
			Value:            "#this is a comment\n[the section]\n",
			ExpectedValue:    "#this is a comment\n",
		},
		{
			Value:            "#this is a comment\n\r[the section]\n",
			ExpectedValue:    "#this is a comment\n\r",
		},
		{
			Value:            "#this is a comment\r[the section]\r",
			ExpectedValue:    "#this is a comment\r",
		},
		{
			Value:            "#this is a comment\r\n[the section]\r\n",
			ExpectedValue:    "#this is a comment\r\n",
		},
		{
			Value:            "#this is a comment\u0085[the section]\u0085",
			ExpectedValue:    "#this is a comment\u0085",
		},
		{
			Value:            "#this is a comment\u2028[the section]\u2028",
			ExpectedValue:    "#this is a comment\u2028",
		},



		{
			Value:            ";this is a comment\n[the section]\napple=one",
			ExpectedValue:    ";this is a comment\n",
		},
		{
			Value:            ";this is a comment\n\r[the section]\napple=one",
			ExpectedValue:    ";this is a comment\n\r",
		},
		{
			Value:            ";this is a comment\r[the section]\rapple=one",
			ExpectedValue:    ";this is a comment\r",
		},
		{
			Value:            ";this is a comment\r\n[the section]\r\napple=one",
			ExpectedValue:    ";this is a comment\r\n",
		},
		{
			Value:            ";this is a comment\u0085[the section]\u0085apple=one",
			ExpectedValue:    ";this is a comment\u0085",
		},
		{
			Value:            ";this is a comment\u2028[the section]\u2028apple=one",
			ExpectedValue:    ";this is a comment\u2028",
		},



		{
			Value:            "#this is a comment\n[the section]\napple=one",
			ExpectedValue:    "#this is a comment\n",
		},
		{
			Value:            "#this is a comment\n\r[the section]\napple=one",
			ExpectedValue:    "#this is a comment\n\r",
		},
		{
			Value:            "#this is a comment\r[the section]\rapple=one",
			ExpectedValue:    "#this is a comment\r",
		},
		{
			Value:            "#this is a comment\r\n[the section]\r\napple=one",
			ExpectedValue:    "#this is a comment\r\n",
		},
		{
			Value:            "#this is a comment\u0085[the section]\u0085apple=one",
			ExpectedValue:    "#this is a comment\u0085",
		},
		{
			Value:            "#this is a comment\u2028[the section]\u2028apple=one",
			ExpectedValue:    "#this is a comment\u2028",
		},



		{
			Value:            ";this is a comment\n[the section]\napple=one\n",
			ExpectedValue:    ";this is a comment\n",
		},
		{
			Value:            ";this is a comment\n\r[the section]\napple=one\n",
			ExpectedValue:    ";this is a comment\n\r",
		},
		{
			Value:            ";this is a comment\r[the section]\rapple=one\r",
			ExpectedValue:    ";this is a comment\r",
		},
		{
			Value:            ";this is a comment\r\n[the section]\r\napple=one\r\n",
			ExpectedValue:    ";this is a comment\r\n",
		},
		{
			Value:            ";this is a comment\u0085[the section]\u0085apple=one\u0085",
			ExpectedValue:    ";this is a comment\u0085",
		},
		{
			Value:            ";this is a comment\u2028[the section]\u2028apple=one\u2028",
			ExpectedValue:    ";this is a comment\u2028",
		},



		{
			Value:            "#this is a comment\n[the section]\napple=one\n",
			ExpectedValue:    "#this is a comment\n",
		},
		{
			Value:            "#this is a comment\n\r[the section]\napple=one\n",
			ExpectedValue:    "#this is a comment\n\r",
		},
		{
			Value:            "#this is a comment\r[the section]\rapple=one\r",
			ExpectedValue:    "#this is a comment\r",
		},
		{
			Value:            "#this is a comment\r\n[the section]\r\napple=one\r\n",
			ExpectedValue:    "#this is a comment\r\n",
		},
		{
			Value:            "#this is a comment\u0085[the section]\u0085apple=one\u0085",
			ExpectedValue:    "#this is a comment\u0085",
		},
		{
			Value:            "#this is a comment\u2028[the section]\u2028apple=one\u2028",
			ExpectedValue:    "#this is a comment\u2028",
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
			t.Errorf("For test #%d, the actual copied value is not what was expected.", testNumber)
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
			Value: " ; This comment has whitespace in front of it.",
		},
		{
			Value: " # This comment has whitespace in front of it.",
		},
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
