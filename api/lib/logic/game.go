package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gomock/api/lib/svr"
	"gomock/api/lib/types"
)

// post
func NewKingdom(c *gin.Context) {
	var newKingdomReq types.NewKingdomReq

	c.Bind(&newKingdomReq)

	re, err := svr.NewKingdom(newKingdomReq)

	if err != nil {
		log.Err(err)
		c.JSON(200, types.CodeInternalServiceErr.WithErr(err))
		return
	}

	c.JSON(200,
		types.NewKingdomResp{
			KingdomId: re.KingdomId,
		},
	)
}
