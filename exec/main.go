package main

import ("fmt"
"unicode"
"strings"
"strconv"
"bufio"
"os"
)

type item struct {
	index int
	name string
	quantite int
	price int
}

type Inventaire struct {
    objets string
    capaciteMax int
    utilisations int
}

type Personnage struct {
	nom string
	classe string
	level int
	pvmax int
	pv int
	inventaire map[string]int
	Money int
	inventairecount int
	up Inventaire
	golds int
	equipement equipement
}

type equipement struct {
	EquipementTête string
	EquipementTorse string
	EquipementPieds string

}

func main() {
	var p1 Personnage
	p1.CharCreation()
	p1.Menu()
}

func (p *Personnage) Init(nom string, classe string, level int, pvmax int, pv int, inventaire map[string]int, gold int) {
	p.nom = nom
	p.classe = classe
	p.level = level 
	p.pvmax = pvmax
	p.pv = pv
	p.inventaire = inventaire
	p.golds = gold
}

func (p Personnage) DisplayInfo() {
	fmt.Println("Nom :", p.nom)
	if p.classe == "1"{
		fmt.Println("Classe : Demon")
	} else if p.classe == "2"{
		fmt.Println("Classe : Lampadaire")
	} else if p.classe == "3"{
		fmt.Println("Classe : Capybara")
	}

	fmt.Println("Level", p.level)
	fmt.Println("pvmax", p.pvmax)
	fmt.Println("pvcurrent", p.pv)
	fmt.Println("gold :", p.golds)
	fmt.Println("Pour quitter le menu appuyer sur 1")
	var choice int
	fmt.Scan(&choice)
	fmt.Println(choice)
	switch choice {
	case 1:
		fmt.Print("\033[H\033[2J")
		p.Menu()
	case 2 :
		fmt.Println("Pour quitter le menu appuyer sur 1")
	default :
	fmt.Println("commande imcomprise")
	}
	
}


func (p Personnage) AccessInventory() {
	for item, nb := range p.inventaire {
		fmt.Println(item, ":", nb)
		fmt.Println("")
	}
	fmt.Println("Pour quitter le menu appuyer sur 1")
	fmt.Println("Appuyer sur 2 pour equiper l'equipement")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Print("\033[H\033[2J")
		p.Menu()
	case 2 :
		if keyExists("Chapeau de l'aventurier", p.inventaire){
			p.pvmax += 50
			p.pv += 50
			p.removeInventory2("Chapeau de l’aventurier")
			p.removeInventory2("Tunique de l’aventurier")
			p.removeInventory2("Bottes de l’aventurier")
			fmt.Print("\033[H\033[2J")
			fmt.Println("Vous avez équipé l'équipement")
			p.Menu()
		} else{
			fmt.Print("\033[H\033[2J")
			fmt.Println("Vous avez déjà equiper cet équipement")
			p.Menu()
		}
	}
}


