/*
Package intoken provides types useful for representing tokens that result from parsing INI data.

Note that parsing, scanning, and tokenizing does NOT happen in this package.

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
