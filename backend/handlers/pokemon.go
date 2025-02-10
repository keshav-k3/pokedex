package handlers

import (
	"net/http"
	"pokedex-backend/models"

	"github.com/gin-gonic/gin"
)

type PokemonHandler struct {
	// TO DO : Add database cache
}

func NewPokemonHandler() *PokemonHandler {
	return &PokemonHandler{}
}

func (h *PokemonHandler) GetAllPokemon(c *gin.Context) {
	// Bulba bulbaaa
	pokemon := []models.Pokemon{
		{
			ID:       1,
			Name:     "Bulbasaur",
			Types:    []string{"Grass", "Poison"},
			ImageURL: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png",
		},
	}

	c.JSON(http.StatusOK, pokemon)
}
