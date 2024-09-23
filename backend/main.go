package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var todolists []tasks

type tasks struct {
	id   int
	task string
}

func main() {
	mux := http.NewServeMux()
	// Create an http.HandlerFunc from handlePost
	handler := http.HandlerFunc(handlePost)

	// Wrap it with the CORS middleware
	corsHandler := corsMiddleware(handler)
	mux.HandleFunc("/", corsHandler)
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal("cannot listen on port 8000")
	}
}

func corsMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Method)
	switch r.Method {
	case "POST":
		log.Println("innn")

		var todo map[string]string
		var task tasks
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(500)
			log.Println("couldn't  read request body", err)
		}
		err = json.Unmarshal([]byte(body), &todo)
		if err != nil {
			w.WriteHeader(500)
			log.Println("couldn't  unmarshal respone json", err)

		}
		task.id = 1
		if value, ok := todo["task"]; ok {

			task.task = value
			todolists = append(todolists, task)
			log.Println(todolists)
			w.WriteHeader(200)
		} else {
			w.WriteHeader(400)
			w.Write([]byte("send valid key"))
			log.Println("send valid key. valid key is task")

		}
	}
}
