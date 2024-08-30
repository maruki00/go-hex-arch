package main

import (
	routers "go-hex-arch/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routers.ItemRouter(router)
	router.Run(":3000")

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	// 	controller.Insert(&domain_dtos.InsertItemDTO{
	// 		Name:  "item1",
	// 		No:    1,
	// 		Qty:   1,
	// 		Price: 1,
	// 	})
	// })

	// fmt.Println("SErver running on http://127.0.0.1:3000")
	// http.ListenAndServe(":3000", nil)

}
