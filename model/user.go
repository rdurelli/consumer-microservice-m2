package model

import "fmt"

type User struct {
	Name     string `json:"name,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Email    string `json:"email,omitempty"`
}

func (u User) String() string {
	return fmt.Sprintf("USER Name: [%v], LastName: [%v], EMAIL: [%v]", u.Name, u.LastName, u.Email)
}
