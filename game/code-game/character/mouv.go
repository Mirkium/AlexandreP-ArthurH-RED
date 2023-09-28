package sao 

import (
	"fmt"
)

func  Guilde_for_ability() {
	var input_yes_or_no int
	fmt.Println("Aller voir la guilde pour commencer les quêtes et débloquer les skills de classes")
		fmt.Println("        1- OUI                             0- NON")
		fmt.Scanln(&input_yes_or_no)
		if input_yes_or_no == 1 {
			Quete1()
		} else if input_yes_or_no == 0 {
			Clear()
			fmt.Print("Ok")
		} else {
			fmt.Println("Rentrez 1 ou 0 pour choisir")
			Guilde_for_ability()
		}
}


func Dongeons_de_larche_perdue() {
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
	fmt.Println("                                                                                                               ")
}