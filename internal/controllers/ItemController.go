package controllers

import (
	domain_dtos "go-hex-arch/internal/domain/DTOS"
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

	dto := &domain_dtos.InsertItemDTO{}
	if err := ctx.BindJSON(dto); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid data",
		})
		return
	}

	res := obj.InputPort.Add(dto)

	ctx.JSON(200, gin.H{
		"data": res.GetResponse(),
	})

}

func (obj *ItemController) Update(ctx *gin.Context) {
	dto := &domain_dtos.UpdateItemDTO{}
	if err := ctx.BindJSON(dto); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid data",
		})
		return
	}
	res := obj.InputPort.Update(dto.Id, dto)

	ctx.JSON(200, gin.H{
		"data": res.GetResponse(),
	})
}

func (obj *ItemController) Delete(ctx *gin.Context) {
	dto := &domain_dtos.DeleteItemDTO{}
	if err := ctx.BindJSON(dto); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid data",
		})
		return
	}
	res := obj.InputPort.Delete(dto.Id)

	ctx.JSON(200, gin.H{
		"data": res.GetResponse(),
	})
}

func (obj *ItemController) Show(ctx *gin.Context) {
	dto := &domain_dtos.GetItemByIdDTO{}
	if err := ctx.BindJSON(dto); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid data",
		})
		return
	}
	res := obj.InputPort.SearchById(dto.Id)

	ctx.JSON(200, gin.H{
		"data": res.GetResponse(),
	})
}

func (obj *ItemController) Search(ctx *gin.Context) {

	dto := &domain_dtos.SearchItemDTO{}
	if err := ctx.BindJSON(dto); err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid data",
		})
		return
	}

	res := obj.InputPort.Search(dto.Key, dto.Query)

	ctx.JSON(200, gin.H{
		"data": res.GetResponse(),
	})
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
