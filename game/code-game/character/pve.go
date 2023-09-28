package sao

import (
	"fmt"
	"math/rand"
	"time"
)

func PVE() {
	fmt.Printf("%s entre en combat avec %s!\n", stat.Pseudo, mob1.nom)

	// Boucle de combat
	for stat.PV_actuelle > 0 && mob1.pointsDeVieMax > 0 {
		// Joueur attaque le monstre
		playerDamage := rand.Intn(stat.AD)
		mob1.pointsDeVieMax -= playerDamage
		fmt.Printf("%s inflige %d points de dégâts à %s!\n", stat.Pseudo, playerDamage, mob1.nom)

		// Vérifier si le monstre est vaincu
		if mob1.pointsDeVieMax <= 0 {
			fmt.Printf("%s a vaincu le monstre!\n", stat.Pseudo)
			break
		}

		// Monstre attaque le joueur
		mob1Damage := rand.Intn(mob1.degats)
		stat.PV_actuelle -= mob1Damage
		fmt.Printf("%s inflige %d points de dégâts à %s!\n", mob1.nom, mob1Damage, stat.Pseudo)

		// Vérifier si le joueur est vaincu
		if stat.PV_actuelle <= 0 {
			fmt.Printf("%s a été vaincu par le monstre!\n", stat.Pseudo)
			break
		}

		// Pause avant le prochain tour
		time.Sleep(1 * time.Second)
	}
}
