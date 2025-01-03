package ini_test

import (
	"bytes"
	"fmt"
	"io"

	"github.com/reiver/go-ini"
)

func ExampleAppendSectionHeader() {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = ini.AppendSectionHeader(p, "apple", "Banana", "CHERRY") // <---------

	var result string = string(p)

	fmt.Print(result)

	// Output: [apple.Banana.CHERRY]
}

func ExampleSectionHeaderToString() {
	var result string = ini.SectionHeaderToString("apple", "Banana", "CHERRY") // <---------

	fmt.Print(result)

	// Output: [apple.Banana.CHERRY]
}

func ExampleWriteSectionHeader() {
	var buffer bytes.Buffer
	var writer io.Writer = &buffer

	err := ini.WriteSectionHeader(writer, "apple", "Banana", "CHERRY") // <---------

	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	var result string = buffer.String()

	fmt.Print(result)

	// Output: [apple.Banana.CHERRY]
}
