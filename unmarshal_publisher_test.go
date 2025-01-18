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
			INI: `[this.is.the.name]`,
			Expected: []string{
				`INI-SECTION-HEADER: name="this"."is"."the"."name"`,
			},
		},
		{
			INI: `[this.is\.the\.name]`,
			Expected: []string{
				`INI-SECTION-HEADER: name="this"."is.the.name"`,
			},
		},
		{
			INI: `[this\.is.the\.name]`,
			Expected: []string{
				`INI-SECTION-HEADER: name="this.is"."the.name"`,
			},
		},
		{
			INI: `[this\.is\.the.name]`,
			Expected: []string{
				`INI-SECTION-HEADER: name="this.is.the"."name"`,
			},
		},
		{
			INI: `[this\.is\.the\.name]`,
			Expected: []string{
				`INI-SECTION-HEADER: name="this.is.the.name"`,
			},
		},


		{
			INI: `[this\.is\.the\.name]`,
			Expected: []string{
				`INI-SECTION-HEADER: name="this.is.the.name"`,
			},
		},



		{
			INI: `[[this.is.the.name]]`,
			Expected: []string{
				`INI-SEQUENCE-HEADER: name="this"."is"."the"."name"`,
			},
		},









		{
			INI:
`apple = ONE
Banana = TWO
Cherry = THREE

[A]

once
twice
thrice
fource

[B.C]

ONCE   = 1234
TWICE  = 2345
THRICE = 3456
FOURCE = 4567

[[C]]

id = 1
name = Joe Blow

[[C]]

id = 2
name = Jane Doe

[[C]]

id = 3
name = John Doe

[[D.EF.GHI]]

toy = GraySkull
`,
			Expected: []string{
				`INI-KEY-VALUE: key="apple" value="ONE"`,
				`INI-KEY-VALUE: key="banana" value="TWO"`,
				`INI-KEY-VALUE: key="cherry" value="THREE"`,
				`INI-SECTION-HEADER: name="a"`,
				`INI-KEY-VALUE: key="once" value=""`,
				`INI-KEY-VALUE: key="twice" value=""`,
				`INI-KEY-VALUE: key="thrice" value=""`,
				`INI-KEY-VALUE: key="fource" value=""`,
				`INI-SECTION-HEADER: name="b"."c"`,
				`INI-KEY-VALUE: key="once" value="1234"`,
				`INI-KEY-VALUE: key="twice" value="2345"`,
				`INI-KEY-VALUE: key="thrice" value="3456"`,
				`INI-KEY-VALUE: key="fource" value="4567"`,
				`INI-SEQUENCE-HEADER: name="c"`,
				`INI-KEY-VALUE: key="id" value="1"`,
				`INI-KEY-VALUE: key="name" value="Joe Blow"`,
				`INI-SEQUENCE-HEADER: name="c"`,
				`INI-KEY-VALUE: key="id" value="2"`,
				`INI-KEY-VALUE: key="name" value="Jane Doe"`,
				`INI-SEQUENCE-HEADER: name="c"`,
				`INI-KEY-VALUE: key="id" value="3"`,
				`INI-KEY-VALUE: key="name" value="John Doe"`,
				`INI-SEQUENCE-HEADER: name="d"."ef"."ghi"`,
				`INI-KEY-VALUE: key="toy" value="GraySkull"`,
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
			PublishINISequenceHeaderFunc: func(name ...string) error {
				var str string = "INI-SEQUENCE-HEADER: name="
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
