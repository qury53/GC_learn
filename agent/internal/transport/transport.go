package agent_transport

import (
	"context"

	"github.com/qury53/GC_learn/agent/internal/domain"
)

type AgentTransport struct {
	agentService AgentService
}

type AgentService interface {
	GenerateCompletion(
		ctx context.Context,
		request domain.ChatRequest,
	) (domain.AuthResponse, error)
}

func NewAgentTransport(service AgentService) *AgentTransport {
	return &AgentTransport{
		agentService: service,
	}
}
