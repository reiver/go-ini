package ini_test

import (
	"bytes"
	"fmt"
	"io"

	"github.com/reiver/go-ini"
)

func ExampleAppendKeyValue() {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	var key   string = "Banana"
	var value string = "TWO"

	p = ini.AppendKeyValue(p, key, value) // <---------

	var result string = string(p)

	fmt.Print(result)

	// Output: Banana = TWO
}

func ExampleKeyValueToString() {
	var key   string = "Banana"
	var value string = "TWO"

	var result string = ini.KeyValueToString(key, value) // <---------

	fmt.Print(result)

	// Output: Banana = TWO
}

func ExampleWriteKeyValue() {
	var buffer bytes.Buffer
	var writer io.Writer = &buffer

	var key   string = "Banana"
	var value string = "TWO"

	err := ini.WriteKeyValue(writer, key, value) // <---------
	if nil != err {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	var result string = buffer.String()

	fmt.Print(result)

	// Output: Banana = TWO
}
