package inisection

import (
	"testing"
)

func TestParseString(t *testing.T) {

	tests := []struct{
		Value string
		Expected string
	}{
		{
			Value:   "[",
			Expected: "",
		},
		{
			Value:   "[\n",
			Expected: "",
		},
		{
			Value:   "[\n\r",
			Expected: "",
		},
		{
			Value:   "[\r",
			Expected: "",
		},
		{
			Value:   "[\r\n",
			Expected: "",
		},
		{
			Value:   "[\u0085",
			Expected: "",
		},
		{
			Value:   "[\u2028",
			Expected: "",
		},



		{
			Value:   "[hello world! 🙂",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[[hello world! 🙂\n",
			Expected: "[hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂\n\r",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂\r",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂\r\n",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂\u0085",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂\u2028",
			Expected: "hello world! 🙂",
		},



		{
			Value:   "[hello world! 🙂]",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂]\n",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂]\n\r",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂]\r",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂]\r\n",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂]\u0085",
			Expected: "hello world! 🙂",
		},
		{
			Value:   "[hello world! 🙂]\u2028",
			Expected: "hello world! 🙂",
		},
	}

	for testNumber, test := range tests {

		actual, err := ParseString(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: %q", test.Value)
			continue
		}

		{
			expected := test.Expected

			if expected != actual {
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
