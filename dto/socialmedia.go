package dto

import "time"

type NewSocialMediaRequest struct {
	Name string `json:"name" valid:"required~name cannot be empty" example:"instagram"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~social media url cannot be empty" example:"http://instagram.com"`
}

type NewSocialMediaResponse struct {
	Result     string `json:"result" example:"success"`
	Message    string `json:"message" example:"new social media data successfully created"`
	StatusCode int    `json:"statusCode" example:"201"`
}

type SocialMediaResponse struct {
	Id          int       `json:"id" example:"1"`
	Name string `json:"name" example:"instagram"`
	SocialMediaUrl string `json:"social_media_url" example:"http://instagram.com"`
	UserId      int       `json:"user_id" example:"1"`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-01-01"`
}

type GetSocialMediaResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data []SocialMediaResponse `json:"data"`
}