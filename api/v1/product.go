package v1

import (
	"github.com/gin-gonic/gin"
	util "mall/pkg/utils"
	"mall/service"
)

// CreateProduct
//  @Description: 创建商品接口
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	//c.SaveUploadedFile()
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListProducts
//  @Description: 商品列表接口
func ListProducts(c *gin.Context) {
	listProductsService := service.ProductService{}
	if err := c.ShouldBind(&listProductsService); err != nil {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
	res := listProductsService.List(c.Request.Context())
	c.JSON(200, res)
}

// ShowProduct
//  @Description: 商品详情接口
func ShowProduct(c *gin.Context) {
	showProductService := service.ProductService{}
	res := showProductService.Show(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

// DeleteProduct
//  @Description: 删除商品接口
func DeleteProduct(c *gin.Context) {
	deleteProductService := service.ProductService{}
	res := deleteProductService.Delete(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

// UpdateProduct
//  @Description: 更新商品接口
func UpdateProduct(c *gin.Context) {
	updateProductService := service.ProductService{}
	if err := c.ShouldBind(&updateProductService); err == nil {
		res := updateProductService.Update(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// SearchProducts
//  @Description: 搜索商品
func SearchProducts(c *gin.Context) {
	searchProductsService := service.ProductService{}
	if err := c.ShouldBind(&searchProductsService); err != nil {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
	res := searchProductsService.Search(c.Request.Context())
	c.JSON(200, res)
}

// ListProductImg
//  @Description: 商品图片接口
func ListProductImg(c *gin.Context) {
	var listProductImgService service.ListProductImgService
	if err := c.ShouldBind(&listProductImgService); err != nil {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
	res := listProductImgService.List(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}
