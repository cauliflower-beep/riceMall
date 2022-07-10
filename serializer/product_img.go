package serializer

import (
	"mall/conf"
	"mall/model"
)

type ProductImg struct {
	ProductID uint   `json:"product_id" form:"product_id"`
	ImgPath   string `json:"img_path" form:"img_path"`
}

func BuildProductImg(item model.ProductImg) ProductImg {
	return ProductImg{
		ProductID: item.ProductID,
		ImgPath:   conf.ProductPhotoHost + conf.HttpPort + item.ImgPath[1:],
	}
}

func BuildProductImgs(items []model.ProductImg) (productImgs []ProductImg) {
	for _, item := range items {
		product := BuildProductImg(item)
		productImgs = append(productImgs, product)
	}
	return productImgs
}
