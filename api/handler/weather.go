package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/Asrez/WeatherAPIGo/api/dto"
	"github.com/Asrez/WeatherAPIGo/api/helper"
	"github.com/Asrez/WeatherAPIGo/config"
	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	
}

func NewWeatherHandler(cfg *config.Config) *WeatherHandler {
	return &WeatherHandler{}
}

func (w *WeatherHandler) Current(c *gin.Context) {
	req := new(dto.Weather)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=b49199ccd7464afa8a691054231107&q=London&aqi=no")
	response, err := http.Get(url)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(responseData, true, helper.Success))
}
