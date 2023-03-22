package v1

import (
	"github.com/gin-gonic/gin"
	util "mall/pkg/utils"
	"mall/service"
)

// ListCategories
//  @Description: 分类列表接口
//  @param c
func ListCategories(c *gin.Context) {
	listCategoriesService := service.ListCategoriesService{}
	if err := c.ShouldBind(&listCategoriesService); err != nil {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
	res := listCategoriesService.List(c.Request.Context())
	c.JSON(200, res)
}
