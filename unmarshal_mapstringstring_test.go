package ini_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-humane"
	"github.com/reiver/go-ini"
)

func TestUnmarshal_mapStringString(t *testing.T) {

	mapKeys := func(m map[string]string) (keys []string) {
		for key, _ := range m {
			keys = append(keys, key)
		}

		humane.SortStrings(keys)

		return keys
	}

	tests := []struct{
		INI string
		Expected map[string]string
	}{
		{
			Expected: map[string]string{},
		},



		{
			INI: "[abcdefg.hijk.lmnop]",
			Expected: map[string]string{},
		},



		{
			INI: "abc = 123",
			Expected: map[string]string{
				"abc":"123",
			},
		},



		{
			INI: `
a   = 1
ab  = 12
abc = 123
`,
			Expected: map[string]string{
				"a":"1",
				"ab":"12",
				"abc":"123",
			},
		},
		{
			INI: `
[once]

a   = 1
ab  = 12
abc = 123
`,
			Expected: map[string]string{
				"once.a":"1",
				"once.ab":"12",
				"once.abc":"123",
			},
		},
		{
			INI: `
[once.twice]

a   = 1
ab  = 12
abc = 123
`,
			Expected: map[string]string{
				"once.twice.a":"1",
				"once.twice.ab":"12",
				"once.twice.abc":"123",
			},
		},
		{
			INI: `
[once.twice.thrice]

a   = 1
ab  = 12
abc = 123
`,
			Expected: map[string]string{
				"once.twice.thrice.a":"1",
				"once.twice.thrice.ab":"12",
				"once.twice.thrice.abc":"123",
			},
		},
		{
			INI: `
[once.twice.thrice.fource]

a   = 1
ab  = 12
abc = 123
`,
			Expected: map[string]string{
				"once.twice.thrice.fource.a":"1",
				"once.twice.thrice.fource.ab":"12",
				"once.twice.thrice.fource.abc":"123",
			},
		},
	}

	for testNumber, test := range tests {

		var actual map[string]string

		err := ini.Unmarshal([]byte(test.INI), &actual)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("INI:\n%s", test.INI)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual unmarshaled-INI (into a map[string]string) is not what was expected.", testNumber)
			t.Log("EXPECTED:")
			for index, key := range mapKeys(expected) {
				t.Logf("\t- [%d] %q -> %q", index, key, expected[key])
			}
			t.Log("ACTUAL:")
			for index, key := range mapKeys(actual) {
				t.Logf("\t- [%d] %q -> %q", index, key, actual[key])
			}
			t.Logf("INI:\n%s", test.INI)
			continue
		}
	}
}

func TestUnmarshal_mapStringString_error(t *testing.T) {

	tests := []struct{
		INI string
		ExpectedError string
	}{
		{
			INI: "[[abcdefg.hijk.lmnop]]",
			ExpectedError: "ini: problem publishing INI sequence: ini: setter cannot handle sequence-headers — there was attempt to set sequence header [[abcdefg.hijk.lmnop]]",
		},


		{
			INI: `
[[fruit]]

id = 1
color = green

[[fruit]]

id = 2
color = yellow

[[fruit]]

id = 3
color = red
`,
			ExpectedError: "ini: problem publishing INI sequence: ini: setter cannot handle sequence-headers — there was attempt to set sequence header [[fruit]]",
		},
	}

	for testNumber, test := range tests {

		var dst map[string]string

		err := ini.Unmarshal([]byte(test.INI), &dst)
		if nil == err {
			t.Errorf("For test #%d, expected an error but actually get one.", testNumber)
			t.Logf("INI:\n%s", test.INI)
			continue
		}

		actual := err.Error()

		expected := test.ExpectedError

		if expected != actual {
			t.Errorf("For test #%d, the actual error-message is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("INI:\n%s", test.INI)
			continue
		}
	}


}
