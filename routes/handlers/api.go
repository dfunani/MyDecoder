package handlers

import (
	"dfunani/mydecoder/models"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleEncodeMessage(c *gin.Context) {
	var data models.EncodeMessage
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	json_data, err := json.Marshal(data.Data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message := []byte(json_data)
	encoded := base64.StdEncoding.EncodeToString(message)
	c.JSON(http.StatusOK, gin.H{"message": encoded})
}

func HandleDecodeMessage(c *gin.Context) {
	var data models.DecodeMessage
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	decoded, err := base64.StdEncoding.DecodeString(data.Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var decoded_data any
	err = json.Unmarshal(decoded, &decoded_data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": decoded_data})
}
