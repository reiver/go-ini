package ini

import (
	"testing"

	"reflect"
)

func TestInternalMapStringString_INIContent(t *testing.T) {

	tests := []struct{
		Value internalMapStringString
		Nesting []string
		Expected []byte
	}{
		{
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringString{},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringString{nil},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringString{map[string]string{}},
			Expected: []byte(nil),
		},



		{
			Expected: []byte(nil),
			Nesting: []string{"parent"},
		},
		{
			Value: internalMapStringString{},
			Nesting: []string{"parent"},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringString{nil},
			Nesting: []string{"parent"},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringString{map[string]string{}},
			Nesting: []string{"parent"},
			Expected: []byte(nil),
		},



		{
			Expected: []byte(nil),
			Nesting: []string{"p-1","p-2","p-3"},
		},
		{
			Value: internalMapStringString{},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringString{nil},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringString{map[string]string{}},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(nil),
		},



		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE"}},
			Expected: []byte(
				`apple = ONCE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE"}},
			Expected: []byte(
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Expected: []byte(
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				`CHERRY = THRICE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
			Expected: []byte(
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				`CHERRY = THRICE` +"\n"+
				`dAtE = FOURCE` +"\n"+
				"",
			),
		},



		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE"}},
			Nesting: []string{"parent"},
			Expected: []byte(
				`[parent]`+"\n"+
				`apple = ONCE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE"}},
			Nesting: []string{"parent"},
			Expected: []byte(
				`[parent]`+"\n"+
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Nesting: []string{"parent"},
			Expected: []byte(
				`[parent]`+"\n"+
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				`CHERRY = THRICE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
			Nesting: []string{"parent"},
			Expected: []byte(
				`[parent]`+"\n"+
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				`CHERRY = THRICE` +"\n"+
				`dAtE = FOURCE` +"\n"+
				"",
			),
		},



		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE"}},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(
				`[p-1.p-2.p-3]`+"\n"+
				`apple = ONCE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE"}},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(
				`[p-1.p-2.p-3]`+"\n"+
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(
				`[p-1.p-2.p-3]`+"\n"+
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				`CHERRY = THRICE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringString{map[string]string{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(
				`[p-1.p-2.p-3]`+"\n"+
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				`CHERRY = THRICE` +"\n"+
				`dAtE = FOURCE` +"\n"+
				"",
			),
		},
	}

	for testNumber, test := range tests {

		actual, err := test.Value.INIContent(test.Nesting...)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual keys is not what was expected." , testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}
	}
}
