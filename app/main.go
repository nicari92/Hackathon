package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
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
	vrfyClient     pb.VrfyServiceClient = nil
	pubSubClient   *pubsub.Client       = nil
	pubSubTopic    *pubsub.Topic        = nil
	pubSubEncoding pubsub.SchemaEncoding
)

func (s *server) SyntaxVerification(ctx context.Context, in *pb.VerificationRequest) (*pb.VerificationResponse, error) {
	response, err := InternalSyntaxVerification(in.Email)
	go PubSubPublish(ctx, &pb.VerificationOutput{Email: in.Email, Valid: response.Valid, ErrorMessage: response.ErrorMessage})
	return response, err
}

func InternalSyntaxVerification(email string) (*pb.VerificationResponse, error) {
	if !matcher.MatchString(email) {
		return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
	}
	return &pb.VerificationResponse{Valid: true}, nil
}

func (s *server) SimpleVerification(ctx context.Context, in *pb.VerificationRequest) (*pb.VerificationResponse, error) {
	response, err := InternalSimpleVerification(in.Email)
	go PubSubPublish(ctx, &pb.VerificationOutput{Email: in.Email, Valid: response.Valid, ErrorMessage: response.ErrorMessage})
	return response, err
}

func InternalSimpleVerification(email string) (*pb.VerificationResponse, error) {
	response, syntaxErr := InternalSyntaxVerification(email)
	if response.Valid {
		_, err := net.LookupIP(strings.Split(email, "@")[1])
		if err != nil {
			return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
		}
		return &pb.VerificationResponse{Valid: true}, nil
	}
	return response, syntaxErr
}

func (s *server) FullVerification(ctx context.Context, in *pb.VerificationRequest) (*pb.VerificationResponse, error) {
	response, err := InternalFullVerification(in.Email)
	go PubSubPublish(ctx, &pb.VerificationOutput{Email: in.Email, Valid: response.Valid, ErrorMessage: response.ErrorMessage})
	return response, err
}

func InternalFullVerification(email string) (*pb.VerificationResponse, error) {
	response, syntaxErr := InternalSimpleVerification(email)
	if response.Valid {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		vrfyResponse, err := vrfyClient.Verify(ctx, &pb.VrfyRequest{Email: email})
		if err != nil {
			return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
		}
		if vrfyResponse.StatusCode != 0 {
			return &pb.VerificationResponse{Valid: false, ErrorMessage: &errorMessage}, nil
		}
		return &pb.VerificationResponse{Valid: true}, nil
	} else {
		return response, syntaxErr
	}
}

func PubSubPublish(ctx context.Context, in *pb.VerificationOutput) {
	var msg []byte
	switch pubSubEncoding {
	case pubsub.EncodingBinary:
		msg, _ = proto.Marshal(in)
	default:
		msg, _ = protojson.Marshal(in)
	}
	pubSubTopic.Publish(ctx, &pubsub.Message{
		Data: msg,
	})
}

func main() {
	log.Printf("Starting server")
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("failed to parse SERVER_PORT: %v", err)
	}

	vrfyHost := os.Getenv("VRFY_ENDPOINT")
	vrfyPort, err := strconv.Atoi(os.Getenv("VRFY_PORT"))
	if err != nil {
		log.Fatalf("failed to parse VRFY_PORT: %v", err)
	}

	pubSubTopicName := os.Getenv("PUBSUB_TOPIC")
	pubSubProject := os.Getenv("PUBSUB_PROJECT")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMailVerifierServer(s, &server{})
	reflection.Register(s)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", vrfyHost, vrfyPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	vrfyClient = pb.NewVrfyServiceClient(conn)

	pubSubClient, err = pubsub.NewClient(context.Background(), pubSubProject)
	pubSubTopic = pubSubClient.Topic(pubSubTopicName)
	cfg, err := pubSubTopic.Config(context.Background())
	if cfg.SchemaSettings != nil {
		pubSubEncoding = cfg.SchemaSettings.Encoding
	}

	log.Println("Server listening at %v", lis.Addr())

	errorSrv := s.Serve(lis)
	if errorSrv != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
