package main

import (
	"context"
	"fmt"
	_ "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	pb "hackathon/protos"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type server struct {
	pb.UnimplementedMailVerifierServer
}

var matcher, _ = regexp.Compile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])")
var errorMessage = "Email not valid"

var (
	vrfyClient pb.VrfyServiceClient = nil
)

func (s *server) SyntaxVerification(_ context.Context, in *pb.VerificationRequest) (*pb.VerificationResponse, error) {

	if !matcher.MatchString(in.Email) {
		return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
	}
	return &pb.VerificationResponse{Valid: true}, nil
}

func (s *server) SimpleVerification(_ context.Context, in *pb.VerificationRequest) (*pb.VerificationResponse, error) {
	if matcher.MatchString(in.GetEmail()) {
		_, err := net.LookupIP(strings.Split(in.Email, "@")[1])
		if err != nil {
			return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
		}
		return &pb.VerificationResponse{Valid: true}, nil
	}
	return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
}

func (s *server) FullVerification(ctx context.Context, in *pb.VerificationRequest) (*pb.VerificationResponse, error) {
	r, _ := s.SimpleVerification(ctx, in)
	if !r.Valid {
		return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
	defer cancel()
	vrfyResponse, err := vrfyClient.Verify(ctx, &pb.VrfyRequest{Email: in.Email})
	if err != nil {
		return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
	}
	if vrfyResponse.StatusCode != 0 {
		return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
	}
	return &pb.VerificationResponse{Valid: true, ErrorMessage: &errorMessage}, nil
}

func main() {
	log.Printf("Starting server")
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("failed to parse SERVER_PORT: %v", err)
	}

	vrfyAddr := os.Getenv("VRFY_ENDPOINT")
	vrfyPort, err := strconv.Atoi(os.Getenv("VRFY_PORT"))
	if err != nil {
		log.Fatalf("failed to parse VRFY_PORT: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMailVerifierServer(s, &server{})
	reflection.Register(s)

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", vrfyAddr, vrfyPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	vrfyClient = pb.NewVrfyServiceClient(conn)
	log.Printf("Server listening at %v", lis.Addr())

	errorSrv := s.Serve(lis)
	if errorSrv != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
