package ini_test

import (
	"github.com/reiver/go-ini"

	"fmt"
)

func ExampleSectionHeaderToString() {
	var result string = ini.SectionHeaderToString("apple", "Banana", "CHERRY")

	fmt.Print(result)

	// Output: [apple.Banana.CHERRY]
}
