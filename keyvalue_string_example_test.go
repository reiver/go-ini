package ini_test

import (
	"github.com/reiver/go-ini"

	"fmt"
)

func ExampleKeyValueToString() {
	var key   string = "Banana"
	var value string = "TWO"

	var result string = ini.KeyValueToString(key, value)

	fmt.Print(result)

	// Output: Banana = TWO
}
