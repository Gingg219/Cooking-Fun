package models

import (
	"github.com/upper/db/v4"
)

type Image struct {
	ID        int    `db:"id,omitempty"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Activated bool   `db:"activated"`
	CreatedAt bool   `db:"created_at"`
}

func (m ImagesModel) getNameTable() string {
	return "images"
}

type ImagesModel struct {
	db db.Session
}
