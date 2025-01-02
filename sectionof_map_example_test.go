package ini_test

import (
	"fmt"

	"github.com/reiver/go-ini"
)

func ExampleSectionOf_mapString() {

	var m = map[string]string {
		"apple":  "ONCE",
		"Banana": "TWICE",
		"CHERRY": "THRICE",
		"dAtE":   "FOURCE",
	}

	result, err := ini.SectionOf(m)

	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	fmt.Printf("%s", result)

	// Output:
	// apple = ONCE
	// Banana = TWICE
	// CHERRY = THRICE
	// dAtE = FOURCE
}
