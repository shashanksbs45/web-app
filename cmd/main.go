package main

import (
	"fmt"
	"log"

	"github.com/shashanksbs/web-app/api"
)

func main() {
	fmt.Println("hello")
	srv := api.NewServer()
	log.Fatal(srv.ListenAndServe())
}
