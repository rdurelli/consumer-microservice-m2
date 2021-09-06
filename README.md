





<p align="center">
  

<h3 align="center">Microservice M1</h3>

The second micro-service (M2) was implemented using Go and some nice libs such as: (i) cobra to create a command line interface, (ii) Viper to read .env files, (iii) GIN to provide a web-service, and (iv) Gocron to trigger a task periodically. 

M2 consumes the EMAIL queue as a message is ready. Then, the message is unmarshalled and a welcome email is sent to the added user. Then, information weather the email was correctly send or not is persisted into the database. Using Gocron a task periodically check if there is any email that wasn’t sent - if so, we try to send the email again.

</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>




<!-- ABOUT THE PROJECT -->
## About The Project
![alt text](https://github.com/rdurelli/save-user-microservice-m1/blob/main/image/arch.png?raw=true)
Nowadays there are a lot of micro-services out there. One way to accomplish the communication between them is by using message-broker, such as RabbitMQ Kafka.

I have created two isolate micro-services as can be seen in the figure. The first micro-service (M2) was devised using Java Spring Boot.

M1 is just responsible to validade and persist a user into the database. Then, a message is sent to the RabbitMQ (message-broker). All messages are put into a queue (FIFO).

The second micro-service (M2) was implemented using Go and some nice libs such as: (i) cobra to create a command line interface, (ii) Viper to read .env files, (iii) GIN to provide a web-service, and (iv) Gocron to trigger a task periodically.

M2 consumes the EMAIL queue as a message is ready. Then, the message is unmarshalled and a welcome email is sent to the added user. Then, information weather the email was correctly send or not is persisted into the database. Using Gocron a task periodically check if there is any email that wasn’t sent - if so, we try to send the email again.

#golang #Java #SpringBoot #go #docker #rabbitmq

Important links:

M1 GitHub: [link](https://github.com/rdurelli/save-user-microservice-m1.git)

M2 GitHub: [link](https://github.com/rdurelli/consumer-microservice-m2.git)


### Built With

* [Go](https://golang.org/)
* [Cobra CLI]()
* [Viper]()
* [Migration]()
* [GIN]()
* [HTML Template]()
* [GoCron]()




<!-- GETTING STARTED -->
## Getting Started

Firstly install Go and Docker in your machine then install the following:

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* docker compose
  ```sh
  docker-compose up 
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/rdurelli/save-user-microservice-m1.git
   ```
2. Install go mod
  ```sh
  go get -u github.com/spf13/cobra
  ```
  ```sh
  go get -u github.com/gin-gonic/gin
  ```
  ```sh
  go get github.com/spf13/viper
  ```
  ```sh
  go get -u github.com/go-sql-driver/mysql
  ```
  ```sh
  go get -u github.com/jasonlvhit/gocron
  ```




<!-- USAGE EXAMPLES -->
## Usage

Install all dependencies listed before

  ```sh
  go run main.go
  ```
![alt text](https://github.com/rdurelli/consumer-microservice-m2/blob/main/image/terminal.png?raw=true)

As can be seen we have available four commands: (i) completion, (ii) help, (iii) queue, and (iv) server. The first one is generated automatically by COBRA CLI. Command help show how to use the app. Command queue and server are the important ones. If I run:
```sh
  go run main.go queue --help
  ```
![alt text](https://github.com/rdurelli/consumer-microservice-m2/blob/main/image/terminal2.png?raw=true)

As can be noted we can specify some flags like address to connect the AMQP (RabbitMQ), SMTP Port etc. If none information is provided we have default values.
```sh
  go run main.go queue 
  ```
Calling the command queue will start the RabbitMQ
![alt text](https://github.com/rdurelli/consumer-microservice-m2/blob/main/image/terminal3.png?raw=true)

The next command is the server command. By running you can define the port to create the web-service, default is 8080
```sh
  go run main.go server --help 
  ```
![alt text](https://github.com/rdurelli/consumer-microservice-m2/blob/main/image/terminal4.png?raw=true)

To create the web-service just call the following command:

```sh
  go run main.go server 
  ```

![alt text](https://github.com/rdurelli/consumer-microservice-m2/blob/main/image/terminal5.png?raw=true)

Two end-point is created as depicted in figure: (i) GET     /api/emails and (ii) GET    /api/emails/showAll

The first end-point show all emails that have been failed to be sent. The second end-point list all emails.


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Rafael S. Durelli - [@rafaeldurelli](https://twitter.com/rafaeldurelli) - rafael.durelli@ufla.br

Project Link: [https://github.com/rdurelli/save-user-microservice-m1.git](https://github.com/rdurelli/save-user-microservice-m1.git)
