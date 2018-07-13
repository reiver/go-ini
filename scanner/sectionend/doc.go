/*
Package iniscanner_sectionend provides tools for scanning an INI sectionend.


Reading

Getting an INI sectionend token is done with code like the following:

	var runeScanner io.RuneScanner
	
	// ...
	
	token, n, err := iniscanner_sectionend.Read(runeScanner)

If what was next in the io.RuneScanner was not an INI sectionend, then iniscanner_sectionend.Read() will return an error.

Else, it will return the token, and the number of bytes read from the io.RuneScanner.


Peeking

Sometimes it is useful to be able to tell if what is next in INI data ia a sectionend or not.

This is done with the iniscanner_sectionend.Peek() func.

For example:

	var r rune
	
	// ...
	
	if iniscanner_sectionend.Peek(r) {
		//@TODO
	}

*/
package iniscanner_sectionend
