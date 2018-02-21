package controller

import (
	"github.com/dafian47/manda-api/model"
	"github.com/dafian47/manda-api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *BaseController) GetUserAll(context *gin.Context) {

	var itemList []model.MandaUser

	page := context.DefaultQuery("page", "1")
	perPage := context.DefaultQuery("per_page", "5")
	orderBy := context.DefaultQuery("order_by", "full_name ASC")

	// Search by full_name & nick_name
	search := context.DefaultQuery("q", "")

	limit, offset := util.GetLimitAndOffset(perPage, page)

	if search != "" {

		c.DB.Limit(limit).Offset(offset).Order(orderBy).Where("full_name = ?", search).Or("nick_name = ?", search).Find(&itemList)

		if len(itemList) == 0 {
			responseJSON(context, http.StatusNotFound, "Not found data", nil)
			return
		}

		responseJSON(context, http.StatusOK, "Success search data", itemList)
		return
	}

	c.DB.Limit(limit).Offset(offset).Order(orderBy).Find(&itemList)

	if len(itemList) == 0 {
		responseJSON(context, http.StatusNotFound, "Empty data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", itemList)
}

func (c *BaseController) GetUser(context *gin.Context) {

	var item model.MandaUser

	id := context.Param("id")

	c.DB.Where(&model.MandaUser{ID: id}).First(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data", item)
}

func (c *BaseController) UpdateUser(context *gin.Context) {

	var item model.MandaUser

	id := context.Param("id")

	err := context.BindJSON(&item)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	c.DB.Where(&model.MandaUser{ID:id}).Save(&item)

	if item.ID == "" {
		responseJSON(context, http.StatusBadRequest, "Failed update data", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data", item)
}

func (c *BaseController) DeleteUser(context *gin.Context) {

	var auth model.MandaAuth
	var user model.MandaUser

	id := context.Param("id")

	c.DB.Where(&model.MandaAuth{UserID:id}).First(&auth)
	c.DB.Where(&model.MandaUser{ID:id}).First(&user)

	if auth.UserID == "" || user.ID == "" {
		responseJSON(context, http.StatusNotFound, "Not found data", nil)
		return
	}

	c.DB.Where(&model.MandaAuth{UserID:id}).Delete(&auth)
	c.DB.Where(&model.MandaUser{ID:id}).Delete(&user)

	responseJSON(context, http.StatusOK, "Success delete data", nil)
}
