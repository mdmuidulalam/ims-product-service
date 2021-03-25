package routes

import (
	"github.com/gin-gonic/gin"

	"product-service/data"
	"product-service/logics"
)

func ProductsRoutes(router *gin.Engine) {
	router.POST("/create", create)
}

// @Summary Create a product
// @Description It will create a product from provided information
// @ID create-product
// @Router /create [post]
// @Accept json
// @Param displayId body string true "DisplayId of a product"
// @Param name body string true "Name of a product"
// @Param description body string true "Description of a product"
// @Success 250 {object} object{errorTypes=[]string} "A product is created. ErrorTypes array will be empty."
// @success 251 {object} object{errorTypes=[]string} "All ready a product with displayId or name. ErrorType{1 => displayId, 2 => name}"
func create(c *gin.Context) {
	var productInfo productInformation
	if err := c.ShouldBind(&productInfo); err != nil {
		panic(err)
	}

	createProductLogic := &logics.CreateProductLogic{
		ProductData: &data.ProductsData{
			PostgresData: &data.PostgresData{},
		},
	}

	createProductLogic.SetDisplayId(productInfo.DisplayId)
	createProductLogic.SetName(productInfo.Name)
	createProductLogic.SetDescription(productInfo.Description)

	statusCode, errorTypes := createProductLogic.CreateProduct()

	c.JSON(250+statusCode, gin.H{
		"errorTypes": errorTypes,
	})
}

//* Data Classes

//* Dealing product information
type productInformation struct {
	DisplayId   string `json:"displayId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
