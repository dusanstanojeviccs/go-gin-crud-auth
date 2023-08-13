package users

import (
	"database/sql"
	"go-gin-crud-auth/utils/db"
)

type userRepository struct{}

func userMapper(rows *sql.Rows, u *User) error {
	return rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password)
}

func (this *userRepository) findByEmail(email string) (*User, error) {
	return db.SelectSingle[User](
		userMapper,
		"SELECT id, name, email, password FROM users WHERE email = ?",
		email,
	)
}

func (this *userRepository) findById(id int) (*User, error) {
	return db.SelectSingle[User](
		userMapper,
		"SELECT id, name, email, password FROM users WHERE id = ?",
		id,
	)
}

func (this *userRepository) create(user *User) error {
	id, error := db.Insert(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, user.Password,
	)
	user.Id = id

	return error
}

var UserRepository = userRepository{}
