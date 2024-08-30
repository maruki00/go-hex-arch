package presenters

import (
	viewmodels "go-hex-arch/internal/adapters/ViewModels"
	domain_interactions "go-hex-arch/internal/domain/Interactions"
)

type ItemPresenter struct {
}

func (obj *ItemPresenter) Success(response domain_interactions.ResponseModel) any {
	return viewmodels.JsonViewModel{
		Data: response.Data,
	}
}
func (obj *ItemPresenter) Error(response domain_interactions.ResponseModel) any {
	return viewmodels.JsonViewModel{
		Data: response.Data,
	}
}
