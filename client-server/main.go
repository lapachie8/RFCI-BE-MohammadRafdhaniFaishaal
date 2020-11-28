package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type form struct {
	ID     int    `json:"id"`
	author string `json:"author"`
	title  string `json:"title"`
}

func client(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var clt form
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&clt); err != nil {
				log.Fatal(err)
			}
		} else {
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			author := r.PostFormValue("author")
			title := r.PostFormValue("title")
			clt = form{
				ID:     id,
				author: author,
				title:  title,
			}
		}

		dataClient, _ := json.Marshal(clt)
		w.Write(dataClient)
		return
	}

	http.Error(w, "error", http.StatusNotFound)
	return
}

func main() {
	http.HandleFunc("/post", client)
	fmt.Println("server running...")
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}
