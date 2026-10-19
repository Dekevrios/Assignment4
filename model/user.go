package model

import (
	
	"github.com/go-crypt/x/bcrypt"
)

type User struct {
	ID       int
	Name     string
	Username string
	Password string // password disini di isi hash
}

var Users []User //data store menggunakan array, di memori

func SeedUsers() {
	hash, err := bcrypt.GenerateFromPassword([]byte("StrongPassword"), 10)
	if err!= nil{
		panic(err)
	}
	Users = []User{
		{ID: 1,Name: "User 1", Username: "user1", Password: string(hash)},
	}
}

func FindbyUsername(username string) (*User, bool){
	for i := range Users{
		if Users[i].Username == username{
			return &Users[i], true
		}
	}
	return nil, false
}