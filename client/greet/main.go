package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/hunter32292/weeklyProject/grpcapi/protos/greet"

	"github.com/hunter32292/randnames"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greet.NewGreetServiceClient(cc)

	doServerStreaming(c)
}

func doServerStreaming(c greet.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greet.GreetManyTimesRequest{
		Greeting: &greet.Greeting{
			FirstName: randnames.GiveName(),
			LastName:  randnames.GiveName(),
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		fmt.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}
}
