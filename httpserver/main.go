package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	Sender    string    `json:"sender"`
	Recipient string    `json:"recipient"`
	Message   string    `json:"message"`
	Time      time.Time `json:"time"`
}

var db *sql.DB

func main() {
	var err error
	db, err := sql.Open("mysql", "your-username:your-password@tcp(database:3306)/InstantMessagingDB")
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}
	defer db.Close()

	http.HandleFunc("/messages", handleMessages)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		messages, err := fetchMessagesFromDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(messages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func fetchMessagesFromDB() ([]Message, error) {
	rows, err := db.Query("SELECT sender, recipient, message, time FROM InstantMessagingDB")
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
