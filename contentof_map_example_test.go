package ini_test

import (
	"fmt"

	"github.com/reiver/go-ini"
)

func ExampleContentOf_mapAny() {

	var m = map[string]any {
		"aardvark":"adult",
		"meese":"pair",
		"moose":"single",
		"dairy":map[string]any{
			"cheese":"20",
			"milk":map[string]string{
				"1%":"2",
				"2%":"1",
			},
			"yogurt":"12",
		},
		"fruits":map[string]string{
			"apple":  "ONCE",
			"Banana": "TWICE",
			"CHERRY": "THRICE",
			"dAtE":   "FOURCE",
		},
	}

	result, err := ini.ContentOf(m)

	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("%s", result)

	// Output:
	// aardvark = adult
	// meese = pair
	// moose = single
	//
	// [dairy]
	//
	// cheese = 20
	// yogurt = 12
	//
	// [dairy.milk]
	//
	// 1% = 2
	// 2% = 1
	//
	// [fruits]
	//
	// apple = ONCE
	// Banana = TWICE
	// CHERRY = THRICE
	// dAtE = FOURCE
	// --TODO--
}
