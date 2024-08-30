package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

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
