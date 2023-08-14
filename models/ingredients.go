package models

import (
	"github.com/upper/db/v4"
)

type Ingredient struct {
	ID        int    `db:"id,omitempty"`
	Name      string `db:"name"`
	CreatedAt bool   `db:"created_at"`
}

func (m IngredientsModel) getNameTable() string {
	return "ingredients"
}

type IngredientsModel struct {
	db db.Session
}
