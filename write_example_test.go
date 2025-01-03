package ini_test

import (
	"bytes"
	"fmt"
	"io"

	"github.com/reiver/go-ini"
)

func ExampleWrite_mapStringString() {

	var src = map[string]string{
		"apple":"ONE",
		"Banana":"TWO",
		"CHERRY":"THREE",
	}

	var buffer bytes.Buffer
	var writer io.Writer = &buffer

	err := ini.Write(writer, src)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	var result string = buffer.String()

	fmt.Print(result)

	// Output:
	// apple = ONE
	// Banana = TWO
	// CHERRY = THREE
}

func ExampleWrite_mapStringAny() {

	var src = map[string]any{
		"cookies":"2",
		"crackers":"1",
		"dairy": map[string]string{
			"cheese":"20",
			"milk":"2",
			"yogurt":"12",
		},
		"fruits": map[string]string{
			"apple":"ONE",
			"Banana":"TWO",
			"CHERRY":"THREE",
		},
	}

	var buffer bytes.Buffer
	var writer io.Writer = &buffer

	err := ini.Write(writer, src)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	var result string = buffer.String()

	fmt.Print(result)

	// Output:
	// cookies = 2
	// crackers = 1
	// [dairy]
	// cheese = 20
	// milk = 2
	// yogurt = 12
	// [fruits]
	// apple = ONE
	// Banana = TWO
	// CHERRY = THREE
}
