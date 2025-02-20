package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/martian/log"
)

var userList = []string{"Вася", "Петя", "Дмитрий", "Тимур", "Евгений"}

func userGetHandler(w http.ResponseWriter, req *http.Request) {

	userIdStr := req.URL.Query().Get("userid")

	if userIdStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Request error not changed userid", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		log.Errorf("%#v", err)
		return
	}

	if userId < 0 || userId >= len(userList) {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "userid bigest", http.StatusBadRequest)
		return
	}
	selectedUser := userList[userId]

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(selectedUser))

}

func getAllUser(w http.ResponseWriter, req *http.Request) {

	answer := ""
	for id, userName := range userList {
		answer += fmt.Sprintf("id=%d, имя=%s \n", id, userName)
	}

	w.Write([]byte(answer))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/user", userGetHandler)
	mux.HandleFunc("/getalluser", getAllUser)
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		panic(err)
	}

}
