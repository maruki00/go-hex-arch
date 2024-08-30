package domain_ports

import (
	domain_contructs "go-hex-arch/internal/domain/Contructs"
	domain_dtos "go-hex-arch/internal/domain/DTOS"
)

type InputPort interface {
	Add(dto *domain_dtos.InsertItemDTO) domain_contructs.ViewModel
	Delete(id int) domain_contructs.ViewModel
	Update(id int, dto *domain_dtos.UpdateItemDTO) domain_contructs.ViewModel
	Search(key, query string) domain_contructs.ViewModel
	SearchById(id int) domain_contructs.ViewModel
}
