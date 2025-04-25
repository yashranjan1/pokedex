package main

type Pokedex struct {
	pokemon map[string]Pokemon
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		pokemon: make(map[string]Pokemon),
	}
}

func (p Pokedex) Add(name string, pokemon Pokemon) {
	p.pokemon[name] = pokemon
}
