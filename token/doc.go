/*
Package intoken provides types useful for representing tokens that result from parsing INI data.

Note that parsing, scanning, and tokenizing does NOT happen in this package.


INI Tokens

Conceptually, there are 6 different kinds of INI tokens:

• comments,

• keys,

• sections,

• separators,

• values, and

• whitespace.


An example INI file with all of these tokens would be:

	# This is a comment.
	
	the_key=the value

The tokens in this example INI file would be:

• comment: "# This is a comment."

• whitespace: "\n\n"

• key: "the_key"

• separator: "="

• value: "the value"

• whitespace: "\n"

From the perspective of this package, these would be encoded as:

• initoken.Somecomment("# This is a comment.")

• initoken.SomeWhitespace("\n\n")

• initoken.SOmeKey("the_key")

• initoken.SomeSeparator("=")

• initoken.SomeValue("the value")

• initoken.SomeWhitespace("\n")


Creation

The way these are created are with the "some funcs" or the "none func",

For example, the "some funcs" look like this:

	var token initoken.Type = initoken.SomeComment("; This is a comment.")
.
	var token initoken.Type = initoken.SomeKey("the_key")
.
	var token initoken.Type = initoken.SomeSection("[the section]")
.
	var token initoken.Type = initoken.SomeSeparator("=")
.
	var token initoken.Type = initoken.SomeUndefined("")
.
	var token initoken.Type = initoken.SomeValue("the value")
.
	var token initoken.Type = initoken.SomeWhitespace("\n\n    \n")

And, for example the "none func" looks like this:

	var token initoken.Type = initoken.None()

Probably Errors

If you received a token from these:

	var token initoken.Type = initoken.SomeUndefined("")
.
	var token initoken.Type = initoken.None()

Then likely some kind of error has happened, and you should handle it in whatever is the appropriate way for your application.


Discerning

Being to discern between the different kinds of tokens is important.

That can be done using the Golang type-switch.

Here are some example of that in use:

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
