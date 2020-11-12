package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rommel96/torre-information-manager/backend/middleware"
	"github.com/rommel96/torre-information-manager/backend/models"
	"github.com/rommel96/torre-information-manager/backend/repository"
	"github.com/rommel96/torre-information-manager/backend/utils"
)

func login(c *gin.Context) {
	var loginModel models.LoginModel
	if err := c.ShouldBindJSON(&loginModel); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.MsgResponse{
			Status:  models.StatusError,
			Message: models.BodyRequest,
		})
		return
	}
	//find user
	user, err := repository.FindUserFromLogin(loginModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.MsgResponse{
			Status:  models.StatusError,
			Message: models.UserNotFound,
		})
		return
	}
	//validate credentials
	if utils.ValidPassword(user.Password, loginModel.Password) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.MsgResponse{
			Status:  models.StatusError,
			Message: models.PasswordIncorrect,
		})
		return
	}
	//Generate token
	token, err := middleware.GenerateToken(user.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.MsgResponse{
			Status:  models.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.MsgResponse{
		Status:  models.StatusOk,
		Message: token,
	})
}

func signup(c *gin.Context) {
	var signupModel models.SignupModel
	if err := c.ShouldBindJSON(&signupModel); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.MsgResponse{
			Status:  models.StatusError,
			Message: models.BodyRequest,
		})
		return
	}
	hashedPassword, err := utils.Hash(signupModel.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.MsgResponse{
			Status:  models.StatusError,
			Message: err.Error(),
		})
		return
	}
	signupModel.Password = string(hashedPassword)
	insertID, err := repository.InsertUser(signupModel)
	if err != nil {
		if isDup := repository.IsDuplicateKey(err); isDup {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.MsgResponse{
				Status:  models.StatusError,
				Message: models.UserIsTaken,
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.MsgResponse{
			Status:  models.StatusError,
			Message: err,
		})
		return
	}
	token, err := middleware.GenerateToken(insertID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.MsgResponse{
			Status:  models.StatusError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.MsgResponse{
		Status:  models.StatusOk,
		Message: token,
	})
}
