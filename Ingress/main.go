
package main

import (
	"math/rand"
	"net/http"
	"time"
	"os"
	"log"
	"fmt"
	"encoding/json"
)

type Insult struct {
	Content string
	CreatedAt time.Time
}

func main() {

	mux := http.NewServeMux();
	mux.HandleFunc("/", insultHandler)

	if err := http.ListenAndServe(":7999", mux); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func insultHandler(w http.ResponseWriter, r *http.Request){
	log.Printf("Received request from ingress")

	insult := Insult{}

	insult.CreatedAt = time.Now().Local()
	insult.Content = return_insult()

	insultJson, err := json.Marshal(insult)
	dieIf(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(insultJson)
}


func return_insult() string {
	insult_map := map[int]string{
		1  :  "What language are you speaking? Cause it sounds like bullshit.",
		2  :  "Aha, I see the Fuck-Up Fairy has visited us again!",
		3  :  "Karma takes too long, I'd rather beat the shit out of you just now.",
		4  :  "I'd like to see things from your point of view, but I can't seem to get my head that far up your ass.",
	}

	randNumber := random(1, 4)
	// set the insult to random
	c := insult_map[randNumber]

	return c

}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func dieIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}