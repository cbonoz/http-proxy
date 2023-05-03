package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cbonoz/http-proxy/models"
	"github.com/cbonoz/http-proxy/util"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

var client = &http.Client{}

func ProxyRequest(c *gin.Context) {
	var body models.CustomRequest
	var err error
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	origin := c.Request.Header.Get("Origin")
	originHash := util.GetMD5Hash(origin)

	msg := fmt.Sprintf("Proxy attempt %s %s %v", origin, originHash, body)
	log.Logger.Info().Msg(msg)
	if originHash != body.Hash {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	req, err := http.NewRequest(body.Type, body.Url, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Body != nil {
		payload, err := json.Marshal(body.Body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Body = io.NopCloser(bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		// Copy any x- prefixed headers
		for k, v := range c.Request.Header {
			if len(k) > 2 && (k[0:2] == "X-" || k[0:2] == "x-") {
				req.Header.Set(k, v[0])
			}
		}

	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, data)
}
