package controllers

import (
	"fmt"
	domain_dtos "go-hex-arch/internal/domain/DTOS"
	domain_ports "go-hex-arch/internal/domain/ports"
)

type ItemController struct {
	InputPort domain_ports.InputPort
}

func NewItemController(inPort domain_ports.InputPort) *ItemController {
	return &ItemController{
		InputPort: inPort,
	}
}

func (obj *ItemController) Insert(dto *domain_dtos.InsertItemDTO) {

	res := obj.InputPort.Add(dto)

	fmt.Println(res)
}

func (obj *ItemController) Update(dto *domain_dtos.UpdateItemDTO) {
	res := obj.InputPort.Update(dto.Id, dto)

	fmt.Println(res)
}

func (obj *ItemController) Delete(dto *domain_dtos.DeleteItemDTO) {
	res := obj.InputPort.Delete(dto.Id)

	fmt.Println(res)
}

func (obj *ItemController) Show(dto *domain_dtos.GetItemByIdDTO) {
	res := obj.InputPort.SearchById(dto.Id)

	fmt.Println(res)
}

func (obj *ItemController) Search(dto *domain_dtos.SearchItemDTO) {
	res := obj.InputPort.Search(dto.Key, dto.Query)

	fmt.Println(res)
}

// func (obj *ItemController) Insert(rw http.ResponseWriter, r *http.Request) {

// 	var
// 	r.Body

// 	res := obj.InputPort.Add()
// }

// func (obj *ItemController) Update(rw http.ResponseWriter, r *http.Request) {

// }

// func (obj *ItemController) Delete(rw http.ResponseWriter, r *http.Request)) {

// }

// func (obj *ItemController) Show(rw http.ResponseWriter, r *http.Request) {

// }

// func (obj *ItemController) Search(rw http.ResponseWriter, r *http.Request) {

// }
