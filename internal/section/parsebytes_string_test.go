package inisection

import (
	"testing"
)

func TestParseBytes(t *testing.T) {

	tests := []struct{
		String string
		ExpectedSection string
		ExpectedSize int
	}{
		{
			String:           "[",
			ExpectedSection:   "",
			ExpectedSize: len("["),
		},
		{
			String:           "[\n",
			ExpectedSection:   "",
			ExpectedSize: len("[\n"),
		},
		{
			String:           "[\n\r",
			ExpectedSection:   "",
			ExpectedSize: len("[\n\r"),
		},
		{
			String:           "[\r",
			ExpectedSection:   "",
			ExpectedSize: len("[\r"),
		},
		{
			String:           "[\r\n",
			ExpectedSection:   "",
			ExpectedSize: len("[\r\n"),
		},
		{
			String:           "[\u0085",
			ExpectedSection:   "",
			ExpectedSize: len("[\u0085"),
		},
		{
			String:           "[\u2028",
			ExpectedSection:   "",
			ExpectedSize: len("[\u2028"),
		},



		{
			String:           "[hello world! 🙂",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂"),
		},
		{
			String:           "[hello world! 🙂\n",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\n"),
		},
		{
			String:           "[hello world! 🙂\n\r",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\n\r"),
		},
		{
			String:           "[hello world! 🙂\r",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\r"),
		},
		{
			String:           "[hello world! 🙂\r\n",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\r\n"),
		},
		{
			String:           "[hello world! 🙂\u0085",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\u0085"),
		},
		{
			String:           "[hello world! 🙂\u2028",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\u2028"),
		},



		{
			String:           "[hello world! 🙂]",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]"),
		},
		{
			String:           "[hello world! 🙂]\n",
			ExpectedSection  : "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\n"),
		},
		{
			String:           "[hello world! 🙂]\n\r",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\n\r"),
		},
		{
			String:           "[hello world! 🙂]\r",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\r"),
		},
		{
			String:           "[hello world! 🙂]\r\n",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\r\n"),
		},
		{
			String:           "[hello world! 🙂]\u0085",
			ExpectedSection  : "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\u0085"),
		},
		{
			String:           "[hello world! 🙂]\u2028",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\u2028"),
		},









		{
			String:           "[\nline2=something",
			ExpectedSection:   "",
			ExpectedSize: len("[\n"),
		},
		{
			String:           "[\n\rline2=something",
			ExpectedSection:   "",
			ExpectedSize: len("[\n\r"),
		},
		{
			String:           "[\rline2=something",
			ExpectedSection:   "",
			ExpectedSize: len("[\r"),
		},
		{
			String:           "[\r\nline2=something",
			ExpectedSection:   "",
			ExpectedSize: len("[\r\n"),
		},
		{
			String:           "[\u0085line2=something",
			ExpectedSection:   "",
			ExpectedSize: len("[\u0085"),
		},
		{
			String:           "[\u2028line2=something",
			ExpectedSection:   "",
			ExpectedSize: len("[\u2028"),
		},



		{
			String:           "[hello world! 🙂\nline2=something",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\n"),
		},
		{
			String:           "[hello world! 🙂\n\rline2=something",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\n\r"),
		},
		{
			String:           "[hello world! 🙂\rline2=something",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\r"),
		},
		{
			String:           "[hello world! 🙂\r\nline2=something",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\r\n"),
		},
		{
			String:           "[hello world! 🙂\u0085line2=something",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\u0085"),
		},
		{
			String:           "[hello world! 🙂\u2028line2=something",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂\u2028"),
		},



		{
			String:           "[hello world! 🙂]\nline2=something",
			ExpectedSection  : "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\n"),
		},
		{
			String:           "[hello world! 🙂]\n\rline2=something",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\n\r"),
		},
		{
			String:           "[hello world! 🙂]\rline2=something",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\r"),
		},
		{
			String:           "[hello world! 🙂]\r\nline2=something",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\r\n"),
		},
		{
			String:           "[hello world! 🙂]\u0085line2=something",
			ExpectedSection  : "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\u0085"),
		},
		{
			String:           "[hello world! 🙂]\u2028line2=something",
			ExpectedSection:   "hello world! 🙂",
			ExpectedSize: len("[hello world! 🙂]\u2028"),
		},
	}

	for testNumber, test := range tests {

		var p []byte = []byte(test.String)

		actualSection, actualSize, err := parseBytes(p)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.ExpectedSection
			actual   := actualSection

			if expected != actual {
				t.Errorf("For test #%d, the actual 'section' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING:  %q", test.String)
				continue
			}
		}

		{
			expected := test.ExpectedSize
			actual   := actualSize

			if expected != actual {
				t.Errorf("For test #%d, the actual 'size' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}
