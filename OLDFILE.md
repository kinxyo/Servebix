
# go old file

```go
package Servebix

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// gotta split this

var clients = make(map[string]map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)                      // broadcast channel
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Type   string          `json:"type"`
	Data   json.RawMessage `json:"data"`
	RoomID string          `json:"roomID"`
}

type Meeting struct {
	DoctorID  string `json:"doctorID"`
	PatientID string `json:"patientID"`
	SlotTime  string `json:"slotTime"`
	RoomID    string `json:"-"`
}

func oldmain() {
	router := http.NewServeMux()

	// Configure websocket route
	router.HandleFunc("/ws/{roomID}", handleConnections)
	router.HandleFunc("/schedule", scheduleMeeting)
	router.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Server Running </h1>")
	})

	http.Handle("/", router)

	// Starting & listening for incoming chat messages
	go handleMessages()

	// Starting the server
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func scheduleMeeting(w http.ResponseWriter, r *http.Request) {

	log.Println("Scheduling a meet...")

	// log.Printf("Request Recieved: %v", r.Body)

	var meeting Meeting
	err := json.NewDecoder(r.Body).Decode(&meeting)
	if err != nil {
		log.Println("Invalid Request :(")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Println("Meeting:", meeting)

	// layout := "2006-01-02T15:04:05.000"
	// startTime, err := time.Parse(layout, meeting.SlotTime)

	// startTime, err := time.Parse(time.RFC3339, meeting.SlotTime)

	layout := "2006-01-02 15:04:05 -0700 MST"
	startTime, err := time.Parse(layout, meeting.SlotTime)

	if err != nil {
		log.Println("Parsing failed :(")
		http.Error(w, "Could not parse the slot entryTime", http.StatusBadRequest)
		return
	}

	log.Println("startTime:", startTime)

	// location, err := time.LoadLocation("Asia/Kolkata")
	// if err != nil {
	// log.Fatal(err)
	// }

	// startTime = startTime.In(location)

	// log.Println("after location:", startTime)

	duration := time.Until(startTime)

	if duration < 0 {
		log.Println("The start time is in the past")
		http.Error(w, "The start time is in the past", http.StatusBadRequest)
		return
	}

	log.Println("duration", duration)

	meeting.RoomID = uuid.New().String()
	time.AfterFunc(duration, func() {
		clients[meeting.RoomID] = make(map[*websocket.Conn]bool)
		log.Println("creating 	room", meeting)
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"RoomID": meeting.RoomID,
	})
}

func handleMessages() {
	for {
		msg := <-broadcast

		for client := range clients[msg.RoomID] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error %v", err)
				client.Close()
				delete(clients[msg.RoomID], client)
			}
		}
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	roomID := r.PathValue("roomID")

	if _, ok := clients[roomID]; !ok {
		http.Error(w, "Invalid room ID", http.StatusNotFound)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Could not upgrade the connection :(", err)
		return
	}

	defer ws.Close()

	clients[roomID][ws] = true

	log.Printf("New connection from %s", r.Header.Get("User-Agent"))

	for {
		var msg Message

		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("error %v", err)
			delete(clients[roomID], ws)
			break
		}
		broadcast <- msg
	}
}
```