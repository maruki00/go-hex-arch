package domain_ports

import domain_contructs "go-hex-arch/internal/domain/Contructs"

type InputPort interface {
	Add(items map[string]any) domain_contructs.ViewModel
	Delete(id int) domain_contructs.ViewModel
	Update(id int, items map[string]any) domain_contructs.ViewModel
	Search(key, query string) domain_contructs.ViewModel
	SearchById(id int) domain_contructs.ViewModel
}
