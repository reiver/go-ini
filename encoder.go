package ini

import (
	"io"
)

// Encoder is the interface implemented by types that can encode themselves into valid INI.
//
// What is being written, by EncodeINI, is the value part in a key-value pair.
//
// This means that EncodeINI would do something a bit different if the custom type that implemented
// Encoder is a simple INI value or complex INI value.
//
// I.e., EncodeINI would do something a bit different if the custom type that implemented Encoder
// has a set of key-value pairs of its own or not.
//
//
// Simple Values Versus Complex Values
//
// A simple INI value is something such as:
//
//	22cm
//
// A complex INI value is something such as:
//
//	{
//		border-style = solid
//		border-width = 1px
//		border-color = rgb(22,22,29)
//	}
//
// Note that a complex value uses the extended INI format.
//
// So, in the context of a key-value pair, a simple INI value is something such as:
//
//	length = 22cm
//
// Likewise, in the context of a key-value pair, a complex INI value is something such as:
//
//	square = {
//		border-style = solid
//		border-width = 1px
//		border-color = rgb(22,22,29)
//	}
//
// (Of course, the name and separator in the key-value pair (such as “length” and “square” in the examples) are generated elsewhere.)
//
//
// Simple INI Value
//
// Here is an example custom type that generates a simple INI value.
//
//	type Centimeter64 struct {
//		Value int64
//	}
//	
//	func (receiver Centimeter64) EncodeINI(writer, io.Writer) error {
//		fmt.Fprintf(writer, "%d", receiver.Value)
//		io.WriteString(writer, "cm")
//	}
//
// So (for example), if ‘receiver.Value’ was ‘5’…. I.e.,:…
//
//	var cm64 Centimeter64 = Centimeter64{5}
//
// Then its EncodeINI would write:
//
//	5cm
//
// And (for a different example), if ‘receiver.Value’ was ‘-23’…. I.e.,:…
//
//	var cm64 Centimeter64 = Centimeter64{-23}
//
// Then its EncodeINI would write:
//
//	-23cm
//
// A more full example using this custom type (in our example) might be something such as:
//
//	var data map[string]interface{} = map[string]interface{}}
//	
//	data["apple"] = "one"
//	
//	data["banana"] = 2
//	
//	data["cherry"] = "۳"
//	
//	data["Cube"] = map[string]interface{}{
//		"color":  "red",
//		"height": Centimeter64{2}
//		"length": Centimeter64{3}
//		"width":  Centimeter64{5}
//	}
//
// And then calling ini.Encode for this, i.e.,:
//
//	err := ini.Encode(writer, data)
//
// Would produce INI data like the following:
//
//	apple = one
//	banana = 2
//	cherry = ۳
//	
//	[Cube]
//	color = red
//	height = 2cm
//	length = 3cm
//	width = 5cm
//
//
// Complex INI Value
//
// Now lets look at a new example, where we illustrate a complex INI values.
//
// Here is an example custom type that generates a complex INI value.
//
//	type Border struct {
//		Style      string
//		WidthValue uint64
//		WidthUnit  string
//		Color      string
//	}
//	
//	func (receiver Border) EncodeINI(writer, io.Writer) error {
//		io.WriteString(writer, "{")
//
//		io.WriteString(writer, "\t")
//		io.WriteString(writer, "border-style = ")
//		io.WriteString(writer, receiver.Style)
//		io.WriteString(writer, "\n")
//
//		io.WriteString(writer, "\t")
//		io.WriteString(writer, "border-width = ")
//		fmt.Fprintf(writer, "%d", receiver.WidthValue)
//		io.WriteString(writer, receiver.WidthUnit)
//		io.WriteString(writer, "\n")
//
//		io.WriteString(writer, "\t")
//		io.WriteString(writer, "border-color = ")
//		io.WriteString(writer, receiver.Color)
//		io.WriteString(writer, "\n")
//
//		io.WriteString(writer, "}")
//	}
//
// So (for example), if we had:…
//
//	var border Border = Border{
//		Style:      "solid",
//		WidthValue: 2,
//		WidthUnit: "px",
//		Color:     "rgb(22,22,29)",
//	}
//
// Then its EncodeINI would write:
//
//	{
//		border-style = solid
//		border-width = 2px
//		border-color = rgb(22,22,29)
//	}
//
// And (for a different example), if we had:…
//
//	var border Border = Border{
//		Style:      "dotted",
//		WidthValue: 1,
//		WidthUnit: "rem",
//		Color:     "indigo",
//	}
//
// Then its EncodeINI would write:
//
//	{
//		border-style = dotted
//		border-width = 1rem
//		border-color = indigo
//	}
//
// A more full example using this custom type (in our new example) might be something such as:
//
//	var data map[string]interface{} = map[string]interface{}}
//	
//	data["apple"] = "one"
//	
//	data["banana"] = 2
//	
//	data["cherry"] = "۳"
//	
//	data["Person"] = map[string]interface{}{
//		"age":  22,
//		"fashion": Border{
//			border-style = solid
//			border-width = 2px
//			border-color = rgb(22,22,29)
//		},
//		"name": "Joe Blow",
//		"salutation": "Mr",
//	}
//
// And then calling ini.Encode for this, i.e.,:
//
//	err := ini.Encode(writer, data)
//
// Would produce INI data like the following:
//
//	apple = one
//	banana = 2
//	cherry = ۳
//	
//	[Person]
//	age = 22
//	fashion = {
//		border-style = solid
//		border-width = 2px
//		border-color = rgb(22,22,29)
//	}
//	name = Joe Blow
//	salutation = Mr
type Encoder interface {
	EncodeINI(io.Writer) error
}
