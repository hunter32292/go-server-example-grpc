package server

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hunter32292/weeklyProject/grpcapi/protos/greet"
	"google.golang.org/grpc"
)

type Server struct{}

func (*Server) GreetManyTimes(req *greet.GreetManyTimesRequest, stream greet.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greet.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		log.Printf("Sent: %v", res)

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func RegisterGS(s *grpc.Server) {
	greet.RegisterGreetServiceServer(s, &Server{})
}
