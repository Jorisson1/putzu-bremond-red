package main

import (
	"fmt"
	"time"
)

type personnage struct {
	Nom        string
	Classe     string
	Niveau     int
	PVmax      int
	PVactuel   int
	Inventaire map[string]int
	Skills []string
}

func (p *personnage) Init(Nom string, Classe string, Niveau string, PVmax int, PVactuel int, Inventaire map[string]int, Skills [] 
	string) {
	var p1 personnage

	p1.Nom = "kawaki"
	p1.Classe = "Slayer"
	p1.Niveau = 1
	p1.PVmax = 100
	p1.PVactuel = 40
	p1.Inventaire = map[string]int{"potion": 3, "katana": 1}
	p1.Skills = []string{"coup de poing"}

}

func (p personnage) DisplayInfo() {
	fmt.Println("Nom :", p.Nom)
	fmt.Println("Classe :", p.Classe)
	fmt.Println("Niveau :", p.Niveau)
	fmt.Println("PVmax :", p.PVmax)
	fmt.Println("PVactuel :", p.PVactuel)
	fmt.Println("Inventaire :", p.Inventaire)

}

func (p* personnage) accesInventory() {
	for item, nb := range p.Inventaire {
		fmt.Println(item)
		fmt.Println(nb)

	}
}
func (p *personnage) takepot() {
	if p.Inventaire["potion"] > 0 {
		p.Inventaire["potion"]--
		fmt.Println("potion de soins utilisée")
		p.PVactuel = p.PVactuel + 50
	} else if p.PVactuel >= p.PVmax {
		fmt.Println("tes PV sont déja au max")
	} else if p.PVactuel >= p.PVmax {
		surplus := p.PVactuel - p.PVmax
		fmt.Printf("Etes vous sur de vouloir utiliser une potion, elle vous rendra %d Pv de trop. (Oui / Non) \n", surplus)
		var choice string
		fmt.Scan(&choice)
		if choice == "Oui" {
			p.PVactuel = p.PVmax
		} else if choice == "Non" {
			fmt.Printf("Vous n'avez pas pris de potion de soin. Vous avez %d PV \n", p.PVactuel)
		}

	}
		 
}

type boutique struct {
	item map[string] int
	price int
	quantity int
}
func (m* boutique) Marchand() {
	fmt.Println("Bonjour aventurier, souhaites tu voir mon stock ?")
	fmt.Println("Tu as gagné un potion de soin gratuite")
	addInventory = p.potion (1)
}

func (m* boutique) addInventory() {
		fmt.Printf("Item ajouté a ton inventaire")
}

func (m* boutique) removeInventory() {
	fmt.Printf("Item supprimé de ton inventaire")

}

func (p* personnage) Dead() {
	if p.PVactuel <= 0 
	fmt.Printf("Tu es mort") 
}

func (p* personnage) Revive() {
	var Revive string
	Revive == p.PVactuel + 50
	fmt.Printf("Tu as été réanimé")
}

func (p* personnage) poisonPot() {
	for poisonPot --10 p.PVactuel
		time.Sleep(1* time.Second)
		time.Sleep(1* time.Second)
		time.Sleep(1* time.Second)	
}

func (p* personnage) Livre de sort(skills string) {
	p.skill = append(p.skill, skills)

} 
func (p* personnage) spellBook() {
	spellToAdd := "Boule de feu"
	for _, skill := range c.Skill {
		if skill == spellToAdd {
			fmt.Println("Vous avez déjà appris le sort", spellToAdd)
			return
		}
	}
	c.Skill = append(c.Skill, spellToAdd)
	fmt.Println("Vous avez appris le sort", spellToAdd)
}

func (c *Character) AddToInventory(item string) {
	c.Inventory = append(c.Inventory, item)
	player := NewCharacter("NomDuPersonnage")
	player.AddToInventory("Livre de Sort : Boule de Feu")
	for _, item := range player.Inventory {
		if item == "Livre de Sort : Boule de Feu" {
			player.spellBook()
			player.RemoveFromInventory(item)
		}
	}

	fmt.Println("Compétences du personnage:", player.Skill)
}

type Monstre struct {
Nom string
PVmax int
PVactuel int
Dégats int
}

func(m* Monstre) InitGoblin() {
	m.Nom = "Mannequin"
	m.PVmax = 30
	m.PVactuel = 30
	m.Dégats = 10

}

func(p* personnage) TrainingFight() {
	m := Monstre {"Mannequin", 30 , 30 , 10}
	fmt.Println("Bonjour voyageur, que voulez-vous faire ?")
	fmt.Println("1: Combattre le mannequin d'entrainement")
	fmt.Println("2: Quitter l'entrainement")
	fmt.Printf("Entre ton choix")
	fmt.Scan(&choice)
	fmt.Println(" ")

	switch choice {
	case "1": 
		
	}
}


