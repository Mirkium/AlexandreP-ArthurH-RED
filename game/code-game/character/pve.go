package sao

import (
	"fmt"
	"math/rand"
)

var input_pve int
var input_attaque int
var input_skill int

func PVE_1() {
	fmt.Printf("%s entre en combat avec %s!\n", stat.Pseudo, mob1.nom)
	if stat.PV_actuelle > 0 && mob1.pointsDeVieActuels > 0 {
		Player_PVE1()
		Mob1_attaque()
	} else if stat.PV_actuelle <= 0 {
		Clear()
		fmt.Println("Vous êtes mort !!")
	} else {
		stat.xp = stat.xp + 5
		fmt.Println("Vous avez battu ", mob1.nom)
	}
}

func Player_PVE1() {
	var reponse int
	fmt.Println(stat.Pseudo, "                                                                                                 ", mob1.nom)
	fmt.Println("PV :", stat.PV_actuelle, "                                                                                 ", mob1.pointsDeVieActuels)
	fmt.Println("Mana : ", stat.Mana_actuelle)
	fmt.Println("----------------Choix d'Action----------------")
	fmt.Println(" ")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║                 1- ATTAQUE               ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Println(" ")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║                2- INVENTAIRE             ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Println(" ")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║               3- Defense                 ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║               4- MEDITATION              ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Scanln(&input_pve)
	if input_pve == 1 {
		Clear()
		fmt.Println("1- Attaque de base")
		fmt.Println("2- compétence")
		fmt.Println("3- Retour")
		fmt.Scanln(&input_attaque)
		if input_attaque == 1 {
			mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - stat.AD
		} else if input_attaque == 2 {
			if input_classe == 3 {
				fmt.Println("1- Sonic Leap")
				fmt.Println("")
				fmt.Println("2- Horizontale Square")
				fmt.Println("")
				fmt.Println("3- Sword QCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - SonicLeap()
				} else if input_skill == 2 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - HorizontaleSquare()
				} else if input_skill == 3 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - SwordQCM()
				}
			} else if input_classe == 1 {
				fmt.Println("1- Tir Charger")
				fmt.Println(" ")
				fmt.Println("2- Explode Arrow")
				fmt.Println(" ")
				fmt.Println("3- Tir QCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - TirCharge()
				} else if input_skill == 2 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - ExplodeArrow()
				} else if input_skill == 3 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - TirQCM()
				}
			} else if input_classe == 2 {
				fmt.Println("1- Spike")
				fmt.Println(" ")
				fmt.Println("2- Javelot")
				fmt.Println(" ")
				fmt.Println("3- SpearQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - Spike()
				} else if input_skill == 2 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - Javelot()
				} else if input_skill == 3 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - SpearQCM()
				}
			} else if input_classe == 4 {
				fmt.Println("1- Xcleaves")
				fmt.Println(" ")
				fmt.Println("2- TwistedCleave")
				fmt.Println(" ")
				fmt.Println("3- CleaveQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - Xcleaves()
				} else if input_skill == 2 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - TwistedCleave()
				} else if input_skill == 3 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - CleaveQCM()
				}
			} else if input_classe == 5 {
				fmt.Println("1- Mutilation")
				fmt.Println(" ")
				fmt.Println("2- Assassina")
				fmt.Println(" ")
				fmt.Println("3- KinfeQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - Mutilation()
				} else if input_skill == 2 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - Assassina()
				} else if input_skill == 3 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - KnifeQCM()
				}
			} else if input_classe == 6 {
				fmt.Println("1- Bonk")
				fmt.Println(" ")
				fmt.Println("2- Coup de poile")
				fmt.Println(" ")
				fmt.Println("3- MasseQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - Bonk()
				} else if input_skill == 2 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - Coup_de_poile()
				} else if input_skill == 3 {
					mob1.pointsDeVieActuels = mob1.pointsDeVieActuels - MasseQCM()
				}
			}
		} else if input_attaque == 3 {
			Clear()
			Player_PVE1()
		}
	} else if input_pve == 2 {
		Clear()
		fmt.Println("Voulez vous sortir de l'inventaire")
		fmt.Println("  1- OUI          2- NON")
		fmt.Scanln(&reponse)
		if reponse == 1 {
			Clear()
			//
			var reponse1 int
			var reponse2 int
			fmt.Println(Inventory)
			for k := 0; k < len(Inventory); k++ {
				for i := 0; i < len(Inventory); i++ {
					if Inventory[k] == "potion de 40 PV" {
						fmt.Println("Voulez vous utiliser une potion ?")
						fmt.Println("  1- OUI         2- NON")
						fmt.Scanln(&reponse1)
						if reponse1 == 1 {
							fmt.Println("Vous utiliser une potion de vie 40 PV")
							Takepot_20()
							break
						} else if reponse1 == 2 {
							fmt.Println("Vous n'utiliser pas de potion de vie 40 PV")
							break
						}
					} else if Inventory[i] == "potion de poison" {
						fmt.Println("Voulez vous utiliser une potion de poison ?")
						fmt.Println("  1- OUI         2- NON")
						fmt.Scanln(&reponse2)
						if reponse2 == 1 {
							fmt.Println("Vous utilisez une potion de poison")
						} else if reponse2 == 2 {
							fmt.Println("Vous n'utilisez pas une potion de poison")
						}
					}
				}

			}
			//
		} else if reponse == 2 {
			Clear()
			Player_PVE1()
		}

	} else if input_pve == 3 {
		Clear()
		fmt.Println("Vous êtes défendu contre les dégâts adverse")
		fmt.Println(" ")
	} else if input_pve == 4 {
		stat.Mana_actuelle = stat.Mana_max
	}
}

