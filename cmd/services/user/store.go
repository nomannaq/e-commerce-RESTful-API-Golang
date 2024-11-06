package user

import (
	"database/sql"
	"fmt"

	"github.com/nomannaq/e-commerce-restfulAPI-go/cmd/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUsersByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)

	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {

		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		if u.ID == "" {
			return nil, fmt.Errorf("user not found")
		}
	}
	return u, nil

}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}
