package entity

import (
	"mygram/dto"
	"time"
)

type Comment struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	PhotoId      int       `json:"photo_id"`
	Message string `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m *Comment) EntityToCommentResponseDto() dto.CommentResponse {
	return dto.CommentResponse{
		Id:          m.Id,
		PhotoId: m.PhotoId,
		Message: m.Message,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
