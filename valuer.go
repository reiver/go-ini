package ini

// Valuer is used by [ValueOf]
type Valuer interface {
	INIValue() (string, error)
}
