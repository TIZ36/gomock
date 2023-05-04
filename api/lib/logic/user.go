package logic

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"gomock/api/lib/svr"
	"gomock/api/lib/types"
	"net/http"
	"strconv"
)

func GetUserBasicInfo(c *gin.Context) {
	uidStr := c.Query("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)

	if err != nil {
		c.JSON(http.StatusOK, types.CodeParamErr.WithErr(err))
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, types.CodeInternalServiceErr.WithErr(err))
		return
	}

	re, err := svr.GetUserBasicInfo(uid)

	if err != nil {
		c.JSON(http.StatusOK, types.CodeInternalServiceErr.WithErr(err))
		return
	}

	c.JSON(
		http.StatusOK,
		types.GetUserBasicInfoResp{
			UserBasicInfo: types.UserBasicInfo{
				Uid:           uid,
				CurStage:      re.CurStage,
				MaincityLevel: re.MaincityLevel,
			},
		})
}

func GetUserInfo(c *gin.Context) {
	uid := c.Query("uid")
	lang := c.DefaultQuery("language", "en")

	uidInt64, _ := strconv.ParseInt(uid, 10, 64)

	re, err := svr.GetUserInfo(uidInt64, lang)
	if err != nil {
		c.JSON(http.StatusOK, types.CodeInternalServiceErr.WithErr(err))
	}

	c.JSON(http.StatusOK, types.GetUserInfoResp{
		UserInfo: types.UserInfo{
			Uid:            re.Uid,
			UserName:       re.UserName,
			Avatar:         re.Avatar,
			Kingdom:        types.Kingdom{},
			Guild:          types.Guild{},
			BadgeUrl:       "",
			SubTitleList:   nil,
			AvatarFrameUrl: "",
			BubbleConfigs:  types.BubbleConfigs{},
			VipLevel:       0,
			ShowVip:        false,
			Level:          re.Level,
			GameExtra:      "",
			Sex:            re.Sex,
			EmblemUrls:     re.EmblemUrls,
			CreateTime:     re.CreateTime,
			TextColor:      re.TextColor,
			UserType:       re.UserType,
		},
	})
}

func GetUserBatchInfo(c *gin.Context) {}
func GetUserGroups(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
func GetUserFriends(c *gin.Context)    {}
func GetUserBlackList(c *gin.Context)  {}
func GetUserFriendReqs(c *gin.Context) {}

//-- POSTS -- //

func UserCreate(c *gin.Context) {

}
func PostUserAddFriend(c *gin.Context)    {}
func PostUserAcceptFriend(c *gin.Context) {}
func PostUserDelFriend(c *gin.Context)    {}
func PostUserAddBlackList(c *gin.Context) {}
func PostUserDelBlackList(c *gin.Context) {}
