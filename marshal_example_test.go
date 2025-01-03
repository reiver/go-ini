package ini_test

import (
	"github.com/reiver/go-ini"

	"fmt"
)

func ExampleMarshal_mapStringString() {

	var src = map[string]string{
		"apple":"ONE",
		"Banana":"TWO",
		"CHERRY":"THREE",
	}

	p, err := ini.Marshal(src)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	var result string = string(p)

	fmt.Print(result)

	// Output:
	// apple = ONE
	// Banana = TWO
	// CHERRY = THREE
}

func ExampleMarshal_mapStringAny() {

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

	p, err := ini.Marshal(src)
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	var result string = string(p)

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

func ExampleMarshal_mapStringAny_2() {

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

	result, err := ini.Marshal(m)

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
