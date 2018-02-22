package router

import (
	"github.com/dafian47/manda-api/config"
	"github.com/dafian47/manda-api/controller"
	"github.com/dafian47/manda-api/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/secure"
)

func InitRouter(db *gorm.DB) *gin.Engine {

	secureMiddleware := secure.New(secure.Options{
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
		IsDevelopment:      config.IsDevelopment,
	})

	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {

			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
				c.Abort()
				return
			}

			// Avoid header rewrite if response is a redirection.
			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	// Set DebugMode if you want to enable Log on Rest Server ( Gin )
	// And set ReleaseMode if you want to deploy to Production
	if config.IsDevelopment {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(secureFunc)
	router.Use(util.Limit(1000))
	router.MaxMultipartMemory = 8 << 20
	router.Static("/image", "./resource/images")

	baseController := controller.BaseController{DB: db}

	authV1 := router.Group("/auth")
	{
		authV1.POST("/login", baseController.Login)
		authV1.POST("/register", baseController.Register)
		authV1.POST("/forgetPassword", baseController.ForgetPassword)
		authV1.POST("/emailVerification", baseController.EmailVerification)
	}

	apiV1 := router.Group("/api/v1")
	{
		userRoute := apiV1.Group("/user")
		{
			userRoute.GET("/", baseController.GetUserAll)
			userRoute.GET("/:id", baseController.GetUser)
			userRoute.PUT("/:id", baseController.UpdateUser)
			userRoute.DELETE("/:id", baseController.DeleteUser)
		}

		channelRoute := apiV1.Group("/channel")
		{
			channelRoute.GET("/", baseController.GetChannelAll)
			channelRoute.GET("/:id", baseController.GetChannel)
			channelRoute.POST("/", baseController.CreateChannel)
			channelRoute.PUT("/:id", baseController.UpdateChannel)
			channelRoute.DELETE("/:id", baseController.DeleteChannel)
		}

		threadRoute := apiV1.Group("/thread")
		{
			threadRoute.GET("/", baseController.GetThreadAll)
			threadRoute.GET("/:id", baseController.GetThread)
			threadRoute.POST("/", baseController.CreateThread)
			threadRoute.PUT("/:id", baseController.UpdateThread)
			threadRoute.DELETE("/:id", baseController.DeleteThread)

			threadRoute.GET("/:id/comment", baseController.GetCommentAll)
			threadRoute.GET("/:id/comment/:comment_id", baseController.GetComment)
			threadRoute.POST("/:id/comment/", baseController.CreateComment)
			threadRoute.PUT("/:id/comment/:comment_id", baseController.UpdateComment)
			threadRoute.DELETE("/:id/comment/:comment_id", baseController.DeleteComment)
		}

		masterRoute := apiV1.Group("/master")
		{
			marriageRoute := masterRoute.Group("/marriage")
			{
				marriageRoute.GET("/", baseController.GetMarriageAll)
				marriageRoute.GET("/:id", baseController.GetMarriage)
				marriageRoute.POST("/", baseController.CreateMarriage)
				marriageRoute.PUT("/:id", baseController.UpdateMarriage)
				marriageRoute.DELETE("/:id", baseController.DeleteMarriage)
			}

			workRoute := masterRoute.Group("/work")
			{
				workRoute.GET("/", baseController.GetWorkAll)
				workRoute.GET("/:id", baseController.GetWork)
				workRoute.POST("/", baseController.CreateWork)
				workRoute.PUT("/:id", baseController.UpdateWork)
				workRoute.DELETE("/:id", baseController.DeleteWork)
			}

			majorRoute := masterRoute.Group("/major")
			{
				majorRoute.GET("/", baseController.GetMajorAll)
				majorRoute.GET("/:code", baseController.GetMajor)
				majorRoute.POST("/", baseController.CreateMajor)
				majorRoute.PUT("/:code", baseController.UpdateMajor)
				majorRoute.DELETE("/:code", baseController.DeleteMajor)
			}

			roleRoute := masterRoute.Group("/role")
			{
				roleRoute.GET("/", baseController.GetRoleAll)
				roleRoute.GET("/:code", baseController.GetRole)
				roleRoute.POST("/", baseController.CreateRole)
				roleRoute.PUT("/:code", baseController.UpdateRole)
				roleRoute.DELETE("/:code", baseController.DeleteRole)
			}

			userRoute := masterRoute.Group("/user")
			{
				userRoute.GET("/", baseController.GetStatusUserAll)
				userRoute.GET("/:code", baseController.GetStatusUser)
				userRoute.POST("/", baseController.CreateStatusUser)
				userRoute.PUT("/:code", baseController.UpdateStatusUser)
				userRoute.DELETE("/:code", baseController.DeleteStatusUser)
			}

			threadRoute := masterRoute.Group("/thread")
			{
				threadRoute.GET("/", baseController.GetStatusThreadAll)
				threadRoute.GET("/:code", baseController.GetStatusThread)
				threadRoute.POST("/", baseController.CreateStatusThread)
				threadRoute.PUT("/:code", baseController.UpdateStatusThread)
				threadRoute.DELETE("/:code", baseController.DeleteStatusThread)
			}
		}
	}

	return router
}
