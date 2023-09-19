package main

import "fmt"

type Personnage struct {
	nom string
	classe string
	level int
	pvmax int
	pv int
	inventaire map[string]int
}

func  main() {
	fmt.Printf("test")
	var p1 Personnage
	p1.Init("Alan", "magicien", 1, 1000000, 100, map[string]int{"potion": 3})
	p1.DisplayInfo()
	p1.TakePot()
	p1.DisplayInfo()
}

func (p *Personnage) Init(nom string, classe string, level int, pvmax int, pv int, inventaire map[string]int) {
	p.nom = nom
	p.classe = classe
	p.level = level 
	p.pvmax = pvmax
	p.pv = pv
	p.inventaire = inventaire
}

func (p Personnage) DisplayInfo() {
	fmt.Println("Nom :", p.nom)
	fmt.Println("Classe", p.classe)
	fmt.Println("Level", p.level)
	fmt.Println("pvmax", p.pvmax)
	fmt.Println("pvcurrent", p.pv)
	p.AccessInventory()
}

func (p Personnage) AccessInventory() {
	for item, nb := range p.inventaire {
		fmt.Println(item)
		fmt.Println(nb)
	}
}

func (p *Personnage) TakePot() {
	if p.inventaire["potion"] > 0 {
		p.inventaire["potion"]--
		fmt.Println("potion de soins utilisée")
		p.pv = p.pv +50
	} else if p.pv >= p.pvmax{
		fmt.Println("tes pv sont déja au max")
	} else if p.pv >= p.pvmax {
		surplus := p.pv - p.pvmax
		fmt.Printf("etes vous sur de vouloir utiliser une potion, elle vous rendra %d pv de trop. (oui / non) \n", surplus)
		var choice string
		fmt.Scan(&choice)
		if choice == "oui" {
			p.pv = p.pvmax
		} else if choice == "non" {
			fmt.Printf("Vous n'avez pas pris de potion de soin. vous avez %d pv \n", p.pv)
		}
	}
}

func (p *Personnage) Menu() {
	var choice string
	fmt.Println("1) Afficher les informations du personnage")
	fmt.Println("2) Afficher l'inventaire")
	fmt.Println("3) Allez au marchand")
	fmt.Println("4) ")
	fmt.Scan(choice)

	switch choice {
	case "1":
		p.DisplayInfo()
	case "2": 
		p.AccessInventory()
	case "3":
		fmt.Println()
	default:
		p.Menu()
	}
}
