package entity

import (
	"mygram/dto"
	"time"
)

type SocialMedia struct {
	Id          int       `json:"id"`
	Name string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m *SocialMedia) EntityToSocialMediaResponseDto() dto.SocialMediaResponse {
	return dto.SocialMediaResponse{
		Id:          m.Id,
		Name: m.Name,
		SocialMediaUrl: m.SocialMediaUrl,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
