package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rommel96/torre-information-manager/backend/middleware"
)

func Run() {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	//config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
	r.GET("/search/:id", searchTorreApi) // search?name=applebo
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", login)
		authRoutes.POST("/signup", signup)
	}

	userRoutes := r.Group("/user")
	{
		userRoutes.Use(middleware.AuthMiddleware())
		/*API Torre */
		userRoutes.GET("/torre-user/:username", getBioInfo)
		userRoutes.GET("/torre-job/:id", getJobInfo)
		/*My API*/
		//userRoutes.GET("/apply/:id", apply) /*https://torre.co/en/jobs/:id*/ // THIS FRONTEND DIRECTLY
		userRoutes.POST("/job", saveJob)
		userRoutes.GET("/job", getFavorites)
		userRoutes.DELETE("/job/:id", removeJob)

	}
	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}

func searchTorreApi(c *gin.Context) {
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
