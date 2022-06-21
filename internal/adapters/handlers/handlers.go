package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"e-comerce/internal/models"
	"e-comerce/internal/service"
)

type Register interface {
	Register(router *gin.Engine)
}

type handler struct {
	editor service.ServiceEditor
	getter service.ServiceGetter
}

func NewHandler(editor service.ServiceEditor, getter service.ServiceGetter) Register {
	return &handler{
		editor: editor,
		getter: getter,
	}
}

var (
	failed       = "failure"
	success      = "success"
	inValidPrice = "Price must be positive number"
	inValidId    = "invalid product id"
)

func (h *handler) Register(router *gin.Engine) {
	router.GET("/products", h.GetAll)

	router.GET("/product/add", h.ParseCreatePage)
	router.POST("/cmd/add-product", h.CreateOne)

	router.POST("/product/edit/:productId", h.ParseUpdatePage)
	router.PUT("/cmd/edit-product", h.UpdateOne)

	router.DELETE("/cmd/delete-product", h.DeleteOne)

	router.GET("/q/product-search-by-name", h.GetOneByName)
}

func (h *handler) GetOneByName(c *gin.Context) {
	var p *models.GetByName
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		if err := c.BindJSON(&p); err != nil {
			return
		}
		product, err := h.getter.GetOneByName(p.SearchName)
		if err != nil {
			log.Printf("Error get one by name method handler: %v\n", err)
			return
		}
		// var answer models.GiveProduct
		c.IndentedJSON(http.StatusOK, gin.H{
			"product": product,
		})
	}
}

func (h *handler) GetAll(c *gin.Context) {
	products, err := h.getter.GetAll()
	if err != nil {
		log.Printf("error in get all method handler: %v\n", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"store": products,
	})
}

func (h *handler) DeleteOne(c *gin.Context) {
	var product *models.Product
	if err := c.BindJSON(&product); err != nil {
		return
	}
	if err := h.editor.DeleteOne(product.Id); err != nil {
		log.Printf("error in delete one method handler: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  inValidId,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": success,
	})
}

func (h *handler) UpdateOne(c *gin.Context) {
	var product *models.Product
	if err := c.BindJSON(&product); err != nil {
		return
	}
	if err := h.editor.UpdateOne(product); err != nil {
		log.Printf("error in delete one method handler: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  inValidId,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": success,
	})
}

func (h *handler) CreateOne(c *gin.Context) {
	var product *models.Product
	if err := c.BindJSON(&product); err != nil {
		return
	}
	id, err := h.editor.CreateOne(product)
	if err != nil {
		log.Printf("error in delete one method handler: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": success,
		"id":     id,
	})
}

func (h *handler) ParseCreatePage(c *gin.Context) {

}

func (h *handler) ParseUpdatePage(c *gin.Context) {

}
