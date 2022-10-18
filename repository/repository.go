package repository

import (
	"context"
	"database/sql"
	"fmt"

	"test-school/models"
)

const (
	selectProfile = `select u.id, u.username, up.first_name, up.last_name, up.city, ud.school 
from user u
         left join user_profile up on u.id = up.user_id
         left join user_data ud on u.id = ud.user_id
where u.id = ?
`
	selectProfileList = `select u.id, u.username, up.first_name, up.last_name, up.city, ud.school 
from user u
         left join user_profile up on u.id = up.user_id
         left join user_data ud on u.id = ud.user_id
`

	getAuthKey = `select a.api_key
from  auth a
where a.api_key = ?
`
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

type Repository interface {
	GetProfileByID(ctx context.Context, userID int) (*models.Profile, error)
	GetProfileList(ctx context.Context) ([]models.Profile, error)

	GetAuthKey(ctx context.Context, authKey string) (bool, error)
}

func (r repository) GetProfileByID(ctx context.Context, userID int) (*models.Profile, error) {
	row := r.db.QueryRowContext(ctx, selectProfile, userID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var userProfile models.Profile

	err := row.Scan(
		&userProfile.ID,
		&userProfile.Username,
		&userProfile.FirstName,
		&userProfile.LastName,
		&userProfile.City,
		&userProfile.School,
	)
	if err != nil {
		return nil, err
	}

	return &userProfile, nil
}

func (r repository) GetProfileList(ctx context.Context) ([]models.Profile, error) {
	rows, err := r.db.QueryContext(ctx, selectProfileList)
	if err != nil {
		return nil, err
	}

	var userProfiles []models.Profile

	for rows.Next() {
		userProfile := models.Profile{}
		err := rows.Scan(
			&userProfile.ID,
			&userProfile.Username,
			&userProfile.FirstName,
			&userProfile.LastName,
			&userProfile.City,
			&userProfile.School,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		userProfiles = append(userProfiles, userProfile)
	}

	return userProfiles, nil
}

func (r repository) GetAuthKey(ctx context.Context, authKey string) (bool, error) {
	row := r.db.QueryRowContext(ctx, getAuthKey, authKey)
	if row.Err() != nil {
		return false, row.Err()
	}

	var key string

	err := row.Scan(&key)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
