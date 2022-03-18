package main

import (
	"context"
	"log"
	"time"

	"github.com/hashwavelab/ample/pb"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:8880"
	defaultName = "tester"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAmpleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	r, err := c.GetAllUniV2Dexs(ctx, &pb.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Received Reserve: %s", r)
	for _, dex := range r.Dexs {
		log.Println(dex)
	}
	// 	r, err := c.GetRawJsonConfig(ctx, &pb.KeyRequest{
	// 		Key: "obexTradingPairs",
	// 	})
	// 	if err != nil {
	// 		log.Fatalf("could not greet: %v", err)
	// 	}
	// 	log.Printf("Received Reserve: %s", r)
}
