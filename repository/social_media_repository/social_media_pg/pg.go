package social_media_pg

import (
	"time"
	"database/sql"
	"errors"
	"fmt"
	"mygram/entity"
	"mygram/pkg/errs"
	"mygram/repository/social_media_repository"
)

const (
	getSocialMediaByIdQuery = `
		SELECT id, name, social_media_url, user_id, created_at, updated_at from "socialmedias"
		WHERE id = $1;
	`

	getSocialMediasQuery = `
		SELECT id, name, social_media_url, user_id, created_at, updated_at from "socialmedias"
	`

	updateSocialMediaByIdQuery = `
		UPDATE "socialmedias"
		SET name = $2,
		social_media_url = $3,
		updated_at = $4
		WHERE id = $1;
	`

	deleteSocialMediaByIdQuery = `
		DELETE FROM "socialmedias"
		WHERE id = $1;
	`
)

type socialMediaPG struct {
	db *sql.DB
}

func NewSocialMediaPG(db *sql.DB) social_media_repository.SocialMediaRepository {
	return &socialMediaPG{
		db: db,
	}
}

func (m *socialMediaPG) DeleteSocialMediaById(socialMediaId int) errs.MessageErr {
	_, err := m.db.Exec(deleteSocialMediaByIdQuery, socialMediaId)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *socialMediaPG) UpdateSocialMediaById(payload entity.SocialMedia) errs.MessageErr {
	_, err := m.db.Exec(updateSocialMediaByIdQuery, payload.Id, payload.Name, payload.SocialMediaUrl, time.Now())

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *socialMediaPG) GetSocialMedias() ([]*entity.SocialMedia, errs.MessageErr) {
	rows, err := m.db.Query(getSocialMediasQuery)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("socialMedia not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	socialMedias := []*entity.SocialMedia{}

	for rows.Next() {
		var socialMedia entity.SocialMedia

		err := rows.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		socialMedias = append(socialMedias, &socialMedia)

	}

	return socialMedias, nil
}

func (m *socialMediaPG) GetSocialMediaById(socialMediaId int) (*entity.SocialMedia, errs.MessageErr) {
	row := m.db.QueryRow(getSocialMediaByIdQuery, socialMediaId)

	var socialMedia entity.SocialMedia

	err := row.Scan(&socialMedia.Id, &socialMedia.Name, &socialMedia.SocialMediaUrl, &socialMedia.UserId, &socialMedia.CreatedAt, &socialMedia.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("social media not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &socialMedia, nil
}

func (m *socialMediaPG) CreateSocialMedia(socialMediaPayload *entity.SocialMedia) (*entity.SocialMedia, errs.MessageErr) {
	createSocialMediaQuery := `
		INSERT INTO "socialmedias"
		(
			name,
			social_media_url,
			user_id
		)
		VALUES($1, $2, $3)
		RETURNING id, name, social_media_url, user_id;
	`
	row := m.db.QueryRow(createSocialMediaQuery, socialMediaPayload.Name, socialMediaPayload.SocialMediaUrl, socialMediaPayload.UserId)

	var socialMedia entity.SocialMedia

	err := row.Scan(&socialMedia.Id, &socialMediaPayload.Name, &socialMediaPayload.SocialMediaUrl, &socialMediaPayload.UserId)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &socialMedia, nil

}
