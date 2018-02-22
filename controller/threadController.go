package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/dafian47/manda-api/model"
	"github.com/dafian47/manda-api/config"
	"github.com/dafian47/manda-api/util"
	"net/http"
)

func (c *BaseController) GetThreadAll(context *gin.Context) {

	var itemList []model.MandaThread

	page := context.DefaultQuery("page", "1")
	perPage := context.DefaultQuery("per_page", "5")
	orderBy := context.DefaultQuery("order_by", "title ASC")

	status := context.DefaultQuery("status", config.APPROVED)
	channelID := context.DefaultQuery("channel_id", "")
	userID := context.DefaultQuery("user_id", "")

	limit, offset := util.GetLimitAndOffset(perPage, page)

	if channelID != "" {

		c.DB.Limit(limit).Offset(offset).Order(orderBy).Where(&model.MandaThread{ChannelID:channelID}).Find(&itemList)

		if len(itemList) == 0 {
			responseJSON(context, http.StatusNotFound, "Empty data", nil)
			return
		}

		responseJSON(context, http.StatusOK, "Success get data", itemList)
		return

	} else if userID != "" {

		c.DB.Limit(limit).Offset(offset).Order(orderBy).Where(&model.MandaThread{UserID:userID}).Find(&itemList)

		if len(itemList) == 0 {
			responseJSON(context, http.StatusNotFound, "Empty data", nil)
			return
		}

		responseJSON(context, http.StatusOK, "Success get data", itemList)
		return
	}

	c.DB.Limit(limit).Offset(offset).Order(orderBy).Where(&model.MandaThread{Status:status}).Find(&itemList)

	if len(itemList) == 0 {
		responseJSON(context, http.StatusNotFound, "Empty data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", itemList)
	return
}

func (c *BaseController) GetThread(context *gin.Context) {

	var item model.MandaThread

	id := context.Param("id")

	c.DB.Where(&model.MandaThread{ID:id}).First(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", item)
}

func (c *BaseController) CreateThread(context *gin.Context) {

	var item model.MandaThread

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data " + err.Error(), nil)
		return
	}

	threadID, err := util.GenerateThreadID()
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed generate id", nil)
		return
	}

	item.ID = threadID
	item.Status = config.WAITING_APPROVAL

	c.DB.Save(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusBadRequest, "Failed save data", nil)
		return
	}

	responseJSON(context, http.StatusCreated, "Success save data", item)
}

func (c *BaseController) UpdateThread(context *gin.Context) {

	var item model.MandaThread

	id := context.Param("id")

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	c.DB.Where(&model.MandaThread{ID:id}).Save(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusBadRequest, "Failed update data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data", item)
}

func (c *BaseController) DeleteThread(context *gin.Context) {

	var item model.MandaThread

	id := context.Param("id")

	c.DB.Where(&model.MandaThread{ID:id}).First(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	c.DB.Where(&model.MandaThread{ID:id}).Delete(&item)

	responseJSON(context, http.StatusOK, "Success delete data", item)
}
