package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/barbosaigor/april/auth"
)

func someRoute(w http.ResponseWriter, r *http.Request) {
}

func main() {
	ath := auth.New("my secret")
	ath.Register("bob", "123")
	serveMux := http.NewServeMux()
	serveMux.Handle("/signin", ath.MwGenerateToken(http.HandlerFunc(someRoute)))
	serveMux.Handle("/", ath.MwAuth(http.HandlerFunc(someRoute)))
	fmt.Println("(HTTP) Listening")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 8081), serveMux))
}