func (p Personnage)MyScan(input int) {
	fmt.Scan(&input)
	fmt.Println(input)
	switch input {
	case 1:
		fmt.Print("\033[H\033[2J")
		p.Menu()
	case 2 :
		fmt.Println("Pour quitter le menu appuyer sur 1")
	default :
	fmt.Println("commande imcomprise")
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

func capitalizeString(s string) string {
    if s == "" {
        return s // Retourne la chaîne vide si elle est vide
    }

    // Met la première lettre en majuscule et le reste en minuscules
    return strings.Title(strings.ToLower(s))
}

func (p *Personnage)CharCreation(){
	var name string
	var classe string	
	fmt.Println("Bienvenue dans le Multivers , vous incarnerez l'une des 3 classes suivantes : démon , lampadaire , capybara ")
	fmt.Println("Mais avant cela veuillez choisir votre nom ")
	fmt.Print("Choisis ton nom : ")
	fmt.Scanln(&name)
	if !(contientUniquementLettres(name)){
		fmt.Println("Veuillez utiliser uniquement des lettres")
		p.CharCreation()
		return
	}
	name = capitalizeString(name)
	fmt.Println("----------------------------------------")
	fmt.Println("Choisis ta classe :")
	fmt.Println("1) Demon")
	fmt.Println("\tLe Demon est& un etre surnaturel avec un grand pouvoir alimenter du royaume des ténébres")
	fmt.Println("\tPV : 180")
	fmt.Println("\n")
	fmt.Println("2) Lampadaire")
	fmt.Println("\tLes lampadaires est un etre lumineux qui est fixe son pouvoir s'active la nuit")
	fmt.Println("\tPV : 120")
	fmt.Println("\n")
	fmt.Println("3) Capybara")
	fmt.Println("\tLe Capybara est fait pour explorer le monde sous marin mais les jours de pleine lune il se transforme en Hydrochoerus hydrochaerisq")
	fmt.Println("\tPV : 100")
	p.inventairecount = 2
	
	fmt.Scanln(&classe)
	p.Init(name, classe, 1, 100, 50	, map[string]int{"potion": 3, "Chapeau de l’aventurier": 1, "Tunique de l’aventurier": 1,  "Bottes de l’aventurier": 1}, 100)
	fmt.Print("\033[H\033[2J")
}

func keyExists(key string, myMap map[string]int) bool {
	_, exists := myMap[key]
	return exists
}

func contientUniquementLettres(s string) bool {
    for _, char := range s {
        if !unicode.IsLetter(char) {
            return false
        }
    }
    return true
}

func (p *Personnage) Menu() {
	fmt.Println("----------------------------------------")
	fmt.Println("1) Afficher les informations du personnage")
	fmt.Println("2) Afficher l'inventaire")
	fmt.Println("3) Allez au marchand")
	fmt.Println("4) Manequin combat")
	fmt.Println("5) Retour au menu")
	var choice int
	fmt.Scan(&choice)
	fmt.Println(choice)
	switch choice {
	case 1:
		fmt.Print("\033[H\033[2J")
		p.DisplayInfo()
	case 2: 
		fmt.Print("\033[H\033[2J")
		p.AccessInventory()
	case 3:
		fmt.Print("\033[H\033[2J")
		p.Marchand()
	case 4:
		fmt.Print("\033[H\033[2J")
		p.TrainingFight()
	default:
	}
	p.Menu()
}

func (p *Personnage) AjouterObjet(objet string , quantite int) bool {
	if len(p.inventaire) >= 10 {
		fmt.Println("L inventaire est plein ,impossible d'ajouter plus d'objets")
		return false 
	} else {
		p.inventaire[objet] = quantite
		fmt.Println(objet + "ajouté à l'inventaire")
		return true
	}	
}

type boutique struct {
    item map[string] int
    price int
    quantity int
}
func (p *Personnage) Marchand() {
    fmt.Println("Bonjour aventurier, voici mon stock ?")
    fmt.Println("\t1) Potion de vie, 2€")
    fmt.Println("\t2) Potion de poison, 5€")
    fmt.Println("\t3) Livre de sort: Boule de feu, 20€")
	fmt.Println("\t4) Fourrure de loup, 5€")
	fmt.Println("\t5) Peau de troll, 6€")
	fmt.Println("\t6) Cuir de sanglier, 2€")
	fmt.Println("\t7) Plume de corbeau, 1€")
	fmt.Println("\t8) Augmenter ton inventaire, 30€")
	var choice int
	fmt.Scan(&choice)
    switch choice {
    case 1:
        if p.golds >= 2 {
            p.golds -= 2
			fmt.Println("Il vous reste", p.golds)
            p.addInventory("Potion de vie")
        } else {
            fmt.Println("tu n'as pas d'argent")
        }
    case 2:
        if p.golds >= 5 {
            p.golds -= 5
			fmt.Println("Il vous reste", p.golds)
            p.addInventory("Potion de poison")
        } else {
            fmt.Println("tu n'as pas d'argent")
        }
    case 3:
        if p.golds >= 20 {
            p.golds -= 20
			fmt.Println("Il vous reste", p.golds)
            p.addInventory("Livre de sort: Boule de feu")
        } else {
            fmt.Println("tu n'as pas d'argent")
        }
    case 4:
        if p.golds >= 5 {
            p.golds -= 5
            p.addInventory("Fourrure de loup")
        } else {
            fmt.Println("tu n'as pas d'argent")
        }
    case 5:
        if p.golds >= 6 {
            p.golds -= 6
            p.addInventory("Peau de troll")
        } else {
            fmt.Println("tu n'as pas d'argent")
        }
    case 6:
        if p.golds >= 2 {
            p.golds -= 2
            p.addInventory("Cuir de sanglier")
        } else {
            fmt.Println("tu n'as pas d'argent")
        }

    case 7:
        if p.golds >= 1 {
            p.golds -= 1
            p.addInventory("Plume de corbeau")
        } else {
            fmt.Println("tu n'as pas d'argent")
        }
    case 15:
        if p.golds >= 30 {
            p.golds -= 30
            p.UpgradeInventorySlot()

        } else {
            fmt.Println("tu n'as pas d'argent")
        }
    default:
        fmt.Println("Pour quitter le menu appuyer sur 1")
    }

}

func (p *Personnage) addInventory(item string) {
	p.inventaire [item] = p.inventaire[item] +1
	p.inventairecount++
    fmt.Println(item, "ajouté a ton inventaire")
}

func (p *Personnage) removeInventory(item string) {
	p.inventaire [item] = p.inventaire[item] -1
	p.inventairecount--
    fmt.Printf(item, "supprimé de ton inventaire")

}

func (p *Personnage) UpgradeInventorySlot() bool {
    if len(p.inventaire) >= 3 {
        fmt.Println("Limite d'augmentation atteinte. Impossible d'augmenter davantage la capacité de l'inventaire.")
        return false
    }
	p.up.utilisations++
	p.up.capaciteMax += 10
    fmt.Printf("La capacité maximale de l'inventaire a été augmentée à %d.\n", p.up.capaciteMax)
	return true
}

func (p *Personnage) removeInventory2(itemName string) {
	delete(p.inventaire, itemName)
}

func Inputint() (int, error) {
    fmt.Print(">> ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()
    chiffre, err := strconv.Atoi(input)
    if err != nil {
        return 0, err
    }
    return chiffre, nil
}

func ClearConsole() {
    const clearScreen = "\033[H\033[2J"
    fmt.Print(clearScreen)
}

type Monstre struct {
nom string
pvmax int
pv int
Dégats int
}

func(m* Monstre) InitGoblin() {
    m.nom = "Mannequin"
    m.pvmax = 30
    m.pv = 30
    m.Dégats = 10
}

func (m Monstre) EnnemyRound(p Personnage){
    ClearConsole()
    fmt.Println("HP du joueur : ", m.pv)
    fmt.Println("HP de l'ennemi : ", m.pv)
    degats := m.Dégats
    fmt.Println("L'ennemi vous a infigé", degats, " dégats")
    p.pv -= degats
}

func (p Personnage) PlayerRound(m *Monstre){
    ClearConsole()
    fmt.Println("HP du joueur : ", p.pv)
    fmt.Println("HP de l'ennemi : ", m.pv)
    fmt.Println("Bonjour voyageur, que voulez-vous faire ?")
    fmt.Println("1: Attaquer")
    fmt.Println("2: Abilités")
    choice, _ := Inputint()
    switch choice{
    case 1:
        fmt.Println("Vous avez infligé 10 de dégats")
        m.pv -= 10
    case 2:
        fmt.Println("abilités")
    }
}

func isDead(p Personnage, m *Monstre) bool {
    if p.pv <= 0 || m.pv <= 0{
        return true
    }
    return false
}

func(p *Personnage) TrainingFight() {
    m := Monstre {"Mannequin", 30 , 30 , 10}
    for !(isDead(*p, &m)){
        p.PlayerRound(&m)
        if !(isDead(*p, &m)){
            m.EnnemyRound(*p)
        }
    }
    if p.pv <= 0{
        fmt.Println("Vous avez perdu")
    } else {
        fmt.Println("Vous avez gagné")
    }
}
