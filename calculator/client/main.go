package main

import (
	"log"

	pb "github.com/Shashwat5522/go_gRPC/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)

	// findPrime(c)
	// findAvg(c)
	// findMax(c)
	doSqrt(c,10)

}
