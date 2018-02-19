package router

import (
	"github.com/dafian47/manda-api/config"
	"github.com/dafian47/manda-api/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/secure"
	"github.com/dafian47/manda-api/controller"
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

	baseController := controller.BaseController{DB:db}

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
			userRoute.PUT("/", baseController.UpdateUser)
			userRoute.DELETE("/:id", baseController.DeleteUser)
		}

		channelRoute := apiV1.Group("/channel")
		{
			channelRoute.GET("/", baseController.GetChannelAll)
			channelRoute.GET("/:id", baseController.GetChannel)
			channelRoute.POST("/", baseController.CreateChannel)
			channelRoute.PUT("/", baseController.UpdateChannel)
			channelRoute.DELETE("/:id", baseController.DeleteChannel)
		}

		postRoute := apiV1.Group("/post")
		{
			postRoute.GET("/", baseController.GetPostAll)
			postRoute.GET("/:id", baseController.GetPost)
			postRoute.POST("/", baseController.CreatePost)
			postRoute.PUT("/", baseController.UpdatePost)
			postRoute.DELETE("/:id", baseController.DeletePost)
		}

		commentRoute := apiV1.Group("/comment")
		{
			commentRoute.GET("/", baseController.GetCommentAll)
			commentRoute.GET("/:id", baseController.GetComment)
			commentRoute.POST("/", baseController.CreateComment)
			commentRoute.PUT("/", baseController.UpdateComment)
			commentRoute.DELETE("/:id", baseController.DeleteComment)
		}

		masterRoute := apiV1.Group("/master")
		{
			marriageRoute := masterRoute.Group("/marriage")
			{
				marriageRoute.GET("/", baseController.GetMarriageAll)
				marriageRoute.GET("/:id", baseController.GetMarriage)
				marriageRoute.POST("/", baseController.CreateMarriage)
				marriageRoute.PUT("/", baseController.UpdateMarriage)
				marriageRoute.DELETE("/:id", baseController.DeleteMarriage)
			}

			workRoute := masterRoute.Group("/work")
			{
				workRoute.GET("/", baseController.GetWorkAll)
				workRoute.GET("/:id", baseController.GetWork)
				workRoute.POST("/", baseController.CreateWork)
				workRoute.PUT("/", baseController.UpdateWork)
				workRoute.DELETE("/:id", baseController.DeleteWork)
			}

			majorRoute := masterRoute.Group("/major")
			{
				majorRoute.GET("/", baseController.GetMajorAll)
				majorRoute.GET("/:code", baseController.GetMajor)
				majorRoute.POST("/", baseController.CreateMajor)
				majorRoute.PUT("/", baseController.UpdateMajor)
				majorRoute.DELETE("/:code", baseController.DeleteMajor)
			}

			roleRoute := masterRoute.Group("/role")
			{
				roleRoute.GET("/", baseController.GetRoleAll)
				roleRoute.GET("/:code", baseController.GetRole)
				roleRoute.POST("/", baseController.CreateRole)
				roleRoute.PUT("/", baseController.UpdateRole)
				roleRoute.DELETE("/:code", baseController.DeleteRole)
			}

			userRoute := masterRoute.Group("/user")
			{
				userRoute.GET("/", baseController.GetStatusUserAll)
				userRoute.GET("/:code", baseController.GetStatusUser)
				userRoute.POST("/", baseController.CreateStatusUser)
				userRoute.PUT("/", baseController.UpdateStatusUser)
				userRoute.DELETE("/:code", baseController.DeleteStatusUser)
			}

			threadRoute := masterRoute.Group("/thread")
			{
				threadRoute.GET("/", baseController.GetStatusThreadAll)
				threadRoute.GET("/:code", baseController.GetStatusThread)
				threadRoute.POST("/", baseController.CreateStatusThread)
				threadRoute.PUT("/", baseController.UpdateStatusThread)
				threadRoute.DELETE("/:code", baseController.DeleteStatusThread)
			}
		}
	}

	return router
}
