package models

import (
	"time"

	"github.com/upper/db/v4"
)

type RecipeIngredient struct {
	RecipeID     int       `db:"recipe_id"`
	IngredientID int       `db:"ingredient_id"`
	Quantity      int       `db:"quantity"`
	CreatedAt     time.Time `db:"created_at"`
}

func (m RecipeIngredientsModel) getNameTable() string {
	return "recipe_ingredient"
}

type RecipeIngredientsModel struct {
	db db.Session
}

func (m RecipeIngredientsModel) Insert(ri *RecipeIngredient) error {

	ri.CreatedAt = time.Now()

	col := m.db.Collection(m.getNameTable())
	_, err := col.Insert(ri)
	if err != nil {
		return err
	}
	return nil
}
