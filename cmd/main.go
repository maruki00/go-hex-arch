package main

import (
	"fmt"
	repositories "go-hex-arch/internal/Repositories"
	presenters "go-hex-arch/internal/adapters/Presenters"
	"go-hex-arch/internal/controllers"
	"go-hex-arch/internal/services"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		controllers.NewItemController(&services.NewItemService(&repositories.NewItemsRepository()), &presenters.ItemPresenter{})
	})

	fmt.Println("SErver running on http://127.0.0.1:3000")
	http.ListenAndServe(":3000", nil)

}
