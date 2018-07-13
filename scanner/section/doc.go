/*
Package iniscanner_section provides tools for scanning an INI section.


Reading

Getting an INI section token is done with code like the following:

	var runeScanner io.RuneScanner
	
	// ...
	
	token, n, err := iniscanner_section.Read(runeScanner)

If what was next in the io.RuneScanner was not an INI section, then iniscanner_section.Read() will return an error.

Else, it will return the token, and the number of bytes read from the io.RuneScanner.


Peeking

Sometimes it is useful to be able to tell if what is next in INI data ia a section or not.

This is done with the iniscanner_section.Peek() func.

For example:

	var r rune
	
	// ...
	
	if iniscanner_section.Peek(r) {
		//@TODO
	}

*/
package iniscanner_section
