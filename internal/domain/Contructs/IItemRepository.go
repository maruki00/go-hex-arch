package domain_contructs

import "go-hex-arch/internal/models"

type IItemRepository interface {
	Add(item *models.Item) (*models.Item, error)
	SearchById(id int) (*models.Item, error)
	Delete(id int) (*models.Item, error)
	Udapte(id int, item *models.Item) (*models.Item, error)
	Search(key, query string) ([]*models.Item, error)
}
