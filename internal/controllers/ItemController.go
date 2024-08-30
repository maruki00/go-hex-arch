package controllers

import (
	"fmt"
	domain_ports "go-hex-arch/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	InputPort domain_ports.InputPort
}

func NewItemController(inPort domain_ports.InputPort) *ItemController {
	return &ItemController{
		InputPort: inPort,
	}
}

func (obj *ItemController) Insert(ctx *gin.Context) {

	res := obj.InputPort.Add(dto)

	fmt.Println(res.GetResponse(), res.GetResponse())
}

func (obj *ItemController) Update(ctx *gin.Context) {
	res := obj.InputPort.Update(dto.Id, dto)

	fmt.Println(res)
}

func (obj *ItemController) Delete(ctx *gin.Context) {
	res := obj.InputPort.Delete(dto.Id)

	fmt.Println(res)
}

func (obj *ItemController) Show(ctx *gin.Context) {
	res := obj.InputPort.SearchById(dto.Id)

	fmt.Println(res)
}

func (obj *ItemController) Search(ctx *gin.Context) {
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
