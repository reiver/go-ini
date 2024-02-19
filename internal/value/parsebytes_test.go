package inivalue_test

import (
	"testing"

	"github.com/reiver/go-ini/internal/value"
)

func TestParse(t *testing.T) {

	tests := []struct{
		String       string
		ExpectedValue string
		ExpectedSize int
	}{
		{
			String:        "value",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\n",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\n\r",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\r",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\r\n",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\u0085",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\u2028",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},



		{
			String:        "value\nline2=wow",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\n\rline2=wow",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\rline2=wow",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\r\nline2=wow",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\u0085line2=wow",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\u2028line2=wow",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},



		{
			String:        "value\nline2=wow\n",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\n\rline2=wow\n\r",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\rline2=wow\r",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\r\nline2=wow\r\n",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\u0085line2=wow\u0085",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},
		{
			String:        "value\u2028line2=wow\u2028",
			ExpectedValue: "value",
			ExpectedSize:  5,
		},









		{
			String:        "value\u06F5",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\n",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\n\r",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\r",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\r\n",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\u0085",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\u2028",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},



		{
			String:        "value\u06F5\nline2=wow",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\n\rline2=wow",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\rline2=wow",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\r\nline2=wow",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\u0085line2=wow",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\u2028line2=wow",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},



		{
			String:        "value\u06F5\nline2=wow\n",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\n\rline2=wow\n\r",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\rline2=wow\r",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\r\nline2=wow\r\n",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\u0085line2=wow\u0085",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},
		{
			String:        "value\u06F5\u2028line2=wow\u2028",
			ExpectedValue: "value\u06F5",
			ExpectedSize:  7,
		},









		{
			String:        "value\u06F5 🙂 \t ",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \n",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \n\r",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \r",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \r\n",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \u0085",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \u2028",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},



		{
			String:        "value\u06F5 🙂 \t \nline2=wow",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \n\rline2=wow",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \rline2=wow",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \r\nline2=wow",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \u0085line2=wow",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \u2028line2=wow",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},



		{
			String:        "value\u06F5 🙂 \t \nline2=wow\n",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \n\rline2=wow\n\r",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \rline2=wow\r",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \r\nline2=wow\r\n",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \u0085line2=wow\u0085",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
		{
			String:        "value\u06F5 🙂 \t \u2028line2=wow\u2028",
			ExpectedValue: "value\u06F5 🙂 \t ",
			ExpectedSize:  15,
		},
	}


	for testNumber, test := range tests {

		var bytes []byte = []byte(test.String)

		actualName, actualSize, err := inivalue.ParseBytes(bytes)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one", testNumber)
			t.Logf("ERROR: (%T) %q.", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.ExpectedValue
			actual   := actualName

			if expected != actual {
				t.Errorf("For test #%d, the actual 'value' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
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
