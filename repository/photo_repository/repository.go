package photo_repository

import (
	"mygram/entity"
	"mygram/pkg/errs"
)

type PhotoRepository interface {
	CreatePhoto(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr)
	UpdatePhotoById(payload entity.Photo) errs.MessageErr
	GetPhotos() ([]*entity.Photo, errs.MessageErr)
	DeletePhotoById(photoId int) errs.MessageErr
}
