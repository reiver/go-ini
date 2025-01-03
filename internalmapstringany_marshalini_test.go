package ini

import (
	"testing"

	"reflect"
)

func TestInternalMapStringAny_MarshalINI(t *testing.T) {

	tests := []struct{
		Value internalMapStringAny
		Nesting []string
		Expected []byte
	}{
		{
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringAny{},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringAny{nil},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringAny{map[string]any{}},
			Expected: []byte(nil),
		},



		{
			Expected: []byte(nil),
			Nesting: []string{"parent"},
		},
		{
			Value: internalMapStringAny{},
			Nesting: []string{"parent"},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringAny{nil},
			Nesting: []string{"parent"},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringAny{map[string]any{}},
			Nesting: []string{"parent"},
			Expected: []byte(nil),
		},



		{
			Expected: []byte(nil),
			Nesting: []string{"p-1","p-2","p-3"},
		},
		{
			Value: internalMapStringAny{},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringAny{nil},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(nil),
		},
		{
			Value: internalMapStringAny{map[string]any{}},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(nil),
		},



		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE"}},
			Expected: []byte(
				`apple = ONCE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE"}},
			Expected: []byte(
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
			Expected: []byte(
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				`CHERRY = THRICE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
			Expected: []byte(
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				`CHERRY = THRICE` +"\n"+
				`dAtE = FOURCE` +"\n"+
				"",
			),
		},



		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE"}},
			Nesting: []string{"parent"},
			Expected: []byte(
				`[parent]`+"\n"+
				`apple = ONCE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE"}},
			Nesting: []string{"parent"},
			Expected: []byte(
				`[parent]`+"\n"+
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
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
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
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
			Value: internalMapStringAny{map[string]any{"apple":"ONCE"}},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(
				`[p-1.p-2.p-3]`+"\n"+
				`apple = ONCE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE"}},
			Nesting: []string{"p-1","p-2","p-3"},
			Expected: []byte(
				`[p-1.p-2.p-3]`+"\n"+
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				"",
			),
		},
		{
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE"}},
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
			Value: internalMapStringAny{map[string]any{"apple":"ONCE","Banana":"TWICE","CHERRY":"THRICE","dAtE":"FOURCE"}},
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



		{
			Value: internalMapStringAny{map[string]any{
				"apple":"ONCE",
				"Banana":"TWICE",
				"CHERRY":"THRICE",
				"dAtE":"FOURCE",
				"something":map[string]string{
					"a":"1",
					"b":"2",
					"c":"3",
				},
				"a":map[string]any{
					"p":map[string]any{
						"t":map[string]any{
							"face":"kissy",
						},
					},
				},
			}},
			Expected: []byte(
				`apple = ONCE` +"\n"+
				`Banana = TWICE` +"\n"+
				`CHERRY = THRICE` +"\n"+
				`dAtE = FOURCE` +"\n"+
				`[a]`+"\n"+
				`[a.p]`+"\n"+
				`[a.p.t]`+"\n"+
				`face = kissy`+"\n"+
				`[something]`+"\n"+
				`a = 1`+"\n"+
				`b = 2`+"\n"+
				`c = 3`+"\n"+
				"",
			),
		},
	}

	for testNumber, test := range tests {

		var actual []byte
		var err error

		actual, err = test.Value.MarshalINI(actual, test.Nesting...)
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
