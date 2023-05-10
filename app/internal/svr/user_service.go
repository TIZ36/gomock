package svr

import (
	"github.com/rs/zerolog/log"
	"gomock/app/base"
	mysql2 "gomock/app/internal/model/mysql"
	"gomock/types"
)

func NewUserBasicInfo(uid int64, curStage int32, maincityLevel int32) error {

	return mysql2.UBIR.InsertBasicInfo(mysql2.BasicInfo{
		Uid:           uid,
		CurStage:      curStage,
		MaincityLevel: maincityLevel,
	})
}

func GetUserBasicInfo(uid int64) (*types.UserBasicInfo, error) {
	userBasicInfoSvc := &base.Service[mysql2.BasicInfo]{}

	userBasicInfo, err := userBasicInfoSvc.Proxy(func() (mysql2.BasicInfo, error) {
		re, err := mysql2.UBIR.FindBasicInfoByUid(uid)

		return re, err
	}).Exec()

	if err != nil {
		return nil, err
	}

	return &types.UserBasicInfo{
		Uid:           userBasicInfo.Uid,
		CurStage:      userBasicInfo.CurStage,
		MaincityLevel: userBasicInfo.MaincityLevel,
	}, nil
}

func GetUserInfo(uid int64, lang string) (*types.UserInfo, error) {

	userInfoSvc := &base.Service[mysql2.UserInfo]{}

	log.Info().Msg("xxxx")
	userInfo, err := userInfoSvc.Proxy(func() (mysql2.UserInfo, error) {
		re, err := mysql2.UIR.GetUserInfo(uid)

		if err != nil {
			return mysql2.UserInfo{}, err
		}

		return re, nil
	}).Exec()

	if err != nil {
		log.Err(err).Msg("GetUserInfo")
		return nil, err
	}

	return &types.UserInfo{
		Uid:            userInfo.Uid,
		UserName:       "",
		Avatar:         "",
		Kingdom:        types.Kingdom{},
		Guild:          types.Guild{},
		BadgeUrl:       "",
		SubTitleList:   nil,
		AvatarFrameUrl: "",
		BubbleConfigs:  types.BubbleConfigs{},
		VipLevel:       0,
		ShowVip:        false,
		Level:          0,
		GameExtra:      "",
		Sex:            0,
		EmblemUrls:     nil,
		CreateTime:     0,
		TextColor:      "",
		UserType:       0,
	}, nil
}
