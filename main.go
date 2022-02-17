package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/isyarah/models"
)

type Data struct {
	Results []Bounds `json:"results"`
}

type Bounds struct {
	Northeast LatLng `json:"northeast"`
	Southwest LatLng `json:"southwest"`
}

type LatLng struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/api/locations/:location", func(c *gin.Context) {
		location := c.Param("location")
		if location == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid url parameter",
			})
			return
		}

		key := os.Getenv("OPENCAGEDATA_KEY")
		url := fmt.Sprintf("https://api.opencagedata.com/geocode/v1/json?key=%s&q=%s&pretty=1&no_annotations=1", key, location)
		resp, _ := http.Get(url)

		var query models.Query
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &query)

		c.JSON(http.StatusOK, query)
	})

	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
