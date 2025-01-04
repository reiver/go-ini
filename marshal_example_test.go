package ini_test

import (
	"github.com/reiver/go-ini"

	"fmt"
)

func ExampleMarshal_MapStringSliceString() {

	var src = map[string][]string{
		"apple":[]string{"1"},
		"Banana":[]string{"2","22"},
		"CHERRY":[]string{"3","33","333"},
		"dAtE":[]string{"4","44","444","4444"},
	}

	p, err := ini.Marshal(src, "fruits") // <---------

	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	var result string = string(p)

	fmt.Print(result)

	// Output:
	// [fruits]
	// apple = 1
	// Banana = 2
	// Banana = 22
	// CHERRY = 3
	// CHERRY = 33
	// CHERRY = 333
	// dAtE = 4
	// dAtE = 44
	// dAtE = 444
	// dAtE = 4444
}

func ExampleMarshal_sliceMapStringString() {

	var src = []map[string]string{
		map[string]string{
			"apple":"ONE",
			"Banana":"TWO",
			"CHERRY":"THREE",
			"dAtE":"FOUR",
		},
		map[string]string{
			"apple":"1",
			"Banana":"2",
			"CHERRY":"3",
			"dAtE":"4",
		},
		map[string]string{
			"apple":"once",
			"Banana":"twice",
			"CHERRY":"thrice",
			"dAtE":"fource",
		},
	}

	p, err := ini.Marshal(src, "fruits") // <---------

	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	var result string = string(p)

	fmt.Print(result)

	// Output:
	// [[fruits]]
	// apple = ONE
	// Banana = TWO
	// CHERRY = THREE
	// dAtE = FOUR
	// [[fruits]]
	// apple = 1
	// Banana = 2
	// CHERRY = 3
	// dAtE = 4
	// [[fruits]]
	// apple = once
	// Banana = twice
	// CHERRY = thrice
	// dAtE = fource
}

func ExampleMarshal_mapStringString() {

	var src = map[string]string{
		"apple":"ONE",
		"Banana":"TWO",
		"CHERRY":"THREE",
	}

	p, err := ini.Marshal(src) // <---------

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

	p, err := ini.Marshal(src) // <---------

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

	result, err := ini.Marshal(m) // <---------

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
