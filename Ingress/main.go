
package main

import (
	proto "github.com/joecanto/kickMeWhenImDown/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"fmt"
	"os"
	"net/http"
	"log"
)

const (
	aggregatorAddress	= "7999"
)

func dieIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", insultHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func insultHandler(w http.ResponseWriter, r *http.Request){
	// Set up a connection to the insult engine.
	conn, err := grpc.Dial(aggregatorAddress, grpc.WithInsecure())
	dieIf(err)
	defer conn.Close()

	proxyClient := proto.NewInsultClient(conn)
	request := &proto.InsultRequest{}

	resp, err := proxyClient.GoForIt(context.Background(), request)
	dieIf(err)
	fmt.Printf("Insult: %s", resp.Message)
}