package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func schedule(w http.ResponseWriter, r *http.Request) {
	_, user := toPost(w, r)
	userBytes, err := json.Marshal(user)
	if err != nil{
		return
	}
	userBuf := bytes.NewBuffer(userBytes)
	StartSchedule(1, "day", 12, toRun, "post", "", userBuf)
	//toRun("post", "", userBuf)
}
