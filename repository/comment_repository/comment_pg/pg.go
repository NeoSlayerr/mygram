package comment_pg

import (
	"time"
	"database/sql"
	"errors"
	"fmt"
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/repository/comment_repository"
)

const (
	getCommentByIdQuery = `
		SELECT id, user_id, photo_id, message, created_at, updated_at from "comments"
		WHERE id = $1;
	`

	getCommentsQuery = `
		SELECT id, user_id, photo_id, message, created_at, updated_at from "comments"
	`

	updateCommentByIdQuery = `
		UPDATE "comments"
		SET message = $2,
		photo_id = $3,
		updated_at = $4
		WHERE id = $1;
	`

	deleteCommentByIdQuery = `
		DELETE FROM "comments"
		WHERE id = $1;
	`
)

type commentPG struct {
	db *sql.DB
}

func NewCommentPG(db *sql.DB) comment_repository.CommentRepository {
	return &commentPG{
		db: db,
	}
}

func (m *commentPG) DeleteCommentById(commentId int) errs.MessageErr {
	_, err := m.db.Exec(deleteCommentByIdQuery, commentId)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *commentPG) UpdateCommentById(payload entity.Comment) errs.MessageErr {
	_, err := m.db.Exec(updateCommentByIdQuery, payload.Id, payload.Message, payload.PhotoId, time.Now())

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *commentPG) GetComments() ([]*entity.Comment, errs.MessageErr) {
	rows, err := m.db.Query(getCommentsQuery)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("comment not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	comments := []*entity.Comment{}

	for rows.Next() {
		var comment entity.Comment

		err := rows.Scan(&comment.Id, &comment.UserId, &comment.PhotoId, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		comments = append(comments, &comment)

	}

	return comments, nil
}

func (m *commentPG) GetCommentById(commentId int) (*entity.Comment, errs.MessageErr) {
	row := m.db.QueryRow(getCommentByIdQuery, commentId)

	var comment entity.Comment

	err := row.Scan(&comment.Id, &comment.UserId, &comment.PhotoId, &comment.Message, &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("comment not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &comment, nil
}

func (m *commentPG) CreateComment(commentPayload *entity.Comment) (*entity.Comment, errs.MessageErr) {
	createCommentQuery := `
		INSERT INTO "comments"
		(
			user_id,
			photo_id,
			message
		)
		VALUES($1, $2, $3)
		RETURNING id, user_id, photo_id, message;
	`
	row := m.db.QueryRow(createCommentQuery, commentPayload.UserId, commentPayload.PhotoId, commentPayload.Message)

	var comment entity.Comment

	err := row.Scan(&comment.Id, &commentPayload.UserId, &commentPayload.PhotoId, &commentPayload.Message)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &comment, nil

}
