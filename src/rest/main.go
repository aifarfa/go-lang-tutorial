package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// routes
	http.HandleFunc("/", homeHandler)
	// start server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(wr http.ResponseWriter, r *http.Request) {
	data := getData()
	bytes, err := json.Marshal(data)
	if err == nil {
		fmt.Fprintf(wr, "%s", bytes)
	} else {
		fmt.Fprintf(wr, "%s", err)
	}
}

func getData() []map[string]string {
	data := []map[string]string{
		{
			"firstname": "John",
			"lastname":  "Doe",
		},
	}
	return data
}
