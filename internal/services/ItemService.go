package services

import (
	domain_contructs "go-hex-arch/internal/domain/Contructs"
	domain_interactions "go-hex-arch/internal/domain/Interactions"
	domain_ports "go-hex-arch/internal/domain/ports"
	"go-hex-arch/internal/models"
	"strings"
)

type ItemService struct {
	repo    domain_contructs.IItemRepository
	outPort domain_ports.OutputPort
}

func (s *ItemService) Add(items map[string]any) domain_contructs.ViewModel {

	item := ItemFromMap(items)
	res, err := s.repo.Add(item)
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

func (s *ItemService) Update(id int, items map[string]any) domain_contructs.ViewModel {

	res, err := s.repo.Udapte(id, ItemFromMap(items))
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

func ItemFromMap(items map[string]any) *models.Item {
	item := &models.Item{}
	for key, value := range items {
		switch strings.ToLower(key) {
		case "id":
			item.Id = value.(int)

		case "name":
			item.Name = value.(string)

		case "no":
			item.No = value.(int)

		case "qty":
			item.Qty = value.(int)

		case "price":
			item.Price = value.(float32)
		}
	}
	return item
}
