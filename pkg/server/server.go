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

var ErrUnsetApiKey = errors.New("no METADATA_API_KEY environment variable set")

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
		_, err := strconv.Atoi(tokenId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token ID"})
			return
		}
		entry, _ := m.Get(tokenId)
		// TODO handle error above
		c.JSON(http.StatusOK, entry)
	})

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.POST("/:tokenId", func(c *gin.Context) {
		tokenId := c.Param("tokenId")
		_, err := strconv.Atoi(tokenId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token ID"})
			return
		}
		if c.Request.Header.Get("Content-Type") != "application/json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be application/json"})
			return
		}

		auth := c.Request.Header.Get("Authorization")
		if auth == "" || auth != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		body, _ := c.GetRawData()

		entry := metadata.Entry{}
		// handling of the below can be dropped for the sake of handling m.Add
		if err := json.Unmarshal(body, &entry); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = m.Add(tokenId, entry)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid entry"})
			return
		}
		c.Status(http.StatusCreated)
	})

	return
}
