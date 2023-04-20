package social_media_repository

import (
	"mygram/entity"
	"mygram/pkg/errs"
)

type SocialMediaRepository interface {
	CreateSocialMedia(socialMediaPayload *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr)
	GetSocialMediaById(socialMediaId int) (*entity.SocialMedia, errs.MessageErr)
	UpdateSocialMediaById(payload entity.SocialMedia) errs.MessageErr
	GetSocialMedias() ([]*entity.SocialMedia, errs.MessageErr)
	DeleteSocialMediaById(socialMediaId int) errs.MessageErr
}
