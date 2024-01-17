package main

import (
	"context"
	"fmt"
	_ "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hackathon/protos"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
)

type server struct {
	pb.UnimplementedMailVerifierServer
}

var matcher, _ = regexp.Compile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])")

func (s *server) SyntaxVerification(ctx context.Context, in *pb.VerificationRequest) (*pb.VerificationResponse, error) {
	log.Printf("Received: %v", in.GetEmail())
	errorMessage := "Email not valid"

	if !matcher.MatchString(in.GetEmail()) {
		return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
	}
	return &pb.VerificationResponse{Valid: true}, nil

}

func main() {
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("failed to parse port: %v", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	// Register Greeter on the server.
	pb.RegisterMailVerifierServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
