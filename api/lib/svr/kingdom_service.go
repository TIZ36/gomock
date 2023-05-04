package svr

import (
	"fmt"
	"gomock/api/lib/base"
	"gomock/api/lib/types"
)

func NewKingdom(req types.NewKingdomReq) (*types.NewKingdomResp, error) {

	KingdomSvc := &base.Service[types.NewKingdomResp]{}
	re, err :=
		KingdomSvc.
			Proxy(
				func() (types.NewKingdomResp, error) {
					return types.NewKingdomResp{
						KingdomId: 1,
					}, nil
				}).
			Exec()

	fmt.Println(re)
	if err != nil {
		return nil, err
	}

	return &types.NewKingdomResp{
		KingdomId: re.KingdomId,
	}, nil
}
