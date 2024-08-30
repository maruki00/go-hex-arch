package repositories

import (
	"errors"
	"fmt"
	"go-hex-arch/internal/models"
)

type ItemsRepository struct {
	items map[int]*models.Item
	index int
}

func NewItemsRepository() *ItemsRepository {
	return &ItemsRepository{
		items: make(map[int]*models.Item, 0),
		index: 1,
	}
}

func RepositoryFactory() *ItemsRepository {
	return &ItemsRepository{
		items: make(map[int]*models.Item, 0),
	}
}

func (obj *ItemsRepository) Add(item *models.Item) (*models.Item, error) {
	fmt.Println("result : ", obj.items, item.Id)
	if _, ok := obj.items[item.Id]; ok {
		return nil, errors.New("item already exists")
	}

	item.Id = obj.index

	obj.items[obj.index] = item
	obj.index++
	return obj.items[item.Id], nil
}

func (obj *ItemsRepository) SearchById(id int) (*models.Item, error) {
	item, ok := obj.items[id]
	if !ok {
		return nil, errors.New("item doesnt exists")
	}
	return item, nil
}

func (obj *ItemsRepository) Delete(id int) (*models.Item, error) {
	item, ok := obj.items[id]

	if !ok {
		return nil, errors.New("item doesnt exists")
	}
	delete(obj.items, item.Id)
	return item, nil
}

func (obj *ItemsRepository) Udapte(id int, item *models.Item) (*models.Item, error) {
	_, ok := obj.items[id]
	if !ok {
		return nil, errors.New("item doest exists")
	}
	obj.items[id] = item
	return obj.items[item.Id], nil
}

func (obj *ItemsRepository) Search(key, query string) ([]*models.Item, error) {
	items := make([]*models.Item, 0)

	for _, item := range obj.items {
		val, err := getValueByKey(key, item)
		if err != nil {
			return nil, errors.New("column not found")
		}
		if val == query {
			items = append(items, item)
		}
	}
	return items, nil
}

func getValueByKey(key string, obj *models.Item) (string, error) {
	switch key {
	case "Id":
		return string(obj.Id), nil

	case "Name":
		return obj.Name, nil

	case "No":
		return string(obj.No), nil

	case "Qty":
		return string(obj.Qty), nil

	case "Price":
		return fmt.Sprintf("%f", obj.Price), nil
	default:
		return "", errors.New("key not found")

	}
}
