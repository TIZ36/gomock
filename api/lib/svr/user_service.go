package svr

import (
	"gomock/api/lib/base"
	"gomock/api/lib/types"
	"gomock/model/mysql"
)

func NewUserBasicInfo(uid int64, curStage int32, maincityLevel int32) error {

	return mysql.UBIR.InsertBasicInfo(mysql.BasicInfo{
		Uid:           uid,
		CurStage:      curStage,
		MaincityLevel: maincityLevel,
	})
}

func GetUserBasicInfo(uid int64) (*types.UserBasicInfo, error) {
	userBasicInfoSvc := &base.Service[mysql.BasicInfo]{}

	userBasicInfo, err := userBasicInfoSvc.Proxy(func() (mysql.BasicInfo, error) {
		re, err := mysql.UBIR.FindBasicInfoByUid(uid)

		return *re, err
	}).Cached("UserBasicInfo", string(uid)).Exec()

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

	//var userInfo mysql.UserInfo
	//var serviceUserInfo types.UserInfo
	//
	//re, err := UserSvc.
	//	Proxy(func() (any, error) { return mysql.UIR.GetUserInfo(uid) }).
	//	Cached("UserInfo", string(uid)).
	//	Exec()
	//
	//userInfoBin, _ := re.([]byte)
	//err = json.Unmarshal(userInfoBin, &userInfo)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = json.Unmarshal(userInfo.Data, &serviceUserInfo)
	//
	//return &serviceUserInfo, nil

	userInfoSvc := &base.Service[mysql.UserInfo]{}

	userInfo, err := userInfoSvc.Proxy(func() (mysql.UserInfo, error) {
		re, err := mysql.UIR.GetUserInfo(uid)
		return *re, err
	}).Cached("UserInfo", string(uid)).Exec()

	if err != nil {
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
