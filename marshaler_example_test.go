package ini_test

import (
	"fmt"
	"strconv"

	"github.com/reiver/go-ini"
)

type MyCustomType struct {
	color string
	number uint64
}

func (receiver MyCustomType) MarshalINI(p []byte, nesting ...string) ([]byte, error) {
	if 0 < len(nesting) {
		p = ini.AppendSectionHeader(p, nesting...)
	}

	p = ini.AppendKeyValue(p, "color", receiver.color)
	p = ini.AppendKeyValue(p, "number", strconv.FormatUint(receiver.number,10))

	return p, nil
}

func ExampleMarshaler() {

	var value = MyCustomType{
		color: "red",
		number: 18,
	}

	result, err := ini.ToString(value)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Print(result)

	// Output:
	//
	// color = red
	// number = 18
}
