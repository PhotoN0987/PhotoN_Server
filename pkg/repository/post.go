package repository

import (
	"database/sql"
	"fmt"
	"photon-server/pkg/database"
	"photon-server/pkg/model"
)

// PostRepository repository
type PostRepository interface {
	GetAll() ([]model.PostedPhotosTable, error)
	GetByUserID(userid int) ([]model.PostedPhotosTable, error)
	Create(user model.PostedPhotosTable) (int64, error)
	UpdateByPhotoIDAndUserID(user model.PostedPhotosTable) error
	Delete(photoID int, userID int) error
}

type postRepository struct {
	db *sql.DB
}

// NewPostRepository is init for PostController
func NewPostRepository(db *database.DB) PostRepository {
	return &postRepository{
		db: db.Connection,
	}
}

// GetAll is GetAll
func (r *postRepository) GetAll() ([]model.PostedPhotosTable, error) {
	photos := []model.PostedPhotosTable{}

	query := `
		SELECT photo_id, user_id, photo_url, photo_comment, photo_category, delete_flag, update_date, create_date 
		FROM posted_photos
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return photos, err
	}

	for rows.Next() {
		var photo model.PostedPhotosTable
		err := rows.Scan(&photo.PhotoID, &photo.UserID, &photo.PhotoURL, &photo.PhotoComment, &photo.PhotoCategory, &photo.DeleteFlag, &photo.UpdatedAt, &photo.CreatedAt)

		if err != nil {
			return photos, err
		}

		photos = append(photos, photo)
	}

	return photos, err
}

// GetByUserID is GetByUserID
func (r *postRepository) GetByUserID(userid int) ([]model.PostedPhotosTable, error) {
	photos := []model.PostedPhotosTable{}

	query := `
		SELECT photo_id, user_id, photo_url, photo_comment, photo_category, delete_flag, update_date, create_date 
		FROM posted_photos
		WHERE  user_id=?
	`
	rows, err := r.db.Query(query, userid)
	if err != nil {
		return photos, err
	}

	for rows.Next() {
		var photo model.PostedPhotosTable
		err := rows.Scan(&photo.PhotoID, &photo.UserID, &photo.PhotoURL, &photo.PhotoComment, &photo.PhotoCategory, &photo.DeleteFlag, &photo.UpdatedAt, &photo.CreatedAt)

		if err != nil {
			return photos, err
		}

		photos = append(photos, photo)
	}

	return photos, err
}

// Create postedPhotos
func (r *postRepository) Create(postedPhotos model.PostedPhotosTable) (int64, error) {
	query := `
		INSERT INTO 
		posted_photos(photo_id, user_id, photo_url, photo_category, photo_comment, delete_flag) 
		VALUES(?, ?, ?, ?, ?, ?)
	`
	stmtInsert, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(nil, postedPhotos.UserID, postedPhotos.PhotoURL, postedPhotos.PhotoCategory,
		postedPhotos.PhotoComment, postedPhotos.DeleteFlag)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println(lastInsertID)

	return lastInsertID, err
}

// Update postedPhotos
func (r *postRepository) UpdateByPhotoIDAndUserID(postedPhotos model.PostedPhotosTable) error {
	query := `
		UPDATE posted_photos 
		SET photo_url=?, photo_category=?, photo_comment=? 
		WHERE photo_id=? AND user_id=?
	`
	stmtUpdate, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmtUpdate.Close()

	result, err := stmtUpdate.Exec(postedPhotos.PhotoURL, postedPhotos.PhotoCategory, postedPhotos.PhotoComment,
		postedPhotos.PhotoID, postedPhotos.UserID)
	if err != nil {
		return err
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffect == 0 {
		return sql.ErrNoRows
	}

	return err
}

// Delete Delete userdata
func (r *postRepository) Delete(photoID int, userID int) error {
	query := `
		DELETE 
		FROM posted_photos 
		WHERE photo_id=? AND user_id=?
	`
	stmtDelete, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmtDelete.Close()

	result, err := stmtDelete.Exec(photoID, userID)
	if err != nil {
		return err
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffect == 0 {
		return sql.ErrNoRows
	}

	return err
}
