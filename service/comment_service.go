package service

import (
	"mygram/dto"
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/pkg/helpers"
	"mygram/repository/comment_repository"
	"net/http"
)

type CommentService interface {
	CreateComment(userId int, payload dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr)
	UpdateCommentById(commentId int, commentRequest dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr)
	GetCommentById(commentId int) (*dto.CommentResponse, errs.MessageErr)
	GetComments() (*dto.GetCommentResponse, errs.MessageErr)
	DeleteCommentById(commentId int) (*dto.NewCommentResponse, errs.MessageErr)
}

type commentService struct {
	commentRepo comment_repository.CommentRepository
}

func NewCommentService(commentRepo comment_repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

func (m *commentService) GetComments() (*dto.GetCommentResponse, errs.MessageErr) {
	result, err := m.commentRepo.GetComments()

	if err != nil {
		return nil, err
	}

	commentResponse := []dto.CommentResponse{}

	for _, eachComment := range result {
		commentResponse = append(commentResponse, eachComment.EntityToCommentResponseDto())
	}

	response := dto.GetCommentResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "comment data have been sent successfully",
		Data:       commentResponse,
	}

	return &response, nil
}

func (m *commentService) DeleteCommentById(commentId int) (*dto.NewCommentResponse, errs.MessageErr) {

	err := m.commentRepo.DeleteCommentById(commentId)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "comment data successfully deleted",
	}

	return &response, nil
}

func (m *commentService) GetCommentById(commentId int) (*dto.CommentResponse, errs.MessageErr) {

	result, err := m.commentRepo.GetCommentById(commentId)

	if err != nil {
		return nil, err
	}

	response := result.EntityToCommentResponseDto()

	return &response, nil
}

func (m *commentService) UpdateCommentById(commentId int, commentRequest dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(commentRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Comment{
		Id:          commentId,
		Message: commentRequest.Message,
		PhotoId: commentRequest.PhotoId,
	}

	err = m.commentRepo.UpdateCommentById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "comment data successfully updated",
	}

	return &response, nil
}

func (m *commentService) CreateComment(userId int, payload dto.NewCommentRequest) (*dto.NewCommentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	commentRequest := &entity.Comment{
		Message: payload.Message,
		UserId:      userId,
		PhotoId: payload.PhotoId,
	}

	_, err = m.commentRepo.CreateComment(commentRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewCommentResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new comment data successfully created",
	}

	return &response, err
}