func Mob1_attaque() {
	action := rand.Int31n(3)
	if action == 0 {
		if input_pve == 3 {
			fmt.Println("L'attaque à été barré !!")
			fmt.Println(" ")
		} else {
			fmt.Println("attaque ", mob1.nom)
			stat.PV_actuelle = stat.PV_actuelle - mob1.degats + armor.armor
		}
	} else if action == 1 {
		if input_pve == 3 {
			fmt.Println("L'attaque à été barré !!")
			fmt.Println(" ")
		} else {
			fmt.Println("attaque ", mob1.nom)
			fmt.Println(" ")
			stat.PV_actuelle = stat.PV_actuelle - mob1.degats + armor.armor
		}
	} else {
		if input_pve == 1 {
			mob1.pointsDeVieActuels = mob1.pointsDeVieActuels + stat.AD
		} else {
			fmt.Println(mob1.nom, " a utilisé défense")
			fmt.Println(" ")
		}
	}
	PVE_1()
}

func Takepot_20() int {
	stat.PV_actuelle = stat.PV_actuelle + 40
	if stat.PV_actuelle > stat.PV_max {
		stat.PV_actuelle = stat.PV_max
	}
	return stat.PV_actuelle
}

//=================================================================================================================================================
//=================================================================================================================================================

func PVE_2() {
	fmt.Printf("%s entre en combat avec %s!\n", stat.Pseudo, mob2.nom)
	if stat.PV_actuelle > 0 && mob2.pointsDeVieActuels > 0 {
		Player_PVE2()
		Mob2_attaque()
	} else if stat.PV_actuelle <= 0 {
		Clear()
		fmt.Println("Vous êtes mort !!")
	} else {
		stat.xp = stat.xp + 5
		fmt.Println("Vous avez battu ", mob2.nom)
	}
}

