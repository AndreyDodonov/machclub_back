package handler

import (
	"github.com/AndreyDodonov/machclub_back/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

// new handler ...
func NewHandler(services *service.Service) *Handler  {
	return &Handler{services: services}
}

// init routes ...
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/news")
		{
			lists.POST("/", h.userIdentity, h.createNews)
			lists.GET("/", h.getAllNews)
			lists.GET("/:id", h.getNewsById)
			lists.PUT("/:id",h.userIdentity, h.updateNews)
			lists.DELETE("/:id",h.userIdentity, h.deleteNews)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
