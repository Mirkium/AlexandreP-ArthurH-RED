package sao

import (
	"fmt"
)

//===================================================================================================================================

func TirCharge() int {
	if stat.Mana_actuelle < 51 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		degat := 0
		for i := 0; i < 2; i++ {
			degat = stat.AD * stat.AD
		}
		stat.Mana_actuelle = stat.Mana_actuelle - 50
		return degat
	}
}

func TirQCM() int {
	var reponse int
	if stat.Mana_actuelle < 101 {
		fmt.Println("Mana insuffiasant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 100
		fmt.Println("Qui a inventée le premier ordinateur ?")
		fmt.Println(" 1- Thomas Edison        2- Bill Gate        3- Allan Turing        4- Jérôme Jouan")
		fmt.Scanln(&reponse)
		if reponse == 3 {
			return stat.AD * 10
		} else {
			return stat.AD / 2
		}
	}
}

func ExplodeArrow() int {
	if stat.Mana_actuelle < 31 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		degat := stat.AD * 2
		stat.Mana_actuelle = stat.Mana_actuelle - 30
		return degat
	}
}

//===================================================================================================================================

func SonicLeap() int {
	fmt.Println("Une compétence d'épée de type charge, l'utilisateur se précipite vers l'ennemi pour lui donner un coup vers le bas.")
	if stat.Mana_actuelle < 51 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 50
		return stat.AD * 3
	}
}

func HorizontaleSquare() int {
	if stat.Mana_actuelle < 31 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		for i := 0; i < 3; i++ {
			stat.Mana_actuelle = stat.Mana_actuelle - 30
			return stat.AD * 2
		}
		return stat.AD * 2
	}
}

func SwordQCM() int {
	var reponse int
	degat := 0
	if stat.Mana_actuelle < 101 {
		Clear()
		fmt.Println("Mana insuffisant")
		fmt.Println(" ")
		return 0	
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 100
		fmt.Println("Avec qu'elle langage a été codée Minecraft ?")
		fmt.Println(" -1 Java        -2 JavaScript        3- C        4- C++")
		fmt.Scanln(&reponse)
		if reponse == 1 {
			degat = stat.AD * 10
		} else {
			degat = stat.AD / 2
		}
		return degat
	}
}

//===================================================================================================================================

func Xcleaves() int {
	if stat.Mana_actuelle < 36 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 35
		return stat.AD * 2
	}
}

func TwistedCleave() int {
	if stat.Mana_actuelle < 51 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 50
		return stat.AD * 4
	}
}

func CleaveQCM() int {
	var reponse int
	if stat.Mana_actuelle < 101 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		fmt.Println("Dans quelle pays 'les pâtes' ont été inventés ?")
		fmt.Println("  1- France    2- Chine    3- Italie    4- Japon")
		fmt.Scanln(&reponse)
		if reponse == 2 {
			return stat.AD * 10
		} else {
			return stat.AD / 2
		}
	}
}


//================================================================================================================================================

func Spike() int{
	if stat.Mana_actuelle < 31 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 30
		return stat.AD * 2
	}
}

func Javelot() int{
	if stat.Mana_actuelle < 51 {
		fmt.Println("Mana Insuffisant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 50
		return stat.AD * 4
	}
}

func SpearQCM() int {
	var reponse int
	if stat.Mana_actuelle < 101 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 100
		fmt.Println("Quelle armée de la Grèce Antique est-elle connue pour son maniement de la lance ?")
		fmt.Println("  1- Athène   2- Percy Jackson   3- Méliodas    4- Sparte")
		fmt.Scanln(&reponse)
		if reponse == 4 {
			return stat.AD * 10
		} else {
			return stat.AD /2
		}
	}
}

//=============================================================================================================================================


func Mutilation() int {
	if stat.Mana_actuelle < 40 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		return stat.AD * 3
	}
}

func Assassina() int {
	if stat.Mana_actuelle < 60 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 60 
		return stat.AD * 5
	}
}

func KnifeQCM() int {
	var reponse int
	if stat.Mana_actuelle < 100 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 100
		fmt.Println("Quelle est la date de création du projet éducative privé 'Ynov' ?")
		fmt.Println("   1- 2010     2- 2012    3- 2015   4- 2014")
		fmt.Scanln(&reponse)
		if reponse == 1 {
			return stat.AD * 10
		} else {
			return stat.AD / 2
		}
	}
}

//===============================================================================================================================================

func Bonk() int {
	if stat.Mana_actuelle < 40 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 40
		return stat.AD * 4
	}
}


func Coup_de_poile() int {
	if stat.Mana_actuelle < 30 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		return stat.AD *2
	}
}

func MasseQCM() int {
	var reponse int
	if stat.Mana_actuelle < 100 {
		fmt.Println("Mana insuffisant")
		return 0
	} else {
		stat.Mana_actuelle = stat.Mana_actuelle - 100
		fmt.Println("Quelle est la date précise du débarquement de Normandie pendant la Seconde Guerre Mondial ?")
		fmt.Println("   1- 10 Juin 1944    2- 5 Juin 1944    3- 6 Juin 1944   4- 9 Juin 1944")
		fmt.Scanln(&reponse)
		if reponse == 3 {
			return stat.AD * 10
		} else {
			return stat.AD / 2
		}
	}
}


//============================================================================================================================================
