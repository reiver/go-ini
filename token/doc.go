/*
Package intoken provides types useful for representing tokens that result from parsing INI data.

Example

	var token interface{}
	
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

	case initoken.Value:
		//@TODO

	case initoken.Undefined:
		//@TODO: invalid INI data

	default:
		//@TODO: internal error
	}
*/
package initoken
