package storage

var instance *Storage


type Storage struct {
	user *User
}

func (s *Storage) User() *User{
	return s.user
}

func UseStorage() *Storage{
	if instance == nil{
		db := NewDB()
		instance = &Storage{
			NewUser(db),
		}
	}
	return instance
}