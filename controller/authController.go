package controller

import (
	"github.com/dafian47/manda-api/config"
	"github.com/dafian47/manda-api/model"
	"github.com/dafian47/manda-api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *BaseController) Login(context *gin.Context) {

	var user model.MandaUser
	var auth model.MandaAuth

	err := context.BindJSON(&auth)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data", nil)
		return
	}

	plainPassword := auth.Password

	c.DB.Where(&model.MandaAuth{Username: auth.Username}).First(&auth)

	if auth.UserID == "" {
		responseJSON(context, http.StatusNotFound, "Not found auth", nil)
		return
	}

	isMatch := util.MatchString(auth.Password, plainPassword)
	if !isMatch {
		responseJSON(context, http.StatusBadRequest, "Password wrong", nil)
		return
	}

	c.DB.Where(&model.MandaUser{ID: auth.UserID}).First(&user)

	if auth.UserID == "" {
		responseJSON(context, http.StatusNotFound, "Not found user", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Login success", user)
}

func (c *BaseController) Register(context *gin.Context) {

	var user model.MandaUser
	var auth model.MandaAuth

	err := context.BindJSON(&user)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userID, err := util.GenerateUserID()
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed generate id", nil)
	}

	hashPassword, err := util.HashString(user.Password)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed hashing password", nil)
		return
	}

	auth.UserID = userID
	auth.Username = user.Username
	auth.Password = hashPassword

	user.ID = userID
	user.Type = config.USER
	user.Status = config.USER_NOT_VERIFIED

	c.DB.Save(&user)
	c.DB.Save(&auth)

	if user.ID == "" || auth.UserID == "" {
		responseJSON(context, http.StatusBadRequest, "Failed register data", nil)
		return
	}

	user.Password = ""

	responseJSON(context, http.StatusOK, "Success register data", user)
}

func (c *BaseController) ForgetPassword(context *gin.Context) {

}

func (c *BaseController) EmailVerification(context *gin.Context) {

}
