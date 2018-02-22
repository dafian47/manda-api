package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/dafian47/manda-api/util"
	"github.com/dafian47/manda-api/model"
	"net/http"
)

func (c *BaseController) GetCommentAll(context *gin.Context) {

	var itemList []model.MandaComment

	threadID := context.Param("id")

	page := context.DefaultQuery("page", "1")
	perPage := context.DefaultQuery("per_page", "5")
	orderBy := context.DefaultQuery("order_by", "created_at ASC")

	userID := context.DefaultQuery("user_id", "")

	limit, offset := util.GetLimitAndOffset(perPage, page)

	if userID != "" {

		c.DB.Limit(limit).Offset(offset).Order(orderBy).Where(&model.MandaComment{UserID:userID}).Find(&itemList)

		if len(itemList) == 0 {
			responseJSON(context, http.StatusNotFound, "Empty data", nil)
			return
		}

		responseJSON(context, http.StatusOK, "Success get data", itemList)
		return
	}

	c.DB.Limit(limit).Offset(offset).Order(orderBy).Where(&model.MandaComment{ThreadID:threadID}).Find(&itemList)

	if len(itemList) == 0 {
		responseJSON(context, http.StatusNotFound, "Empty data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", itemList)
}

func (c *BaseController) GetComment(context *gin.Context) {

	var item model.MandaComment

	threadID := context.Param("id")
	id := context.Param("comment_id")

	c.DB.Where(&model.MandaComment{ID:id, ThreadID:threadID}).First(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", item)
}

func (c *BaseController) CreateComment(context *gin.Context) {

	var item model.MandaComment

	threadID := context.Param("id")

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	commentID, err := util.GenerateCommentID()
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed generate data", nil)
		return
	}

	item.ID = commentID
	item.ThreadID = threadID

	c.DB.Save(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusBadRequest, "Failed create data", nil)
		return
	}

	responseJSON(context, http.StatusCreated, "Success create data", item)
}

func (c *BaseController) UpdateComment(context *gin.Context) {

	var item model.MandaComment

	threadID := context.Param("id")
	id := context.Param("comment_id")

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	c.DB.Where(&model.MandaComment{ID:id, ThreadID:threadID}).Save(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusBadRequest, "Failed update data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data", item)
}

func (c *BaseController) DeleteComment(context *gin.Context) {

	var item model.MandaComment

	threadID := context.Param("id")
	id := context.Param("comment_id")

	c.DB.Where(&model.MandaComment{ID:id, ThreadID:threadID}).First(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	c.DB.Where(&model.MandaComment{ID:id, ThreadID:threadID}).Delete(&item)

	responseJSON(context, http.StatusOK, "Success delete data", nil)
}
