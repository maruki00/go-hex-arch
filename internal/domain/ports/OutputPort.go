package domain_ports

import (
	domain_contructs "go-hex-arch/internal/domain/Contructs"
	domain_interactions "go-hex-arch/internal/domain/Interactions"
)

type OutputPort interface {
	Success(response domain_interactions.ResponseModel) domain_contructs.ViewModel
	Error(response domain_interactions.ResponseModel) domain_contructs.ViewModel
}
