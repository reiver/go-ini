package ini

import (
	"fmt"
	"sort"
	"strings"

	"github.com/reiver/go-val"
)

type Record struct {
	data map[string]*val.Values[string]
}

func EmptyRecord() Record {
	return Record{}
}

func NewEmptyRecord() *Record {
	return new(Record)
}

func (receiver *Record) init() {
	if nil == receiver {
		return
	}

	if nil == receiver.data {
		receiver.data = map[string]*val.Values[string]{}
	}
}

func (receiver Record) get(name string) (*val.Values[string], bool) {
	var data map[string]*val.Values[string] = receiver.data
	if len(data) <= 0 {
		return nil, false
	}

	values, found := data[name]
	if !found {
		return nil, false
	}

	return values, true
}

func (receiver *Record) Append(name string, value string) {
	if nil == receiver {
		return
	}

	receiver.init()

	if nil == receiver.data[name] {
		receiver.data[name] = new(val.Values[string])
	}

	var values *val.Values[string] = receiver.data[name]

	values.Append(value)
}

func (receiver *Record) Chain(fn func(*Record)) *Record {
	if nil == fn {
		return receiver
	}

	fn(receiver)
	return receiver
}

func (receiver *Record) ChainAppend(name string, value string) *Record {
	return receiver.Chain(func(record *Record) {
		record.Append(name, value)
	})
}

func (receiver *Record) ChainSet(name string, values ...string) *Record {
	return receiver.Chain(func(record *Record) {
		record.Set(name, values...)
	})
}

func (receiver *Record) ChainUnset(name string) *Record {
	return receiver.Chain(func(record *Record) {
		record.Unset(name)
	})
}

func (receiver Record) First(name string) (string, bool) {
	values, found := receiver.get(name)
	if !found {
		return "", false
	}

	return values.First()
}

func (receiver Record) FirstElse(name string, alternative string) string {
	values, found := receiver.get(name)
	if !found {
		return alternative
	}

	return values.FirstElse(alternative)
}

func (receiver Record) GoString() string {
	if receiver.IsEmpty() {
		return "ini.EmptyRecord()"
	}

	{
		var buffer [256]byte
		var p []byte = buffer[0:0]

		p = append(p, "ini.EmptyRecord()"...)
		for _, name := range receiver.Keys() {
			var quotedValues []string
			for _, value := range receiver.Values(name) {
				var quotedValue string = fmt.Sprintf("%q", value)
				quotedValues = append(quotedValues, quotedValue)
			}

			p = append(p, ".ChainSet("...)
			p = append(p, fmt.Sprintf("%q", name)...)
			p = append(p, ", "...)
			p = append(p, strings.Join(quotedValues, ",")...)
			p = append(p, "...)"...)
		}

		return string(p)
	}
}

func (receiver Record) Has(name string) bool {
	if receiver.IsEmpty() {
		return false
	}

	_, found := receiver.data[name]
	return found
}

func (receiver Record) IsEmpty() bool {
	return len(receiver.data) <= 0
}

func (receiver Record) Keys() []string {
	if receiver.IsEmpty() {
		return []string{}
	}

	var keys []string
	for key, _ := range receiver.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func (receiver Record) Last(name string) (string, bool) {
	values, found := receiver.get(name)
	if !found {
		return "", false
	}

	return values.Last()
}

func (receiver Record) LastElse(name string, alternative string) string {
	values, found := receiver.get(name)
	if !found {
		return alternative
	}

	return values.LastElse(alternative)
}

func (receiver Record) Len() int {
	return len(receiver.data)
}

func (receiver *Record) Record() Record {
	if nil == receiver {
		return EmptyRecord()
	}

	return *receiver
}

func (receiver *Record) Set(name string, values ...string) {
	if nil == receiver {
		return
	}
	if len(values) <= 0 {
		return
	}

	receiver.init()

	receiver.data[name] = val.NewValues(values...)
}

func (receiver *Record) Unset(name string) {
	if nil == receiver {
		return
	}

	if receiver.IsEmpty() {
		return
	}

	delete(receiver.data, name)
}

func (receiver Record) Values(name string) []string {
	values, found := receiver.get(name)
	if !found {
		return []string{}
	}

	return values.All()
}
