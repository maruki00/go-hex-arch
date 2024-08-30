package services

import (
	domain_contructs "go-hex-arch/internal/domain/Contructs"
	domain_dtos "go-hex-arch/internal/domain/DTOS"
	domain_interactions "go-hex-arch/internal/domain/Interactions"
	domain_ports "go-hex-arch/internal/domain/ports"
	"go-hex-arch/internal/models"
)

type ItemService struct {
	repo    domain_contructs.IItemRepository
	outPort domain_ports.OutputPort
}

func NewItemService(repo domain_contructs.IItemRepository, outPort domain_ports.OutputPort) *ItemService {
	return &ItemService{
		repo:    repo,
		outPort: outPort,
	}
}

func (s *ItemService) Add(dto *domain_dtos.InsertItemDTO) domain_contructs.ViewModel {

	res, err := s.repo.Add(&models.Item{
		Id:    0,
		Name:  dto.Name,
		No:    dto.No,
		Qty:   dto.Qty,
		Price: dto.Price,
	})
	if err != nil {
		return s.outPort.Error(domain_interactions.ResponseModel{
			Data: map[string]any{
				"message": "Error",
				"error":   err.Error(),
				"data":    nil,
			},
		})
	}

	return s.outPort.Error(domain_interactions.ResponseModel{
		Data: map[string]any{
			"message": "Success",
			"error":   nil,
			"data":    res,
		},
	})
}

func (s *ItemService) Delete(id int) domain_contructs.ViewModel {

	res, err := s.repo.Delete(id)
	if err != nil {
		return s.outPort.Error(domain_interactions.ResponseModel{
			Data: map[string]any{
				"message": "Error",
				"error":   err.Error(),
				"data":    nil,
			},
		})
	}

	return s.outPort.Error(domain_interactions.ResponseModel{
		Data: map[string]any{
			"message": "Success",
			"error":   nil,
			"data":    res,
		},
	})

}

func (s *ItemService) Update(id int, dto *domain_dtos.UpdateItemDTO) domain_contructs.ViewModel {

	res, err := s.repo.Udapte(id, &models.Item{
		Id:    id,
		Name:  dto.Name,
		No:    dto.No,
		Qty:   dto.Qty,
		Price: dto.Price,
	})
	if err != nil {
		return s.outPort.Error(domain_interactions.ResponseModel{
			Data: map[string]any{
				"message": "Error",
				"error":   err.Error(),
				"data":    nil,
			},
		})
	}

	return s.outPort.Error(domain_interactions.ResponseModel{
		Data: map[string]any{
			"message": "Success",
			"error":   nil,
			"data":    res,
		},
	})

}

func (s *ItemService) Search(key, query string) domain_contructs.ViewModel {

	res, err := s.repo.Search(key, query)
	if err != nil {
		return s.outPort.Error(domain_interactions.ResponseModel{
			Data: map[string]any{
				"message": "Error",
				"error":   err.Error(),
				"data":    nil,
			},
		})
	}

	return s.outPort.Error(domain_interactions.ResponseModel{
		Data: map[string]any{
			"message": "Success",
			"error":   nil,
			"data":    res,
		},
	})

}

func (s *ItemService) SearchById(id int) domain_contructs.ViewModel {

	res, err := s.repo.SearchById(id)
	if err != nil {
		return s.outPort.Error(domain_interactions.ResponseModel{
			Data: map[string]any{
				"message": "Error",
				"error":   err.Error(),
				"data":    nil,
			},
		})
	}

	return s.outPort.Error(domain_interactions.ResponseModel{
		Data: map[string]any{
			"message": "Success",
			"error":   nil,
			"data":    res,
		},
	})

}
