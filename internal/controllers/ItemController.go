package controllers

import domain_ports "go-hex-arch/internal/domain/ports"

type ItemController struct {
	InputPort domain_ports.InputPort
}

func (obj *ItemController) Insert() {
	res := obj.InputPort.
}

func (obj *ItemController) Update() {

}

func (obj *ItemController) Delete() {

}

func (obj *ItemController) Show() {

}

func (obj *ItemController) Search() {

}
