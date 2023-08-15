package models

import (
	"errors"

	"github.com/upper/db/v4"
)

var (
	ErrNoMoreRows = errors.New("No record ma")
)

type Models struct {
	Users             UsersModel
	Recipes           RecipesModel
	Ingredients       IngredientsModel
	RecipeIngredients RecipeIngredientsModel
	Images            ImagesModel
}

func New(db db.Session) Models {
	return Models{
		Users:             UsersModel{db: db},
		Recipes:           RecipesModel{db: db},
		RecipeIngredients: RecipeIngredientsModel{db: db},
		Ingredients:       IngredientsModel{db: db},
		Images:            ImagesModel{db: db},
	}
}
