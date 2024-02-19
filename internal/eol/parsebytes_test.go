package inieol_test

import (
	"testing"

	"github.com/reiver/go-ini/internal/eol"
)

func TestParseBytes(t *testing.T) {

	tests := []struct{
		String string
		ExpectedSize int
	}{
		{
			String: "",
			ExpectedSize: 0,
		},



		{
			String: "apple",
			ExpectedSize: 0,
		},
		{
			String: "banana",
			ExpectedSize: 0,
		},
		{
			String: "cherry",
			ExpectedSize: 0,
		},



		{
			String: "\n",
			ExpectedSize: 1,
		},
		{
			String: "\n\r",
			ExpectedSize: 2,
		},
		{
			String: "\r",
			ExpectedSize: 1,
		},
		{
			String: "\r\n",
			ExpectedSize: 2,
		},
		{
			String: "\u0085",
			ExpectedSize: 2,
		},
		{
			String: "\u2028",
			ExpectedSize: 3,
		},



		{
			String: "\nline2=something",
			ExpectedSize: 1,
		},
		{
			String: "\n\rline2=something",
			ExpectedSize: 2,
		},
		{
			String: "\rline2=something",
			ExpectedSize: 1,
		},
		{
			String: "\r\nline2=something",
			ExpectedSize: 2,
		},
		{
			String: "\u0085line2=something",
			ExpectedSize: 2,
		},
		{
			String: "\u2028line2=something",
			ExpectedSize: 3,
		},



		{
			String: "\nline2=something\n",
			ExpectedSize: 1,
		},
		{
			String: "\n\rline2=something\n\r",
			ExpectedSize: 2,
		},
		{
			String: "\rline2=something\r",
			ExpectedSize: 1,
		},
		{
			String: "\r\nline2=something\r\n",
			ExpectedSize: 2,
		},
		{
			String: "\u0085line2=something\u0085",
			ExpectedSize: 2,
		},
		{
			String: "\u2028line2=something\u2028",
			ExpectedSize: 3,
		},









		{
			String: "\n\n",
			ExpectedSize: 1,
		},
		{
			String: "\n\r\n",
			ExpectedSize: 2,
		},
		//{
		//	String: "\r\n",
		//	ExpectedSize: 1,
		//},
		{
			String: "\r\n\n",
			ExpectedSize: 2,
		},
		{
			String: "\u0085\n",
			ExpectedSize: 2,
		},
		{
			String: "\u2028\n",
			ExpectedSize: 3,
		},



		{
			String: "\n\n\r",
			ExpectedSize: 1,
		},
		{
			String: "\n\r\n\r",
			ExpectedSize: 2,
		},
		//{
		//	String: "\r\n\r",
		//	ExpectedSize: 1,
		//},
		{
			String: "\r\n\n\r",
			ExpectedSize: 2,
		},
		{
			String: "\u0085\n\r",
			ExpectedSize: 2,
		},
		{
			String: "\u2028\n\r",
			ExpectedSize: 3,
		},



		//{
		//	String: "\n\r",
		//	ExpectedSize: 1,
		//},
		{
			String: "\n\r\r",
			ExpectedSize: 2,
		},
		{
			String: "\r\r",
			ExpectedSize: 1,
		},
		{
			String: "\r\n\r",
			ExpectedSize: 2,
		},
		{
			String: "\u0085\r",
			ExpectedSize: 2,
		},
		{
			String: "\u2028\r",
			ExpectedSize: 3,
		},



		//{
		//	String: "\n\r\n",
		//	ExpectedSize: 1,
		//},
		{
			String: "\n\r\r\n",
			ExpectedSize: 2,
		},
		{
			String: "\r\r\n",
			ExpectedSize: 1,
		},
		{
			String: "\r\n\r\n",
			ExpectedSize: 2,
		},
		{
			String: "\u0085\r\n",
			ExpectedSize: 2,
		},
		{
			String: "\u2028\r\n",
			ExpectedSize: 3,
		},



		{
			String: "\n\u0085",
			ExpectedSize: 1,
		},
		{
			String: "\n\r\u0085",
			ExpectedSize: 2,
		},
		{
			String: "\r\u0085",
			ExpectedSize: 1,
		},
		{
			String: "\r\n\u0085",
			ExpectedSize: 2,
		},
		{
			String: "\u0085\u0085",
			ExpectedSize: 2,
		},
		{
			String: "\u2028\u0085",
			ExpectedSize: 3,
		},



		{
			String: "\n\u2028",
			ExpectedSize: 1,
		},
		{
			String: "\n\r\u0085",
			ExpectedSize: 2,
		},
		{
			String: "\r\u0085",
			ExpectedSize: 1,
		},
		{
			String: "\r\n\u0085",
			ExpectedSize: 2,
		},
		{
			String: "\u0085\u0085",
			ExpectedSize: 2,
		},
		{
			String: "\u2028\u0085",
			ExpectedSize: 3,
		},
	}

	for testNumber, test := range tests {

		var p []byte = []byte(test.String)

		actual, err := inieol.ParseBytes(p)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.ExpectedSize

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
