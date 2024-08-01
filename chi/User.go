package main

type User struct {
	Name string
	Age  int
}

func (u User) getName() string {
	return u.Name
}
