package agent_service

import "github.com/qury53/GC_learn/agent/pkg/ai_provider"

type AgentService struct {
	aiProvider ai_provider.AIProvider
}

func NewAgentService(provider ai_provider.AIProvider) *AgentService {
	return &AgentService{
		aiProvider: provider,
	}
}
