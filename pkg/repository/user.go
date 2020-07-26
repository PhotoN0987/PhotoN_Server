package repository

import (
	"Photon_Server/pkg/database"
	"Photon_Server/pkg/model"
	"database/sql"
	"fmt"
	"log"
)

// UserRepository repository
type UserRepository interface {
	GetAll() ([]model.UserTable, error)
	GetByEmailAndPassword(email string, password string) (model.UserTable, error)
	Create(user model.User) (int64, error)
	Update(user model.User) error
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

// NewUserRepository is init for UserController
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{
		db: db.Connection,
	}
}

// GetAll Get all usersdata
func (r *userRepository) GetAll() ([]model.UserTable, error) {
	users := []model.UserTable{}

	query := `
		SELECT user_id,user_password,user_name,user_email,user_introduce,update_date,create_date FROM users 
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user model.UserTable
		err := rows.Scan(&user.ID, &user.Password, &user.Name, &user.Email, &user.Introduce, &user.UpdatedAt, &user.CreatedAt)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, err
}

// GetByEmailAndPassword Get single usersdata
func (r *userRepository) GetByEmailAndPassword(email string, password string) (model.UserTable, error) {
	user := model.UserTable{}

	query := `
		SELECT user_id, user_name, user_email, user_introduce, update_date, create_date FROM users 
		WHERE user_email=? AND user_password=?
	`
	row := r.db.QueryRow(query, email, password)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Introduce, &user.UpdatedAt, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			return user, err
		}
	}

	return user, err
}

// Create Create user
func (r *userRepository) Create(user model.User) (int64, error) {
	query := `
		INSERT INTO 
		users(user_password, user_name, user_email) 
		VALUES(?, ?, ?)
	`
	stmtInsert, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(user.Password, user.Name, user.Email)
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

// Update Update user
func (r *userRepository) Update(user model.User) error {
	query := `
		UPDATE users 
		SET name=?, email=? 
		WHERE id=?
	`
	stmtUpdate, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmtUpdate.Close()

	result, err := stmtUpdate.Exec(user.Name, user.Email, user.ID)
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
func (r *userRepository) Delete(id int) error {
	query := `
		DELETE 
		FROM users 
		WHERE id=?
	`
	stmtDelete, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	defer stmtDelete.Close()

	result, err := stmtDelete.Exec(id)
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
