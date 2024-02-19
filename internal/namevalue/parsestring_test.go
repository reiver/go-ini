package ininamevalue

import (
	"io"

	"testing"
)

func TestParseString(t *testing.T) {

	tests := []struct{
		String        string
		ExpectedName  string
		ExpectedValue string
	}{
		{
			String:       "name value",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name value\n",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name value\n\r",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name value\r",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name value\r\n",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name value\u0085",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name value\u2028",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},



		{
			String:       "name     value",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name     value\n",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name     value\n\r",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name     value\r",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name     value\r\n",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name     value\u0085",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name     value\u2028",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},



		{
			String:       "name\tvalue",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\tvalue\n",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\tvalue\n\r",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\tvalue\r",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\tvalue\r\n",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\tvalue\u0085",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\tvalue\u2028",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},



		{
			String:       "name\t\t\tvalue",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\t\t\tvalue\n",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\t\t\tvalue\n\r",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\t\t\tvalue\r",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\t\t\tvalue\r\n",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\t\t\tvalue\u0085",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name\t\t\tvalue\u2028",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},



		{
			String:       "name: value",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name: value\n",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name: value\n\r",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name: value\r",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name: value\r\n",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name: value\u0085",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},
		{
			String:       "name: value\u2028",
			ExpectedName: "name",
			ExpectedValue:      "value",
		},



		{
			String:       "name : value",
			ExpectedName: "name",
			ExpectedValue:       "value",
		},
		{
			String:       "name : value\n",
			ExpectedName: "name",
			ExpectedValue:       "value",
		},
		{
			String:       "name : value\n\r",
			ExpectedName: "name",
			ExpectedValue:       "value",
		},
		{
			String:       "name : value\r",
			ExpectedName: "name",
			ExpectedValue:       "value",
		},
		{
			String:       "name : value\r\n",
			ExpectedName: "name",
			ExpectedValue:       "value",
		},
		{
			String:       "name : value\u0085",
			ExpectedName: "name",
			ExpectedValue:       "value",
		},
		{
			String:       "name : value\u2028",
			ExpectedName: "name",
			ExpectedValue:       "value",
		},



		{
			String:       "name=value",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name=value\n",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name=value\n\r",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name=value\r",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name=value\r\n",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name=value\u0085",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name=value\u2028",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},



		{
			String:       "name = value",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name = value\n",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name = value\n\r",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name = value\r",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name = value\r\n",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name = value\u0085",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
		{
			String:       "name = value\u2028",
			ExpectedName: "name",
			ExpectedValue:     "value",
		},
	}


	for testNumber, test := range tests {

		actualName, actualValue, err := ParseString(test.String)
		if nil != err && io.EOF != err {
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
	}
}
