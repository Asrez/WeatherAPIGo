package dto

type Weather struct {
	City string `json:"city" binding:"required,min=3"`
}