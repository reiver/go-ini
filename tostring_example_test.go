package ini_test

import (
	"github.com/reiver/go-ini"

	"fmt"
)

func ExampleToString_mapStringString() {

	var src = map[string]string{
		"apple":"ONE",
		"Banana":"TWO",
		"CHERRY":"THREE",
	}

	var result string
	var err error

	result, err = ini.ToString(src)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Print(result)

	// Output:
	// apple = ONE
	// Banana = TWO
	// CHERRY = THREE
}

func ExampleToString_mapStringAny() {

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

	var result string
	var err error

	result, err = ini.ToString(src)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

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
