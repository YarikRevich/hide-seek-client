package storage

import "strings"

type Query struct {
	Fields []string
	Values []interface{}
}

func (q *Query) AddField(field string) {
	q.Fields = append(q.Fields, field)
}

func (q *Query) AddValue(value interface{}) {
	q.Values = append(q.Values, value)
}

func (q *Query) GetFieldsAsString() string {
	return strings.Join(q.Fields, ",")[1:]
}

func NewQuery()*Query{
	return new(Query)
}