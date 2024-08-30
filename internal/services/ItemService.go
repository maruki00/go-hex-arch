package services

import (
	domain_contructs "go-hex-arch/internal/domain/Contructs"
	domain_ports "go-hex-arch/internal/domain/ports"
	"go-hex-arch/internal/models"
)

type ItemService struct {
	repo    domain_contructs.IItemRepository
	outPort domain_ports.OutputPort
}

func (s *ItemService) Add() *models.Item {

}

func (s *ItemService) Delete() *models.Item {

}

func (s *ItemService) Update() *models.Item {

}

func (s *ItemService) Search() []*models.Item {

}

func (s *ItemService) SearchById() *models.Item {

}
