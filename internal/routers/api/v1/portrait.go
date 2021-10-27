package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xielizyh/ctid_service/global"
	"github.com/xielizyh/ctid_service/internal/service"
	"github.com/xielizyh/ctid_service/pkg/app"
	"github.com/xielizyh/ctid_service/pkg/errcode"
)

type Portrait struct{}

// NewOrder 新建人像
func NewPortrait() Portrait {
	return Portrait{}
}

// @Summary 实人认证
// @Produce  json
// @Success 200 {object} model.Order "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/orders [get]
func (p Portrait) Auth(c *gin.Context) {
	// 入参校验和参数绑定
	param := service.PortraitAuthRequest{}
	// 创建响应
	response := app.NewResponse(c)
	// 校验
	valid, errs := app.BindAndValid(c, &param, "json")
	if !valid {
		global.Logger.Errorf("app.BindAndValid error: %v", errs)
		// 回应错误
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 正常回应
	code, err := service.PortraitAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.PortraitAuth error: %v", err)
		response.ToErrorResponse(errcode.ErrorPortraitAuthFail)
		return
	}
	log.Println("认证结果:", code)

	switch code {
	case "00XX":
		response.ToResponse(gin.H{})
	default:
		response.ToErrorResponse(errcode.ErrorPortraitAuthFail)
	}
}
