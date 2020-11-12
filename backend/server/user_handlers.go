package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rommel96/torre-information-manager/backend/models"
	"github.com/rommel96/torre-information-manager/backend/repository"
)

const apiBios = "https://torre.bio/api/bios/"
const apiOpportunities = "https://torre.co/api/opportunities/"

func getBioInfo(c *gin.Context) {
	username := c.Param("username")
	response, err := http.Get(apiBios + username)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	defer reader.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func getJobInfo(c *gin.Context) {
	idJob := c.Param("id")
	response, err := http.Get(apiOpportunities + idJob)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	defer reader.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="gopher.png"`,
	}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func saveJob(c *gin.Context) {
	email := c.MustGet("email").(string)
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.MsgResponse{
				Status:  models.StatusError,
				Message: models.BodyRequest,
			})
			return
		}
	}
	user, err := repository.FindUserFromLogin(models.LoginModel{
		Email: email,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.MsgResponse{
			Status:  models.StatusError,
			Message: models.UserNotFound,
		})
		return
	}
	job.UserId = user.Id
	err = repository.InsertJob(&job)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.MsgResponse{
			Status:  models.StatusError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.MsgResponse{
		Status:  models.StatusOk,
		Message: job,
	})
}

func removeJob(c *gin.Context) {
	email := c.MustGet("email").(string)
	jobId := c.Param("id")
	user, err := repository.FindUserFromLogin(models.LoginModel{
		Email: email,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.MsgResponse{
			Status:  models.StatusError,
			Message: models.UserNotFound,
		})
		return
	}
	job, err := repository.FindJobById(jobId)
	if user.Id != job.UserId {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.MsgResponse{
			Status:  models.StatusError,
			Message: models.NotPermissions,
		})
		return
	}
	err = repository.DeleteJob(job)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.MsgResponse{
			Status:  models.StatusError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.MsgResponse{
		Status:  models.StatusOk,
		Message: job,
	})
}

func getFavorites(c *gin.Context) {
	email := c.MustGet("email").(string)
	user, err := repository.FindUserFromLogin(models.LoginModel{
		Email: email,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.MsgResponse{
			Status:  models.StatusError,
			Message: models.UserNotFound,
		})
		return
	}
	jobs, err := repository.FindFavorites(user.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.MsgResponse{
			Status:  models.StatusError,
			Message: err,
		})
		return
	}
	c.JSON(http.StatusOK, models.MsgResponse{
		Status:  models.StatusOk,
		Message: jobs,
	})
}
