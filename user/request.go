package user

import "encoding/json"

type UserRequest struct {
	Name_user  string      `json:"name_user" binding:"required"`
	Email_user string      `json:"email_user" binding:"required"`
	Image_url  string      `json:"image_url" binding:"required"`
	Following  json.Number `json:"following" binding:"required,number"`
	Followers  json.Number `json:"followers" binding:"required,number"`
}
