module github.com/megarage9000/pokedex-project

go 1.23.4

require internal/pokeapi v0.0.0
replace internal/pokeapi => ./internal/pokeapi


require internal/pokecache v0.0.0
replace internal/pokecache => ./internal/pokecache