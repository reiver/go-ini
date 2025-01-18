package ini_test

import (
	"fmt"

	"github.com/reiver/go-ini"
)

func ExampleInvalidUnmarshalError() {

	var bytes []byte = []byte(
`[stuff]

apple = once
Banana = twice
CHERRY = thrice
dAtE = fource
`)

	var dst float64

	err := ini.Unmarshal(bytes, &dst)

	if nil != err {
		switch casted := err.(type) {
		case ini.InvalidUnmarshalError:
			fmt.Printf("ini.Unmarshal() cannot unmarshal into something of type %s\n", casted.InvalidType())
		default:
			fmt.Println("unknown error:", err)
		}
	}

	// Output:
	// ini.Unmarshal() cannot unmarshal into something of type *float64
}
