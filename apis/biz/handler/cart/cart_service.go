package cart

import (
	"apis/biz/utils"
	cart "apis/hertz_gen/api/cart"
	common "apis/hertz_gen/api/common"
	"apis/rpc"
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth"
	cart_kitex "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
)

// AddCartItem .
// @router /cart [POST]
//func AddCartItem(ctx context.Context, c *app.RequestContext) {
//	var err error
//	var req cart.AddItemReq
//	err = c.BindAndValidate(&req)
//	if err != nil {
//		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
//		return
//	}
//
//	//获取请求头的token
//	token := c.Request.Header.Get("Authorization")
//	if token == "" {
//		utils.SendErrResponse(ctx, c, consts.StatusUnauthorized, err)
//		return
//	}
//	//获取用户id
//	rawID, err := rpc.AuthClient.DecodeToken(ctx, &auth.DecodeTokenReq{Token: token})
//	if err != nil {
//		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
//		return
//	}
//	//获取商品信息
//	resp, err := rpc.CartClient.AddItem(ctx, &cart_kitex.AddItemReq{
//		UserId: uint32(rawID.UserId),
//		Item: &cart_kitex.CartItem{
//			ProductId: req.ProductId,
//			Quantity:  req.Quantity,
//		},
//	})
//
//	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
//}

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}
	//获取请求头的token
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		utils.SendErrResponse(ctx, c, consts.StatusUnauthorized, err)
		return
	}

	// 验证token是否有效
	verify, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{Token: token})

	if verify.Res == false || err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusUnauthorized, err)
	}

	//获取用户id
	rawID, err := rpc.AuthClient.DecodeToken(ctx, &auth.DecodeTokenReq{Token: token})
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}

	resp, err := rpc.CartClient.GetCart(ctx, &cart_kitex.GetCartReq{UserId: uint32(rawID.UserId)})
	if err != nil {
		klog.Error(err)
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// EmptyCart .
// @router /api/cart/del [DELETE]
func EmptyCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		hlog.Error(err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//获取请求头的token
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		hlog.Error(err)
		utils.SendErrResponse(ctx, c, consts.StatusUnauthorized, err)
		return
	}
	// 验证token是否有效
	verify, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{Token: token})

	if verify.Res == false || err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusUnauthorized, err)
	}

	//获取用户id
	rawID, err := rpc.AuthClient.DecodeToken(ctx, &auth.DecodeTokenReq{Token: token})
	if err != nil {
		hlog.Error(err)
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}

	resp, err := rpc.CartClient.EmptyCart(ctx, &cart_kitex.EmptyCartReq{UserId: uint32(rawID.UserId)})

	if err != nil {
		hlog.Error(err)
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AddItem .
// @router /api/cart/add [POST]
func AddItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddItemReq
	err = c.BindAndValidate(&req)
	if err != nil {
		hlog.Error(err)
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}

	//获取请求头的token
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		hlog.Error(err)
		utils.SendErrResponse(ctx, c, consts.StatusUnauthorized, err)
		return
	}

	// 验证token是否有效
	verify, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{Token: token})

	if verify.Res == false || err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusUnauthorized, err)
	}

	//获取用户id
	rawID, err := rpc.AuthClient.DecodeToken(ctx, &auth.DecodeTokenReq{Token: token})
	if err != nil {
		hlog.Error(err)
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}
	//获取商品信息
	resp, err := rpc.CartClient.AddItem(ctx, &cart_kitex.AddItemReq{
		UserId: uint32(rawID.UserId),
		Item: &cart_kitex.CartItem{
			ProductId: req.ProductId,
			Quantity:  req.Quantity,
		},
	})
	if err != nil {
		hlog.Error(err)
		utils.SendErrResponse(ctx, c, consts.StatusInternalServerError, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
