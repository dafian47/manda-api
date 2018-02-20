package controller

import (
	"github.com/dafian47/manda-api/model"
	"github.com/dafian47/manda-api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *BaseController) GetStatusThreadAll(context *gin.Context) {

	var itemList []model.MandaThreadStatus

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

func (c *BaseController) GetStatusThread(context *gin.Context) {

	var item model.MandaThreadStatus

	code := context.Param("code")

	c.DB.Where(&model.MandaThreadStatus{Code: code}).First(&item)

	if item.Code == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", item)
}

func (c *BaseController) CreateStatusThread(context *gin.Context) {

	var item model.MandaThreadStatus

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

func (c *BaseController) UpdateStatusThread(context *gin.Context) {

	var item model.MandaThreadStatus

	code := context.Param("code")

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	c.DB.Where(&model.MandaThreadStatus{Code: code}).Save(&item)

	if item.Code == "" {
		responseJSON(context, http.StatusBadRequest, "Failed update data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data", item)
}

func (c *BaseController) DeleteStatusThread(context *gin.Context) {

	var item model.MandaThreadStatus

	code := context.Param("code")

	c.DB.Where(&model.MandaThreadStatus{Code: code}).First(&item)

	if item.Code == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	c.DB.Delete(&item)

	responseJSON(context, http.StatusOK, "Success delete data", nil)
}
