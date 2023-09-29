package sao

import (
	"fmt"
	"time"
)

func Guilde_for_ability() {
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
	var input int

	for i := 0; i < 3; i++ {
		fmt.Println("CHARGEMENT : ", "\n")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔ ")
		Clear()
	}
	fmt.Println("              █                                                                       █                          ")
	fmt.Println("              █                                                                       █                          ")
	fmt.Println("              █                                                      ├┴┬              █                          ")
	fmt.Println("              █                                                                       █                          ")
	fmt.Println("              █          ┬─├                                                          █                          ")
	fmt.Println("              █                              ████████████                             █                                      ")
	fmt.Println("              █                           ███     ▒▒     ███                          █                              ")
	fmt.Println("              █                        ███        ▒▒        ███                       █                             ")
	fmt.Println("              █                    ████           ▒▒           ████                   █                             ")
	fmt.Println("              █                   █               ▒▒               █                  █                            ")
	fmt.Println("              █                 █                 ▒▒                █                 █                            ")
	fmt.Println("              █                █                  ▒▒                 █                █                            ")
	fmt.Println("              █                █                  ▒▒         ▓        █               █                             ")
	fmt.Println("              █               █           //      ▒▒         ▓        █               █                           ")
	fmt.Println("              █               █                   ▒▒                  █               █                            ")
	fmt.Println("              █               █                   ▒▒         /        █         ┬     █                            ")
	fmt.Println("              █               █                   ▒▒                  █        ├┴     █                            ")
	fmt.Println("              █              █           ▓        ▒▒                   █              █                            ")
	fmt.Println("              █              █                    ▒▒                   █              █                            ")
	fmt.Println("              █              █                    ▒▒      ▓            █              █                            ")
	fmt.Println("              █              █   /                ▒▒                   █              █                            ")
	fmt.Println("              █              █                    ▒▒                   █              █                            ")
	fmt.Println("              █              █                    ▒▒                   █              █                            ")
	fmt.Println("              █              █                   θ▒▒θ          ▓       █              █                             ")
	fmt.Println("              █              █                    ▒▒                   █              █                            ")
	fmt.Println("              █              █                    ▒▒                   █              █                            ")
	fmt.Println("              █              █               /    ▒▒                   █              █                           ")
	fmt.Println("              █              █                    ▒▒                   █              █                            ")
	fmt.Println("              █              █                    ▒▒                   █              █                            ")
	fmt.Println("              █              █                    ▒▒                   █              █                            ")
	fmt.Println("              █              █                    ▒▒                   █              █                     ")
	fmt.Println(" ")
	fmt.Println("Vous êtes devant le donjons 'L'arche Perdue'. Voulez le commencer ?")
	fmt.Println("   1- OUI                                    0-NON")
	fmt.Scanln(&input)
	if input == 1 {
		Clear()
		fmt.Println("Vous entrez dans le donjons")
		time.Sleep(1 * time.Second)
		Clear()
		for i := 0; i < 3; i++ {
			fmt.Println("CHARGEMENT : ", "\n")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("⚔ ")
			Clear()
		}
		InitMob1()
		InitMob8()
		InitMob3()
	} else if input == 0 {
		fmt.Println("Vous retournez à la ville")
	}
}





