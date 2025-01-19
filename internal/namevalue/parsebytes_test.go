package ininamevalue_test

import (
	"testing"

	"github.com/reiver/go-ini/internal/namevalue"
)

func TestParseBytes(t *testing.T) {

	tests := []struct{
		String        string
		ExpectedName  string
		ExpectedValue string
		ExpectedSize  int
		ExpectedValueSize int
	}{
		{
			String:           "name value",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value"),
		},
		{
			String:           "name value\n",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\n"),
		},
		{
			String:           "name value\n\r",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\n\r"),
		},
		{
			String:           "name value\r",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\r"),
		},
		{
			String:           "name value\r\n",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\r\n"),
		},
		{
			String:           "name value\u0085",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\u0085"),
		},
		{
			String:           "name value\u2028",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\u2028"),
		},



		{
			String:           "name     value",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value"),
		},
		{
			String:           "name     value\n",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\n"),
		},
		{
			String:           "name     value\n\r",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\n\r"),
		},
		{
			String:           "name     value\r",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\r"),
		},
		{
			String:           "name     value\r\n",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\r\n"),
		},
		{
			String:           "name     value\u0085",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\u0085"),
		},
		{
			String:           "name     value\u2028",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\u2028"),
		},



		{
			String:           "name\tvalue",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue"),
		},
		{
			String:           "name\tvalue\n",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\n"),
		},
		{
			String:           "name\tvalue\n\r",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\n\r"),
		},
		{
			String:           "name\tvalue\r",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\r"),
		},
		{
			String:           "name\tvalue\r\n",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\r\n"),
		},
		{
			String:           "name\tvalue\u0085",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\u0085"),
		},
		{
			String:           "name\tvalue\u2028",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\u2028"),
		},



		{
			String:           "name\t\t\tvalue",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue"),
		},
		{
			String:           "name\t\t\tvalue\n",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\n"),
		},
		{
			String:           "name\t\t\tvalue\n\r",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\n\r"),
		},
		{
			String:           "name\t\t\tvalue\r",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\r"),
		},
		{
			String:           "name\t\t\tvalue\r\n",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\r\n"),
		},
		{
			String:           "name\t\t\tvalue\u0085",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\u0085"),
		},
		{
			String:           "name\t\t\tvalue\u2028",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\u2028"),
		},



		{
			String:           "name: value",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value"),
		},
		{
			String:           "name: value\n",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\n"),
		},
		{
			String:           "name: value\n\r",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\n\r"),
		},
		{
			String:           "name: value\r",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\r"),
		},
		{
			String:           "name: value\r\n",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\r\n"),
		},
		{
			String:           "name: value\u0085",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\u0085"),
		},
		{
			String:           "name: value\u2028",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\u2028"),
		},



		{
			String:           "name : value",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value"),
		},
		{
			String:           "name : value\n",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\n"),
		},
		{
			String:           "name : value\n\r",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\n\r"),
		},
		{
			String:           "name : value\r",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\r"),
		},
		{
			String:           "name : value\r\n",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\r\n"),
		},
		{
			String:           "name : value\u0085",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\u0085"),
		},
		{
			String:           "name : value\u2028",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\u2028"),
		},



		{
			String:           "name=value",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value"),
		},
		{
			String:           "name=value\n",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\n"),
		},
		{
			String:           "name=value\n\r",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\n\r"),
		},
		{
			String:           "name=value\r",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\r"),
		},
		{
			String:           "name=value\r\n",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\r\n"),
		},
		{
			String:           "name=value\u0085",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\u0085"),
		},
		{
			String:           "name=value\u2028",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\u2028"),
		},



		{
			String:           "name = value",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value"),
		},
		{
			String:           "name = value\n",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\n"),
		},
		{
			String:           "name = value\n\r",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\n\r"),
		},
		{
			String:           "name = value\r",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\r"),
		},
		{
			String:           "name = value\r\n",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\r\n"),
		},
		{
			String:           "name = value\u0085",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\u0085"),
		},
		{
			String:           "name = value\u2028",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\u2028"),
		},



		{
			String:           "name&END\nvalue",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\nvalue"),
		},
		{
			String:           "name&END\nvalueEND",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\nvalueEND"),
		},
		{
			String:           "name&END\nvalue\n",
			ExpectedName:     "name",
			ExpectedValue:              "value\n",
			ExpectedSize: len("name&END\nvalue\n"),
		},
		{
			String:           "name&END\nvalue\nEND",
			ExpectedName:     "name",
			ExpectedValue:              "value\n",
			ExpectedSize: len("name&END\nvalue\nEND"),
		},
		{
			String:           "name&END\nvalue\nEND\n",
			ExpectedName:     "name",
			ExpectedValue:              "value\n",
			ExpectedSize: len("name&END\nvalue\nEND\n"),
		},
		{
			String:           "name&END\nvalue\nEND\nwow",
			ExpectedName:     "name",
			ExpectedValue:              "value\n",
			ExpectedSize: len("name&END\nvalue\nEND\n"),
		},


		{
			String:           "name&END\n\rvalue",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\n\rvalue"),
		},
		{
			String:           "name&END\n\rvalueEND",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\n\rvalueEND"),
		},
		{
			String:           "name&END\n\rvalue\n\r",
			ExpectedName:     "name",
			ExpectedValue:              "value\n\r",
			ExpectedSize: len("name&END\n\rvalue\n\r"),
		},
		{
			String:           "name&END\n\rvalue\n\rEND",
			ExpectedName:     "name",
			ExpectedValue:              "value\n\r",
			ExpectedSize: len("name&END\n\rvalue\n\rEND"),
		},
		{
			String:           "name&END\n\rvalue\n\rEND\n\r",
			ExpectedName:     "name",
			ExpectedValue:              "value\n\r",
			ExpectedSize: len("name&END\n\rvalue\n\rEND\n\r"),
		},
		{
			String:           "name&END\n\rvalue\n\rEND\n\rwow",
			ExpectedName:     "name",
			ExpectedValue:              "value\n\r",
			ExpectedSize: len("name&END\n\rvalue\n\rEND\n\r"),
		},


		{
			String:           "name&END\rvalue",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\rvalue"),
		},
		{
			String:           "name&END\rvalueEND",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\rvalueEND"),
		},
		{
			String:           "name&END\rvalue\r",
			ExpectedName:     "name",
			ExpectedValue:              "value\r",
			ExpectedSize: len("name&END\rvalue\r"),
		},
		{
			String:           "name&END\rvalue\rEND",
			ExpectedName:     "name",
			ExpectedValue:              "value\r",
			ExpectedSize: len("name&END\rvalue\rEND"),
		},
		{
			String:           "name&END\rvalue\rEND\r",
			ExpectedName:     "name",
			ExpectedValue:              "value\r",
			ExpectedSize: len("name&END\rvalue\rEND\r"),
		},
		{
			String:           "name&END\rvalue\rEND\rwow",
			ExpectedName:     "name",
			ExpectedValue:              "value\r",
			ExpectedSize: len("name&END\rvalue\rEND\r"),
		},


		{
			String:           "name&END\r\nvalue",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\r\nvalue"),
		},
		{
			String:           "name&END\r\nvalueEND",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\r\nvalueEND"),
		},
		{
			String:           "name&END\r\nvalue\r\n",
			ExpectedName:     "name",
			ExpectedValue:              "value\r\n",
			ExpectedSize: len("name&END\r\nvalue\r\n"),
		},
		{
			String:           "name&END\r\nvalue\r\nEND",
			ExpectedName:     "name",
			ExpectedValue:              "value\r\n",
			ExpectedSize: len("name&END\r\nvalue\r\nEND"),
		},
		{
			String:           "name&END\r\nvalue\r\nEND\r\n",
			ExpectedName:     "name",
			ExpectedValue:              "value\r\n",
			ExpectedSize: len("name&END\r\nvalue\r\nEND\r\n"),
		},
		{
			String:           "name&END\r\nvalue\r\nEND\r\nwow",
			ExpectedName:     "name",
			ExpectedValue:              "value\r\n",
			ExpectedSize: len("name&END\r\nvalue\r\nEND\r\n"),
		},


		{
			String:           "name&END\u0085value",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\u0085value"),
		},
		{
			String:           "name&END\u0085valueEND",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\u0085valueEND"),
		},
		{
			String:           "name&END\u0085value\u0085",
			ExpectedName:     "name",
			ExpectedValue:              "value\u0085",
			ExpectedSize: len("name&END\u0085value\u0085"),
		},
		{
			String:           "name&END\u0085value\u0085END",
			ExpectedName:     "name",
			ExpectedValue:              "value\u0085",
			ExpectedSize: len("name&END\u0085value\u0085END"),
		},
		{
			String:           "name&END\u0085value\u0085END\u0085",
			ExpectedName:     "name",
			ExpectedValue:              "value\u0085",
			ExpectedSize: len("name&END\u0085value\u0085END\u0085"),
		},
		{
			String:           "name&END\u0085value\u0085END\u0085wow",
			ExpectedName:     "name",
			ExpectedValue:              "value\u0085",
			ExpectedSize: len("name&END\u0085value\u0085END\u0085"),
		},


		{
			String:           "name&END\u2028value",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\u2028value"),
		},
		{
			String:           "name&END\u2028valueEND",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name&END\u2028valueEND"),
		},
		{
			String:           "name&END\u2028value\u2028",
			ExpectedName:     "name",
			ExpectedValue:              "value\u2028",
			ExpectedSize: len("name&END\u2028value\u2028"),
		},
		{
			String:           "name&END\u2028value\u2028END",
			ExpectedName:     "name",
			ExpectedValue:              "value\u2028",
			ExpectedSize: len("name&END\u2028value\u2028END"),
		},
		{
			String:           "name&END\u2028value\u2028END\u2028",
			ExpectedName:     "name",
			ExpectedValue:              "value\u2028",
			ExpectedSize: len("name&END\u2028value\u2028END\u2028"),
		},
		{
			String:           "name&END\u2028value\u2028END\u2028wow",
			ExpectedName:     "name",
			ExpectedValue:              "value\u2028",
			ExpectedSize: len("name&END\u2028value\u2028END\u2028"),
		},









		{
			String:           "name value\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\n"),
		},
		{
			String:           "name value\n\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\n\r"),
		},
		{
			String:           "name value\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\r"),
		},
		{
			String:           "name value\r\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\r\n"),
		},
		{
			String:           "name value\u0085line2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\u0085"),
		},
		{
			String:           "name value\u2028line2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name value\u2028"),
		},



		{
			String:           "name     value\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\n"),
		},
		{
			String:           "name     value\n\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\n\r"),
		},
		{
			String:           "name     value\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\r"),
		},
		{
			String:           "name     value\r\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\r\n"),
		},
		{
			String:           "name     value\u0085line2=something",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\u0085"),
		},
		{
			String:           "name     value\u2028line2=something",
			ExpectedName:     "name",
			ExpectedValue:             "value",
			ExpectedSize: len("name     value\u2028"),
		},



		{
			String:           "name\tvalue\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\n"),
		},
		{
			String:           "name\tvalue\n\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\n\r"),
		},
		{
			String:           "name\tvalue\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\r"),
		},
		{
			String:           "name\tvalue\r\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\r\n"),
		},
		{
			String:           "name\tvalue\u0085line2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\u0085"),
		},
		{
			String:           "name\tvalue\u2028line2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name\tvalue\u2028"),
		},



		{
			String:           "name\t\t\tvalue\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\n"),
		},
		{
			String:           "name\t\t\tvalue\n\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\n\r"),
		},
		{
			String:           "name\t\t\tvalue\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\r"),
		},
		{
			String:           "name\t\t\tvalue\r\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\r\n"),
		},
		{
			String:           "name\t\t\tvalue\u0085line2=something",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\u0085"),
		},
		{
			String:           "name\t\t\tvalue\u2028line2=something",
			ExpectedName:     "name",
			ExpectedValue:              "value",
			ExpectedSize: len("name\t\t\tvalue\u2028"),
		},



		{
			String:           "name: value\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\n"),
		},
		{
			String:           "name: value\n\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\n\r"),
		},
		{
			String:           "name: value\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\r"),
		},
		{
			String:           "name: value\r\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\r\n"),
		},
		{
			String:           "name: value\u0085line2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\u0085"),
		},
		{
			String:           "name: value\u2028line2=something",
			ExpectedName:     "name",
			ExpectedValue:          "value",
			ExpectedSize: len("name: value\u2028"),
		},



		{
			String:           "name : value\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\n"),
		},
		{
			String:           "name : value\n\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\n\r"),
		},
		{
			String:           "name : value\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\r"),
		},
		{
			String:           "name : value\r\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\r\n"),
		},
		{
			String:           "name : value\u0085line2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\u0085"),
		},
		{
			String:           "name : value\u2028line2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name : value\u2028"),
		},



		{
			String:           "name=value\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\n"),
		},
		{
			String:           "name=value\n\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\n\r"),
		},
		{
			String:           "name=value\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\r"),
		},
		{
			String:           "name=value\r\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\r\n"),
		},
		{
			String:           "name=value\u0085line2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\u0085"),
		},
		{
			String:           "name=value\u2028line2=something",
			ExpectedName:     "name",
			ExpectedValue:         "value",
			ExpectedSize: len("name=value\u2028"),
		},



		{
			String:           "name = value\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\n"),
		},
		{
			String:           "name = value\n\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\n\r"),
		},
		{
			String:           "name = value\rline2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\r"),
		},
		{
			String:           "name = value\r\nline2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\r\n"),
		},
		{
			String:           "name = value\u0085line2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\u0085"),
		},
		{
			String:           "name = value\u2028line2=something",
			ExpectedName:     "name",
			ExpectedValue:           "value",
			ExpectedSize: len("name = value\u2028"),
		},









		{
			String:           "\t name value\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name value\n"),
		},
		{
			String:           "\t name value\n\rline2=something",
			ExpectedName   :     "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name value\n\r"),
		},
		{
			String:           "\t name value\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name value\r"),
		},
		{
			String:           "\t name value\r\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name value\r\n"),
		},
		{
			String:           "\t name value\u0085line2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name value\u0085"),
		},
		{
			String:           "\t name value\u2028line2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name value\u2028"),
		},



		{
			String:           "\t name     value\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:                "value",
			ExpectedSize: len("\t name     value\n"),
		},
		{
			String:           "\t name     value\n\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:                "value",
			ExpectedSize: len("\t name     value\n\r"),
		},
		{
			String:           "\t name     value\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:                "value",
			ExpectedSize: len("\t name     value\r"),
		},
		{
			String:           "\t name     value\r\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:                "value",
			ExpectedSize: len("\t name     value\r\n"),
		},
		{
			String:           "\t name     value\u0085line2=something",
			ExpectedName:        "name",
			ExpectedValue:                "value",
			ExpectedSize: len("\t name     value\u0085"),
		},
		{
			String:           "\t name     value\u2028line2=something",
			ExpectedName:        "name",
			ExpectedValue:                "value",
			ExpectedSize: len("\t name     value\u2028"),
		},



		{
			String:           "\t name\tvalue\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name\tvalue\n"),
		},
		{
			String:           "\t name\tvalue\n\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name\tvalue\n\r"),
		},
		{
			String:           "\t name\tvalue\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name\tvalue\r"),
		},
		{
			String:           "\t name\tvalue\r\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name\tvalue\r\n"),
		},
		{
			String:           "\t name\tvalue\u0085line2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name\tvalue\u0085"),
		},
		{
			String:           "\t name\tvalue\u2028line2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name\tvalue\u2028"),
		},



		{
			String:           "\t name\t\t\tvalue\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:                 "value",
			ExpectedSize: len("\t name\t\t\tvalue\n"),
		},
		{
			String:           "\t name\t\t\tvalue\n\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:                 "value",
			ExpectedSize: len("\t name\t\t\tvalue\n\r"),
		},
		{
			String:           "\t name\t\t\tvalue\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:                 "value",
			ExpectedSize: len("\t name\t\t\tvalue\r"),
		},
		{
			String:           "\t name\t\t\tvalue\r\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:                 "value",
			ExpectedSize: len("\t name\t\t\tvalue\r\n"),
		},
		{
			String:           "\t name\t\t\tvalue\u0085line2=something",
			ExpectedName:        "name",
			ExpectedValue:                 "value",
			ExpectedSize: len("\t name\t\t\tvalue\u0085"),
		},
		{
			String:           "\t name\t\t\tvalue\u2028line2=something",
			ExpectedName:        "name",
			ExpectedValue:                 "value",
			ExpectedSize: len("\t name\t\t\tvalue\u2028"),
		},



		{
			String:           "\t name: value\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name: value\n"),
		},
		{
			String:           "\t name: value\n\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name: value\n\r"),
		},
		{
			String:           "\t name: value\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name: value\r"),
		},
		{
			String:           "\t name: value\r\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name: value\r\n"),
		},
		{
			String:           "\t name: value\u0085line2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name: value\u0085"),
		},
		{
			String:           "\t name: value\u2028line2=something",
			ExpectedName:        "name",
			ExpectedValue:             "value",
			ExpectedSize: len("\t name: value\u2028"),
		},



		{
			String:           "\t name : value\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name : value\n"),
		},
		{
			String:           "\t name : value\n\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name : value\n\r"),
		},
		{
			String:           "\t name : value\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name : value\r"),
		},
		{
			String:           "\t name : value\r\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name : value\r\n"),
		},
		{
			String:           "\t name : value\u0085line2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name : value\u0085"),
		},
		{
			String:           "\t name : value\u2028line2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name : value\u2028"),
		},



		{
			String:           "\t name=value\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name=value\n"),
		},
		{
			String:           "\t name=value\n\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name=value\n\r"),
		},
		{
			String:           "\t name=value\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name=value\r"),
		},
		{
			String:           "\t name=value\r\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name=value\r\n"),
		},
		{
			String:           "\t name=value\u0085line2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name=value\u0085"),
		},
		{
			String:           "\t name=value\u2028line2=something",
			ExpectedName:        "name",
			ExpectedValue:            "value",
			ExpectedSize: len("\t name=value\u2028"),
		},



		{
			String:           "\t name = value\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name = value\n"),
		},
		{
			String:           "\t name = value\n\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name = value\n\r"),
		},
		{
			String:           "\t name = value\rline2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name = value\r"),
		},
		{
			String:           "\t name = value\r\nline2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name = value\r\n"),
		},
		{
			String:           "\t name = value\u0085line2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name = value\u0085"),
		},
		{
			String:           "\t name = value\u2028line2=something",
			ExpectedName:        "name",
			ExpectedValue:              "value",
			ExpectedSize: len("\t name = value\u2028"),
		},
	}

	for testNumber, test := range tests {

		var p []byte = []byte(test.String)

		actualName, actualValue, actualSize, err := ininamevalue.ParseBytes(p)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one", testNumber)
			t.Logf("ERROR: (%T) %q.", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.ExpectedName
			actual   := actualName

			if expected != actual {
				t.Errorf("For test #%d, the actual 'name' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			expected := test.ExpectedValue
			actual   := actualValue

			if expected != actual {
				t.Errorf("For test #%d, the actual 'value' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}

		{
			expected := test.ExpectedSize
			actual   := actualSize

			if expected != actual {
				t.Errorf("For test #%d, the actual 'size' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}
