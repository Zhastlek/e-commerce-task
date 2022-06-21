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
	inValidName  = "invalid product name"
)

func (h *handler) Register(router *gin.Engine) {
	router.GET("/products", h.GetAll)
	router.POST("/cmd/add-product", h.CreateOne)
	router.POST("/product/edit/:productId", h.GetOneById)
	router.PUT("/cmd/edit-product", h.UpdateOne)
	router.DELETE("/cmd/delete-product", h.DeleteOne)
	router.GET("/q/product-search-by-name", h.GetOneByName)
}

func (h *handler) GetOneByName(c *gin.Context) {
	var p *models.GetByName
	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	// log.Println("handler--get by name---->", p, p.SearchName)
	product, err := h.getter.GetOneByName(p.SearchName)
	if err != nil {
		log.Printf("Error get one by name method handler: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  inValidName,
		})
		return
	}
	// log.Println("get by name handler:----->", product)
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
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
		log.Printf("error in delete one method handler invalid json: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
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
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	if product.Price <= 0 {
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  inValidPrice,
		})
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
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	if product.Price <= 0 {
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  inValidPrice,
		})
		return
	}
	id, err := h.editor.CreateOne(product)
	if err != nil {
		log.Printf("error in delete one method handler: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": success,
		"id":     id,
	})
}

func (h *handler) GetOneById(c *gin.Context) {
	id := c.Param("productId")
	// log.Println("id---->", id)
	product, err := h.getter.GetOneById(id)
	if err != nil || product == nil {
		log.Printf("error in get one by id method handler: %v\n", err)
		c.AbortWithStatusJSON(400, gin.H{
			"status": failed,
			"error":  inValidId,
		})
		return
	}
	c.JSON(200, gin.H{
		"store": product,
	})
}
