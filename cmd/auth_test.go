package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/barbosaigor/april/auth"
)

func someRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey there"))
}

func main() {
	ath := auth.New("my secret")
	ath.Register("bob", "123")
	fmt.Println("Token: ", auth.EncryptUser("bob", "123"))
	serveMux := http.NewServeMux()
	serveMux.Handle("/", ath.MwAuth(http.HandlerFunc(someRoute)))
	fmt.Println("(HTTP) Listening on port 8081")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 8081), serveMux))
}
