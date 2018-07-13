package iniscanner_sectionend

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
			Value:            "}",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},



		{
			Value:            "}\n",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\r",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\r\n",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\v",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\u0085",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\u2028",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\u2029",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},



		{
			Value:            "}\napple  = one\nbanana = two\ncherry = three\n}\n",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\rapple  = one\rbanana = two\rcherry = three\r}\r",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\r\napple  = one\r\nbanana = two\r\ncherry = three\r\n}\r\n",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\vapple  = one\vbanana = two\vcherry = three\v}\v",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\u0085apple  = one\u0085banana = two\u0085cherry = three\u0085}\u0085",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\u2028apple  = one\u2028banana = two\u2028cherry = three\u2028}\u2028",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
		},
		{
			Value:            "}\u2029apple  = one\u2029banana = two\u2029cherry = three\u2029}\u2029",
			ExpectedValue:    "}",
			ExpectedSize: len("}"),
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
			Value: "{",
		},



		{
			Value: "=value",
		},
		{
			Value: ":value",
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
