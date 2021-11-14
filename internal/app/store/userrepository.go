package store

import (
	"github.com/AnnDutova/database_course_work/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	sqlStatment := `INSERT INTO users(email, encrypted_password) VALUE (?, ?);`
	if res, err := r.store.db.Exec(
		sqlStatment,
		u.Email,
		u.EncryptedPassword); err != nil {
		return nil, err
	} else {
		id, _ := res.LastInsertId()
		u.ID = int(id)
	}
	return u, nil
}
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	sqlRequest := "SELECT id, email, encrypted_password FROM users WHERE email=?"
	if err := r.store.db.QueryRow(sqlRequest, email).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}
