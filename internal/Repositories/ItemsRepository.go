package repositories

import "errors"

type Item struct {
	Id    int
	Name  string
	No    int
	Qty   int
	Price float32
}
type ItemsRepository struct {
	items map[int]Item
}

func RepositoryFactory() *ItemsRepository {
	return &ItemsRepository{
		items: make(map[int]Item, 0),
	}
}

func (obj *ItemsRepository) Add(item Item) error {
	if _, ok := obj.items[item.Id]; ok {
		return errors.New("item already exists")
	}
	obj.items[item.Id] = item
	return nil
}

func (obj *ItemsRepository) SearchById(id int) (*Item, error) {
	item, ok := obj.items[id]
	if !ok {
		return nil, errors.New("item doesnt exists")
	}
	return &item, nil
}

func (obj *ItemsRepository) Delete(id int) error {
	item, ok := obj.items[id]

	if !ok {
		return errors.New("item doesnt exists")
	}
	delete(obj.items, item.Id)
	return nil
}

func (obj *ItemsRepository) Udapte(item Item) error {
	item, ok := obj.items[item.Id]
	if !ok {
		return errors.New("item doest exists")
	}
	obj.items[item.Id] = item
	return nil
}
