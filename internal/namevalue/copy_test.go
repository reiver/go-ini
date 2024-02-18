package ininamevalue

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
			Value:            "name value",
			ExpectedValue:    "name value",
		},
		{
			Value:            "name value\n",
			ExpectedValue:    "name value\n",
		},
		{
			Value:            "name value\n\r",
			ExpectedValue:    "name value\n\r",
		},
		{
			Value:            "name value\r",
			ExpectedValue:    "name value\r",
		},
		{
			Value:            "name value\r\n",
			ExpectedValue:    "name value\r\n",
		},
		{
			Value:            "name value\u0085",
			ExpectedValue:    "name value\u0085",
		},
		{
			Value:            "name value\u2028",
			ExpectedValue:    "name value\u2028",
		},



		{
			Value:            "name: value",
			ExpectedValue:    "name: value",
		},
		{
			Value:            "name: value\n",
			ExpectedValue:    "name: value\n",
		},
		{
			Value:            "name: value\n\r",
			ExpectedValue:    "name: value\n\r",
		},
		{
			Value:            "name: value\r",
			ExpectedValue:    "name: value\r",
		},
		{
			Value:            "name: value\r\n",
			ExpectedValue:    "name: value\r\n",
		},
		{
			Value:            "name: value\u0085",
			ExpectedValue:    "name: value\u0085",
		},
		{
			Value:            "name: value\u2028",
			ExpectedValue:    "name: value\u2028",
		},



		{
			Value:            "name=value",
			ExpectedValue:    "name=value",
		},
		{
			Value:            "name=value\n",
			ExpectedValue:    "name=value\n",
		},
		{
			Value:            "name=value\n\r",
			ExpectedValue:    "name=value\n\r",
		},
		{
			Value:            "name=value\r",
			ExpectedValue:    "name=value\r",
		},
		{
			Value:            "name=value\r\n",
			ExpectedValue:    "name=value\r\n",
		},
		{
			Value:            "name=value\u0085",
			ExpectedValue:    "name=value\u0085",
		},
		{
			Value:            "name=value\u2028",
			ExpectedValue:    "name=value\u2028",
		},



		{
			Value:            "name = value",
			ExpectedValue:    "name = value",
		},
		{
			Value:            "name = value\n",
			ExpectedValue:    "name = value\n",
		},
		{
			Value:            "name = value\n\r",
			ExpectedValue:    "name = value\n\r",
		},
		{
			Value:            "name = value\r",
			ExpectedValue:    "name = value\r",
		},
		{
			Value:            "name = value\r\n",
			ExpectedValue:    "name = value\r\n",
		},
		{
			Value:            "name = value\u0085",
			ExpectedValue:    "name = value\u0085",
		},
		{
			Value:            "name = value\u2028",
			ExpectedValue:    "name = value\u2028",
		},



		{
			Value:            "name value\n[the section]",
			ExpectedValue:    "name value\n",
		},
		{
			Value:            "name value\n\r[the section]",
			ExpectedValue:    "name value\n\r",
		},
		{
			Value:            "name value\r[the section]",
			ExpectedValue:    "name value\r",
		},
		{
			Value:            "name value\r\n[the section]",
			ExpectedValue:    "name value\r\n",
		},
		{
			Value:            "name value\u0085[the section]",
			ExpectedValue:    "name value\u0085",
		},
		{
			Value:            "name value\u2028[the section]",
			ExpectedValue:    "name value\u2028",
		},



		{
			Value:            "name: value\n[the section]",
			ExpectedValue:    "name: value\n",
		},
		{
			Value:            "name: value\n\r[the section]",
			ExpectedValue:    "name: value\n\r",
		},
		{
			Value:            "name: value\r[the section]",
			ExpectedValue:    "name: value\r",
		},
		{
			Value:            "name: value\r\n[the section]",
			ExpectedValue:    "name: value\r\n",
		},
		{
			Value:            "name: value\u0085[the section]",
			ExpectedValue:    "name: value\u0085",
		},
		{
			Value:            "name: value\u2028[the section]",
			ExpectedValue:    "name: value\u2028",
		},



		{
			Value:            "name=value\n[the section]",
			ExpectedValue:    "name=value\n",
		},
		{
			Value:            "name=value\n\r[the section]",
			ExpectedValue:    "name=value\n\r",
		},
		{
			Value:            "name=value\r[the section]",
			ExpectedValue:    "name=value\r",
		},
		{
			Value:            "name=value\r\n[the section]",
			ExpectedValue:    "name=value\r\n",
		},
		{
			Value:            "name=value\u0085[the section]",
			ExpectedValue:    "name=value\u0085",
		},
		{
			Value:            "name=value\u2028[the section]",
			ExpectedValue:    "name=value\u2028",
		},



		{
			Value:            "name = value\n[the section]",
			ExpectedValue:    "name = value\n",
		},
		{
			Value:            "name = value\n\r[the section]",
			ExpectedValue:    "name = value\n\r",
		},
		{
			Value:            "name = value\r[the section]",
			ExpectedValue:    "name = value\r",
		},
		{
			Value:            "name = value\r\n[the section]",
			ExpectedValue:    "name = value\r\n",
		},
		{
			Value:            "name = value\u0085[the section]",
			ExpectedValue:    "name = value\u0085",
		},
		{
			Value:            "name = value\u2028[the section]",
			ExpectedValue:    "name = value\u2028",
		},



		{
			Value:            "name value\n[the section]\n",
			ExpectedValue   : "name value\n",
		},
		{
			Value:            "name value\n\r[the section]\n",
			ExpectedValue   : "name value\n\r",
		},
		{
			Value:            "name value\r[the section]\r",
			ExpectedValue:    "name value\r",
		},
		{
			Value:            "name value\r\n[the section]\r\n",
			ExpectedValue:    "name value\r\n",
		},
		{
			Value:            "name value\u0085[the section]\u0085",
			ExpectedValue:    "name value\u0085",
		},
		{
			Value:            "name value\u2028[the section]\u2028",
			ExpectedValue:    "name value\u2028",
		},



		{
			Value:            "name: value\n[the section]\n",
			ExpectedValue   : "name: value\n",
		},
		{
			Value:            "name: value\n\r[the section]\n",
			ExpectedValue   : "name: value\n\r",
		},
		{
			Value:            "name: value\r[the section]\r",
			ExpectedValue:    "name: value\r",
		},
		{
			Value:            "name: value\r\n[the section]\r\n",
			ExpectedValue:    "name: value\r\n",
		},
		{
			Value:            "name: value\u0085[the section]\u0085",
			ExpectedValue:    "name: value\u0085",
		},
		{
			Value:            "name: value\u2028[the section]\u2028",
			ExpectedValue:    "name: value\u2028",
		},



		{
			Value:            "name=value\n[the section]\n",
			ExpectedValue:    "name=value\n",
		},
		{
			Value:            "name=value\n\r[the section]\n",
			ExpectedValue:    "name=value\n\r",
		},
		{
			Value:            "name=value\r[the section]\r",
			ExpectedValue:    "name=value\r",
		},
		{
			Value:            "name=value\r\n[the section]\r\n",
			ExpectedValue:    "name=value\r\n",
		},
		{
			Value:            "name=value\u0085[the section]\u0085",
			ExpectedValue:    "name=value\u0085",
		},
		{
			Value:            "name=value\u2028[the section]\u2028",
			ExpectedValue:    "name=value\u2028",
		},



		{
			Value:            "name = value\n[the section]\n",
			ExpectedValue:    "name = value\n",
		},
		{
			Value:            "name = value\n\r[the section]\n",
			ExpectedValue:    "name = value\n\r",
		},
		{
			Value:            "name = value\r[the section]\r",
			ExpectedValue:    "name = value\r",
		},
		{
			Value:            "name = value\r\n[the section]\r\n",
			ExpectedValue:    "name = value\r\n",
		},
		{
			Value:            "name = value\u0085[the section]\u0085",
			ExpectedValue:    "name = value\u0085",
		},
		{
			Value:            "name = value\u2028[the section]\u2028",
			ExpectedValue:    "name = value\u2028",
		},



		{
			Value:            "name value\n[the section]\napple=one",
			ExpectedValue:    "name value\n",
		},
		{
			Value:            "name value\n\r[the section]\napple=one",
			ExpectedValue:    "name value\n\r",
		},
		{
			Value:            "name value\r[the section]\rapple=one",
			ExpectedValue:    "name value\r",
		},
		{
			Value:            "name value\r\n[the section]\r\napple=one",
			ExpectedValue:    "name value\r\n",
		},
		{
			Value:            "name value\u0085[the section]\u0085apple=one",
			ExpectedValue:    "name value\u0085",
		},
		{
			Value:            "name value\u2028[the section]\u2028apple=one",
			ExpectedValue:    "name value\u2028",
		},



		{
			Value:            "name: value\n[the section]\napple=one",
			ExpectedValue:    "name: value\n",
		},
		{
			Value:            "name: value\n\r[the section]\napple=one",
			ExpectedValue:    "name: value\n\r",
		},
		{
			Value:            "name: value\r[the section]\rapple=one",
			ExpectedValue:    "name: value\r",
		},
		{
			Value:            "name: value\r\n[the section]\r\napple=one",
			ExpectedValue:    "name: value\r\n",
		},
		{
			Value:            "name: value\u0085[the section]\u0085apple=one",
			ExpectedValue:    "name: value\u0085",
		},
		{
			Value:            "name: value\u2028[the section]\u2028apple=one",
			ExpectedValue:    "name: value\u2028",
		},



		{
			Value:            "name=value\n[the section]\napple=one",
			ExpectedValue:    "name=value\n",
		},
		{
			Value:            "name=value\n\r[the section]\napple=one",
			ExpectedValue:    "name=value\n\r",
		},
		{
			Value:            "name=value\r[the section]\rapple=one",
			ExpectedValue:    "name=value\r",
		},
		{
			Value:            "name=value\r\n[the section]\r\napple=one",
			ExpectedValue:    "name=value\r\n",
		},
		{
			Value:            "name=value\u0085[the section]\u0085apple=one",
			ExpectedValue:    "name=value\u0085",
		},
		{
			Value:            "name=value\u2028[the section]\u2028apple=one",
			ExpectedValue:    "name=value\u2028",
		},



		{
			Value:            "name = value\n[the section]\napple = one",
			ExpectedValue:    "name = value\n",
		},
		{
			Value:            "name = value\n\r[the section]\napple = one",
			ExpectedValue:    "name = value\n\r",
		},
		{
			Value:            "name = value\r[the section]\rapple = one",
			ExpectedValue:    "name = value\r",
		},
		{
			Value:            "name = value\r\n[the section]\r\napple = one",
			ExpectedValue:    "name = value\r\n",
		},
		{
			Value:            "name = value\u0085[the section]\u0085apple = one",
			ExpectedValue:    "name = value\u0085",
		},
		{
			Value:            "name = value\u2028[the section]\u2028apple = one",
			ExpectedValue:    "name = value\u2028",
		},



		{
			Value:            "name value\n[the section]\napple=one\n",
			ExpectedValue:    "name value\n",
		},
		{
			Value:            "name value\n\r[the section]\napple=one\n",
			ExpectedValue:    "name value\n\r",
		},
		{
			Value:            "name value\r[the section]\rapple=one\r",
			ExpectedValue:    "name value\r",
		},
		{
			Value:            "name value\r\n[the section]\r\napple=one\r\n",
			ExpectedValue:    "name value\r\n",
		},
		{
			Value:            "name value\u0085[the section]\u0085apple=one\u0085",
			ExpectedValue:    "name value\u0085",
		},
		{
			Value:            "name value\u2028[the section]\u2028apple=one\u2028",
			ExpectedValue:    "name value\u2028",
		},



		{
			Value:            "name: value\n[the section]\napple=one\n",
			ExpectedValue:    "name: value\n",
		},
		{
			Value:            "name: value\n\r[the section]\napple=one\n",
			ExpectedValue:    "name: value\n\r",
		},
		{
			Value:            "name: value\r[the section]\rapple=one\r",
			ExpectedValue:    "name: value\r",
		},
		{
			Value:            "name: value\r\n[the section]\r\napple=one\r\n",
			ExpectedValue:    "name: value\r\n",
		},
		{
			Value:            "name: value\u0085[the section]\u0085apple=one\u0085",
			ExpectedValue:    "name: value\u0085",
		},
		{
			Value:            "name: value\u2028[the section]\u2028apple=one\u2028",
			ExpectedValue:    "name: value\u2028",
		},



		{
			Value:            "name=value\n[the section]\napple=one\n",
			ExpectedValue:    "name=value\n",
		},
		{
			Value:            "name=value\n\r[the section]\napple=one\n",
			ExpectedValue:    "name=value\n\r",
		},
		{
			Value:            "name=value\r[the section]\rapple=one\r",
			ExpectedValue:    "name=value\r",
		},
		{
			Value:            "name=value\r\n[the section]\r\napple=one\r\n",
			ExpectedValue:    "name=value\r\n",
		},
		{
			Value:            "name=value\u0085[the section]\u0085apple=one\u0085",
			ExpectedValue:    "name=value\u0085",
		},
		{
			Value:            "name=value\u2028[the section]\u2028apple=one\u2028",
			ExpectedValue:    "name=value\u2028",
		},



		{
			Value:            "name = value\n[the section]\napple = one\n",
			ExpectedValue:    "name = value\n",
		},
		{
			Value:            "name = value\n\r[the section]\napple = one\n",
			ExpectedValue:    "name = value\n\r",
		},
		{
			Value:            "name = value\r[the section]\rapple = one\r",
			ExpectedValue:    "name = value\r",
		},
		{
			Value:            "name = value\r\n[the section]\r\napple = one\r\n",
			ExpectedValue:    "name = value\r\n",
		},
		{
			Value:            "name = value\u0085[the section]\u0085apple = one\u0085",
			ExpectedValue:    "name = value\u0085",
		},
		{
			Value:            "name = value\u2028[the section]\u2028apple = one\u2028",
			ExpectedValue:    "name = value\u2028",
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
