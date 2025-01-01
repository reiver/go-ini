package ini

type KeyValueIterator interface {
	For(func(string,string)error)error
}
