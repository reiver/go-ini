package ini

// Unmarshaler is the interface implemented by types that can unmarshal an INI description of themselves.
// The input can be assumed to be a valid encoding of an INI value.
// UnmarshalINI must copy the INI data if it wishes to retain the data after returning.
type Unmarshaler interface {
	UnmarshalINI([]byte) error
}
