package controllers

import (
	domain_ports "go-hex-arch/internal/domain/ports"
	"go-hex-arch/internal/services"
)

type SearchController struct {
	inPort  domain_ports.InputPort
	service services.ItemService
}

func (c *SearchController) handel() {

}
