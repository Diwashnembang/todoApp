package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var todolists []tasks

type doc []byte

var files []doc

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
	corsHandlerUpload := corsMiddleware(http.HandlerFunc(handleUpload))
	mux.HandleFunc("/", corsHandler)
	mux.HandleFunc("/addfile", corsHandlerUpload)
	mux.HandleFunc("/login", corsMiddleware(handleLogin))
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal("cannot listen on port 8000")
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
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

func handleUpload(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			w.WriteHeader(500)
			log.Println("couldn't parse data")
			return
		}

		// var data map[string][]byte
		file, _, err := r.FormFile("files")
		if err != nil {
			w.WriteHeader(400)
			log.Println("send file with valid key")
			return
		}
		buffer := make([]byte, 1024)
		for {

			data, err := file.Read(buffer)
			if err == io.EOF {
				log.Println("end of the file")
				break
			}
			if err != nil {
				w.WriteHeader(500)
				log.Println("couln't read file")
				return
			}
			fmt.Printf("Read %d bytes: %s\n", data, buffer[:data])
		}
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		//start by creating file
		body := make(map[string]string)
		err := r.ParseForm()
		// if err != nil {
		// 	log.Println("error reading body data ")
		// }

		// err = json.Unmarshal(bodyData, &body)
		// if err != nil {
		// 	log.Println("error unmarshalling data")
		// 	log.Println("error unmarshalling data", err)
		// }
		body["username"] = r.FormValue("username")
		file, err := os.OpenFile("sessionsinfo", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			log.Println("error with file handling")
		}
		session := newSession(body["username"], body["password"])
		err = addSession(w, file, session)
		if err != nil {
			log.Println("error adding session")
			w.WriteHeader(500)
		}
	}
}
func handlePost(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Method)
	switch r.Method {
	case "POST":

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
