package iniscanner_comment

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
			Value:            ";this is a comment",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\n",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\r",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\r\n",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\v",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u0085",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u2028",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u2029",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},



		{
			Value:            "#this is a comment",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\n",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\r",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\r\n",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\v",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u0085",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u2028",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u2029",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},



		{
			Value:            ";this is a comment\n[the section]",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\r[the section]",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\r\n[the section]",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\v[the section]",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u0085[the section]",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u2028[the section]",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u2029[the section]",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},



		{
			Value:            "#this is a comment\n[the section]",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\r[the section]",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\r\n[the section]",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\v[the section]",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u0085[the section]",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u2028[the section]",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u2029[the section]",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},



		{
			Value:            ";this is a comment\n[the section]\n",
			ExpectedValue   : ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\r[the section]\r",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\r\n[the section]\r\n",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\v[the section]\v",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u0085[the section]\u0085",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u2028[the section]\u2028",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u2029[the section]\u2029",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},



		{
			Value:            "#this is a comment\n[the section]\n",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\r[the section]\r",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\r\n[the section]\r\n",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\v[the section]\v",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u0085[the section]\u0085",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u2028[the section]\u2028",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u2029[the section]\u2029",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},



		{
			Value:            ";this is a comment\n[the section]\napple=one",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\r[the section]\rapple=one",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\r\n[the section]\r\napple=one",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\v[the section]\vapple=one",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u0085[the section]\u0085apple=one",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u2028[the section]\u2028apple=one",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u2029[the section]\u2029apple=one",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},



		{
			Value:            "#this is a comment\n[the section]\napple=one",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\r[the section]\rapple=one",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\r\n[the section]\r\napple=one",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\v[the section]\vapple=one",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u0085[the section]\u0085apple=one",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u2028[the section]\u2028apple=one",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u2029[the section]\u2029apple=one",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},



		{
			Value:            ";this is a comment\n[the section]\napple=one\n",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\r[the section]\rapple=one\r",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\r\n[the section]\r\napple=one\r\n",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\v[the section]\vapple=one\v",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u0085[the section]\u0085apple=one\u0085",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u2028[the section]\u2028apple=one\u2028",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			Value:            ";this is a comment\u2029[the section]\u2029apple=one\u2029",
			ExpectedValue:    ";this is a comment",
			ExpectedSize: len(";this is a comment"),
		},



		{
			Value:            "#this is a comment\n[the section]\napple=one\n",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\r[the section]\rapple=one\r",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\r\n[the section]\r\napple=one\r\n",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\v[the section]\vapple=one\v",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u0085[the section]\u0085apple=one\u0085",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u2028[the section]\u2028apple=one\u2028",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			Value:            "#this is a comment\u2029[the section]\u2029apple=one\u2029",
			ExpectedValue:    "#this is a comment",
			ExpectedSize: len("#this is a comment"),
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
			Value: " ; This comment has whitespace in front of it.",
		},
		{
			Value: " # This comment has whitespace in front of it.",
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
