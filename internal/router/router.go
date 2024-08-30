package router

import (
	repositories "go-hex-arch/internal/Repositories"
	presenters "go-hex-arch/internal/adapters/Presenters"
	"go-hex-arch/internal/controllers"
	"go-hex-arch/internal/services"

	"github.com/gin-gonic/gin"
)

func ItemRouter(router *gin.Engine) {
	service := services.NewItemService(repositories.NewItemsRepository(), &presenters.ItemPresenter{})
	controller := controllers.NewItemController(
		service,
	)

	router.POST("/item/insert", controller.Insert)
	router.POST("/item/update", controller.Update)
	router.POST("/item/delete", controller.Delete)
	router.POST("/item/get", controller.Show)
	router.POST("/item/search", controller.Search)

}
