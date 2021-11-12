package storage

type IUser interface {
	GetUsername() string
	SetUsername(string) 
}