package sao

import "fmt"

type Mob1 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
	degats             int
}

var mob1 Mob1
var mob2 Mob2
var mob3 Mob3
var mob4 Mob4
var mob5 Mob5

func InitMob1() {
	mob1.nom = "apprenti chevalier"
	mob1.classe = "Traitre"
	mob1.niveau = 1
	mob1.pointsDeVieMax = 50
	mob1.pointsDeVieActuels = 50
	mob1.inventaire = map[string]int{"épée fragile": 1, "casque": 1}
	mob1.degats = 10
	for i:= 0; i<4; i++ {
		fmt.Println("Un autre ", mob1.nom)
		PVE_1()
	}
	InitMob8()
}

type Mob2 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
	degats             int
}

func InitMob8() {
	mob2.nom = "chevalier déchu"
	mob2.classe = "Traitre"
	mob2.niveau = 5
	mob2.pointsDeVieMax = 75
	mob2.pointsDeVieActuels = 75
	mob2.inventaire = map[string]int{"épée en pierre": 1, "casque": 1, "bouclier": 1}
	mob2.degats = 20
	PVE_2()
}

type Mob3 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
	degats             int
}

func InitMob3() {
	mob3.nom = "sergent a double face"
	mob3.classe = "Traitre"
	mob3.niveau = 10
	mob3.pointsDeVieMax = 100
	mob3.pointsDeVieActuels = 100
	mob3.inventaire = map[string]int{"Lance": 1, "casque": 1, "bouclier": 1}
	mob3.degats = 30
	PVE_3()
}

type Mob4 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
	degats             int
}

func InitMob14() {
	mob4.nom = "Capitaine mort"
	mob4.classe = "Traitre"
	mob4.niveau = 15
	mob4.pointsDeVieMax = 125
	mob4.pointsDeVieActuels = 125
	mob4.inventaire = map[string]int{"épée en fer": 1, "casque": 1, "bouclier": 1}
	mob4.degats = 40
}

type Mob5 struct {
	nom                string
	classe             string
	niveau             int
	pointsDeVieMax     int
	pointsDeVieActuels int
	inventaire         map[string]int
	degats             int
}

func InitMob5() {
	mob5.nom = "roi de l'ombre"
	mob5.classe = "Traitre"
	mob5.niveau = 20
	mob5.pointsDeVieMax = 150
	mob5.pointsDeVieActuels = 150
	mob5.inventaire = map[string]int{"épée en obsidienne": 2, "Armure complète": 1, "potion": 1}
	mob5.degats = 50
}
