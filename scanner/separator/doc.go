/*
Package iniscanner_separator provides tools for scanning an INI separator.


Reading

Getting an INI separator token is done with code like the following:

	var runeScanner io.RuneScanner
	
	// ...
	
	token, n, err := iniscanner_separator.Read(runeScanner)

If what was next in the io.RuneScanner was not an INI separator, then iniscanner_separator.Read() will return an error.

Else, it will return the token, and the number of bytes read from the io.RuneScanner.


Peeking

Sometimes it is useful to be able to tell if what is next in INI data ia a separator or not.

This is done with the iniscanner_separator.Peek() func.

For example:

	var r rune
	
	// ...
	
	if iniscanner_separator.Peek(r) {
		//@TODO
	}

*/
package iniscanner_separator
