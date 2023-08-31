package main

import (
	"log"
	"time"

	pb "github.com/Shashwat5522/go_gRPC/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal()
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	// doGreet(c)
	// doSum(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// bidireactionalGreet(c)
	doGreetWithDeadline(c, 10*time.Second)
}
