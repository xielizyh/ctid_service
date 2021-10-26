package app

import "github.com/xielizyh/ctid_service/global"

func GetAuthServerAddr() string {
	return global.AppSetting.AuthServerAddr
}
