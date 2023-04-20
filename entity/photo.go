package entity

import (
	"mygram/dto"
	"time"
)

type Photo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Caption string    `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m *Photo) EntityToPhotoResponseDto() dto.PhotoResponse {
	return dto.PhotoResponse{
		Id:          m.Id,
		Title:       m.Title,
		Caption: 	 m.Caption,
		PhotoUrl: 	 m.PhotoUrl,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
