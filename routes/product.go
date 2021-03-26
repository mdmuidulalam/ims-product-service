package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"

	routesinterface "product-service/routes/interfaces"

	"product-service/data"
	"product-service/logics"
)

func ProductsRoutes(router *gin.Engine) {
	router.POST("/create", create)
	router.GET("/readdisplaycard", readDisplayCard)
}

// @Summary Create a product
// @Description It will create a product from provided information
// @ID create-product
// @Router /create [post]
// @Accept json
// @Param displayId body string true "DisplayId of a product"
// @Param name body string true "Name of a product"
// @Param description body string true "Description of a product"
// @Success 250 {object} object{errorTypes=[]int} "A product is created. ErrorTypes array will be empty."
// @success 251 {object} object{errorTypes=[]int} "All ready a product with displayId or name. ErrorType{1 => displayId, 2 => name}"
func create(c *gin.Context) {
	var productInfo productInformation
	if err := c.ShouldBind(&productInfo); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request (Data representation not correct). Please, observe the API doc.",
		})
		return
	}

	var createProductLogic routesinterface.ICreateProductLogic
	createProductLogic = &logics.CreateProductLogic{
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

// @Summary Read data for display card for a product
// @Description It will collect primary data for display card for a product
// @ID read-display-card-product
// @Router /readdisplaycard [get]
// @Accept json
// @Param Id body int true "Id of a product"
// @Success 250 {object} object{product=object} "Display card data of a product"
// @Success 251 {object} object{errorTypes=[]int} "No product found with this id"
func readDisplayCard(c *gin.Context) {
	var id int
	var err error

	if id, err = strconv.Atoi(c.Query("id")); id <= 0 || err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request (Data representation not correct). Please, observe the API doc.",
		})
		return
	}

	var readProductLogic routesinterface.IReadProductLogic
	readProductLogic = &logics.ReadProductLogic{
		ProductData: &data.ProductsData{
			PostgresData: &data.PostgresData{},
		},
	}

	readProductLogic.SetId(id)
	statusCode := readProductLogic.ReadDisplayCard()

	productJson := gin.H{
		"product": gin.H{},
	}
	if readProductLogic.GetId() != 0 {
		productJson = gin.H{
			"product": gin.H{
				"displayId": readProductLogic.GetDisplayId(),
				"name":      readProductLogic.GetName(),
			},
		}
	}

	c.JSON(251+statusCode, gin.H{
		"product": productJson,
	})
}

//* Data Classes

//* Dealing product information
type productInformation struct {
	DisplayId   string `json:"displayId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
