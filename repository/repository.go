package repository

import "consumer-rabbitmq/model"

type Reader interface {
	Find(offset string, limit string) (*[]model.EmailInformation, error)
	FindById(id int) (*model.EmailInformation, error)
	ShowAll(offset string, limit string) (*[]model.EmailInformation, error)
}

type Writer interface {
	Save(user *model.User, sent bool) (*model.User, error)
}

type Repository interface {
	Writer
	Reader
}
