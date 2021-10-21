package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/xielizyh/ctid_service/docs"
	"github.com/xielizyh/ctid_service/internal/middleware"
	v1 "github.com/xielizyh/ctid_service/internal/routers/api/v1"
)

// NewRouter 新建路由
func NewRouter() *gin.Engine {
	// 创建gin引擎实例
	r := gin.New()
	// 使用Logger中间件
	r.Use(gin.Logger())
	// 使用Recovery中间件
	r.Use(gin.Recovery())
	// 使用Translations注册
	r.Use(middleware.Translations())
	// 注册针对swagger的路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	order := v1.NewOrder()
	// 创建api/v1路由组
	apiv1 := r.Group("/api/v1")
	// 注册Handler到对应的路由规则
	{
		// 增
		apiv1.POST("/orders", order.Create)
		// 删
		apiv1.DELETE("/orders/:id", order.Delete)
		// 改
		// apiv1.PUT("/orders/:id", order.Update)
		// 查
		apiv1.GET("/orders", order.List)
		// 查(这个冒号:表示参数id)，示例：http://127.0.0.1:8000/api/v1/orders/8
		apiv1.GET("/orders/:id", order.Get)
	}

	return r
}
