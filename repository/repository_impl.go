package repository

import (
	"consumer-rabbitmq/database"
	"consumer-rabbitmq/model"
	"log"
	"time"
)

type repositoryImplementation struct {
	Db database.Db
}

func NewRepository(db database.Db) Repository {
	return &repositoryImplementation{
		Db: db,
	}
}

func (rI repositoryImplementation) Update(id int, information model.EmailInformation) (*model.EmailInformation, error) {
	smt, err := rI.Db.Db.Prepare("UPDATE email_information SET sent = ? WHERE id = ?")
	if err != nil {
		log.Fatal("Not possible Create statement to update email_information", err)
		return nil, err
	}
	_, err = smt.Exec(information.Sent, id)
	if err != nil {
		log.Fatal("Not possible to save into the table", err)
		return nil, err
	}

	return &information, nil
}

func (rI repositoryImplementation) Save(user *model.User, sent bool) (*model.User, error) {
	smt, err := rI.Db.Db.Prepare("INSERT INTO email_information (sent, created_at, updated_at, email, name, last_name) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Not possible Create statement to save LOG", err)
		return nil, err
	}
	_, err = smt.Exec(sent, time.Now(), time.Now(), user.Email, user.Name, user.LastName)
	if err != nil {
		log.Fatal("Not possible to save into the table", err)
		return nil, err
	}

	return user, nil
}

func (rI repositoryImplementation) FindById(id int) (*model.EmailInformation, error) {
	var emailInformation model.EmailInformation
	err := rI.Db.Db.QueryRow("SELECT id, sent, created_at, updated_at, email, name, last_name FROM email_information where id = ? ", id).Scan(&emailInformation.Id, &emailInformation.Sent, &emailInformation.CreatedAt, &emailInformation.UpdatedAt, &emailInformation.Email, &emailInformation.Name, &emailInformation.LastName)
	if err != nil {
		log.Fatal("Not possible to query row", err)
		return nil, err
	}
	return &emailInformation, nil
}

func (rI repositoryImplementation) Find(offset string, limit string) (*[]model.EmailInformation, error) {
	var emailInformations = make([]model.EmailInformation, 0)
	stmt, err := rI.Db.Db.Prepare("SELECT id, sent, created_at, updated_at, email, name, last_name FROM email_information WHERE sent = ? LIMIT ? OFFSET ? ")
	if err != nil {
		log.Fatal("Not possible Create statement", err)
		return nil, err
	}
	rows, err := stmt.Query(false, limit, offset)
	defer rows.Close()
	if err != nil {
		log.Fatal("Not possible to save into LOG", err)
		return nil, err
	}
	for rows.Next() {
		var emailInformation model.EmailInformation
		if err := rows.Scan(&emailInformation.Id, &emailInformation.Sent, &emailInformation.CreatedAt, &emailInformation.UpdatedAt,
			&emailInformation.Email, &emailInformation.Name, &emailInformation.LastName); err != nil {
			return &emailInformations, err
		}
		emailInformations = append(emailInformations, emailInformation)
	}
	if err = rows.Err(); err != nil {
		return &emailInformations, err
	}
	return &emailInformations, nil
}

func (rI repositoryImplementation) ShowAll(offset string, limit string) (*[]model.EmailInformation, error) {
	var emailInformations = make([]model.EmailInformation, 0)
	stmt, err := rI.Db.Db.Prepare("SELECT id, sent, created_at, updated_at, email, name, last_name FROM email_information LIMIT ? OFFSET ? ")
	if err != nil {
		log.Fatal("Not possible Create statement", err)
		return nil, err
	}
	rows, err := stmt.Query(limit, offset)
	defer rows.Close()
	if err != nil {
		log.Fatal("Not possible to save into LOG", err)
		return nil, err
	}
	for rows.Next() {
		var emailInformation model.EmailInformation
		if err := rows.Scan(&emailInformation.Id, &emailInformation.Sent, &emailInformation.CreatedAt, &emailInformation.UpdatedAt,
			&emailInformation.Email, &emailInformation.Name, &emailInformation.LastName); err != nil {
			return &emailInformations, err
		}
		emailInformations = append(emailInformations, emailInformation)
	}
	if err = rows.Err(); err != nil {
		return &emailInformations, err
	}
	return &emailInformations, nil
}

func (rI repositoryImplementation) ShowAllSentFlag(sent bool) (*[]model.EmailInformation, error) {
	var emailInformations = make([]model.EmailInformation, 0)
	stmt, err := rI.Db.Db.Prepare("SELECT id, sent, created_at, updated_at, email, name, last_name FROM email_information WHERE sent = ? ")
	if err != nil {
		log.Fatal("Not possible Create statement", err)
		return nil, err
	}
	rows, err := stmt.Query(sent)
	defer rows.Close()
	if err != nil {
		log.Fatal("Not possible to save into LOG", err)
		return nil, err
	}
	for rows.Next() {
		var emailInformation model.EmailInformation
		if err := rows.Scan(&emailInformation.Id, &emailInformation.Sent, &emailInformation.CreatedAt, &emailInformation.UpdatedAt,
			&emailInformation.Email, &emailInformation.Name, &emailInformation.LastName); err != nil {
			return &emailInformations, err
		}
		emailInformations = append(emailInformations, emailInformation)
	}
	if err = rows.Err(); err != nil {
		return &emailInformations, err
	}
	return &emailInformations, nil
}
