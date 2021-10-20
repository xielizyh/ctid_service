package global

import (
	"github.com/xielizyh/ctid_service/pkg/logger"
	"github.com/xielizyh/ctid_service/pkg/setting"
)

// 全局区段配置
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
