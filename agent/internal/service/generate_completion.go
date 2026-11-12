package agent_service

import (
	"context"
	"log"

	"github.com/qury53/GC_learn/agent/internal/domain"
)

func (s *AgentService) GenerateCompletion(
	ctx context.Context,
	request domain.ChatRequest,
) (domain.ChatResponse, error) {
	log.Printf("generate completion service")
	resp, err := s.aiProvider.GenerateCompletion(ctx, request)
	if err != nil {
		return domain.ChatResponse{}, err
	}

	return resp, nil
}
