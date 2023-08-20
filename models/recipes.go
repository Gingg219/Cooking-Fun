package models

import (
	"errors"
	"log"
	"time"

	"github.com/upper/db/v4"
)

type Recipe struct {
	ID        int       `db:"id,omitempty"`
	Name      string    `db:"name"`
	Content   string    `db:"content"`
	Image     string    `db:"image"`
	CreatedAt time.Time `db:"created_at"`
}

func (m RecipesModel) getNameTable() string {
	return "recipes"
}

type RecipesModel struct {
	db db.Session
}

func (m RecipesModel) Get(id int) (*Recipe, error) {

	var u Recipe

	err := m.db.Collection(m.getNameTable()).Find(db.Cond{"id": id}).One(&u)
	if err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			return nil, ErrNoMoreRows
		}
		log.Fatal(err)
	}

	return &u, nil
}

func (m RecipesModel) Insert(r *Recipe) (error, int) {

	r.CreatedAt = time.Now()

	col := m.db.Collection(m.getNameTable())
	_, err := col.Insert(r)
	if err != nil {
		log.Fatal(err)
	}

    newRecipeID := r.ID
	
	return nil, newRecipeID
}
