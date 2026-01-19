package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Day 1
type User struct {
	Name  string
	Email string
}

func UpdateEmail(u *User, email string) {
	u.Email = email
}

// Day 2
type User2 struct {
	Id   int
	Name string
}

type UserRepo interface {
	FindById(Id int) (*User2, error)
}

type InMemoryRepo struct{}

func (r InMemoryRepo) FindById(Id int) (*User2, error) {
	if Id != 1 {
		return nil, errors.New("User not found")
	}

	return &User2{Id: 1, Name: "Yoku"}, nil
}

// Day 3
type User3 struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user3 User3
	err := json.NewDecoder(r.Body).Decode(&user3)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user3)
}

func main() {
	user := User{
		Name:  "Yoku",
		Email: "yoku@gmail.com",
	}
	fmt.Println(user.Email)

	UpdateEmail(&user, "notyoku@gmail.com")

	fmt.Println(user.Email)

	fmt.Println("---------------Day 2-----------------------")

	repo := InMemoryRepo{}
	userRepo, err := repo.FindById(2)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("User : ", userRepo.Name)

	fmt.Println("---------------Day 3-----------------------")
	fmt.Println("Server running on http://localhost:3000")
	http.HandleFunc("/users", CreateUser)
	http.ListenAndServe(":3000", nil)
}
