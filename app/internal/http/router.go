package http

import (
	"github.com/gin-gonic/gin"
	logic2 "gomock/app/internal/http/handler"
	"gomock/types"
	"net/http"
)

func GetRouter(c *gin.Context) {
	t := c.Query("type")

	switch t {
	// -- users -- /
	case "user.basicinfo":
		logic2.GetUserBasicInfo(c)
	case "user.info":
		logic2.GetUserInfo(c)
	case "user.batchInfo":
		logic2.GetUserBatchInfo(c)
	case "user.groups":
		logic2.GetUserGroups(c)
	case "user.friends":
		logic2.GetUserFriends(c)
	case "user.blacklist":
		logic2.GetUserBlackList(c)
	case "user.friendReqs":
		logic2.GetUserFriendReqs(c)
		// -- groups -- /
	case "group.info":
		logic2.GetGroupInfo(c)
	case "group.config":
		logic2.GetGroupConfig(c)
	default:
		c.JSON(http.StatusNotFound, types.CodeParamErr)
		return
	}
}

func NewKingdom(c *gin.Context) {
	logic2.NewKingdom(c)
}
