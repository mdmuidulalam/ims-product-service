package routes

import (
	"github.com/gin-gonic/gin"

	routesinterface "product-service/routes/interfaces"
)

type Products struct {
	R                  *gin.Engine
	CreateProductLogic routesinterface.ICreateProductLogic
}

func (prod *Products) New() {
	prod.R.POST("create", prod.Create)
}

// @Summary Create a product
// @Description It will create a product from provided information
// @ID create-product
// @Router /create [post]
// @Accept json
// @Param displayId body string true "DisplayId of a product"
// @Param name body string true "Name of a product"
// @Param description body string true "Description of a product"
// @Success 250 {boolean} boolean "A product is created"
// @Success 251 {boolean} boolean "DisplayId of this product exists"
// @Success 252 {boolean} boolean "Name of this product exists"
func (prod *Products) Create(c *gin.Context) {
	var productInfo productInformation
	if err := c.ShouldBind(&productInfo); err != nil {
		panic(err)
	}

	prod.CreateProductLogic.SetDisplayId(productInfo.DisplayId)
	prod.CreateProductLogic.SetName(productInfo.Name)
	prod.CreateProductLogic.SetDescription(productInfo.Description)

	c.Writer.WriteHeader(250 + prod.CreateProductLogic.CreateProduct())
}

//* Data Classes

//* Dealing product information
type productInformation struct {
	DisplayId   string `json:"displayId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
