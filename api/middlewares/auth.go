package middlewares

import (
	"net/http"
	"strings"
	"github.com/Asrez/WeatherAPIGo/api/helper"
	"github.com/Asrez/WeatherAPIGo/config"
	"github.com/Asrez/WeatherAPIGo/constants"
	"github.com/Asrez/WeatherAPIGo/packge/service_errors"
	"github.com/Asrez/WeatherAPIGo/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var tokenService = services.NewTokenService(cfg)

	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := c.GetHeader(constants.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if auth == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		} else {
			claimMap, err = tokenService.GetClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				}
			}
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(
				nil, false, helper.AuthError, err,
			))
			return
		}

		c.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		c.Set(constants.FirstNameKey, claimMap[constants.FirstNameKey])
		c.Set(constants.LastNameKey, claimMap[constants.LastNameKey])
		c.Set(constants.UsernameKey, claimMap[constants.UsernameKey])
		c.Set(constants.EmailKey, claimMap[constants.EmailKey])
		c.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])

		c.Next()
	}
}