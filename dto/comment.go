package dto

import "time"

type NewCommentRequest struct {
	Message string `json:"message" valid:"required~message cannot be empty" example:"ini komen"`
	PhotoId int `json:"photo_id" example:"1"`
}

type NewCommentResponse struct {
	Result     string `json:"result" example:"success"`
	Message    string `json:"message" example:"new comment data successfully created"`
	StatusCode int    `json:"statusCode" example:"201"`
}

type CommentResponse struct {
	Id          int       `json:"id" example:"1"`
	UserId      int       `json:"user_id" example:"1"`
	PhotoId      int       `json:"photo_id" example:"1"`
	Message string `json:"message" example:"ini komen"`
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-01-01"`
}

type GetCommentResponse struct {
	Result     string `json:"result" example:"success"`
	Message    string `json:"message" example:"comment data have been sent successfully"`
	StatusCode int    `json:"statusCode" example:"200"`
	Data []CommentResponse `json:"data"`
}