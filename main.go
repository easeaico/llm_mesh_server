package main

import (
	"fmt"
	"log"
	"net"

	"github.com/yelcat/llm-mesh/internal/server"
	"github.com/yelcat/llm-mesh/pkg/llm_mesh"
	"google.golang.org/grpc"
)

const (
	port = 1984
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening success: %d\n", port)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	llm_mesh.RegisterChatCompletionServiceServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}