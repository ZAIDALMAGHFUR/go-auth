package response

import "github.com/zaidalmaghfur/go-app/internal/master/agama/domain"

type AgamaResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FromDomain(agama *domain.Agama) AgamaResponse {
	return AgamaResponse{
		ID:   agama.ID,
		Name: agama.Name,
	}
}

func FromDomainList(agamaList []domain.Agama) []AgamaResponse {
	var responses []AgamaResponse
	for _, agama := range agamaList {
		responses = append(responses, FromDomain(&agama))
	}
	return responses
}
