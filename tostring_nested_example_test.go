package ini_test

import (
	"github.com/reiver/go-ini"

	"fmt"
)

func ExampleNestedToString_mapStringString() {

	var src = map[string]string{
		"apple":"ONE",
		"Banana":"TWO",
		"CHERRY":"THREE",
	}

	var result string
	var err error

	result, err = ini.NestedToString(src, "fruits")
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Print(result)

	// Output:
	// [fruits]
	//
	// apple = ONE
	// Banana = TWO
	// CHERRY = THREE
}

func ExampleNestedToString_mapStringAny() {

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

	result, err = ini.NestedToString(src, "root", "node")
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Print(result)

	// Output:
	//
	// [root.node]
	//
	// cookies = 2
	// crackers = 1
	//
	// [root.node.dairy]
	//
	// cheese = 20
	// milk = 2
	// yogurt = 12
	//
	// [root.node.fruits]
	//
	// apple = ONE
	// Banana = TWO
	// CHERRY = THREE
}
