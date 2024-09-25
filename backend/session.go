package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type sessionManger struct {
	id       []byte
	username string
	password string
}

type Sessions struct {
	Username string `json:"username"`
	Id       []byte `json:"id"`
}

func addSession(w http.ResponseWriter, file io.Writer, session sessionManger) error{
	encoder := json.NewEncoder(file)
	sessions := &Sessions{}
	cookie := http.Cookie{Name: session.username, Value: string(session.id)}
	http.SetCookie(w, &cookie)
	sessions.Username = session.username
	sessions.Id = session.id

	if err := encoder.Encode(sessions);err != nil{
		log.Println("couldn't add session")
		return errors.New("couldn't adss sessions to the file")

	}
	return nil
}

func newSession(username string, password string) sessionManger {
	return sessionManger{md5.New().Sum(nil), username, password}
}
