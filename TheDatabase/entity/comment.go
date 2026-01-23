package entity

type Comment struct {
	Id      int
	Email   string
	Comment string
}

type User struct {
	Id       int
	Username string
	Password string
}
