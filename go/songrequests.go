package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var port string
var topic string
var route string

func init() {
	port = "8080"
	topic = "songs"
	route = "playlist"
}

func main() {

	// Return a collection of topic subscriptions
	http.HandleFunc("/dapr/subscribe", func(w http.ResponseWriter, r *http.Request) {
		json, _ := json.Marshal([]struct {
			Topic string `json:"topic"`
			Route string `json:"route"`
		}{{
			topic, route}})

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})

	http.HandleFunc("/"+route, func(w http.ResponseWriter, r *http.Request) {

		type SongRequest struct {
			ID     int    `json:"id"`
			Artist string `json:"artist"`
			Name   string `json:"name"`
		}

		type CloudEvent struct {
			ID      string      `json:"id"`
			Subject string      `json:"subject"`
			Source  string      `json:"source"`
			Type    string      `json:"type"`
			Data    SongRequest `json:"data"`
		}

		event := &CloudEvent{}
		var requestBody []byte
		requestBody, _ = ioutil.ReadAll((r.Body))

		err := json.Unmarshal(requestBody, &event)
		if err != nil {
			fmt.Println("Could not decode message")
		}
		log.Printf("New song request: %s - %s", event.Data.Artist, event.Data.Name)
	})

	fmt.Println("starting HTTP server....")
	http.ListenAndServe(":"+port, nil)
}
