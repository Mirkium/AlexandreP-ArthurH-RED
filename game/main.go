package main

import (
	sao "projet-red/code-game/character"
	graphiste "projet-red/code-game/menu"
)

func main() {
	graphiste.Menu()
	sao.Create_Player()
	sao.Guilde_for_ability()
	sao.Quete()
}