func Player_PVE2() {
	var reponse int
	fmt.Println(stat.Pseudo, "                                                                                                 ", mob2.nom)
	fmt.Println("PV :", stat.PV_actuelle, "                                                                                 ", mob2.pointsDeVieActuels)
	fmt.Println("Mana : ", stat.Mana_actuelle)
	fmt.Println("----------------Choix d'Action----------------")
	fmt.Println(" ")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║                 1- ATTAQUE               ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Println(" ")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║                2- INVENTAIRE             ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Println(" ")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║               3- Defense                 ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║               4- MEDITATION              ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Scanln(&input_pve)
	if input_pve == 1 {
		Clear()
		fmt.Println("1- Attaque de base")
		fmt.Println("2- compétence")
		fmt.Println("3- Retour")
		fmt.Scanln(&input_attaque)
		if input_attaque == 1 {
			mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - stat.AD
		} else if input_attaque == 2 {
			if input_classe == 3 {
				fmt.Println("1- Sonic Leap")
				fmt.Println("")
				fmt.Println("2- Horizontale Square")
				fmt.Println("")
				fmt.Println("3- Sword QCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - SonicLeap()
				} else if input_skill == 2 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - HorizontaleSquare()
				} else if input_skill == 3 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - SwordQCM()
				}
			} else if input_classe == 1 {
				fmt.Println("1- Tir Charger")
				fmt.Println(" ")
				fmt.Println("2- Explode Arrow")
				fmt.Println(" ")
				fmt.Println("3- Tir QCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - TirCharge()
				} else if input_skill == 2 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - ExplodeArrow()
				} else if input_skill == 3 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - TirQCM()
				}
			} else if input_classe == 2 {
				fmt.Println("1- Spike")
				fmt.Println(" ")
				fmt.Println("2- Javelot")
				fmt.Println(" ")
				fmt.Println("3- SpearQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - Spike()
				} else if input_skill == 2 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - Javelot()
				} else if input_skill == 3 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - SpearQCM()
				}
			} else if input_classe == 4 {
				fmt.Println("1- Xcleaveq")
				fmt.Println(" ")
				fmt.Println("2- TwistedCleave")
				fmt.Println(" ")
				fmt.Println("3- CleaveQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - Xcleaves()
				} else if input_skill == 2 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - TwistedCleave()
				} else if input_skill == 3 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - CleaveQCM()
				}
			} else if input_classe == 5 {
				fmt.Println("1- Mutilation")
				fmt.Println(" ")
				fmt.Println("2- Assassina")
				fmt.Println(" ")
				fmt.Println("3- KinfeQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - Mutilation()
				} else if input_skill == 2 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - Assassina()
				} else if input_skill == 3 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - KnifeQCM()
				}
			} else if input_classe == 6 {
				fmt.Println("1- Bonk")
				fmt.Println(" ")
				fmt.Println("2- Coup de poile")
				fmt.Println(" ")
				fmt.Println("3- MasseQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - Bonk()
				} else if input_skill == 2 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - Coup_de_poile()
				} else if input_skill == 3 {
					mob2.pointsDeVieActuels = mob2.pointsDeVieActuels - MasseQCM()
				}
			}
		} else if input_attaque == 3 {
			Clear()
			Player_PVE2()
		}
	} else if input_pve == 2 {
		Clear()
		fmt.Println("Voulez vous sortir de l'inventaire")
		fmt.Println("  1- OUI          2- NON")
		fmt.Scanln(&reponse)
		if reponse == 1 {
			Clear()
			//
			var reponse1 int
			var reponse2 int
			fmt.Println(Inventory)
			for k := 0; k < len(Inventory); k++ {
				for i := 0; i < len(Inventory); i++ {
					if Inventory[k] == "potion de 40 PV" {
						fmt.Println("Voulez vous utiliser une potion ?")
						fmt.Println("  1- OUI         2- NON")
						fmt.Scanln(&reponse1)
						if reponse1 == 1 {
							fmt.Println("Vous utiliser une potion de vie 40 PV")
							Takepot_20()
							break
						} else if reponse1 == 2 {
							fmt.Println("Vous n'utiliser pas de potion de vie 40 PV")
							break
						}
					} else if Inventory[i] == "potion de poison" {
						fmt.Println("Voulez vous utiliser une potion de poison ?")
						fmt.Println("  1- OUI         2- NON")
						fmt.Scanln(&reponse2)
						if reponse2 == 1 {
							fmt.Println("Vous utilisez une potion de poison")
						} else if reponse2 == 2 {
							fmt.Println("Vous n'utilisez pas une potion de poison")
						}
					}
				}

			}
			//
		} else if reponse == 2 {
			Clear()
			Player_PVE2()
		}

	} else if input_pve == 3 {
		Clear()
		fmt.Println("Vous êtes défendu contre les dégâts adverse")
		fmt.Println(" ")
	} else if input_pve == 4 {
		stat.Mana_actuelle = stat.Mana_max
	}
}

