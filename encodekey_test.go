package ini

import (
	"testing"
)

func TestEncodeKey(t *testing.T) {

	tests := []struct{
		Key string
		Expected string
	}{
		{
			Key:      "",
			Expected: "",
		},


		{
			Key:      "apple",
			Expected: "apple",
		},
		{
			Key:      "Banana",
			Expected: "Banana",
		},
		{
			Key:      "CHERRY",
			Expected: "CHERRY",
		},



		{
			Key:      "Hello world!",
			Expected: "Hello world!",
		},



		{
			Key:      `a"b""c`,
			Expected: `a\"b\"\"c`,
		},
		{
			Key:      `a#b##c`,
			Expected: `a\#b\#\#c`,
		},
		{
			Key:      `a'b''c`,
			Expected: `a\'b\'\'c`,
		},
		{
			Key:      `a:b::c`,
			Expected: `a\:b\:\:c`,
		},
		{
			Key:      `a=b==c`,
			Expected: `a\=b\=\=c`,
		},
		{
			Key:      `a;b;;c`,
			Expected: `a\;b\;\;c`,
		},
		{
			Key:      `a\b\\c`,
			Expected: `a\\b\\\\c`,
		},



		{
			Key:      `apple = one ; Banana : two ; CHERRY "three" ; ===`,
			Expected: `apple \= one \; Banana \: two \; CHERRY \"three\" \; \=\=\=`,
		},
	}

	for testNumber, test := range tests {

		actual := encodeKey(test.Key)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual encoded-key is not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("KEY:      %q", test.Key)
			continue
		}
	}
}
