
package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
)

const (
	aggregatorAddress	= ":7999"
)

func dieIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {
	mux := http.NewServeMux();
	mux.HandleFunc("/", insultHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func insultHandler(w http.ResponseWriter, r *http.Request){
	log.Printf("Received request from load balancer")



}