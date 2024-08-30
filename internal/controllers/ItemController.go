package controllers

import (
	"fmt"
	domain_dtos "go-hex-arch/internal/domain/DTOS"
	domain_ports "go-hex-arch/internal/domain/ports"
	"net/http"
)

type ItemController struct {
	InputPort domain_ports.InputPort
}


func (obj *ItemController) Insert(dto domain_dtos.InsertItemDTO) {

	res := obj.InputPort.Add()



	fmt.Println(res)
}

func (obj *ItemController) Update(rw http.ResponseWriter, r *http.Request) {

}

func (obj *ItemController) Delete(rw http.ResponseWriter, r *http.Request)) {

}

func (obj *ItemController) Show(rw http.ResponseWriter, r *http.Request) {

}

func (obj *ItemController) Search(rw http.ResponseWriter, r *http.Request) {

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
