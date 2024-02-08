package routers

import (
	"github.com/slamancn/gmon-server/routers/ws/v1"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/slamancn/gmon-server/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/slamancn/gmon-server/middleware/jwt"
	"github.com/slamancn/gmon-server/pkg/export"
	"github.com/slamancn/gmon-server/pkg/qrcode"
	"github.com/slamancn/gmon-server/pkg/upload"
	"github.com/slamancn/gmon-server/routers/api"
	"github.com/slamancn/gmon-server/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	// 初始化ws配置
	// html页面位置
	r.LoadHTMLGlob("templates/*")
	// 静态文件位置
	r.Static("/static", "./static")
	// ws长连接业务
	r.GET("/", ws.Index)
	r.GET("/ws", ws.InitWebSocket)
	go r.Run("0.0.0.0:8080")

	return r
}
