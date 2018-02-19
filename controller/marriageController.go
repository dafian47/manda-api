package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dafian47/manda-api/model"
	"github.com/dafian47/manda-api/util"
)

func (c *BaseController) GetMarriageAll(context *gin.Context) {

	var itemList []model.MandaMarriageStatus

	page := context.DefaultQuery("page", "1")
	perPage := context.DefaultQuery("per_page", "5")
	orderBy := context.DefaultQuery("order_by", "id ASC")

	limit, offset := util.GetLimitAndOffset(perPage, page)

	c.DB.Limit(limit).Offset(offset).Order(orderBy).Find(&itemList)

	if len(itemList) == 0 {
		responseJSON(context, http.StatusBadRequest, "Empty Data", nil)
	}

	responseJSON(context, http.StatusOK, "Success get data", itemList)
}

func (c *BaseController) GetMarriage(context *gin.Context) {

	var item model.MandaMarriageStatus

	id := util.ConvertStringToInt(context.Param("id"))

	c.DB.Where(&model.MandaMarriageStatus{ID:id}).First(&item)

	if item.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", item)
}

func (c *BaseController) CreateMarriage(context *gin.Context) {

	var item model.MandaMarriageStatus

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	c.DB.Save(&item)

	if item.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	responseJSON(context, http.StatusCreated, "Success create data", item)
}

func (c *BaseController) UpdateMarriage(context *gin.Context) {

	var item model.MandaMarriageStatus

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	id := util.ConvertStringToInt(context.Param("id"))

	c.DB.Where(&model.MandaMarriageStatus{ID:id}).Save(&item)

	if item.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data", item)
}

func (c *BaseController) DeleteMarriage(context *gin.Context) {

	var item model.MandaMarriageStatus

	id := util.ConvertStringToInt(context.Param("id"))

	c.DB.Where(&model.MandaMarriageStatus{ID:id}).First(&item)

	if item.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	c.DB.Delete(&item)

	responseJSON(context, http.StatusOK, "Success delete data", nil)
}