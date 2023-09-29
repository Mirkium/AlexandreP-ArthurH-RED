package sao

import (
	"fmt"
	"os"
	"os/exec"
)

type Stat_character struct {
	Pseudo        string
	Classe        string
	level         int
	xp            int
	PV_max        int
	PV_actuelle   int
	AD            int
	Mana_max      int
	Mana_actuelle int
	point_stat    int
	ability       []string
	monnaie       int
}

type Weapon struct {
	bow               string
	spear             string
	rapier            string
	katana            string
	sword             string
	longsword         string
	knife             string
	axe               string
	mace              string
	durability_weapon int
	stat_bonus        string
}

type Armor struct {
	armor            int
	head_armor       string
	chest_armor      string
	leg_armor        string
	foot_armor       string
	durability_heat  int
	durability_chest int
	durability_leg   int
	durability_foot  int
	stat_bonus       string
}

var Equipement1 [4]string
var Equipement2 [2]string
var Inventory [10]string

var input_classe int
var input_player string
var input_epee int
var input_axe int

var armor Armor
var weapon Weapon
var stat Stat_character

// =============================================================================================================================================
func Create_Player() {
	fmt.Println("Bonjour player et bienvenu sur SAO ! Comment vous appeler vous ? ")
	fmt.Scanln(&input_player)
	stat.Pseudo = input_player
	stat.level = 1
	fmt.Println("Salutation", stat.Pseudo, "veuillez choisr votre classe d'arme : ", "\n")
	Choose_Classe()
}

// =============================================================================================================================================
func Choose_Classe() {
	fmt.Println("Vous avez les classes :", "\n")
	fmt.Println("   O    |----------------------------------------------------------------------------------------------------------------------------------|    O")
	fmt.Println("   ▒    | 1- archer : Peut utiliser l'arc ou l'arbalète comme arme.                                                                        |    ▒")
	fmt.Println("   ▒    | statistique de base :   PV_max = 50   //   AD = 15   //   Mana = 150                                                             |    ▒")
	fmt.Println("o==▓==o |----------------------------------------------------------------------------------------------------------------------------------| o==▓==o")
	fmt.Println("   █    | 2- Lancier : Utilise la lance ou la halbarde comme arme, La lance permet de réduire l'armure en fonction de la stat Perfo_armor. |    █")
	fmt.Println("   █    | statistique de base :   PV_max = 80   //   AD = 13   //   Mana = 100                                                             |    █")
	fmt.Println("   █    |----------------------------------------------------------------------------------------------------------------------------------|    █")
	fmt.Println("   █    | 3- épéiste : Peut utiliser comme arme la rapière, le katana, l'épée à une main et l'épée à deux main.                            |    █")
	fmt.Println("   ▼    | statistique de base :   PV_max = 100   //   AD = 10   //   Mana = 100                                                            |    ▼")
	fmt.Println("  ynov  |----------------------------------------------------------------------------------------------------------------------------------|   ynov")
	fmt.Println("   ▲    | 4- hacheur : L'arme utilisée est la hache rapide ou hache lourde.                                                                |    ▲")
	fmt.Println("   █    | statistique de base :   PV_max = 120   //   AD = 11   //   Mana = 100                                                            |    █")
	fmt.Println("   █    |----------------------------------------------------------------------------------------------------------------------------------|    █")
	fmt.Println("   █    | 5- Joueur de couteau : Utilisateur des dagues et couteau qui peut avoir des propriétées de saignement.                           |    █")
	fmt.Println("   █    | statistique de base :   Pv_max = 50   //   AD = 14   //   Mana =150                                                              |    █")
	fmt.Println("o==▓==o |----------------------------------------------------------------------------------------------------------------------------------| o==▓==o")
	fmt.Println("   ▒    | 6- masseur : Utilise et manie les masses et marteaux d'armes pouvant étourdir les ennemies.                                      |    ▒")
	fmt.Println("   ▒    | statistique de base :   PV_max = 130   //   AD = 9   //   Mana = 1                                                               |    ▒")
	fmt.Println("   O    |----------------------------------------------------------------------------------------------------------------------------------|    O")
	fmt.Scanln(&input_classe)
	Clear()
	Inventory[0] = "pain"
	Inventory[1] = "potion de 40 PV"
	Inventory[2] = "potion de 40 PV"
	if input_classe == 1 {
		stat.Classe = "archer"
		stat.PV_max = 50
		stat.PV_actuelle = 50
		stat.AD = 15
		stat.Mana_max = 150
		stat.Mana_actuelle = 150
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		Equipement()
	} else if input_classe == 2 {
		stat.Classe = "lancier"
		stat.PV_max = 80
		stat.PV_actuelle = 80
		stat.AD = 13
		stat.Mana_max = 100
		stat.Mana_actuelle = 100
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		Equipement()
	} else if input_classe == 3 {
		stat.Classe = "épéiste"
		stat.PV_max = 100
		stat.PV_actuelle = 100
		stat.AD = 10
		stat.Mana_max = 150
		stat.Mana_actuelle = 150
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		Equipement()
	} else if input_classe == 4 {
		stat.Classe = "hacheur"
		stat.PV_max = 120
		stat.PV_actuelle = 120
		stat.AD = 11
		stat.Mana_max = 100
		stat.Mana_actuelle = 100
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		Equipement()
	} else if input_classe == 5 {
		stat.Classe = "joueur de couteau"
		stat.PV_max = 50
		stat.PV_actuelle = 50
		stat.AD = 14
		stat.Mana_max = 150
		stat.Mana_actuelle = 150
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		Equipement()
	} else if input_classe == 6 {
		stat.Classe = "masseur"
		stat.PV_max = 130
		stat.PV_actuelle = 130
		stat.AD = 9
		stat.Mana_max = 100
		stat.Mana_actuelle = 100
		stat.xp = 0
		stat.monnaie = 0
		armor.head_armor = "casque de départ"
		armor.chest_armor = "buste de départ"
		armor.leg_armor = "pantalon de départ"
		armor.foot_armor = "botte de départ"
		armor.durability_heat = 40
		armor.durability_chest = 70
		armor.durability_leg = 60
		armor.durability_foot = 50
		Equipement()
	} else {
		fmt.Println("Tu as pas rentrer le numéro de l'une des classe")
		Choose_Classe()
	}
}

