package logic

import (
	"context"

	"zore-goods/internal/svc"
	"zore-goods/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsLogic {
	return &GetGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc方法
func (l *GetGoodsLogic) GetGoods(in *goods.GoodsRequest) (res *goods.GoodsResponse, err error) {
	//根据订单id获取商品信息
	goodsId := in.GoodsId
	res = new(goods.GoodsResponse)
	res.GoodsId = goodsId
	res.Name = "茅台"
	return
}
