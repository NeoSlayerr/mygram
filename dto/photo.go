package dto

import "time"

type NewPhotoRequest struct {
	Title       string `json:"title" valid:"required~title cannot be empty" example:"Doraemon"`
	Caption string `json:"caption" example:"Ini foto doraemon"`
	PhotoUrl string `json:"photo_url" valid:"required~photo url cannot be empty" example:"http://imageurl.com"`
}

type NewPhotoResponse struct {
	Result     string `json:"result" example:"success"`
	Message    string `json:"message" example:"new photo data successfully created"`
	StatusCode int    `json:"statusCode" example:"201"`
}

type PhotoResponse struct {
	Id          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Doraemon"`
	Caption string    `json:"caption" example:"Ini foto doraemon"`
	PhotoUrl string `json:"photo_url" example:"http://imageurl.com"`
	UserId      int       `json:"user_id" example:"1"`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-01-01"`
}

type GetPhotoResponse struct {
	Result     string `json:"result" example:"success"`
	Message    string `json:"message" example:"photo data have been sent successfully"`
	StatusCode int    `json:"statusCode" example:"200"`
	Data []PhotoResponse `json:"data"`
}