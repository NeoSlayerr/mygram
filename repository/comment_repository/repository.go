package comment_repository

import (
	"mygram/entity"
	"mygram/pkg/errs"
)

type CommentRepository interface {
	CreateComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr)
	GetCommentById(commentId int) (*entity.Comment, errs.MessageErr)
	UpdateCommentById(payload entity.Comment) errs.MessageErr
	GetComments() ([]*entity.Comment, errs.MessageErr)
	DeleteCommentById(commentId int) errs.MessageErr
}
