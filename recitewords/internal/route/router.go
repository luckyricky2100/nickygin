package route

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"nickygin.com/global"
	"nickygin.com/pkg/limiter"
	"nickygin.com/pkg/middleware"
	"nickygin.com/recitewords/internal/route/api"
	v1 "nickygin.com/recitewords/internal/route/api/v1"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())

	//article := v1.NewArticle()
	classes := v1.NewUserClass()
	//upload := api.NewUpload()
	r.GET("/debug/vars", api.Expvar)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.POST("/upload/file", upload.UploadFile)
	r.POST("/auth", api.GetAuth)
	r.POST("/register", api.Register)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/classes", classes.List)
		apiv1.POST("/classes", classes.Create)
	}

	return r
}
