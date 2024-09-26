package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type sessionManger struct {
	id       []byte
	username string
	password string
}

// type Sessions struct {
// 	Username string `json:"username"`
// 	Id       []byte `json:"id"`
// }

func addSession(w http.ResponseWriter, file *os.File, session sessionManger) error {
	defer file.Close()
	sessionInfo := make(map[string]string)
	cookie := http.Cookie{Name: session.username, Value: string(session.id)}
	sessionInfo[string(session.id)] = session.username
	http.SetCookie(w, &cookie)
	jsonSessionInfo, err := json.Marshal(sessionInfo)
	if err != nil {
		log.Println("error marahslling sessionInfo")
		return errors.New("error in addsession")
	}

	offset, err := file.Seek(-1, 2)
	if err != nil {
		log.Println("errror seeking")
		return errors.New("error in addsession")
	}
	_, err = file.WriteAt([]byte(jsonSessionInfo), offset)
	if err != nil {
		log.Println("error writing")
		return errors.New("error in addsession")
	}
	return nil
}

func newSession(username string, password string) sessionManger {
	return sessionManger{md5.New().Sum(nil), username, password}
}
