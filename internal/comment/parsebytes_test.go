package inicomment_test

import (
	"testing"

	"github.com/reiver/go-ini/internal/comment"
)

func TestParseBytes(t *testing.T) {

	tests := []struct{
		String          string
		ExpectedComment string
		ExpectedSize    int
	}{
		{
			String:           ";",
			ExpectedComment:   "",
			ExpectedSize: len(";"),
		},
		{
			String:           ";\n",
			ExpectedComment:   "",
			ExpectedSize: len(";\n"),
		},
		{
			String:           ";\n\r",
			ExpectedComment:   "",
			ExpectedSize: len(";\n\r"),
		},
		{
			String:           ";\r",
			ExpectedComment:   "",
			ExpectedSize: len(";\r"),
		},
		{
			String:           ";\r\n",
			ExpectedComment:   "",
			ExpectedSize: len(";\r\n"),
		},
		{
			String:           ";\u0085",
			ExpectedComment:   "",
			ExpectedSize: len(";\u0085"),
		},
		{
			String:           ";\u2028",
			ExpectedComment:   "",
			ExpectedSize: len(";\u2028"),
		},



		{
			String:           "#",
			ExpectedComment:   "",
			ExpectedSize: len("#"),
		},
		{
			String:           "#\n",
			ExpectedComment:   "",
			ExpectedSize: len("#\n"),
		},
		{
			String:           "#\n\r",
			ExpectedComment:   "",
			ExpectedSize: len("#\n\r"),
		},
		{
			String:           "#\r",
			ExpectedComment:   "",
			ExpectedSize: len("#\r"),
		},
		{
			String:           "#\r\n",
			ExpectedComment:   "",
			ExpectedSize: len("#\r\n"),
		},
		{
			String:           "#\u0085",
			ExpectedComment:   "",
			ExpectedSize: len("#\u0085"),
		},
		{
			String:           "#\u2028",
			ExpectedComment:   "",
			ExpectedSize: len("#\u2028"),
		},



		{
			String:          ";\n[section]",
			ExpectedComment: "",
			ExpectedSize: len(";\n"),
		},
		{
			String:          ";\n\r[section]",
			ExpectedComment: "",
			ExpectedSize: len(";\r\n"),
		},
		{
			String:          ";\r[section]",
			ExpectedComment: "",
			ExpectedSize: len(";\r"),
		},
		{
			String:          ";\r\n[section]",
			ExpectedComment: "",
			ExpectedSize: len(";\r\n"),
		},
		{
			String:          ";\u0085[section]",
			ExpectedComment: "",
			ExpectedSize: len(";\u0085"),
		},
		{
			String:          ";\u2028[section]",
			ExpectedComment: "",
			ExpectedSize: len(";\u2028"),
		},



		{
			String:          "#\n[section]",
			ExpectedComment: "",
			ExpectedSize: len("#\n"),
		},
		{
			String:          "#\n\r[section]",
			ExpectedComment: "",
			ExpectedSize: len("#\n\r"),
		},
		{
			String:          "#\r[section]",
			ExpectedComment: "",
			ExpectedSize: len("#\r"),
		},
		{
			String:          "#\r\n[section]",
			ExpectedComment: "",
			ExpectedSize: len("#\r\n"),
		},
		{
			String:          "#\u0085[section]",
			ExpectedComment: "",
			ExpectedSize: len("#\u0085"),
		},
		{
			String:          "#\u2028[section]",
			ExpectedComment: "",
			ExpectedSize: len("#\u2028"),
		},









		{
			String:           ";this is a comment",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment"),
		},
		{
			String:           ";this is a comment\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\n"),
		},
		{
			String:           ";this is a comment\n\r",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\n\r"),
		},
		{
			String:           ";this is a comment\r",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\r"),
		},
		{
			String:           ";this is a comment\r\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\r\n"),
		},
		{
			String:           ";this is a comment\u0085",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\u0085"),
		},
		{
			String:           ";this is a comment\u2028",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\u2028"),
		},



		{
			String:           "#this is a comment",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment"),
		},
		{
			String:           "#this is a comment\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\n"),
		},
		{
			String:           "#this is a comment\n\r",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\n\r"),
		},
		{
			String:           "#this is a comment\r",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\r"),
		},
		{
			String:           "#this is a comment\r\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\r\n"),
		},
		{
			String:           "#this is a comment\u0085",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\u0085"),
		},
		{
			String:           "#this is a comment\u2028",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\u2028"),
		},



		{
			String:           ";this is a comment\n[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\n"),
		},
		{
			String:           ";this is a comment\n\r[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\n\r"),
		},
		{
			String:           ";this is a comment\r[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\r"),
		},
		{
			String:           ";this is a comment\r\n[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\r\n"),
		},
		{
			String:           ";this is a comment\u0085[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\u0085"),
		},
		{
			String:           ";this is a comment\u2028[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\u2028"),
		},



		{
			String:           "#this is a comment\n[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\n"),
		},
		{
			String:           "#this is a comment\n\r[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\n\r"),
		},
		{
			String:           "#this is a comment\r[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\r"),
		},
		{
			String:           "#this is a comment\r\n[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\r\n"),
		},
		{
			String:           "#this is a comment\u0085[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\u0085"),
		},
		{
			String:           "#this is a comment\u2028[the section]",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\u2028"),
		},



		{
			String:           ";this is a comment\n[the section]\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\n"),
		},
		{
			String:           ";this is a comment\n\r[the section]\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\n\r"),
		},
		{
			String:           ";this is a comment\r[the section]\r",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\r"),
		},
		{
			String:           ";this is a comment\r\n[the section]\r\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\r\n"),
		},
		{
			String:           ";this is a comment\u0085[the section]\u0085",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\u0085"),
		},
		{
			String:           ";this is a comment\u2028[the section]\u2028",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\u2028"),
		},



		{
			String:           "#this is a comment\n[the section]\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\n"),
		},
		{
			String:           "#this is a comment\n\r[the section]\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\n\r"),
		},
		{
			String:           "#this is a comment\r[the section]\r",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\r"),
		},
		{
			String:           "#this is a comment\r\n[the section]\r\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\r\n"),
		},
		{
			String:           "#this is a comment\u0085[the section]\u0085",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\u0085"),
		},
		{
			String:           "#this is a comment\u2028[the section]\u2028",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\u2028"),
		},



		{
			String:           ";this is a comment\n[the section]\napple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\n"),
		},
		{
			String:           ";this is a comment\n\r[the section]\napple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\n\n"),
		},
		{
			String:           ";this is a comment\r[the section]\rapple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\r"),
		},
		{
			String:           ";this is a comment\r\n[the section]\r\napple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\r\n"),
		},
		{
			String:           ";this is a comment\u0085[the section]\u0085apple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\u0085"),
		},
		{
			String:           ";this is a comment\u2028[the section]\u2028apple=one",
			ExpectedComment  : "this is a comment",
			ExpectedSize: len(";this is a comment\u2028"),
		},



		{
			String:           "#this is a comment\n[the section]\napple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\n"),
		},
		{
			String:           "#this is a comment\n\r[the section]\napple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\n\r"),
		},
		{
			String:           "#this is a comment\r[the section]\rapple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\r"),
		},
		{
			String:           "#this is a comment\r\n[the section]\r\napple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\r\n"),
		},
		{
			String:           "#this is a comment\u0085[the section]\u0085apple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\u0085"),
		},
		{
			String:           "#this is a comment\u2028[the section]\u2028apple=one",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\u2028"),
		},



		{
			String:           ";this is a comment\n[the section]\napple=one\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\n"),
		},
		{
			String:           ";this is a comment\n\r[the section]\napple=one\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\n\r"),
		},
		{
			String:           ";this is a comment\r[the section]\rapple=one\r",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\r"),
		},
		{
			String:           ";this is a comment\r\n[the section]\r\napple=one\r\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\r\n"),
		},
		{
			String:           ";this is a comment\u0085[the section]\u0085apple=one\u0085",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len(";this is a comment\u0085"),
		},
		{
			String:           ";this is a comment\u2028[the section]\u2028apple=one\u2028",
			ExpectedComment: "this is a comment",
			ExpectedSize: len(";this is a comment\u2028"),
		},



		{
			String:           "#this is a comment\n[the section]\napple=one\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\n"),
		},
		{
			String:          "#this is a comment\n\r[the section]\napple=one\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\n\r"),
		},
		{
			String:          "#this is a comment\r[the section]\rapple=one\r",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\r"),
		},
		{
			String:          "#this is a comment\r\n[the section]\r\napple=one\r\n",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\r\n"),
		},
		{
			String:          "#this is a comment\u0085[the section]\u0085apple=one\u0085",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\u0085"),
		},
		{
			String:           "#this is a comment\u2028[the section]\u2028apple=one\u2028",
			ExpectedComment:   "this is a comment",
			ExpectedSize: len("#this is a comment\u2028"),
		},
	}


	for testNumber, test := range tests {

		var p []byte = []byte(test.String)

		actualComment, actualSize, err := inicomment.ParseBytes(p)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one..", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.ExpectedComment
			actual   := actualComment

			if expected != actual {
				t.Errorf("For test #%d, the actual 'comment' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				t.Logf("EXPECTED-SIZE: %d", test.ExpectedSize)
				t.Logf("ACTUAL-SIZE:   %d", actualSize)
				continue
			}
		}

		{
			expected := test.ExpectedSize
			actual   := actualSize

			if expected != actual {
				t.Errorf("For test #%d, the actual 'size' is not what was expected." , testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}

func TestParseBytesError(t *testing.T) {

	tests := []struct{
		String string
	}{
		{
			String:"key=value",
		},
		{
			String:"[section]",
		},
		{
			String:"    \r\n",
		},
		{
			String:"\n\n[section]\napple=one\nbanana=two\ncherry=three\n",
		},



		{
			String:" ; This comment has whitespace in front of it.",
		},
		{
			String:" # This comment has whitespace in front of it.",
		},
	}


	for testNumber, test := range tests {

		var p []byte = []byte(test.String)

		actualComment, actualSize, err := inicomment.ParseBytes(p)

		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually got one.", testNumber)
			t.Logf("STRING: %q", test.String)
			t.Logf("ACTUAL-COMMENT: %q", actualComment)
			t.Logf("ACTUAL-SIZE: %d", actualSize)
			continue
		}
	}
}
