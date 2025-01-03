package ini

import (
	"testing"
)

func TestEncodeValue(t *testing.T) {

	tests := []struct{
		Value string
		Expected string
	}{
		{
			Value:    "",
			Expected: "",
		},


		{
			Value:    "apple",
			Expected: "apple",
		},
		{
			Value:    "Banana",
			Expected: "Banana",
		},
		{
			Value:    "CHERRY",
			Expected: "CHERRY",
		},



		{
			Value:    "Hello world!",
			Expected: "Hello world!",
		},



		{
			Value:    `a"b""c`,
			Expected: `a\"b\"\"c`,
		},
		{
			Value:    `a#b##c`,
			Expected: `a\#b\#\#c`,
		},
		{
			Value:    `a'b''c`,
			Expected: `a\'b\'\'c`,
		},
		{
			Value:    `a:b::c`,
			Expected: `a\:b\:\:c`,
		},
		{
			Value:    `a=b==c`,
			Expected: `a\=b\=\=c`,
		},
		{
			Value:    `a;b;;c`,
			Expected: `a\;b\;\;c`,
		},
		{
			Value:    `a\b\\c`,
			Expected: `a\\b\\\\c`,
		},



		{
			Value:    `apple = one ; Banana : two ; CHERRY "three" ; ===`,
			Expected: `apple \= one \; Banana \: two \; CHERRY \"three\" \; \=\=\=`,
		},



		{
			Value:    `line 1`    +"\n"+`line 2`    +"\n"+`line 3`,
			Expected: `line 1`+`\`+"\n"+`line 2`+`\`+"\n"+`line 3`,
		},
	}

	for testNumber, test := range tests {

		actual := encodeValue(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual encoded-key is not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("KEY:      %q", test.Value)
			continue
		}
	}
}
