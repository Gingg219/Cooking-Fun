package main

import (
	"log"
	"net/http"
	"github.com/CloudyKit/jet/v6"
	"github.com/Gingg219/CookingRecipe/models"
)

func (a *application) getAllIngredient(w http.ResponseWriter, r *http.Request) {
	ingredients, err := a.Models.Ingredients.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	vars := make(jet.VarMap)
	vars.Set("ingredients", ingredients)
	err = a.render(w, r, "admin/recipes/index", vars)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *application) storeRecipe(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	recipe := models.Recipe{
		Name:    r.FormValue("name"),
		Content: r.FormValue("content"),
		Image:   r.FormValue("image"),
	}
	err, newRecipeID := a.Models.Recipes.Insert(&recipe)
	if err != nil {
		log.Fatal(err)
	}
	
	var ingredientID int

	for _, value := range r.Form["ingredients[]"] {

		ingredients, err := a.Models.Ingredients.GetByNameIngredient(value)
		if err != nil {
			log.Fatal(err)
		}

		if len(ingredients) == 0 {
			ingredient := models.Ingredient{
				Name: value,
			}
			err = a.Models.Ingredients.Insert(&ingredient)
			if err != nil {
				log.Fatal(err)
			}
			ingredientID = ingredient.ID
		} else {
			ingredientID = ingredients[0].ID
		}

		recipeIngredient := models.RecipeIngredient{
			RecipeID:     newRecipeID,
			IngredientID: ingredientID,
		}
		err = a.Models.RecipeIngredients.Insert(&recipeIngredient)
		if err != nil {
			log.Fatal(err)
		}
	}

	http.Redirect(w, r, "/admin", http.StatusFound)
}
