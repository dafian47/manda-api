package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/dafian47/manda-api/model"
	"github.com/dafian47/manda-api/util"
	"net/http"
)

func (c *BaseController) GetStatusUserAll(context *gin.Context) {

	var itemList []model.MandaUserStatus

	page := context.DefaultQuery("page", "1")
	perPage := context.DefaultQuery("per_page", "5")
	orderBy := context.DefaultQuery("order_by", "label ASC")

	limit, offset := util.GetLimitAndOffset(perPage, page)

	c.DB.Limit(limit).Offset(offset).Order(orderBy).Find(&itemList)

	if len(itemList) == 0 {
		responseJSON(context, http.StatusNotFound, "Empty data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", itemList)
}

func (c *BaseController) GetStatusUser(context *gin.Context) {

	var item model.MandaUserStatus

	code := context.Param("code")

	c.DB.Where(&model.MandaUserStatus{Code:code}).First(&item)

	if item.Code == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", item)
}

func (c *BaseController) CreateStatusUser(context *gin.Context) {

	var item model.MandaUserStatus

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	c.DB.Save(&item)

	if item.Code == "" {
		responseJSON(context, http.StatusBadRequest, "Failed create data", nil)
		return
	}

	responseJSON(context, http.StatusCreated, "Success create data", item)
}

func (c *BaseController) UpdateStatusUser(context *gin.Context) {

	var item model.MandaUserStatus

	code := context.Param("code")

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	c.DB.Where(&model.MandaUserStatus{Code:code}).Save(&item)

	if item.Code == "" {
		responseJSON(context, http.StatusBadRequest, "Failed update data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data", item)
}

func (c *BaseController) DeleteStatusUser(context *gin.Context) {

	var item model.MandaUserStatus

	code := context.Param("code")

	c.DB.Where(&model.MandaUserStatus{Code:code}).First(&item)

	if item.Code == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	c.DB.Delete(&item)

	responseJSON(context, http.StatusOK, "Success delete data", nil)
}
