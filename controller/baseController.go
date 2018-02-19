package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/dafian47/manda-api/model/entity"
	"github.com/jinzhu/gorm"
)

type BaseController struct {
	DB *gorm.DB
}

func responseJSON(context *gin.Context, status int, message string, data interface{}) {
	context.JSON(status, &entity.Response{
		Status: status,
		Message: message,
		Data: data,
	})
}
