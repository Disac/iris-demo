package services

var UserServ = NewUserService()

func NewUserService() *User {
	return &User{}
}

type User struct {
}

func (serv *User) CheckExist(loginName string) (exist bool, err error) {
	return
}
