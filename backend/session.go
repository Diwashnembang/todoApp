package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

type sessionManger struct {
	id       string
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
	cookie := http.Cookie{Name: session.username, Value: session.id, Expires: time.Now().Add(24 * time.Hour), Quoted: false, Path: "/"}
	sessionInfo[string(session.id)] = session.username
	http.SetCookie(w, &cookie)
	jsonSessionInfo, err := json.Marshal(sessionInfo)
	if err != nil {
		log.Println("error marahslling sessionInfo")
		return errors.New("error in addsession")
	}

	offset, err := file.Seek(-1, 2)
	if err != nil {
		offset, err = file.Seek(0, 0)
		if err != nil {
			log.Println("errror seeking")
			return errors.New("error in addsession")
		}
	}

	_, err = file.WriteAt([]byte(jsonSessionInfo), offset)
	if err != nil {
		log.Println("error writing")
		log.Println(err)
		return errors.New("error in addsession")
	}
	return nil
}

func newSession(username string, password string) sessionManger {
	return sessionManger{hex.EncodeToString(md5.New().Sum([]byte(password + username))), username, password}
}