func Mob2_attaque() {
	action := rand.Int31n(3)
	if action == 0 {
		if input_pve == 3 {
			fmt.Println("L'attaque à été barré !!")
			fmt.Println(" ")
		} else {
			fmt.Println("attaque ", mob1.nom)
			stat.PV_actuelle = stat.PV_actuelle - mob2.degats + armor.armor
		}
	} else if action == 1 {
		if input_pve == 3 {
			fmt.Println("L'attaque à été barré !!")
			fmt.Println(" ")
		} else {
			fmt.Println("attaque ", mob2.nom)
			fmt.Println(" ")
			stat.PV_actuelle = stat.PV_actuelle - mob2.degats + armor.armor
		}
	} else {
		if input_pve == 1 {
			mob2.pointsDeVieActuels = mob2.pointsDeVieActuels + stat.AD
		} else {
			fmt.Println(mob2.nom, " a utilisé défense")
			fmt.Println(" ")
		}
	}
	PVE_2()
}

//
//

func PVE_3() {
	fmt.Printf("%s entre en combat avec %s!\n", stat.Pseudo, mob2.nom)
	if stat.PV_actuelle > 0 && mob2.pointsDeVieActuels > 0 {
		Player_PVE2()
		Mob2_attaque()
	} else if stat.PV_actuelle <= 0 {
		Clear()
		fmt.Println("Vous êtes mort !!")
	} else {
		stat.xp = stat.xp + 5
		fmt.Println("Vous avez battu ", mob2.nom)
	}
}

