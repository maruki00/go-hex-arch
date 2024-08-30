package presenters

type ItemPresenter struct {
	data any
}

func (obj *ItemPresenter) GetResponse() any {
	return obj.data
}
