package service

import (
	"mygram/dto"
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/pkg/helpers"
	"mygram/repository/photo_repository"
	"net/http"
)

type PhotoService interface {
	CreatePhoto(userId int, payload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr)
	UpdatePhotoById(photoId int, photoRequest dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr)
	GetPhotoById(photoId int) (*dto.PhotoResponse, errs.MessageErr)
	GetPhotos() (*dto.GetPhotoResponse, errs.MessageErr)
	DeletePhotoById(photoId int) (*dto.NewPhotoResponse, errs.MessageErr)
}

type photoService struct {
	photoRepo photo_repository.PhotoRepository
}

func NewPhotoService(photoRepo photo_repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepo: photoRepo,
	}
}

func (m *photoService) GetPhotos() (*dto.GetPhotoResponse, errs.MessageErr) {
	result, err := m.photoRepo.GetPhotos()

	if err != nil {
		return nil, err
	}

	photoResponse := []dto.PhotoResponse{}

	for _, eachPhoto := range result {
		photoResponse = append(photoResponse, eachPhoto.EntityToPhotoResponseDto())
	}

	response := dto.GetPhotoResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "photo data have been sent successfully",
		Data:       photoResponse,
	}

	return &response, nil
}

func (m *photoService) DeletePhotoById(photoId int) (*dto.NewPhotoResponse, errs.MessageErr) {

	err := m.photoRepo.DeletePhotoById(photoId)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "photo data successfully deleted",
	}

	return &response, nil
}

func (m *photoService) GetPhotoById(photoId int) (*dto.PhotoResponse, errs.MessageErr) {

	result, err := m.photoRepo.GetPhotoById(photoId)

	if err != nil {
		return nil, err
	}

	response := result.EntityToPhotoResponseDto()

	return &response, nil
}

func (m *photoService) UpdatePhotoById(photoId int, photoRequest dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(photoRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Photo{
		Id:          photoId,
		Title:       photoRequest.Title,
		Caption: photoRequest.Caption,
		PhotoUrl: photoRequest.PhotoUrl,
	}

	err = m.photoRepo.UpdatePhotoById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "photo data successfully updated",
	}

	return &response, nil
}

func (m *photoService) CreatePhoto(userId int, payload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	photoRequest := &entity.Photo{
		Title:       payload.Title,
		Caption: payload.Caption,
		PhotoUrl: payload.PhotoUrl,
		UserId:      userId,
	}

	_, err = m.photoRepo.CreatePhoto(photoRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewPhotoResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new photo data successfully created",
	}

	return &response, err
}
