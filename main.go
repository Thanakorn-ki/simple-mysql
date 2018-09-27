package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/learing-deb/models"
)

type Env struct {
	db models.Datastore
}

func (env *Env) handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	userInfo, err := env.db.AllUserInfo()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, value := range userInfo {
		fmt.Fprintf(w, "%v\n", value)
	}
	// fmt.Printf("%s", userInfo)
}

func main() {
	db, err := models.NewDB("root:root@tcp(127.0.0.1:5555)/thanakorn-db")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db}
	http.HandleFunc("/", env.handler)
	log.Print("Server start :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
