package service

import (
	"mygram/dto"
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/pkg/helpers"
	"mygram/repository/social_media_repository"
	"net/http"
)

type SocialMediaService interface {
	CreateSocialMedia(userId int, payload dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.MessageErr)
	UpdateSocialMediaById(socialMediaId int, socialMediaRequest dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.MessageErr)
	GetSocialMediaById(socialMediaId int) (*dto.SocialMediaResponse, errs.MessageErr)
	GetSocialMedias() (*dto.GetSocialMediaResponse, errs.MessageErr)
	DeleteSocialMediaById(socialMediaId int) (*dto.NewSocialMediaResponse, errs.MessageErr)
}

type socialMediaService struct {
	socialMediaRepo social_media_repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepo social_media_repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{
		socialMediaRepo: socialMediaRepo,
	}
}

func (m *socialMediaService) GetSocialMedias() (*dto.GetSocialMediaResponse, errs.MessageErr) {
	result, err := m.socialMediaRepo.GetSocialMedias()

	if err != nil {
		return nil, err
	}

	socialMediaResponse := []dto.SocialMediaResponse{}

	for _, eachSocialMedia := range result {
		socialMediaResponse = append(socialMediaResponse, eachSocialMedia.EntityToSocialMediaResponseDto())
	}

	response := dto.GetSocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "social media data have been sent successfully",
		Data:       socialMediaResponse,
	}

	return &response, nil
}

func (m *socialMediaService) DeleteSocialMediaById(socialMediaId int) (*dto.NewSocialMediaResponse, errs.MessageErr) {

	err := m.socialMediaRepo.DeleteSocialMediaById(socialMediaId)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediaResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "social media data successfully deleted",
	}

	return &response, nil
}

func (m *socialMediaService) GetSocialMediaById(socialMediaId int) (*dto.SocialMediaResponse, errs.MessageErr) {

	result, err := m.socialMediaRepo.GetSocialMediaById(socialMediaId)

	if err != nil {
		return nil, err
	}

	response := result.EntityToSocialMediaResponseDto()

	return &response, nil
}

func (m *socialMediaService) UpdateSocialMediaById(socialMediaId int, socialMediaRequest dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(socialMediaRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.SocialMedia{
		Id:          socialMediaId,
		Name: socialMediaRequest.Name,
		SocialMediaUrl: socialMediaRequest.SocialMediaUrl,
	}

	err = m.socialMediaRepo.UpdateSocialMediaById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediaResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "social media data successfully updated",
	}

	return &response, nil
}

func (m *socialMediaService) CreateSocialMedia(userId int, payload dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	
	socialMediaRequest := &entity.SocialMedia{
		Name: payload.Name,
		SocialMediaUrl: payload.SocialMediaUrl,
		UserId:      userId,
	}

	_, err = m.socialMediaRepo.CreateSocialMedia(socialMediaRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediaResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new social media data successfully created",
	}

	return &response, err
}
