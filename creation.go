package thanos

import "fmt"


type Personnage struct {
	nom string
	classe string
	level int
	pvmax int
	pvcurrent int
	inventaire map[string]int
}

func Init() {
	var p1 Personnage

	p1.nom = "ainz"
	p1.classe = "magicien"
	p1.level = 1  
	p1.pvmax = 1000000
	p1.pvcurrent = 100
	p1.inventaire = map[string]int{"potion": 3, }
}

func (p Personnage) displayInfo() {
	fmt.Println("Nom :", p.nom)
	fmt.Println("Classe", p.classe)
	fmt.Println("Level", p.level)
	fmt.Println("pvmax", p.pvmax)
	fmt.Println("pvcurrent", p.pvcurrent)
	fmt.Println("inventaire", p.inventaire)
}

func (p Personnage) accessInventory() {
	for item, nb := range p.inventaire {
		fmt.Println(item)
		fmt.Println(nb)
	}
}
