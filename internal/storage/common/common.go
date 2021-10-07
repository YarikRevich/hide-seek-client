package common

import "strings"

type StorageBlock interface {
	Get(string) interface{}

	//If value does not exist it will create it
	//otherwise will modify
	Save(DBQuery)
}

type QueryEntity struct {
	Field string
	Value interface{}
}

type DBQuery []QueryEntity

func (d *DBQuery) FormattedFields() string{
	l := make([]string, len(*d))
	for _, v := range *d{
		l = append(l, v.Field)
	}
	return strings.Join(l, ",")[1:]
}

func (d *DBQuery) FormattedValues() []interface{}{
	l := make([]interface{}, len(*d)-1)
	for _, v := range *d{
		l = append(l, v.Value)
	}
	return l
}