
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

	if err := http.ListenAndServe(":7998", mux); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func insultHandler(w http.ResponseWriter, r *http.Request){
	log.Printf("Received request from ingress")

	timeouts, ok := r.URL.Query()["setTimeout"]

	if ok || len(timeouts) == 1 {
		if timeouts[0] == "long"{
			log.Printf("Sleeping for %s seconds", 6)
			time.Sleep(time.Second * 6)
		} else {
			log.Printf("Sleeping for %s seconds", 3)
			time.Sleep(time.Second * 3)
		}
	}

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
		1  :  "I don't think you are stupid. You just have a bad luck when thinking.",
		2  :  "Calling you an idiot would be an insult to all stupid people.",
		3  :  "Keep talking, someday you'll say something intelligent!",
		4  :  "I love what you've done with your hair. How do you get it to come out of the nostrils like that?",
		5  :  "Ordinarily people live and learn. You just live.",
		6  :  "Support bacteria - they're the only culture some people have.",
		7  :  "you'll never be the man your mother is.",
	}

	randNumber := random(1, 7)
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