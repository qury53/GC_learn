package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	v1 "shared/pkg/proto/agent/v1"
	"syscall"

	"github.com/joho/godotenv"

	agent_service "github.com/qury53/GC_learn/agent/internal/service"
	agent_transport "github.com/qury53/GC_learn/agent/internal/transport"
	ai_provider_gc "github.com/qury53/GC_learn/agent/pkg/ai_provider/gigachat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 5678
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load env", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		slog.Error("failed to create lis", slog.String("error", err.Error()))
		os.Exit(1)
	}
	log.Printf("2")
	s := grpc.NewServer()
	config := ai_provider_gc.NewConfigMust()
	fmt.Println(*config)
	provider := ai_provider_gc.NewGCProvider(config)
	service := agent_service.NewAgentService(provider)
	transport := agent_transport.NewAgentTransport(service)
	log.Printf("3")
	v1.RegisterAgentServiceServer(s, transport)
	reflection.Register(s)

	go func() {
		log.Printf("starting grpc server")
		if err := s.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			slog.Error("serve error", slog.String("error", err.Error()))
			cancel()
		}
	}()

	log.Printf("4")
	<-ctx.Done()

	log.Printf("shutting down grpc server")
}
