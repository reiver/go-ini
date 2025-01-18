package inisection

import (
	"testing"

	"reflect"
)

func TestSplitAndLower(t *testing.T) {

	tests := []struct{
		Value string
		Expected []string
	}{
		{
			Value: "",
			Expected: []string{},
		},


		{
			Value: "A",
			Expected: []string{"a"},
		},
		{
			Value: "Z",
			Expected: []string{"z"},
		},
		{
			Value: "a",
			Expected: []string{"a"},
		},
		{
			Value: "z",
			Expected: []string{"z"},
		},
		{
			Value: "۵",
			Expected: []string{"۵"},
		},



		{
			Value: "۰.۱.۲.۳.۴.۵.۶.۷.۸.۹",
			Expected: []string{"۰", "۱", "۲", "۳", "۴", "۵", "۶", "۷", "۸", "۹"},
		},
		{
			Value: "۰0.۱1.۲2.۳3.۴4.۵5.۶6.۷7.۸8.۹9",
			Expected: []string{"۰0", "۱1", "۲2", "۳3", "۴4", "۵5", "۶6", "۷7", "۸8", "۹9"},
		},




		{
			Value: "apple",
			Expected: []string{"apple"},
		},
		{
			Value: "apple.Banana",
			Expected: []string{"apple", "banana"},
		},
		{
			Value: "apple.banana.CHERRY",
			Expected: []string{"apple", "banana", "cherry"},
		},



		{
			Value: `aBcd\.eFg.HI.j`,
			Expected: []string{"abcd.efg", "hi", "j"},
		},



		{
			Value: `apple\`,
			Expected: []string{"apple"},
		},
		{
			Value: `apple.Banana\`,
			Expected: []string{"apple", "banana"},
		},
		{
			Value: `apple.Banana.CHERRY\`,
			Expected: []string{"apple", "banana", "cherry"},
		},
	}

	for testNumber, test := range tests {

		actual := splitAndLower(test.Value)

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual (split-and-lowered) section-name is not what was expected.", testNumber)
			t.Log("EXPECTED")
			for index, part := range expected {
				t.Logf("\t- [%d] %q", index, part)
			}
			t.Log("ACTUAL:")
			for index, part := range actual {
				t.Logf("\t- [%d] %q", index, part)
			}
			t.Logf("VALUE: %q", test.Value)
			continue
		}
	}
}
