package ini_test

import (
	"github.com/reiver/go-ini"

	"fmt"
)

func ExampleToString() {

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
