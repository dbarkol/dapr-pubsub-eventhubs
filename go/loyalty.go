package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var port string
var topic string
var route string

func init() {
	port = "8080"
	topic = "songs"
	route = "loyalty"
}

func main() {
	http.HandleFunc("/dapr/subscribe", func(w http.ResponseWriter, r *http.Request) {
		j, _ := json.Marshal([]struct {
			Topic string `json:"topic"`
			Route string `json:"route"`
		}{{
			topic, route}})

		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	})

	http.HandleFunc("/"+route, func(w http.ResponseWriter, r *http.Request) {
		payload, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		fmt.Println("Payload - " + string(payload))
	})

	fmt.Println("starting HTTP server....")
	http.ListenAndServe(":"+port, nil)
}
