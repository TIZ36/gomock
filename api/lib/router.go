package lib

import (
	"github.com/gin-gonic/gin"
	"gomock/api/lib/logic"
	"gomock/api/lib/types"
	"net/http"
)

func GetRouter(c *gin.Context) {
	t := c.Query("type")

	switch t {
	// -- users -- /
	case "user.basicinfo":
		logic.GetUserBasicInfo(c)
	case "user.info":
		logic.GetUserInfo(c)
	case "user.batchInfo":
		logic.GetUserBatchInfo(c)
	case "user.groups":
		logic.GetUserGroups(c)
	case "user.friends":
		logic.GetUserFriends(c)
	case "user.blacklist":
		logic.GetUserBlackList(c)
	case "user.friendReqs":
		logic.GetUserFriendReqs(c)
		// -- groups -- /
	case "group.info":
		logic.GetGroupInfo(c)
	case "group.config":
		logic.GetGroupConfig(c)
	default:
		c.JSON(http.StatusNotFound, types.CodeParamErr)
		return
	}
}
