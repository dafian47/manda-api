package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/dafian47/manda-api/model"
	"github.com/dafian47/manda-api/util"
	"net/http"
)

func (c *BaseController) GetWorkAll(context *gin.Context) {

	var itemList []model.MandaWork

	page := context.DefaultQuery("page", "1")
	perPage := context.DefaultQuery("per_page", "5")
	orderBy := context.DefaultQuery("order_by", "id ASC")

	limit, offset := util.GetLimitAndOffset(perPage, page)

	c.DB.Limit(limit).Offset(offset).Order(orderBy).Find(&itemList)

	if len(itemList) == 0 {
		responseJSON(context, http.StatusNotFound, "Empty Data", nil)
	}

	responseJSON(context, http.StatusOK, "Success get data", itemList)
}

func (c *BaseController) GetWork(context *gin.Context) {

	var item model.MandaWork

	id := util.ConvertStringToInt(context.Param("id"))

	c.DB.Where(&model.MandaWork{ID:id}).First(&item)

	if item.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", item)
}

func (c *BaseController) CreateWork(context *gin.Context) {

	var item model.MandaWork

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	c.DB.Save(&item)

	if item.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed create data", nil)
		return
	}

	responseJSON(context, http.StatusCreated, "Success create data", item)
}

func (c *BaseController) UpdateWork(context *gin.Context) {

	var item model.MandaWork

	id := util.ConvertStringToInt(context.Param("id"))

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	c.DB.Where(&model.MandaWork{ID:id}).Save(&item)

	if item.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed update data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data", item)
}

func (c *BaseController) DeleteWork(context *gin.Context) {

	var item model.MandaWork

	id := util.ConvertStringToInt(context.Param("id"))

	c.DB.Where(&model.MandaWork{ID:id}).First(&item)

	if item.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	c.DB.Delete(&item)

	responseJSON(context, http.StatusOK, "Success delete data", nil)
}
