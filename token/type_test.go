package initoken

import (
	"testing"
)

func TestTypeComment(t *testing.T) {

	var x Type = Comment{} // THIS IS THE IMPORTANT PART.
	                       // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestTypeKey(t *testing.T) {

	var x Type = Key{} // THIS IS THE IMPORTANT PART.
	                   // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestTypeSection(t *testing.T) {

	var x Type = Section{} // THIS IS THE IMPORTANT PART.
	                       // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestTypeSeparator(t *testing.T) {

	var x Type = Separator{} // THIS IS THE IMPORTANT PART.
	                         // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestTypeUndefined(t *testing.T) {

	var x Type = Undefined{} // THIS IS THE IMPORTANT PART.
	                       // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestTypeValue(t *testing.T) {

	var x Type = Value{} // THIS IS THE IMPORTANT PART.
	                     // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestTypeWhitespace(t *testing.T) {

	var x Type = Whitespace{} // THIS IS THE IMPORTANT PART.
	                       // If the type of the struct does not fit the interface, it will create a compile time error.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}