func Player_PVE3() {
	var reponse int
	fmt.Println(stat.Pseudo, "                                                                                                 ", mob3.nom)
	fmt.Println("PV :", stat.PV_actuelle, "                                                                                 ", mob3.pointsDeVieActuels)
	fmt.Println("Mana : ", stat.Mana_actuelle)
	fmt.Println("----------------Choix d'Action----------------")
	fmt.Println(" ")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║                 1- ATTAQUE               ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Println(" ")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║                2- INVENTAIRE             ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Println(" ")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║               3- Defense                 ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Println("◢═════════════════════════════════════════◣")
	fmt.Println("║                                          ║")
	fmt.Println("║               4- MEDITATION              ║")
	fmt.Println("║                                          ║")
	fmt.Println("◥═════════════════════════════════════════◤")
	fmt.Scanln(&input_pve)
	if input_pve == 1 {
		Clear()
		fmt.Println("1- Attaque de base")
		fmt.Println("2- compétence")
		fmt.Println("3- Retour")
		fmt.Scanln(&input_attaque)
		if input_attaque == 1 {
			mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - stat.AD
		} else if input_attaque == 2 {
			if input_classe == 3 {
				fmt.Println("1- Sonic Leap")
				fmt.Println("")
				fmt.Println("2- Horizontale Square")
				fmt.Println("")
				fmt.Println("3- Sword QCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - SonicLeap()
				} else if input_skill == 2 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - HorizontaleSquare()
				} else if input_skill == 3 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - SwordQCM()
				}
			} else if input_classe == 1 {
				fmt.Println("1- Tir Charger")
				fmt.Println(" ")
				fmt.Println("2- Explode Arrow")
				fmt.Println(" ")
				fmt.Println("3- Tir QCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - TirCharge()
				} else if input_skill == 2 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - ExplodeArrow()
				} else if input_skill == 3 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - TirQCM()
				}
			} else if input_classe == 2 {
				fmt.Println("1- Spike")
				fmt.Println(" ")
				fmt.Println("2- Javelot")
				fmt.Println(" ")
				fmt.Println("3- SpearQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - Spike()
				} else if input_skill == 2 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - Javelot()
				} else if input_skill == 3 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - SpearQCM()
				}
			} else if input_classe == 4 {
				fmt.Println("1- Xcleaveq")
				fmt.Println(" ")
				fmt.Println("2- TwistedCleave")
				fmt.Println(" ")
				fmt.Println("3- CleaveQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - Xcleaves()
				} else if input_skill == 2 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - TwistedCleave()
				} else if input_skill == 3 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - CleaveQCM()
				}
			} else if input_classe == 5 {
				fmt.Println("1- Mutilation")
				fmt.Println(" ")
				fmt.Println("2- Assassina")
				fmt.Println(" ")
				fmt.Println("3- KinfeQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - Mutilation()
				} else if input_skill == 2 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - Assassina()
				} else if input_skill == 3 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - KnifeQCM()
				}
			} else if input_classe == 6 {
				fmt.Println("1- Bonk")
				fmt.Println(" ")
				fmt.Println("2- Coup de poile")
				fmt.Println(" ")
				fmt.Println("3- MasseQCM")
				fmt.Scanln(&input_skill)
				if input_skill == 1 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - Bonk()
				} else if input_skill == 2 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - Coup_de_poile()
				} else if input_skill == 3 {
					mob3.pointsDeVieActuels = mob3.pointsDeVieActuels - MasseQCM()
				}
			}
		} else if input_attaque == 3 {
			Clear()
			Player_PVE3()
		}
	} else if input_pve == 2 {
		Clear()
		fmt.Println("Voulez vous sortir de l'inventaire")
		fmt.Println("  1- OUI          2- NON")
		fmt.Scanln(&reponse)
		if reponse == 1 {
			Clear()
			//
			var reponse1 int
			var reponse2 int
			fmt.Println(Inventory)
			for k := 0; k < len(Inventory); k++ {
				for i := 0; i < len(Inventory); i++ {
					if Inventory[k] == "potion de 40 PV" {
						fmt.Println("Voulez vous utiliser une potion ?")
						fmt.Println("  1- OUI         2- NON")
						fmt.Scanln(&reponse1)
						if reponse1 == 1 {
							fmt.Println("Vous utiliser une potion de vie 40 PV")
							Takepot_20()
							break
						} else if reponse1 == 2 {
							fmt.Println("Vous n'utiliser pas de potion de vie 40 PV")
							break
						}
					} else if Inventory[i] == "potion de poison" {
						fmt.Println("Voulez vous utiliser une potion de poison ?")
						fmt.Println("  1- OUI         2- NON")
						fmt.Scanln(&reponse2)
						if reponse2 == 1 {
							fmt.Println("Vous utilisez une potion de poison")
						} else if reponse2 == 2 {
							fmt.Println("Vous n'utilisez pas une potion de poison")
						}
					}
				}

			}
			//
		} else if reponse == 2 {
			Clear()
			Player_PVE3()
		}

	} else if input_pve == 3 {
		Clear()
		fmt.Println("Vous êtes défendu contre les dégâts adverse")
		fmt.Println(" ")
	} else if input_pve == 4 {
		stat.Mana_actuelle = stat.Mana_max
	}
}

func Mob3_attaque() {
	action := rand.Int31n(3)
	if action == 0 {
		if input_pve == 3 {
			fmt.Println("L'attaque à été barré !!")
			fmt.Println(" ")
		} else {
			fmt.Println("attaque ", mob3.nom)
			stat.PV_actuelle = stat.PV_actuelle - mob3.degats + armor.armor
		}
	} else if action == 1 {
		if input_pve == 3 {
			fmt.Println("L'attaque à été barré !!")
			fmt.Println(" ")
		} else {
			fmt.Println("attaque ", mob3.nom)
			fmt.Println(" ")
			stat.PV_actuelle = stat.PV_actuelle - mob3.degats + armor.armor
		}
	} else {
		if input_pve == 1 {
			mob3.pointsDeVieActuels = mob3.pointsDeVieActuels + stat.AD
		} else {
			fmt.Println(mob2.nom, " a utilisé défense")
			fmt.Println(" ")
		}
	}
	PVE_3()
}
