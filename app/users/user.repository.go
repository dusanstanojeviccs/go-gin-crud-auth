package users

import (
	"database/sql"
	"go-gin-crud-auth/utils/db"
)

type userRepository struct{}

func userMapper(rows *sql.Rows, u *User) error {
	return rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password)
}

func (this *userRepository) findByEmail(tx *sql.Tx, email string) (*User, error) {
	return db.SelectSingle[User](
		tx,
		userMapper,
		"SELECT id, name, email, password FROM users WHERE email = ?",
		email,
	)
}

func (this *userRepository) findById(tx *sql.Tx, id int) (*User, error) {
	return db.SelectSingle[User](
		tx,
		userMapper,
		"SELECT id, name, email, password FROM users WHERE id = ?",
		id,
	)
}

func (this *userRepository) create(tx *sql.Tx, user *User) error {
	id, error := db.Insert(
		tx,
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, user.Password,
	)
	user.Id = id

	return error
}

var UserRepository = userRepository{}
