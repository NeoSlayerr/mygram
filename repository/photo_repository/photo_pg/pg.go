package photo_pg

import (
	"time"
	"database/sql"
	"errors"
	"fmt"
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/repository/photo_repository"
)

const (
	getPhotoByIdQuery = `
		SELECT id, title, caption, photo_url, user_id, created_at, updated_at from "photos"
		WHERE id = $1;
	`

	getPhotosQuery = `
		SELECT id, title, caption, photo_url, user_id, created_at, updated_at from "photos"
	`

	updatePhotoByIdQuery = `
		UPDATE "photos"
		SET title = $2,
		caption = $3,
		photo_url= $4,
		updated_at = $5
		WHERE id = $1;
	`

	deletePhotoByIdQuery = `
		DELETE FROM "photos"
		WHERE id = $1;
	`
)

type photoPG struct {
	db *sql.DB
}

func NewPhotoPG(db *sql.DB) photo_repository.PhotoRepository {
	return &photoPG{
		db: db,
	}
}

func (m *photoPG) DeletePhotoById(photoId int) errs.MessageErr {
	_, err := m.db.Exec(deletePhotoByIdQuery, photoId)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *photoPG) UpdatePhotoById(payload entity.Photo) errs.MessageErr {
	_, err := m.db.Exec(updatePhotoByIdQuery, payload.Id, payload.Title, payload.Caption, payload.PhotoUrl, time.Now())

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *photoPG) GetPhotos() ([]*entity.Photo, errs.MessageErr) {
	rows, err := m.db.Query(getPhotosQuery)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("photo not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	photos := []*entity.Photo{}

	for rows.Next() {
		var photo entity.Photo

		err := rows.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		photos = append(photos, &photo)

	}

	return photos, nil
}

func (m *photoPG) GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr) {
	row := m.db.QueryRow(getPhotoByIdQuery, photoId)

	var photo entity.Photo

	err := row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId, &photo.CreatedAt, &photo.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("photo not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &photo, nil
}

func (m *photoPG) CreatePhoto(photoPayload *entity.Photo) (*entity.Photo, errs.MessageErr) {
	createPhotoQuery := `
		INSERT INTO "photos"
		(
			title,
			caption,
			photo_url,
			user_id
		)
		VALUES($1, $2, $3, $4)
		RETURNING id,title, caption, photo_url, user_id;
	`
	row := m.db.QueryRow(createPhotoQuery, photoPayload.Title, photoPayload.Caption, photoPayload.PhotoUrl, photoPayload.UserId)

	var photo entity.Photo

	err := row.Scan(&photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &photo, nil

}
