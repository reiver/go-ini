/*
Package iniscanner_value provides tools for scanning an INI value.


Reading

Getting an INI value token is done with code like the following:

	var runeScanner io.RuneScanner
	
	// ...
	
	token, n, err := iniscanner_value.Read(runeScanner)

If what was next in the io.RuneScanner was not an INI value, then iniscanner_value.Read() will return an error.

Else, it will return the token, and the number of bytes read from the io.RuneScanner.


Peeking

Sometimes it is useful to be able to tell if what is next in INI data ia a value or not.

This is done with the iniscanner_value.Peek() func.

For example:

	var r rune
	
	// ...
	
	if iniscanner_value.Peek(r) {
		//@TODO
	}

*/
package iniscanner_value
