/*
Package iniscanner_key provides tools for scanning an INI key.


Reading

Getting an INI key token is done with code like the following:

	var runeScanner io.RuneScanner
	
	// ...
	
	token, n, err := iniscanner_key.Read(runeScanner)

If what was next in the io.RuneScanner was not an INI key, then iniscanner_key.Read() will return an error.

Else, it will return the token, and the number of bytes read from the io.RuneScanner.

*/
package iniscanner_key
