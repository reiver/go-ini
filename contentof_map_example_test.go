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
		"z":map[string]any{
			"z-1":map[string]any{
				"z-1-1":"HERE",
				"z-1-2":map[string]string{
					"z-1-2-1":"(1)",
					"z-1-2-2":"(2)",
					"z-1-2-3":"(3)",
				},
			},
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
	// [dairy]
	// cheese = 20
	// yogurt = 12
	// [dairy.milk]
	// 1% = 2
	// 2% = 1
	// [fruits]
	// apple = ONCE
	// Banana = TWICE
	// CHERRY = THRICE
	// dAtE = FOURCE
	// [z]
	// [z.z-1]
	// z-1-1 = HERE
	// [z.z-1.z-1-2]
	// z-1-2-1 = (1)
	// z-1-2-2 = (2)
	// z-1-2-3 = (3)
}
