package main

import (
	"context"
	"log"
	mainapipb "simlplegrpcclient/proto/gen"
	farewellpb "simlplegrpcclient/proto/gen/farewell"

	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	cert := "cert.pem"

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalln("Failed to load certificates: ", err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal("Couldn't connect:", err)
	}

	defer conn.Close()

	client := mainapipb.NewCalculatorClient(conn)
	client2 := mainapipb.NewGreeterClient(conn)
	client3 := farewellpb.NewAuefWiedersehenClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &mainapipb.AddRequest{
		A: 10,
		B: 20,
	}

	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln("Couldn't add", err)
	}

	reqGreeter := &mainapipb.HelloRequest{
		Name: "Estifanos",
	}

	res1, err := client2.Greet(ctx, reqGreeter)
	if err != nil {
		log.Fatalln("Couldn't greet: ", err)
	}

	reqFarewell := &farewellpb.GoodByeRequest{
		Name: "Estifanos",
	}

	res2, err := client3.BidGoodBye(ctx, reqFarewell)
	if err != nil {
		log.Println("Couldn't greet", err)
	}


	// reqGoodBye := &mainapipb.B



	log.Println("Sum:", res.Sum)
	log.Println("Hello: ", res1)
	log.Println("Goodbye: ", res2)
	log.Println("Printing somthign from the client")
	// Getting the current state of request like ping it returns READY
	state := conn.GetState()
	log.Println("Connection state: ", state)

}
