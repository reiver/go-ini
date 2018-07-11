package initoken

import (
	"testing"
)

func TestSomeComment(t *testing.T) {

	var x SomeType = Comment{} // THIS IS THE IMPORTANT PART.
	                           // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestSomeKey(t *testing.T) {

	var x SomeType = Key{} // THIS IS THE IMPORTANT PART.
	                       // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestSomeSection(t *testing.T) {

	var x SomeType = Section{} // THIS IS THE IMPORTANT PART.
	                           // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestSomeSeparator(t *testing.T) {

	var x SomeType = Separator{} // THIS IS THE IMPORTANT PART.
	                             // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestSomeValue(t *testing.T) {

	var x SomeType = Value{} // THIS IS THE IMPORTANT PART.
	                         // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestSomeUndefined(t *testing.T) {

	var x SomeType = Undefined{} // THIS IS THE IMPORTANT PART.
	                             // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestSomeWhitespace(t *testing.T) {

	var x SomeType = Whitespace{} // THIS IS THE IMPORTANT PART.
	                              // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
