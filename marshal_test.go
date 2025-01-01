package ini_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-ini"
)

func TestMarshal(t *testing.T) {

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: map[string]string{
				"ABC":"XYZ",
				"ABc":"XYz",
				"AbC":"XyZ",
				"Abc":"Xyz",
				"aBC":"xYZ",
				"aBc":"xYz",
				"abC":"xyZ",
				"abc":"xyz",
			},
			Expected: []byte(
				`ABC = XYZ` +"\n"+
				`ABc = XYz` +"\n"+
				`AbC = XyZ` +"\n"+
				`Abc = Xyz` +"\n"+
				`aBC = xYZ` +"\n"+
				`aBc = xYz` +"\n"+
				`abC = xyZ` +"\n"+
				`abc = xyz` +"\n"+
				"",
			),
		},
	}

	for testNumber, test := range tests {

		actual, err := ini.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		var expected []byte = test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual ini-marshaled value is not what was expected." , testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
