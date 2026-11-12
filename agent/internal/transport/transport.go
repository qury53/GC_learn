package agent_transport

import (
	"context"
	agentV1 "shared/pkg/proto/agent/v1"

	"github.com/qury53/GC_learn/agent/internal/domain"
)

type AgentTransport struct {
	agentV1.UnimplementedAgentServiceServer
	agentService AgentService
}

type AgentService interface {
	GenerateCompletion(
		ctx context.Context,
		request domain.ChatRequest,
	) (domain.ChatResponse, error)
}

func NewAgentTransport(service AgentService) *AgentTransport {
	return &AgentTransport{
		agentService: service,
	}
}
