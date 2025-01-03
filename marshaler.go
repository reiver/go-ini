package ini

// Marshaler is used by non-simple custom types to return its INI value.
//
// Behind the scenes, Marshaler is used by [Marshaler]
//
// If you are implmenting your own MarshalINI method for your own custom type, then you are likely going to use one or more of these functions:
// [Append], [AppendKeyValue], [AppendSectionHeader], [AppendSequenceHeader].
//
// Note that if your custom-type is instead a simple custom type, then it should instead implement [Valuer] instead of Marshaler.
type Marshaler interface {
	MarshalINI([]byte, ...string) ([]byte, error)
}
