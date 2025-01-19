package iniampersandvalue_test

import (
	"testing"

	"github.com/reiver/go-ini/internal/ampersandvalue"
)

func TestParseBytes(t *testing.T) {

	tests := []struct{
		End           string
		String        string
		ExpectedValue string
		ExpectedSize  int
	}{
		{
			End: "END",
			String:           "value",
			ExpectedValue:    "value",
			ExpectedSize: len("value"),
		},
		{
			End:                   "END",
			String:           "valueEND",
			ExpectedValue:    "value",
			ExpectedSize: len("valueEND"),
		},
		{
			End: "END",
			String:           "value\n",
			ExpectedValue:    "value\n",
			ExpectedSize: len("value\n"),
		},
		{
			End:                     "END",
			String:           "value\nEND",
			ExpectedValue:    "value\n",
			ExpectedSize: len("value\nEND"),
		},
		{
			End:                     "END",
			String:           "value\nEND\n",
			ExpectedValue:    "value\n",
			ExpectedSize: len("value\nEND"),
		},
		{
			End:                     "END",
			String:           "value\nEND\nwow",
			ExpectedValue:    "value\n",
			ExpectedSize: len("value\nEND"),
		},


		{
			End: "END",
			String:           "value",
			ExpectedValue:    "value",
			ExpectedSize: len("value"),
		},
		{
			End:                   "END",
			String:           "valueEND",
			ExpectedValue:    "value",
			ExpectedSize: len("valueEND"),
		},
		{
			End: "END",
			String:           "value\n\r",
			ExpectedValue:    "value\n\r",
			ExpectedSize: len("value\n\r"),
		},
		{
			End:                       "END",
			String:           "value\n\rEND",
			ExpectedValue:    "value\n\r",
			ExpectedSize: len("value\n\rEND"),
		},
		{
			End:                       "END",
			String:           "value\n\rEND\n\r",
			ExpectedValue:    "value\n\r",
			ExpectedSize: len("value\n\rEND"),
		},
		{
			End:                       "END",
			String:           "value\n\rEND\n\rwow",
			ExpectedValue:    "value\n\r",
			ExpectedSize: len("value\n\rEND"),
		},


		{
			End: "END",
			String:           "value",
			ExpectedValue:    "value",
			ExpectedSize: len("value"),
		},
		{
			End:                   "END",
			String:           "valueEND",
			ExpectedValue:    "value",
			ExpectedSize: len("valueEND"),
		},
		{
			End: "END",
			String:           "value\r",
			ExpectedValue:    "value\r",
			ExpectedSize: len("value\r"),
		},
		{
			End:                     "END",
			String:           "value\rEND",
			ExpectedValue:    "value\r",
			ExpectedSize: len("value\rEND"),
		},
		{
			End:                     "END",
			String:           "value\rEND\r",
			ExpectedValue:    "value\r",
			ExpectedSize: len("value\rEND"),
		},
		{
			End                    : "END",
			String:           "value\rEND\rwow",
			ExpectedValue:    "value\r",
			ExpectedSize: len("value\rEND"),
		},


		{
			End: "END",
			String:           "value",
			ExpectedValue:    "value",
			ExpectedSize: len("value"),
		},
		{
			End:                   "END",
			String:           "valueEND",
			ExpectedValue:    "value",
			ExpectedSize: len("valueEND"),
		},
		{
			End: "END",
			String:           "value\r\n",
			ExpectedValue:    "value\r\n",
			ExpectedSize: len("value\r\n"),
		},
		{
			End:                       "END",
			String:           "value\r\nEND",
			ExpectedValue:    "value\r\n",
			ExpectedSize: len("value\r\nEND"),
		},
		{
			End:                       "END",
			String:           "value\r\nEND\r\n",
			ExpectedValue:    "value\r\n",
			ExpectedSize: len("value\r\nEND"),
		},
		{
			End:                       "END",
			String:           "value\r\nEND\r\nwow",
			ExpectedValue:    "value\r\n",
			ExpectedSize: len("value\r\nEND"),
		},


		{
			End: "END",
			String:           "value",
			ExpectedValue:    "value",
			ExpectedSize: len("value"),
		},
		{
			End:                   "END",
			String:           "valueEND",
			ExpectedValue:    "value",
			ExpectedSize: len("valueEND"),
		},
		{
			End: "END",
			String:           "value\u0085",
			ExpectedValue:    "value\u0085",
			ExpectedSize: len("value\u0085"),
		},
		{
			End:                         "END",
			String:           "value\u0085END",
			ExpectedValue:    "value\u0085",
			ExpectedSize: len("value\u0085END"),
		},
		{
			End:                         "END",
			String:           "value\u0085END\u0085",
			ExpectedValue:    "value\u0085",
			ExpectedSize: len("value\u0085END"),
		},
		{
			End:                         "END",
			String:           "value\u0085END\u0085wow",
			ExpectedValue:    "value\u0085",
			ExpectedSize: len("value\u0085END"),
		},









		{
			End: "Banana",
			String:           "line 1\nline 2\nlin3\n",
			ExpectedValue:    "line 1\nline 2\nlin3\n",
			ExpectedSize: len("line 1\nline 2\nlin3\n"),
		},
		{
			End:                                    "Banana",
			String:           "line 1\nline 2\nlin3\nBanana",
			ExpectedValue:    "line 1\nline 2\nlin3\n",
			ExpectedSize: len("line 1\nline 2\nlin3\nBanana"),
		},
		{
			End:                                    "Banana",
			String:           "line 1\nline 2\nlin3\nBanana\n",
			ExpectedValue:    "line 1\nline 2\nlin3\n",
			ExpectedSize: len("line 1\nline 2\nlin3\nBanana"),
		},
		{
			End:                                    "Banana",
			String:           "line 1\nline 2\nlin3\nBanana\nCHERRY = red",
			ExpectedValue:    "line 1\nline 2\nlin3\n",
			ExpectedSize: len("line 1\nline 2\nlin3\nBanana"),
		},
	}

	for testNumber, test := range tests {

		var p []byte = []byte(test.String)

		actualValue, actualSize, err := iniampersandvalue.ParseBytes(test.End, p)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one", testNumber)
			t.Logf("ERROR: (%T) %q.", err, err)
			t.Logf("STRING: %q", test.String)
			continue
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
