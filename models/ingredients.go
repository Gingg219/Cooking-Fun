package models

import (
	"log"
	"time"
	"github.com/upper/db/v4"
)

type Ingredient struct {
	ID        int       `db:"id,omitempty"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func (m IngredientsModel) getNameTable() string {
	return "ingredients"
}

type IngredientsModel struct {
	db db.Session
}

func (m IngredientsModel) GetAll() ([]Ingredient, error) {

	var ingredients []Ingredient

	err := m.db.SQL().SelectFrom(m.getNameTable()).All(&ingredients)
	if err != nil {
		log.Fatal(err)
	}

	return ingredients, nil
}

func (m IngredientsModel) GetByNameIngredient(name string) ([]Ingredient, error) {
	var ingredients []Ingredient

	err := m.db.SQL().SelectFrom(m.getNameTable()).Where("name = ?", name).All(&ingredients)
	if err != nil {
		log.Fatal(err)
	}

	return ingredients, nil
}

func (m IngredientsModel) Insert(i *Ingredient) error {

	i.CreatedAt = time.Now()

	col := m.db.Collection(m.getNameTable())
	_, err := col.Insert(i)
	if err != nil {
		return err
	}
	return nil
}
