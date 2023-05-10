package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gomock/app/lib/svr"
	types2 "gomock/types"
)

// post
func NewKingdom(c *gin.Context) {
	var newKingdomReq types2.NewKingdomReq

	c.Bind(&newKingdomReq)

	re, err := svr.NewKingdom(newKingdomReq)

	if err != nil {
		log.Err(err)
		c.JSON(200, types2.CodeInternalServiceErr.WithErr(err))
		return
	}

	c.JSON(200,
		types2.NewKingdomResp{
			KingdomId: re.KingdomId,
		},
	)
}
