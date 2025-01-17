package iniconv_test

import (
	"fmt"

	"github.com/reiver/go-ini/conv"
)

func ExampleParseStrings_comma() {

	var text string = "apple pie, Banana Pie, CHERRY PIE"

	var pies []string = iniconv.ParseStrings(text, ",") // <---------

	for _, pie := range pies {
		fmt.Println(pie)
	}

	// Output:
	// apple pie
	// Banana Pie
	// CHERRY PIE
}

func ExampleParseStrings_semicolon() {

	var text string = "Joe Blow, BSc; Jane Doe, PhD"

	var students []string = iniconv.ParseStrings(text, ";") // <---------

	for _, student := range students {
		fmt.Println(student)
	}

	// Output:
	// Joe Blow, BSc
	// Jane Doe, PhD
}
