package svr

import (
	"encoding/json"
	"gomock/api/lib/base"
	"gomock/api/lib/types"
	"gomock/model/mysql"
)

type UserService struct {
	base.Service
}

var (
	UserSvc *UserService
)

func init() {
	UserSvc = NewUserService()
}

func NewUserService() *UserService {
	return &UserService{base.Service{}}
}

func NewUserBasicInfo(uid int64, curStage int32, maincityLevel int32) error {

	return mysql.UBIR.InsertBasicInfo(mysql.BasicInfo{
		Uid:           uid,
		CurStage:      curStage,
		MaincityLevel: maincityLevel,
	})
}

func GetUserBasicInfo(uid int64) (*types.UserBasicInfo, error) {

	var basicInfo mysql.BasicInfo
	re, err := UserSvc.Proxy(func() (any, error) { return mysql.UBIR.FindBasicInfoByUid(uid) }).Cached("UserBasicInfo", string(uid)).Exec()
	basicInfoBin, _ := re.([]byte)
	err = json.Unmarshal(basicInfoBin, &basicInfo)

	if err != nil {
		return nil, err
	}

	return &types.UserBasicInfo{
		Uid:           basicInfo.Uid,
		CurStage:      basicInfo.CurStage,
		MaincityLevel: basicInfo.MaincityLevel,
	}, nil
}

func GetUserInfo(uid int64, lang string) (*types.UserInfo, error) {

	var userInfo mysql.UserInfo
	var serviceUserInfo types.UserInfo

	re, err := UserSvc.
		Proxy(func() (any, error) { return mysql.UIR.GetUserInfo(uid) }).
		Cached("UserInfo", string(uid)).
		Exec()

	userInfoBin, _ := re.([]byte)
	err = json.Unmarshal(userInfoBin, &userInfo)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(userInfo.Data, &serviceUserInfo)

	return &serviceUserInfo, nil
}
