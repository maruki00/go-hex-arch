package presenters

import (
	viewmodels "go-hex-arch/internal/adapters/ViewModels"
	domain_contructs "go-hex-arch/internal/domain/Contructs"
	domain_interactions "go-hex-arch/internal/domain/Interactions"
)

type ItemPresenter struct {
}

func (obj *ItemPresenter) Success(response domain_interactions.ResponseModel) domain_contructs.ViewModel {
	return &viewmodels.JsonViewModel{
		Data: response.Data,
	}
}
func (obj *ItemPresenter) Error(response domain_interactions.ResponseModel) domain_contructs.ViewModel {
	return &viewmodels.JsonViewModel{
		Data: response.Data,
	}
}
