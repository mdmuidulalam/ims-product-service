package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"

	routeinterface "product-service/routes/interfaces"

	"product-service/data"
	"product-service/logics"
)

func ProductsRoutes(router *gin.Engine) {
	router.POST("/create", create)
	router.GET("/readdisplaycard", readDisplayCard)
	router.GET("/readbulk", readBulk)
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

	var createProductLogic routeinterface.ICreateProductLogic
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

	var readProductLogic routeinterface.IReadProductLogic
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

// @Summary Read bulk data for products
// @Description It will collect primary data for products
// @ID read-bulk
// @Router /readbulk [get]
// @Accept json
// @Param pagenumber body int true "pageNumber of products bulk data"
// @Param pagesize body int true "pageSize of products bulk data"
// @Param orderby body string true "orderBy of products bulk data"
// @Param isorderbyincreasing body bool true "isOrderbyIncreasing of products bulk data"
// @Success 250 {object} object{products=[]object} "Products bulk data"
func readBulk(c *gin.Context) {
	pageNumber, err1 := strconv.Atoi(c.Query("pagenumber"))
	pageSize, err2 := strconv.Atoi(c.Query("pagesize"))
	orderBy := c.Query("orderby")
	isOrderbyIncreasing, err3 := strconv.ParseBool(c.Query("isorderbyincreasing"))

	if err1 != nil || pageNumber <= 0 || err2 != nil || pageSize <= 0 || len(orderBy) == 0 || err3 != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request (Data representation not correct). Please, observe the API doc.",
		})
		return
	}

	var readProductLogic routeinterface.IReadProductLogic
	readProductLogic = &logics.ReadProductLogic{
		ProductData: &data.ProductsData{
			PostgresData: &data.PostgresData{},
		},
	}

	products := readProductLogic.ReadBulk(pageNumber, pageSize, orderBy, isOrderbyIncreasing)
	productsJson := []gin.H{}
	for _, element := range products {
		productsJson = append(productsJson, gin.H{
			"id":          (*element).GetId(),
			"displayId":   (*element).GetDisplayId(),
			"name":        (*element).GetName(),
			"description": (*element).GetDescription(),
		})
	}

	c.JSON(250, gin.H{
		"products": productsJson,
	})
}

//* Data Classes

//* Dealing product information
type productInformation struct {
	DisplayId   string `json:"displayId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
