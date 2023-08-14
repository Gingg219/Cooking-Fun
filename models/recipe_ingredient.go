package models

import (
	"time"

	"github.com/upper/db/v4"
)

type RecipeIngredient struct {
	Recipe_ID     int       `db:"recipe_id,omitempty"`
	Ingredient_ID int       `db:"ingredient_id"`
	Quantity      int       `db:"email"`
	CreatedAt     time.Time `db:"created_at"`
}

func (m RecipeIngredientsModel) getNameTable() string {
	return "recipe_ingredient"
}

type RecipeIngredientsModel struct {
	db db.Session
}
