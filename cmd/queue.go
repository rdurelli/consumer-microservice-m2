/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"consumer-rabbitmq/database"
	email2 "consumer-rabbitmq/email"
	"consumer-rabbitmq/model"
	"consumer-rabbitmq/repository"
	"consumer-rabbitmq/service"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"

	"github.com/spf13/cobra"
)

var (
	//RABBITMQ flags
	user     string
	password string
	address  string
	port     string
	//SMTP flags
	portSmtp     string
	from         string
	passwordSmtp string
	smtpHost     string
	userSmtp     string
	emailService service.EmailService
	//DB vars
	db   database.Db
	repo repository.Repository
)

// queueCmd represents the queue command
var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Init the QUEUE EMAIL",
	Run: func(cmd *cobra.Command, args []string) {
		startQueue()
	},
}

func init() {
	cobra.OnInitialize(initDB, initRepository, initEmailService)
	rootCmd.AddCommand(queueCmd)

	//RABBITMQ FLAGS
	queueCmd.Flags().StringVarP(&user, "user-amqp", "u", "admin", "user to connect to amqp")
	queueCmd.Flags().StringVarP(&password, "password-amqp", "p", "123456", "password to connect to amqp")
	queueCmd.Flags().StringVarP(&address, "address-amqp", "d", "localhost", "address to connect to amqp")
	queueCmd.Flags().StringVarP(&port, "port-amqp", "o", "5672", "port to connect to amqp")
	//SMTP FLAGS
	queueCmd.Flags().StringVar(&portSmtp, "port-smtp", "2525", "port to connect to smtp")
	queueCmd.Flags().StringVar(&from, "from-smtp", "piotr@mailtrap.io", "source email")
	queueCmd.Flags().StringVar(&passwordSmtp, "password-smtp", "2509a5263b10ea", "password to connect to smtp")
	queueCmd.Flags().StringVar(&smtpHost, "host-smtp", "smtp.mailtrap.io", "host to connect to smtp")
	queueCmd.Flags().StringVar(&userSmtp, "user-smtp", "f0740004206318", "user to connect to smtp")

}

func initDB() {
	db = database.NewDataBase()
	fmt.Println("Database connected")
}

func initRepository() {
	repo = repository.NewRepository(db)
	fmt.Println("Repository ready")
}

func initEmailService() {
	email := email2.NewEmail(portSmtp, from, passwordSmtp, smtpHost, "Welcome ", userSmtp)
	emailService = service.EmailService{Email: email}
	fmt.Println("Email Service ready")
}

func startQueue() {
	fmt.Println("Go Rabbit ")
	conn, err := amqp.Dial("amqp://" + user + ":" + password + "@" + address + ":" + port + "/")
	if err != nil {
		fmt.Println("Error trying to connect to the RabbitMQ: ", err)
	}
	defer conn.Close()

	fmt.Println("Successfully Connected to our rabbitMQ instance ... time to send email ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error trying to get the Channel: ", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"EMAIL", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		fmt.Println("Error trying to get the QUEUE: ", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		fmt.Println("Error trying to CONSUME: ", err)
	}

	forever := make(chan bool)

	go func() {
		user := model.User{}
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			json.Unmarshal(d.Body, &user)

			log.Println("USER : ", user.String())
			err := emailService.SendEmail(user)
			if err != nil {
				repo.Save(&user, false)
			}
			repo.Save(&user, true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
