/*
Package iniscanner_comment provides tools for scanning an INI comment.


Reading

	var runeScanner io.RuneScanner
	
	// ...
	
	token, n, err := iniscanner_comment.Read(runeScanner)

If what was next in the io.RuneScanner was not an INI comment, then iniscanner_comment.Read() will return an error.

Else, it will return the token, and the number of bytes read from the io.RuneScanner.


Peeking

Sometimes it is useful to be able to tell if what is next in INI data ia a comment or not.

This is done with the iniscanner_comment.Peek() func.

For example:

	var r rune
	
	// ...
	
	if iniscanner_comment.Peek(r) {
		//@TODO
	}

*/
package iniscanner_comment
