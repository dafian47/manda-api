package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/dafian47/manda-api/model"
	"github.com/dafian47/manda-api/util"
	"net/http"
	"github.com/dafian47/manda-api/config"
)

func (c *BaseController) GetChannelAll(context *gin.Context) {

	var itemList []model.MandaChannel

	page := context.DefaultQuery("page", "1")
	perPage := context.DefaultQuery("per_page", "5")
	orderBy := context.DefaultQuery("order_by", "name ASC")

	status := context.DefaultQuery("status", config.APPROVED)
	userID := context.DefaultQuery("user_id", "")

	limit, offset := util.GetLimitAndOffset(perPage, page)

	if userID != "" {

		c.DB.Limit(limit).Offset(offset).Order(orderBy).Where(&model.MandaChannel{UserID:userID}).Find(&itemList)

		if len(itemList) == 0 {
			responseJSON(context, http.StatusNotFound, "Empty data", nil)
			return
		}

		responseJSON(context, http.StatusOK, "Success get data", itemList)
		return
	}

	c.DB.Limit(limit).Offset(offset).Order(orderBy).Where(&model.MandaChannel{Status:status}).Find(&itemList)

	if len(itemList) == 0 {
		responseJSON(context, http.StatusNotFound, "Empty data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", itemList)
}

func (c *BaseController) GetChannel(context *gin.Context) {

	var item model.MandaChannel

	id := context.Param("id")

	c.DB.Where(&model.MandaChannel{ID:id}).First(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data",item)
}

func (c *BaseController) CreateChannel(context *gin.Context) {

	var item model.MandaChannel

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	channelID, err := util.GenerateChannelID()
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed generate id", nil)
		return
	}

	item.ID = channelID
	item.Status = config.WAITING_APPROVAL

	c.DB.Save(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusBadRequest, "Failed save data", nil)
		return
	}

	responseJSON(context, http.StatusCreated, "Success save data", item)
}

func (c *BaseController) UpdateChannel(context *gin.Context) {

	var item model.MandaChannel

	id := context.Param("id")

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	c.DB.Where(&model.MandaChannel{ID:id}).Save(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusBadRequest, "Failed update data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data", item)
}

func (c *BaseController) DeleteChannel(context *gin.Context) {

	var item model.MandaChannel

	id := context.Param("id")

	c.DB.Where(&model.MandaChannel{ID:id}).First(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	c.DB.Where(&model.MandaChannel{ID:id}).Delete(&item)

	responseJSON(context, http.StatusOK, "Success delete data", nil)
}
