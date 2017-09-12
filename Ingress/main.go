
package main

import (
	proto "github.com/joecanto/kickMeWhenImDown/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"fmt"
	"os"
)

const (
	address	= "7999"
)

func dieIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	dieIf(err)
	defer conn.Close()

	proxyClient := proto.NewInsultClient(conn)
	request := &proto.InsultRequest{}

	resp, err := proxyClient.GoForIt(context.Background(), request)
	dieIf(err)
	fmt.Printf("Insult: %s", resp.Message)

}