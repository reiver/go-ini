package ininame_test

import (
	"testing"

	"github.com/reiver/go-ini/internal/name"
)

func TestParse(t *testing.T) {

	tests := []struct{
		String       string
		ExpectedName string
		ExpectedSize int
	}{
		{
			String:       "name value",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\tvalue",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name  value",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\t\tvalue",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name:value",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name: value",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name :value",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name : value",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name=value",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name= value",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name= value",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name = value",
			ExpectedName: "name",
			ExpectedSize: 4,
		},



		{
			String:       "name&END",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name &END",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name  &END",
			ExpectedName: "name",
			ExpectedSize: 4,
		},



		{
			String:       "name",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\n",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\n\r",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\r",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\r\n",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\u0085",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\u2028",
			ExpectedName: "name",
			ExpectedSize: 4,
		},



		{
			String:       "name\nline2=wow",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\n\rline2=wow",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\rline2=wow",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\r\nline2=wow",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\u0085line2=wow",
			ExpectedName: "name",
			ExpectedSize: 4,
		},
		{
			String:       "name\u2028line2=wow",
			ExpectedName: "name",
			ExpectedSize: 4,
		},



		{
			String:       "name\u06F5 value",
			ExpectedName: "name\u06F5",
			ExpectedSize: 6,
		},
		{
			String:       "name\u06F5\tvalue",
			ExpectedName: "name\u06F5",
			ExpectedSize: 6,
		},
		{
			String:       "name\u06F5  value",
			ExpectedName: "name\u06F5",
			ExpectedSize: 6,
		},
		{
			String:       "name\u06F5\t\tvalue",
			ExpectedName: "name\u06F5",
			ExpectedSize: 6,
		},
		{
			String:       "name\u06F5:value",
			ExpectedName: "name\u06F5",
			ExpectedSize: 6,
		},
		{
			String:       "name\u06F5: value",
			ExpectedName: "name\u06F5",
			ExpectedSize: 6,
		},
		{
			String:       "name\u06F5 : value",
			ExpectedName: "name\u06F5",
			ExpectedSize: 6,
		},
		{
			String:       "name\u06F5=value",
			ExpectedName: "name\u06F5",
			ExpectedSize: 6,
		},
		{
			String:       "name\u06F5 = value",
			ExpectedName: "name\u06F5",
			ExpectedSize: 6,
		},
	}


	for testNumber, test := range tests {

		var bytes []byte = []byte(test.String)

		actualName, actualSize, err := ininame.ParseBytes(bytes)
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
