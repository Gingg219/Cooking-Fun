package models

import (
	"errors"
	"log"
	"time"

	"github.com/upper/db/v4"
)

type User struct {
	ID        int       `db:"id,omitempty"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Activated bool      `db:"activated"`
	CreatedAt time.Time `db:"created_at"`
}

func (m UsersModel) getNameTable() string {
	return "users"
}

type UsersModel struct {
	db db.Session
}

func (m UsersModel) Get(id int) (*User, error) {

	var u User

	err := m.db.Collection(m.getNameTable()).Find(db.Cond{"id": id}).One(&u)
	if err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			return nil, ErrNoMoreRows
		}
		log.Fatal(err)
	}

	return &u, nil
}
