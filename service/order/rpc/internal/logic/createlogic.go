package logic

import (
	"context"

	"google.golang.org/grpc/status"

	"mall/service/order/model"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/order"
	"mall/service/product/rpc/product"
	"mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询商品是否存在
	productRes, err := l.svcCtx.ProductRpc.Detail(l.ctx, &product.DetailRequest{
		Id: in.Pid,
	})
	if err != nil {
		return nil, err
	}

	// 判断库存是否充足
	if productRes.Stock <= 0 {
		return nil, status.Error(500, "产品库存不足")
	}

	// 创建订单更新库存
	newOrder := model.Order{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: in.Amount,
		Status: 0,
	}

	res, err := l.svcCtx.OrderModel.Insert(&newOrder)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newOrder.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id:    productRes.Id,
		Stock: productRes.Stock - 1,
	})
	if err != nil {
		return nil, err
	}

	return &order.CreateResponse{
		Id: newOrder.Id,
	}, nil
}
