package iniconv_test

import (
	"testing"

	"reflect"

	"github.com/reiver/go-ini/conv"
)

func TestParseStrings(t *testing.T) {

	tests := []struct{
		Text string
		Separator string
		Expected []string
	}{
		{
			Expected: []string{},
		},



		{
			Text:      "",
			Separator: "",
			Expected: []string{},
		},
		{
			Text:      "",
			Separator: " ",
			Expected: []string{},
		},
		{
			Text:      "",
			Separator: "\t",
			Expected: []string{},
		},
		{
			Text:      "",
			Separator: ",",
			Expected: []string{},
		},
		{
			Text:      "",
			Separator: ";",
			Expected: []string{},
		},
		{
			Text:      "",
			Separator: "/",
			Expected: []string{},
		},
		{
			Text:      "",
			Separator: "😈😈",
			Expected: []string{},
		},



		{
			Text:      "apple Banana CHERRY 🙂🙂",
			Separator: "",
			Expected: []string{"a","p","p","l","e","","B","a","n","a","n","a","","C","H","E","R","R","Y","","🙂","🙂"},
		},
		{
			Text:      "apple Banana CHERRY 🙂🙂",
			Separator: " ",
			Expected: []string{"apple", "Banana", "CHERRY", "🙂🙂"},
		},
		{
			Text:      "apple Banana CHERRY 🙂🙂",
			Separator: "\t",
			Expected: []string{"apple Banana CHERRY 🙂🙂"},
		},
		{
			Text:      "apple Banana CHERRY 🙂🙂",
			Separator: ",",
			Expected: []string{"apple Banana CHERRY 🙂🙂"},
		},
		{
			Text:      "apple Banana CHERRY 🙂🙂",
			Separator: ";",
			Expected: []string{"apple Banana CHERRY 🙂🙂"},
		},
		{
			Text:      "apple Banana CHERRY 🙂🙂",
			Separator: "/",
			Expected: []string{"apple Banana CHERRY 🙂🙂"},
		},
		{
			Text:      "apple Banana CHERRY 🙂🙂",
			Separator: "😈😈",
			Expected: []string{"apple Banana CHERRY 🙂🙂"},
		},
	}

	for testNumber, test := range tests {

		actual := iniconv.ParseStrings(test.Text, test.Separator)

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual parsed-strings it not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("TEXT: %q", test.Text)
			t.Logf("SEPARATOR: %q", test.Separator)
			continue
		}
	}
}
