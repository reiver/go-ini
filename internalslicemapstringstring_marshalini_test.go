package ini

import (
	"testing"

	"reflect"
)

func TestInternalSliceMapStringString_MarshalINI(t *testing.T) {

	tests := []struct{
		Value internalSliceMapStringString
		Nesting []string
		Expected []byte
	}{
		{

			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{nil},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{[]map[string]string(nil)},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{[]map[string]string{}},
			Expected: []byte{},
		},



		{

			Nesting: []string{"parent"},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{},
			Nesting: []string{"parent"},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{nil},
			Nesting: []string{"parent"},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{[]map[string]string(nil)},
			Nesting: []string{"parent"},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{[]map[string]string{}},
			Nesting: []string{"parent"},
			Expected: []byte{},
		},



		{

			Nesting: []string{"yek","do","se"},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{},
			Nesting: []string{"yek","do","se"},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{nil},
			Nesting: []string{"yek","do","se"},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{[]map[string]string(nil)},
			Nesting: []string{"yek","do","se"},
			Expected: []byte{},
		},
		{
			Value: internalSliceMapStringString{[]map[string]string{}},
			Nesting: []string{"yek","do","se"},
			Expected: []byte{},
		},



		{
			Nesting: []string{"parent"},
			Value: internalSliceMapStringString{[]map[string]string{
				map[string]string{
					"apple":"ONCE",
					"Banana":"TWICE",
					"CHERRY":"THRICE",
					"dAtE":"FOURCE",
				},
				map[string]string{
					"apple":"1",
					"Banana":"2",
					"CHERRY":"3",
					"dAtE":"4",
				},
				map[string]string{
					"apple":"one",
					"Banana":"two",
					"CHERRY":"three",
					"dAtE":"four",
				},
			}},
			Expected: []byte(
				`[[parent]]`      +"\n"+
				`apple = ONCE`    +"\n"+
				`Banana = TWICE`  +"\n"+
				`CHERRY = THRICE` +"\n"+
				`dAtE = FOURCE`   +"\n"+
				`[[parent]]`      +"\n"+
				`apple = 1`       +"\n"+
				`Banana = 2`      +"\n"+
				`CHERRY = 3`      +"\n"+
				`dAtE = 4`        +"\n"+
				`[[parent]]`      +"\n"+
				`apple = one`     +"\n"+
				`Banana = two`    +"\n"+
				`CHERRY = three`  +"\n"+
				`dAtE = four`     +"\n"+
				"",
			),
		},
		{
			Nesting: []string{"first", "second"},
			Value: internalSliceMapStringString{[]map[string]string{
				map[string]string{
					"apple":"ONCE",
					"Banana":"TWICE",
					"CHERRY":"THRICE",
					"dAtE":"FOURCE",
				},
				map[string]string{
					"apple":"1",
					"Banana":"2",
					"CHERRY":"3",
					"dAtE":"4",
				},
				map[string]string{
					"apple":"one",
					"Banana":"two",
					"CHERRY":"three",
					"dAtE":"four",
				},
			}},
			Expected: []byte(
				`[[first.second]]` +"\n"+
				`apple = ONCE`     +"\n"+
				`Banana = TWICE`   +"\n"+
				`CHERRY = THRICE`  +"\n"+
				`dAtE = FOURCE`    +"\n"+
				`[[first.second]]` +"\n"+
				`apple = 1`        +"\n"+
				`Banana = 2`       +"\n"+
				`CHERRY = 3`       +"\n"+
				`dAtE = 4`         +"\n"+
				`[[first.second]]` +"\n"+
				`apple = one`      +"\n"+
				`Banana = two`     +"\n"+
				`CHERRY = three`   +"\n"+
				`dAtE = four`      +"\n"+
				"",
			),
		},
		{
			Nesting: []string{"yek","do","se"},
			Value: internalSliceMapStringString{[]map[string]string{
				map[string]string{
					"apple":"ONCE",
					"Banana":"TWICE",
					"CHERRY":"THRICE",
					"dAtE":"FOURCE",
				},
				map[string]string{
					"apple":"1",
					"Banana":"2",
					"CHERRY":"3",
					"dAtE":"4",
				},
				map[string]string{
					"apple":"one",
					"Banana":"two",
					"CHERRY":"three",
					"dAtE":"four",
				},
			}},
			Expected: []byte(
				`[[yek.do.se]]`    +"\n"+
				`apple = ONCE`     +"\n"+
				`Banana = TWICE`   +"\n"+
				`CHERRY = THRICE`  +"\n"+
				`dAtE = FOURCE`    +"\n"+
				`[[yek.do.se]]`    +"\n"+
				`apple = 1`        +"\n"+
				`Banana = 2`       +"\n"+
				`CHERRY = 3`       +"\n"+
				`dAtE = 4`         +"\n"+
				`[[yek.do.se]]`    +"\n"+
				`apple = one`      +"\n"+
				`Banana = two`     +"\n"+
				`CHERRY = three`   +"\n"+
				`dAtE = four`      +"\n"+
				"",
			),
		},
	}

	for testNumber, test := range tests {

		var buffer [256]byte
		var actual []byte = buffer[0:0]
		var err error

		actual, err = test.Value.MarshalINI(actual, test.Nesting...)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("NESTING: %#v", test.Nesting)
			t.Logf("VALUE: %#v", test.Value)
			continue
		}

		{
			expected := test.Expected

			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d, the actual ini-marshaled result is not what was expected.", testNumber)
				t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
				t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
				t.Logf("NESTING: %#v", test.Nesting)
				t.Logf("VALUE: %#v", test.Value)
				continue
			}
		}
	}
}
