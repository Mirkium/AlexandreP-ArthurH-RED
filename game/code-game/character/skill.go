package sao

import (
	"fmt"
)

var degat = stat.AD

//===================================================================================================================================

func TirCharge() {
	fmt.Println("Pendant deux tours, charge une attaque qui causera plus de dégât à la fin")
	for i := 0; i < 2; i++ {
		degat = degat * stat.AD
	}
	fmt.Print(degat)
	fmt.Println("couldown : 2")
}

func TirQCM() {
	var reponse int
	fmt.Println("Qui a inventée le premier ordinateur ?")
	fmt.Println(" 1- Thomas Edison        2- Bill Gate        3- Allan Turing        4- Jérôme Jouan")
	fmt.Scanln(&reponse)
	if reponse == 3 {
		degat = degat * 10
	} else {
		degat = degat / 2
	}
	fmt.Println(degat)
	fmt.Println("Cette compétence peut être utiliser que une fois par combat")
}

func ExplodeArrow() {
	fmt.Println("Tire une flèche de feu causant plus de dégât et applique brulure à la cible pendant 3 tours")
	degat = degat * 2
	brulure := 10
	fmt.Println("degat de la flêche : ", degat, " et ce de brulure ", brulure)
	fmt.Println("couldown : 3")
}

//===================================================================================================================================

func SonicLeap() {
	fmt.Println("Une compétence d'épée de type charge, l'utilisateur se précipite vers l'ennemi pour lui donner un coup vers le bas.")
	fmt.Println(degat / 3 * degat)
	fmt.Println("couldown : 3")
}

func HorizontaleSquare() {
	fmt.Println("Une compétence d'épée traçant la forme d'un carré horizontale.")
	for i := 0; i < 3; i++ {
		fmt.Println(degat * 2)
	}
	fmt.Println(degat * 2)
	fmt.Println("couldown : 2")
}

func SwordQCM() {
	var reponse int
	fmt.Println("Avec qu'elle langage a été codée Minecraft ?")
	fmt.Println(" -1 Java        -2 JavaScript        3- C        4- C++")
	fmt.Scanln(&reponse)
	if reponse == 1 {
		degat = degat * 10
	} else {
		degat = degat / 2
	}
	fmt.Println(degat)
	fmt.Println("Cette compétence peut être utiliser que une fois par combat")
}

//===================================================================================================================================

func Xcleaves() {
	fmt.Println("Compétence de hâche rapide, qui vient dessiner X sur l'ennemie")
	fmt.Println(degat * 4)
	fmt.Println("couldown : 2")
}
