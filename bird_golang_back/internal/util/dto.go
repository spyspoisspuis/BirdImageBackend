package util

type Bird struct {
	ID             string  `json:"ID" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	Description	   string	`json:"description" binding:"required"` 
}