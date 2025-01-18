package ini_test

import (
	"testing"

	"fmt"
	"reflect"

	"github.com/reiver/go-ini"
	"github.com/reiver/go-ini/internal/test"
)

func TestUnmarshal_publisher(t *testing.T) {

	tests := []struct{
		INI string
		Expected []string
	}{
		{

		},



		{
			INI: "# a comment",
			Expected: []string{
				`INI-COMMENT: " a comment"`,
			},
		},
		{
			INI: "; a comment",
			Expected: []string{
				`INI-COMMENT: " a comment"`,
			},
		},



		{
			INI: "the-key=the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: " the-key=the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "  the-key=the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "   the-key=the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key =the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key  =the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key   =the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key= the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key=  the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key=   the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "\tthe-key=the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "\t\tthe-key=the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "\t\t\tthe-key=the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key\t=the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key\t\t=the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key\t\t\t=the-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key=\tthe-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key=\t\tthe-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},
		{
			INI: "the-key=\t\t\tthe-value",
			Expected: []string{
				`INI-KEY-VALUE: key="the-key" value="the-value"`,
			},
		},



		{
			INI: "[this.is.the.name]",
			Expected: []string{
				`INI-SECTION-HEADER: name="this"."is"."the"."name"`,
			},
		},
	}

	for testNumber, test := range tests {

		var actual []string

		var publisher ini.Publisher = initest.Publisher{
			PublishINICommentFunc: func(comment string) error {
				var str string = fmt.Sprintf("INI-COMMENT: %q", comment)
				actual = append(actual, str)
				return nil
			},
			PublishINIKeyValueFunc: func(key string, value string) error {
				var str string = fmt.Sprintf("INI-KEY-VALUE: key=%q value=%q", key, value)
				actual = append(actual, str)
				return nil
			},
			PublishINISectionHeaderFunc: func(name ...string) error {
				var str string = "INI-SECTION-HEADER: name="
				for index, part := range name {
					if 0 < index {
						str += "."
					}
					str += fmt.Sprintf("%q", part)
				}
				actual = append(actual, str)
				return nil
			},
		}

		err := ini.Unmarshal([]byte(test.INI), publisher)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("INI:\n%s", test.INI)
			continue
		}

		expected := test.Expected

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
			t.Log("EXPECTED:")
			for index, str := range expected {
				t.Logf("\t- [%d] %s", index, str)
			}
			t.Log("ACTUAL:")
			for index, str := range actual {
				t.Logf("\t- [%d] %s", index, str)
			}
			t.Logf("INI:\n%s", test.INI)
			continue
		}
	}
}
