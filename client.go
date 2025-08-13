package main

import (
	"context"
	"log"
	mianapipb "simlplegrpcclient/proto/gen"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Couldn't connect:", err)
	}

	defer conn.Close()

	client := mianapipb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &mianapipb.AddRequest{
		A: 10,
		B: 20,
	}
	
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln("Couldn't add", err)
	}	
	log.Println("Sum:", res.Sum)
}
