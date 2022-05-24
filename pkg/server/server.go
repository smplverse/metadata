package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/piotrostr/metadata/pkg/metadata"
)

var ErrUnsetApiKey = errors.New("No METADATA_API_KEY environment variable set")

func SetupRouter() (r *gin.Engine, err error) {
	gin.SetMode(gin.ReleaseMode)

	apiKey := os.Getenv("METADATA_API_KEY")
	if apiKey == "" {
		err = ErrUnsetApiKey
		return
	}

	m := metadata.New()

	r = gin.Default()

	r.GET("/:tokenId", func(c *gin.Context) {
		tokenId := c.Param("tokenId")
		c.JSON(http.StatusOK, m.Get(tokenId))
	})

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.POST("/:tokenId", func(c *gin.Context) {
		tokenId := c.Param("tokenId")
		_, err := strconv.Atoi(tokenId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token ID"})
		}

		auth := c.Request.Header.Get("Authorization")
		if auth == "" || auth != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		body, err := c.GetRawData()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		entry := metadata.Entry{}
		if err := json.Unmarshal(body, &entry); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		m.Add(tokenId, entry)
		c.Status(http.StatusCreated)
	})

	return
}
