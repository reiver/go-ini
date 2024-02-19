package ininamevalue_test

import (
	"testing"

	"github.com/reiver/go-ini/internal/namevalue"
)

func TestParseBytes(t *testing.T) {

	tests := []struct{
		String        string
		ExpectedName  string
		ExpectedValue string
		ExpectedSize  int
		ExpectedValueSize int
	}{
		{
			String:           "name value",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value"),
		},
		{
			String:           "name value\n",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\n"),
		},
		{
			String:           "name value\n\r",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\n\r"),
		},
		{
			String:           "name value\r",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\r"),
		},
		{
			String:           "name value\r\n",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\r\n"),
		},
		{
			String:           "name value\u0085",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\u0085"),
		},
		{
			String:           "name value\u2028",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\u2028"),
		},



		{
			String:           "name     value",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value"),
		},
		{
			String:           "name     value\n",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\n"),
		},
		{
			String:           "name     value\n\r",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\n\r"),
		},
		{
			String:           "name     value\r",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\r"),
		},
		{
			String:           "name     value\r\n",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\r\n"),
		},
		{
			String:           "name     value\u0085",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\u0085"),
		},
		{
			String:           "name     value\u2028",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\u2028"),
		},



		{
			String:           "name\tvalue",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue"),
		},
		{
			String:           "name\tvalue\n",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\n"),
		},
		{
			String:           "name\tvalue\n\r",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\n\r"),
		},
		{
			String:           "name\tvalue\r",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\r"),
		},
		{
			String:           "name\tvalue\r\n",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\r\n"),
		},
		{
			String:           "name\tvalue\u0085",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\u0085"),
		},
		{
			String:           "name\tvalue\u2028",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\u2028"),
		},



		{
			String:           "name\t\t\tvalue",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue"),
		},
		{
			String:           "name\t\t\tvalue\n",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\n"),
		},
		{
			String:           "name\t\t\tvalue\n\r",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\n\r"),
		},
		{
			String:           "name\t\t\tvalue\r",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\r"),
		},
		{
			String:           "name\t\t\tvalue\r\n",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\r\n"),
		},
		{
			String:           "name\t\t\tvalue\u0085",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\u0085"),
		},
		{
			String:           "name\t\t\tvalue\u2028",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\u2028"),
		},



		{
			String:           "name: value",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value"),
		},
		{
			String:           "name: value\n",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\n"),
		},
		{
			String:           "name: value\n\r",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\n\r"),
		},
		{
			String:           "name: value\r",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\r"),
		},
		{
			String:           "name: value\r\n",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\r\n"),
		},
		{
			String:           "name: value\u0085",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\u0085"),
		},
		{
			String:           "name: value\u2028",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\u2028"),
		},



		{
			String:           "name : value",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value"),
		},
		{
			String:           "name : value\n",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\n"),
		},
		{
			String:           "name : value\n\r",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\n\r"),
		},
		{
			String:           "name : value\r",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\r"),
		},
		{
			String:           "name : value\r\n",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\r\n"),
		},
		{
			String:           "name : value\u0085",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\u0085"),
		},
		{
			String:           "name : value\u2028",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\u2028"),
		},



		{
			String:           "name=value",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value"),
		},
		{
			String:           "name=value\n",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\n"),
		},
		{
			String:           "name=value\n\r",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\n\r"),
		},
		{
			String:           "name=value\r",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\r"),
		},
		{
			String:           "name=value\r\n",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\r\n"),
		},
		{
			String:           "name=value\u0085",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\u0085"),
		},
		{
			String:           "name=value\u2028",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\u2028"),
		},



		{
			String:           "name = value",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value"),
		},
		{
			String:           "name = value\n",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\n"),
		},
		{
			String:           "name = value\n\r",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\n\r"),
		},
		{
			String:           "name = value\r",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\r"),
		},
		{
			String:           "name = value\r\n",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\r\n"),
		},
		{
			String:           "name = value\u0085",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\u0085"),
		},
		{
			String:           "name = value\u2028",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\u2028"),
		},









		{
			String:           "name value\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\n"),
		},
		{
			String:           "name value\n\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\n\r"),
		},
		{
			String:           "name value\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\r"),
		},
		{
			String:           "name value\r\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\r\n"),
		},
		{
			String:           "name value\u0085line2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\u0085"),
		},
		{
			String:           "name value\u2028line2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\u2028"),
		},
	}


	for testNumber, test := range tests {

		var p []byte = []byte(test.String)

		actualName, actualValue, actualSize, err := ininamevalue.ParseBytes(p)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one", testNumber)
			t.Logf("ERROR: (%T) %q.", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.ExpectedName
			actual   := actualName

			if expected != actual {
				t.Errorf("For test #%d, the actual 'name' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			expected := test.ExpectedValue
			actual   := actualValue

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
