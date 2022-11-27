package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

// Define port
var PORT string = ":5000"

// Payload struct
type Payload struct {
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

// message to send as json response
type Message struct {
	Msg string
}

// response message struct as json format
func jsonMessageByte(msg string) []byte {
	errrMessage := Message{msg}
	byteContent, _ := json.Marshal(errrMessage)
	return byteContent
}

// Allow only POST method
// Validate Payload - It should contain command and argument
// Parse the post data and execute the command in the payload

func executeShell(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write(jsonMessageByte(r.Method + " - Method not allowed"))
	} else {
		var payloadData Payload
		err := json.NewDecoder(r.Body).Decode(&payloadData)
		if err != nil {
			w.WriteHeader(400)
			w.Write(jsonMessageByte("Bad Request - Failed to parse the payload " + err.Error()))
		} else {
			if payloadData.Command != "" && len(payloadData.Arguments) != 0 {
				out, err := exec.Command(payloadData.Command, payloadData.Arguments...).Output()
				if err != nil {
					w.WriteHeader(500)
					w.Write(jsonMessageByte("Server Error - Failed to execute the command " + err.Error()))
				} else {
					w.Write(jsonMessageByte(string(out)))
				}
			} else {
				w.Write(jsonMessageByte("Please provide command and argument"))
			}

		}

	}

}

// Main function
// Start the server
func main() {
	http.HandleFunc("/", executeShell)

	fmt.Printf("App is listening on %v\n", PORT)

	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Fatal("Failed to start the application")
	}

}
