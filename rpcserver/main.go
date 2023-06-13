package main

import (
	"database/sql"
	"log"
	"net"
	"net/rpc"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	Sender    string    `json:"sender"`
	Recipient string    `json:"recipient"`
	Message   string    `json:"message"`
	Time      time.Time `json:"time"`
}

type MessageService struct {
	db *sql.DB
}

func (s *MessageService) PullMessages(args struct{}, reply *[]Message) error {
	messages, err := s.fetchMessagesFromDB()
	if err != nil {
		return err
	}

	*reply = messages
	return nil
}

func (s *MessageService) fetchMessagesFromDB() ([]Message, error) {
	// Implement your logic to fetch messages from the database and return them as a slice of Message structs.
	// Make the appropriate SELECT query to retrieve the desired data from the InstantMessagingDB table.

	// Example query:
	rows, err := s.db.Query("SELECT sender, recipient, message, time FROM InstantMessagingDB")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []Message{}
	for rows.Next() {
		var sender, recipient, message string
		var time time.Time
		err := rows.Scan(&sender, &recipient, &message, &time)
		if err != nil {
			return nil, err
		}
		messages = append(messages, Message{
			Sender:    sender,
			Recipient: recipient,
			Message:   message,
			Time:      time,
		})
	}

	return messages, nil
}

func main() {
	var err error
	db, err := sql.Open("mysql", "your-username:your-password@tcp(database:3306)/InstantMessagingDB")
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}
	defer db.Close()

	messageService := &MessageService{db: db}
	rpc.Register(messageService)

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}
