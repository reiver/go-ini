/*
Package intoken provides types useful for representing tokens that result from parsing INI data.

Note that parsing, scanning, and tokenizing does NOT happen in this package.


Creation

The way these are created are with the "some funcs" or the "none func",

For example, the "some funcs" look like this:

	var token initoken.Type = initoken.SomeComment("; This is a comment.")

	var token initoken.Type = initoken.SomeKey("the_key")

	var token initoken.Type = initoken.SomeSection("[the section]")

	var token initoken.Type = initoken.SomeSeparator("=")

	var token initoken.Type = initoken.SomeUndefined("")

	var token initoken.Type = initoken.SomeValue("the value")

	var token initoken.Type = initoken.SomeWhitespace("\n\n    \n")

And, for example the "none func" looks like this:

	var token initoken.Type = initoken.None()


Example

	var token initoken.Type
	
	// ...
	
	switch token {
	case initoken.SomeType:
		//@TODO

		// If the code gets here, then `token` is of type:
		// • initoken.Comment
		// • initoken.Key
		// • initoken.Section
		// • initoken.Separator
		// • initoken.Undefined
		// • initoken.Value
		// • initoken.Whitespace

	case initoken.NoneType:
		//@TODO
		
	default:
		//@TODO: This should never happen. This is probably an error.
	}

Example

	var token initoken.Type
	
	// ...
	
	switch token {
	case initoken.SomeType:
		
		switch token {
		case initoken.Comment:
			//@TODO
		
		case initoken.Key:
			//@TODO
		case initoken.Section:
			//@TODO
		
		case initoken.Separator:
			//@TODO
		
		case initoken.Undefined:
			//@TODO: This is probably a syntax error in the INI data.
		
		case initoken.Value:
			//@TODO
		
		case initoken.Whitespace:
			//@TODO
		
		default:
			//@TODO: This should never happen. This is probably an error.
		}
		
	case initoken.NoneType:
		//@TODO
		
	default:
		//@TODO: This should never happen. This is probably an error.
	}

Example

	var token initoken.Type
	
	// ...
	
	switch token {
	case initoken.Comment:
		//@TODO
		
	case initoken.Key:
		//@TODO
		
	case initoken.Section:
		//@TODO
		
	case initoken.Separator:
		//@TODO
		
	case initoken.Undefined:
		//@TODO: This is probably a syntax error in the INI data.
		
	case initoken.Value:
		//@TODO
		
	case initoken.Whitespace:
		//@TODO
		
	case initoken.NoneType:
		//@TODO: This is probably an error.
		
	default:
		//@TODO: This should never happen. This is probably an error.
	}
*/
package initoken
