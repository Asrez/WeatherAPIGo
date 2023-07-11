package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/Asrez/WeatherAPIGo/api/dto"
	"github.com/Asrez/WeatherAPIGo/api/helper"
	"github.com/Asrez/WeatherAPIGo/config"
	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {}

func NewWeatherHandler(cfg *config.Config) *WeatherHandler {
	return &WeatherHandler{}
}

func (w *WeatherHandler) Current(c *gin.Context) {
	cfg := config.GetConfig()
	req := new(dto.Weather)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	url := fmt.Sprintf("%scurrent.json?key=%s&q=%s&aqi=no", cfg.API.BaseUrl, cfg.API.Token, req.City)
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

	var weatherData map[string]interface{}
	err = json.Unmarshal(responseData, &weatherData)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(weatherData, true, helper.Success))
}

func (w *WeatherHandler) Forecast(c *gin.Context){
	cfg := config.GetConfig()
	req := new(dto.Weather)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	url := fmt.Sprintf("%sforecast.json?key=%s&q=%s&aqi=no", cfg.API.BaseUrl, cfg.API.Token, req.City)
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

	var weatherData map[string]interface{}
	err = json.Unmarshal(responseData, &weatherData)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(weatherData, true, helper.Success))
}
