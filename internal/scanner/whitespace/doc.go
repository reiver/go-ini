/*
Package iniscanner_whitespace provides tools for scanning an INI whitespace.


Reading

Getting an INI whitespace token is done with code like the following:

	var runeScanner io.RuneScanner
	
	// ...
	
	token, n, err := iniscanner_whitespace.Read(runeScanner)

If what was next in the io.RuneScanner was not an INI whitespace, then iniscanner_whitespace.Read() will return an error.

Else, it will return the token, and the number of bytes read from the io.RuneScanner.


Peeking

Sometimes it is useful to be able to tell if what is next in INI data ia a whitespace or not.

This is done with the iniscanner_whitespace.Peek() func.

For example:

	var r rune
	
	// ...
	
	if iniscanner_whitespace.Peek(r) {
		//@TODO
	}

*/
package iniscanner_whitespace
