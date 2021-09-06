package model

import "fmt"

type Product struct {
	Id          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Value       float64 `json:"value,omitempty"`
}

func (p Product) String() string {
	return fmt.Sprintf("Product ID %v: Name: [%v], Description: [%v], Value: [%v]", p.Id, p.Name, p.Description, p.Value)
}
