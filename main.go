package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/hunter32292/weeklyProject/grpcapi/pkg/server"
	"google.golang.org/grpc"
)

func init() {
	// Set new rand source to time now
	rand.NewSource(time.Now().UnixNano())
}

func main() {
	log.Println("Starting...")
	Run()
	log.Println("Ending...")
}

func Run() {

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Print("Server started")
	s := grpc.NewServer()

	server.RegisterGS(s)
	server.RegisterHC(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