// =============================================================================================================================================
func Fiche_perso() {
	var input_fermer int
	fmt.Println("                                ~~~~~~FICHE PERSONNAGE~~~~                                            ")
	fmt.Println("╔═════════════════════════════════════════╦══════════════════════════════════════════════════════════╗")
	fmt.Println("║---------------Player--------------------║-----------------------EQUIPEMENT-------------------------║")
	fmt.Print("║ Name   :", stat.Pseudo)
	for i := 0; i < 32-len(stat.Pseudo); i++ {
		fmt.Print(" ")
	}
	fmt.Print("║ head   :", Equipement1[0])
	for i := 0; i < 50-len(Equipement1[0]); i++ {
		fmt.Print(" ")
	}
	fmt.Println("║")
	fmt.Print("║ Classe :", stat.Classe)
	for i := 0; i < 32-len(stat.Classe); i++ {
		fmt.Print(" ")
	}
	fmt.Print("║ chest  :", Equipement1[1])
	for i := 0; i < 22-len(Equipement1[1]); i++ {
		fmt.Print(" ")
	}
	fmt.Print(" hand : ", Equipement2[0])
	for i := 0; i < 21-len(Equipement2[0]); i++ {
		fmt.Print(" ")
	}
	fmt.Println("║")
	fmt.Print("║ Level  :", stat.level)
	for i := 0; i < 30; i++ {
		fmt.Print(" ")
	}
	fmt.Print("║ leg    :", Equipement1[2])
	for i := 0; i < 22-len(Equipement1[2]); i++ {
		fmt.Print(" ")
	}
	fmt.Print(" hand : ", Equipement2[1])
	for i := 0; i < 21-len(Equipement2[1]); i++ {
		fmt.Print(" ")
	}
	fmt.Println("║")
	fmt.Print("║ xp     :", stat.xp)
	for i := 0; i < 32; i++ {
		fmt.Print(" ")
	}
	fmt.Print("║ foot   :", Equipement1[3])
	for i := 0; i < 21-len(Equipement1[3]); i++ {
		fmt.Print(" ")
	}
	fmt.Println("║")
	fmt.Println("╠═════════════════════════════════════════╬══════════════════════════════════════════════════════════╣")
	fmt.Println("║-------------STATISTIQUE-----------------║-----------------------COMPETENCE-------------------------║")
	fmt.Print("║ Pv max      : ", stat.PV_max)
	for i := 0; i < 32; i++ {
		fmt.Print(" ")
	}
	fmt.Print("║", stat.ability)
	for i := 0; i < 32-len(stat.ability); i++ {
		fmt.Print(" ")
	}
	fmt.Println("║")
	fmt.Print("║ Pv actuelle : ", stat.PV_actuelle)
	for i := 0; i < 32; i++ {
		fmt.Print(" ")
	}
	fmt.Println("║                                                          ║")
	fmt.Print("║ AD          : ", stat.AD)
	for i := 0; i < 32; i++ {
		fmt.Print(" ")
	}
	fmt.Println("║                                                          ║")
	fmt.Print("║ Mana     : ", stat.Mana_max)
	for i := 0; i < 32; i++ {
		fmt.Print(" ")
	}
	fmt.Println("║ Mana act : ", stat.Mana_actuelle)
	for i := 0; i < 32; i++ {
		fmt.Print(" ")
	}
	fmt.Println("║                                                          ║")
	fmt.Println("╠═════════════════════════════════════════╩══════════════════════════════════════════════════════════╣")
	fmt.Println("║--------------------------------------------INVENTORY-----------------------------------------------║")
	fmt.Print("║", Inventory)
	for i := 0; i < 99-len(Inventory); i++ {
		fmt.Print(" ")
	}
	fmt.Println("                                                                                     ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println(" ")
	fmt.Println("1- pour utiliser les points de stats sinon; sinon, 0 pour fermer l'inventaire")
	fmt.Scanln(&input_fermer)
	if input_fermer == 1 {
		fmt.Print("Vous utiliser les stats")
	} else if input_fermer == 0 {
		fmt.Println("Vous sortez de l'inventaire")
	} else {
		fmt.Println("Rentrez 1 ou 0")
		Fiche_perso()
	}
}

// =============================================================================================================================================
func Equipement() {
	if input_classe == 1 {
		weapon.bow = "arc de départ"
		weapon.stat_bonus = "+3 AD  //  +10 PV"
		stat.AD = stat.AD + 3
		stat.PV_max = stat.PV_max + 10
		armor.stat_bonus = "set de départ : +25 PV max  //  +4 armor"
		stat.PV_max = stat.PV_max + 25
		armor.armor = 4
		Equipement1[0] = armor.head_armor
		Equipement1[1] = armor.chest_armor
		Equipement1[2] = armor.leg_armor
		Equipement1[3] = armor.foot_armor
		Equipement2[0] = weapon.bow
		Fiche_perso()
	}
	if input_classe == 2 {
		weapon.spear = "lance de départ"
		weapon.stat_bonus = "+4 AD  //  +10 PV max"
		stat.AD = stat.AD + 4
		stat.PV_max = stat.PV_max + 10
		armor.stat_bonus = "set de départ : +25 PV max  //  +4 armor"
		stat.PV_max = stat.PV_max + 25
		armor.armor = 4
		Equipement1[0] = armor.head_armor
		Equipement1[1] = armor.chest_armor
		Equipement1[2] = armor.leg_armor
		Equipement1[3] = armor.foot_armor
		Equipement2[0] = weapon.spear
		Fiche_perso()
	}
	if input_classe == 3 {
		Voix_Epee()
	}

	if input_classe == 4 {
		Clear()
		Voix_hache()
	}
	if input_classe == 5 {
		weapon.knife = "dague de départ"
		weapon.stat_bonus = "+3 AD  //  +20 Mana  //  +10 PV"
		stat.AD = stat.AD + 3
		stat.Mana_max = stat.Mana_max + 20
		stat.PV_max = stat.PV_max + 10
		armor.stat_bonus = "set de départ : +25 PV max  //  +4 armor"
		stat.PV_max = stat.PV_max + 25
		armor.armor = 4
		Equipement1[0] = armor.head_armor
		Equipement1[1] = armor.chest_armor
		Equipement1[2] = armor.leg_armor
		Equipement1[3] = armor.foot_armor
		Equipement2[0] = weapon.knife
		Fiche_perso()
	}
	if input_classe == 6 {
		weapon.mace = "masse de départ"
		weapon.stat_bonus = "+2 AD  //  +30 PV max"
		stat.AD = stat.AD + 2
		stat.PV_max = stat.PV_max + 30
		armor.stat_bonus = "set de départ : +25 PV max  //  +4 armor"
		stat.PV_max = stat.PV_max + 25
		armor.armor = 4
		Equipement1[0] = armor.head_armor
		Equipement1[1] = armor.chest_armor
		Equipement1[2] = armor.leg_armor
		Equipement1[3] = armor.foot_armor
		Equipement2[0] = weapon.mace
		Fiche_perso()
	}
	fmt.Println(Equipement1, " ", Equipement2)
}

// =============================================================================================================================================
func Voix_Epee() {
	fmt.Println("o∷∷{=============▻ Choisiser votre voix de l'épée <ΞΞΞΞΞΞΞΞΞ{o}≠≠≠O: , ,  et ")
	fmt.Println("           ⸸-----------------------Cyril------------------⸸")
	fmt.Println("           |               1 : la rapière                 |")
	fmt.Println("           ⸸----------------------Data-IA-----------------⸸")
	fmt.Println("  ")
	fmt.Println("           🗡----------------------Ethan------------------🗡")
	fmt.Println("           |               2 : le katana                  |")
	fmt.Println("           🗡--------------------Game-Prod----------------🗡")
	fmt.Println("   ")
	fmt.Println("           ⚔️---------------------Kheir------------------⚔️")
	fmt.Println("           |               3 : l'épée à une main          |")
	fmt.Println("           ⚔️--------------------Data-IA-----------------⚔️")
	fmt.Println("  ")
	fmt.Println("           🗡️--------------------Allan-------------------🗡️")
	fmt.Println("           |               4 : l'épée à deux main         |")
	fmt.Println("           🗡️---------------------Web--------------------🗡️")
	fmt.Scanln(&input_epee)
	if input_epee == 1 {
		weapon.rapier = "rapier de départ"
		weapon.stat_bonus = "+5 AD  //  +20 Mana"
		stat.AD = stat.AD + 5
		stat.Mana_max = stat.Mana_max + 20
		armor.stat_bonus = "set de départ : +25 PV max  //  +4 armor"
		stat.PV_max = stat.PV_max + 25
		armor.armor = 4
		Equipement1[0] = armor.head_armor
		Equipement1[1] = armor.chest_armor
		Equipement1[2] = armor.leg_armor
		Equipement1[3] = armor.foot_armor
		Equipement2[0] = weapon.rapier
		Clear()
		Fiche_perso()
	} else if input_epee == 2 {
		weapon.katana = "katana de départ"
		weapon.stat_bonus = "+4 AD  //  +10 Mana  //  +10 PV max"
		stat.AD = stat.AD + 4
		stat.Mana_max = stat.Mana_max + 10
		stat.PV_max = stat.PV_max + 10
		armor.stat_bonus = "set de départ : +25 PV max  //  +4 armor"
		stat.PV_max = stat.PV_max + 25
		armor.armor = 4
		Equipement1[0] = armor.head_armor
		Equipement1[1] = armor.chest_armor
		Equipement1[2] = armor.leg_armor
		Equipement1[3] = armor.foot_armor
		Equipement2[0] = weapon.katana
		Clear()
		Fiche_perso()
	} else if input_epee == 3 {
		weapon.sword = "épée à une main de départ"
		weapon.stat_bonus = "+4 AD  //  +20 PV max"
		stat.AD = stat.AD + 4
		stat.PV_max = stat.PV_max + 20
		armor.stat_bonus = "set de départ : +25 PV max  //  +4 armor"
		stat.PV_max = stat.PV_max + 25
		armor.armor = 4
		Equipement1[0] = armor.head_armor
		Equipement1[1] = armor.chest_armor
		Equipement1[2] = armor.leg_armor
		Equipement1[3] = armor.foot_armor
		Equipement2[0] = weapon.sword
		Clear()
		Fiche_perso()
	} else if input_epee == 4 {
		weapon.longsword = "épée à deux main de départ"
		weapon.stat_bonus = "+3 AD  //  +15 PV max"
		stat.AD = stat.AD + 3
		stat.PV_max = stat.PV_max + 15
		armor.stat_bonus = "set de départ : +25 PV max  //  +4 armor"
		stat.PV_max = stat.PV_max + 25
		armor.armor = 4
		Equipement1[0] = armor.head_armor
		Equipement1[1] = armor.chest_armor
		Equipement1[2] = armor.leg_armor
		Equipement1[3] = armor.foot_armor
		Equipement2[0] = weapon.longsword
		Clear()
		Fiche_perso()
	} else {
		fmt.Println("Rentrer l'un des numéros corespondant à une voix de l'épée")
		Clear()
		Voix_Epee()
	}
}

func Voix_hache() {
	fmt.Println("Pour chosir votre voix du maniement de la hache, entrer 1 pour la hache lourde et 2 pour la hache rapide : ")
	fmt.Scanln(&input_axe)
	if input_axe == 1 {
		weapon.axe = "hache lourde de départ"
		weapon.stat_bonus = "+4 AD  //  +25 PV max"
		stat.AD = stat.AD + 4
		stat.PV_max = stat.PV_max + 25
		armor.stat_bonus = "set de départ : +25 PV max  //  +4 armor"
		stat.PV_max = stat.PV_max + 25
		armor.armor = 4
		Equipement1[0] = armor.head_armor
		Equipement1[1] = armor.chest_armor
		Equipement1[2] = armor.leg_armor
		Equipement1[3] = armor.foot_armor
		Equipement2[0] = weapon.axe
		Clear()
		Fiche_perso()
	} else if input_axe == 2 {
		weapon.axe = "hache rapide de départ"
		weapon.stat_bonus = "+7 AD  //  +10 Mana"
		stat.AD = stat.AD + 7
		stat.Mana_max = stat.Mana_max + 10
		armor.stat_bonus = "set de départ : +25 PV max  //  +4 armor"
		stat.PV_max = stat.PV_max + 25
		armor.armor = 4
		Equipement1[0] = armor.head_armor
		Equipement1[1] = armor.chest_armor
		Equipement1[2] = armor.leg_armor
		Equipement1[3] = armor.foot_armor
		Equipement2[0] = weapon.axe
		Clear()
	} else {
		fmt.Println("ERREUR")
		Clear()
		Voix_hache()
	}
}

// =============================================================================================================================================================
func Quete() {
	var string_quest int
	fmt.Println("Choisiser une quête : ", "\n")
	fmt.Println("[]====================tableau=des=quêtes====================[]")
	fmt.Println("∐               1 -Donjons de l'arche perdu- (D)            ∐")
	fmt.Println("∐                      INDISPONIBLE                         ∐")
	fmt.Println("∐                      INDISPONIBLE                         ∐")
	fmt.Println("∐                      INDISPONIBLE                         ∐")
	fmt.Println("∐                      INDISPONIBLE                         ∐")
	fmt.Println("∐                      INDISPONIBLE                         ∐")
	fmt.Println("[]==========================================================[]", "\n")
	fmt.Scanln(&string_quest)
	if string_quest == 1 {
		fmt.Println("Bonne chance pour votre quête !")
		fmt.Println("Mais je vous conseille de vous entrainer sur le terrain d'entrainement au combat avant.")
		fmt.Println("Vous pouvez aussi visiter la ville si vous le voulez aussi. ")
		Clear()
		fmt.Println("Sur ce bonne chance pour la conquête de 'Sword Art Online' !")
		Dongeons_de_larche_perdue()
	} else if string_quest == 0 {
		fmt.Println("Aurevoir et penser à venir ici pour des quêtes de grandes envergure !")
	} else {
		fmt.Println("Faite un choix sinon l'admin sera pas contente !")
		Quete()
	}

}

func Quete1() {
	Clear()
	fmt.Println("Maintenant aventurier ", stat.Pseudo, "Vous pouvez accéder à vos compétence que vous poccéder liéer à votre classe", "\n")
	if input_classe == 1 {
		stat.ability = append(stat.ability, "Tir charger ,", "Explode Arrow ,", "Tir QCM ,")
	} else if input_classe == 3 {
		stat.ability = append(stat.ability, "Sonic Leap ,", "Horizontale Scare ,", "Sword QCM ,")
	} else if input_classe == 2 {
		stat.ability = append(stat.ability, "Spike ,", "Javelot ,", "SpearQCM ,")
	} else if input_classe == 4 {
		stat.ability = append(stat.ability, "Xcleaves ,", "TwistedCleave", "CleaveQCM")
	} else if input_classe == 5 {
		stat.ability = append(stat.ability, "Mutilation ,", "Assassina ,", "KnifeQCM")
	} else if input_classe == 6 {
		stat.ability = append(stat.ability, "Bonk ,", "Coup de poile ,", "MasseQCM ,")
	}
	Fiche_perso()
}

// =======================================================================================================================================================

// =============================================================================================================================================
func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
