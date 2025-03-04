package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"product/biz/dal/mysql"
	"product/biz/model"
)

type DeleteProductService struct {
	ctx context.Context
} // NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// Finish your business logic.
	resp = &product.DeleteProductResp{
		Success: false,
		Message: "error",
	}
	getProduct, err := model.GetProduct(mysql.DB, req.GetId())
	if err != nil {

		return nil, err
	} else if getProduct.ID <= 0 {
		resp.Message = "product not found"
		klog.Errorf("id=%v,删除失败,未找到该商品\n", req.GetId())
		return resp, nil
	}
	err = model.DeleteProductById(mysql.DB, int(req.GetId()))
	if err != nil {
		resp.Message = resp.Message + err.Error()
		klog.Errorf("id=%v,删除失败\n", req.GetId())
		return resp, err
	}
	resp.Success = true
	resp.Message = "success"
	klog.Errorf("id=%v,删除成功\n", req.GetId())
	return resp, nil
}
