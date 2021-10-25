package app

import "github.com/xielizyh/ctid_service/global"

func GetAuthCmd() string {
	return "java -jar " + global.AppSetting.AuthCmdPath + "/" + global.AppSetting.AuthCmdName
}
