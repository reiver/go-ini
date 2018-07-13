/*
Package iniscanner_sectionbegin provides tools for scanning an INI sectionbegin.


Reading

Getting an INI sectionbegin token is done with code like the following:

	var runeScanner io.RuneScanner
	
	// ...
	
	token, n, err := iniscanner_sectionbegin.Read(runeScanner)

If what was next in the io.RuneScanner was not an INI sectionbegin, then iniscanner_sectionbegin.Read() will return an error.

Else, it will return the token, and the number of bytes read from the io.RuneScanner.


Peeking

Sometimes it is useful to be able to tell if what is next in INI data ia a sectionbegin or not.

This is done with the iniscanner_sectionbegin.Peek() func.

For example:

	var r rune
	
	// ...
	
	if iniscanner_sectionbegin.Peek(r) {
		//@TODO
	}

*/
package iniscanner_sectionbegin
