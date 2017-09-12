

// forms.go
package main

import (
	"math/rand"
	"time"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/joecanto/kickMeWhenImDown/proto"
)

const (
	port = ":8000"
)

type server struct{}

func (s *server) GoForIt(ctx context.Context, in *pb.InsultRequest) (*pb.InsultResponse, error) {
	return &pb.InsultResponse{return_insult()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInsultServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


func return_insult() string {
	insult_map := map[int]string{
		1  :  "I don't think you are stupid. You just have a bad luck when thinking.",
		2  :  "Calling you an idiot would be an insult to all stupid people.",
		3  :  "What language are you speaking? Cause it sounds like bullshit.",
		4  :  "Aha, I see the Fuck-Up Fairy has visited us again!",
		5  :  "Whats the difference between your girlfriend and a walrus? One has a moustache and smells of fish and the other is a walrus.",
		6  :  "Keep talking, someday you'll say something intelligent!",
		7  :  "Karma takes too long, I'd rather beat the shit out of you just now.",
		8  :  "I love what you've done with your hair. How do you get it to come out of the nostrils like that?",
		9  :  "Ordinarily people live and learn. You just live.",
		10 :  "Support bacteria - they're the only culture some people have.",
		11 :  "I'd like to see things from your point of view, but I can't seem to get my head that far up your ass.",
		12 :  "Shut up, you'll never be the man your mother is.",
	}

	randNumber := random(1, 12)
	// set the insult to random
	c := insult_map[randNumber]

	return c

}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}