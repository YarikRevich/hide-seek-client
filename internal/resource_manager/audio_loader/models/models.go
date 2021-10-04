package models

type Controller struct {
	Stop func()
	Start func()
}