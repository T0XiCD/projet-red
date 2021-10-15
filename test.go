package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Personnage struct {
	nom         string
	classe      string
	argent      int
	niveau      int
	maxlife     int
	currentlife int
	attaque 	int
	inventaire  []string
	skill       []string
	equipement Equipement
}

func (c *Personnage) Init(nom string, classe string, argent int, niveau int, maxlife int, currentlife int, attaque int, inventaire []string, skill []string, equipement Equipement) {
	c.nom = nom
	c.classe = classe
	c.argent = argent
	c.niveau = niveau
	c.maxlife = maxlife
	c.currentlife = currentlife
	c.attaque = attaque
	c.inventaire = inventaire
	c.skill = skill
	c.equipement = equipement
}

type Marchand struct {
	nom        string
	argent     int
	inventaire []string
}

func (m *Marchand) Init(nom string, argent int, inventaire []string) {
	m.nom = nom
	m.argent = argent
	m.inventaire = inventaire
}
func main() {
	var c1 Personnage
	var m1 IA
	var e1 Equipement
	c1.Init("nom", "guerrier", 100, 1, 250, 100, 10, []string{}, []string{}, e1)
	m1.InitGoblinn("Goblin", 50, 50, 5)
	c1.Personnagenom()
	c1.Menu()
}
func (c *Personnage) Personnagenom() { // fonction pour que le joueur chosise lui meme le nom du perso 
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Bienvenue veuillez entrer le nom de votre personnage :")
	// reads user input until by default
	scanner.Scan()
	// Holds the string that was scanned
	text := scanner.Text()
	switch text {
	default:
		c.nom = text
		fmt.Println("----------------")
		fmt.Print("votre perso s'appelle :", c.nom)
		fmt.Println()
		fmt.Println("-----------------------")
	}
}
func (c *Personnage) ClasseChoix() { // pas utiliser pour l'instant (permet de choisir sa classe)
	fmt.Println("Veuillez entrer la classe de votre personnage")
	fmt.Println("1-Archer")
	fmt.Println("2-Assassin")
	fmt.Println("3-Guerrier")
	fmt.Println("----------")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := scanner.Text()

	switch a {
	case "1":
		fmt.Println("Archer\n Vous avez choisi la classe archer votre niveau et de 1.")
	case "2":
		fmt.Println("Assassin\n Vous avez choisi la classe assassin votre niveau et de 1.")
	case "3":
		fmt.Println("Guerrier\n Vous avez choisi la classe guerrier votre niveau et de 1.")
	}
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("---------------------------------------------------------------------")
}
func (c *Personnage) Menu() { // fonction d menu
	var i IA 
	fmt.Println("Voici le menu du jeux")
	fmt.Println("1-info")
	fmt.Println("2-inventaire")
	fmt.Println("3-marchand")
	fmt.Println("4-forgeron")
	fmt.Println("5-combat")
	fmt.Println("6-quitter")
	fmt.Println("---------------------")
	fmt.Println("---------------------")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := scanner.Text()

	switch a {
	case "1":
		c.DisplayInfo()
	case "2":
		c.AccesInventaire()
	case "3":
		c.Achat()
	case "4":
		c.Forgeron()
	case "5":
		c.TrainingFight(&i)
	case "6":

	}
}
func (c *Personnage) Achat() { //fonction pour acheter au marchand 
    fmt.Println("Bonjour voulez vous acheter quelque chose")
    fmt.Println("1-Oui")
    fmt.Println("2-Non")
    fmt.Println("---------------------")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    ma := scanner.Text()

    switch ma {
    case "1":
        fmt.Println("1-Potion de Heal GRATUIT")
        fmt.Println("2-Potion de Poison 10€")
        fmt.Println("3-Boule de Feu 50€")
		fmt.Println("4-Peau de troll 30€")
		fmt.Println("5-Cuir de sanglier 20€")
		fmt.Println("6-Plume de corbeau 15€")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        mb := scanner.Text()

        switch mb {
        case "1":
            c.inventaire = append(c.inventaire, "Potion de Heal")
            c.argent -= 0
            fmt.Println("La Potion est Gratuite")
            fmt.Println("---------")
            c.Menu()
        case "2":
            c.inventaire = append(c.inventaire, "Potion de Poison")
            if c.argent >= 10 {
                c.argent -= 10
                fmt.Println("Il vous reste", c.argent)
                fmt.Println("---------")
                c.Menu()
            } else {
                fmt.Println("Vous n'avez pas assez d'argent")
                c.Achat()
            }
        case "3":
            c.skill = append(c.skill, "Boule de feu")
			if c.argent >= 50 { 
                c.argent -= 50
                fmt.Println("Il vous reste", c.argent)
                fmt.Println("---------")
                c.Menu()
            } else {
                fmt.Println("Vous n'avez pas assez d'argent")
                c.Achat()
			}
		case "4":
            c.inventaire = append(c.inventaire, "Peau de troll")
            if c.argent >= 30 {
                c.argent -= 30
                fmt.Println("Il vous reste", c.argent)
                fmt.Println("---------")
                c.Menu()
            } else {
                fmt.Println("Vous n'avez pas assez d'argent")
                c.Achat()
            }
		case "5":
            c.inventaire = append(c.inventaire, "Cuir de sanglier")
            if c.argent >= 20 {
                c.argent -= 20
                fmt.Println("Il vous reste", c.argent)
                fmt.Println("---------")
                c.Menu()
            } else {
                fmt.Println("Vous n'avez pas assez d'argent")
                c.Achat()
            }
		case "6":
            c.inventaire = append(c.inventaire, "Plume de corbeau")
            if c.argent >= 15 {
                c.argent -= 15
                fmt.Println("Il vous reste", c.argent)
                fmt.Println("---------")
                c.Menu()
            } else {
                fmt.Println("Vous n'avez pas assez d'argent")
                c.Achat()
            }
        }
    case "2":
        c.Menu()
    }
}
func (c *Personnage) DisplayInfo() { // function pour les stast du perso 
	fmt.Println("nom :", c.nom)
	fmt.Println("classe :", c.classe)
	fmt.Println("niveau :", c.niveau)
	fmt.Println("maxlife :", c.maxlife)
	fmt.Println("currentlife :", c.currentlife)
	fmt.Println("attaque :", c.attaque)
	fmt.Println("argent :", c.argent)
	fmt.Println("skill :", c.skill)
	fmt.Println("equipement :", c.equipement)
	fmt.Println("------------")
	fmt.Println("1-retour menu")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	r := scanner.Text()

	switch r {
	case "1":
		c.Menu()
	}
}
func (c *Personnage) AccesInventaire() { // fonction qui permet d'ovrire l'inventaire 
	fmt.Println("inventaire de:", c.nom)
	if len(c.inventaire) == 0 {
		fmt.Println("L'inventaire est Vide")
	} else {
		for i := range c.inventaire {
			fmt.Print("objet", i+1, ":", c.inventaire[i], "\n")
		}
	}
	fmt.Println()
	fmt.Println("-------")
	fmt.Println("1-prendre potion de vie")
	fmt.Println("-------")
	fmt.Println("2-prendre potion de poison")
	fmt.Println("-------")
	fmt.Println("3-equiper")
	fmt.Println("-------")
	fmt.Println("4-retour menu")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	in := scanner.Text()

	switch in {
	case "1":
		c.TakePot()
	case "2":
		c.PoisonPot()
	case "3":
		c.Equiper()
	case "4":
		c.Menu()
	}
}
func (c *Personnage) TakePot() { // prendre une potion de heal
	if c.currentlife == c.maxlife {
		fmt.Println(c.nom, "est déjà full HP, il ne peut pas utiliser ceci")
	} else {
		for i := range c.inventaire {
			if c.inventaire[i] == "Potion de Heal" {
				c.RemoveInv("Potion de Heal")
				fmt.Println(c.nom, "Prend une Potion")
				c.currentlife += 50
				if c.currentlife >= c.maxlife {
					c.currentlife = 250
					fmt.Println(c.nom, "est full hp")
				} else {
					fmt.Println(c.nom, "a", c.currentlife, "hp")
				}
			}
		}
	}
	c.AccesInventaire()
}
func (c *Personnage) RemoveInv(obj string) { // Fonction de retrait d'un objet a l'inventaire
    for i := range c.inventaire {
        if i < len(c.inventaire) {
            if c.inventaire[i] == obj {
                c.inventaire = append(c.inventaire[:i], c.inventaire[i+1:]...)
                break
            }
        }
    }
}
func (c *Personnage) Death() { // fonction de mort du perso + revive avec la moitier de ses hp
	if c.currentlife == 0 {
		fmt.Println(c.nom, "est mort", "\n", )
			c.currentlife = c.maxlife/2
			  fmt.Println(c.nom, "est revive avec", c.currentlife, "hp")
		} else {
		fmt.Println()
	}
}
func (c *Personnage) PoisonPot() { // fonction de la potion de poisson 
    for i := range c.inventaire {
        if c.inventaire[i] == "Potion de Poison" {
            c.RemoveInv("Potion de Poison")
	test := 0
	for i := 0; i < 3; i++ {
        if test == 0 {
            c.currentlife -= 10
            fmt.Println(c.nom, "a maintenant", c.currentlife, "Hp")
            if c.currentlife == 0 {
                c.Death()
                test++
            }
            time.Sleep(1 * time.Second)
        		}
   	 		}
		}
	}
	c.Menu()
}
func (c *Personnage) Forgeron() { // fonction du forgeront pour craft 
	fmt.Println("Voici le menu du forgeront")
	fmt.Println("1-Chapeau d'aventurier")
	fmt.Println("2-Tunique d'aventurier")
	fmt.Println("3-Bottes d'aventurier")
	fmt.Println("4-retour menu")
	fmt.Println("---------------------")
	fmt.Println("---------------------")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	f := scanner.Text()

	switch f {
	case "1":
		for i := range c.inventaire {
			if c.inventaire[i] == ("Plume de corbeau") {
				c.RemoveInv("Plume de corbeau")
				c.inventaire = append(c.inventaire, "Chapeau d'aventurier")
		} else {
		fmt.Println("Vous n'avez pas les ressources")
		fmt.Println("-----------------------------")
			}
		}
	case "2":
		for i := range c.inventaire {
			if c.inventaire[i] == ("Cuir de sanglier") {
				c.RemoveInv("Cuir de sanglier")
				c.inventaire = append(c.inventaire, "Tunique d'aventurier")
		} else {
		fmt.Println("Vous n'avez pas les ressources")
		fmt.Println("-----------------------------")
			}
		}
	case "3":
		for i := range c.inventaire {
			if c.inventaire[i] == ("Peau de troll") {
				c.RemoveInv("Peau de troll")
				c.inventaire = append(c.inventaire, "Bottes d'aventurier")
		} else {
		fmt.Println("Vous n'avez pas les ressources")
		fmt.Println("-----------------------------")
			}
		}
	case "4":
		c.Menu()
	}
	c.Menu()
}
type Equipement struct {
    casque string
    torse  string
    pieds  string
}
func (e *Equipement) Init(casque string, torse string, pieds string) {
    e.casque = casque
    e.torse = torse
    e.pieds = pieds
} 

func (c *Personnage)Equiper() { //fonction pour equiper un equipement 
	fmt.Println("1-equiper le casque")
	fmt.Println("2-equiper le torse")
	fmt.Println("3-equiper les bottes")
	fmt.Println("4-retour")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	l := scanner.Text()
	
	switch l {
	case "1":
	for i := range c.inventaire {
        if c.inventaire[i] == "Chapeau d'aventurier" {
			c.RemoveInv("Chapeau d'aventurier")
			c.equipement.casque = "Chapeau d'aventurier"
			c.maxlife += 10
		}
	}
	case "2":
	for i := range c.inventaire {
        if c.inventaire[i] == "Tunique d'aventurier" {
			c.RemoveInv("Tunique d'aventurier")
			c.equipement.torse = "Tunique d'aventurier"
			c.maxlife += 25
		}
	}
	case "3":
	for i := range c.inventaire {
        if c.inventaire[i] == "Bottes d'aventurier" {
			c.RemoveInv("Bottes d'aventurier")
			c.equipement.pieds = "Bottes d'aventurier"
			c.maxlife += 15
			}
		}
	case "4":
		c.AccesInventaire()
	}
	c.AccesInventaire()
}

type IA struct {
		nom string
		cuurentlife  int
		maxlife  int
		attaque int
}
func (i *IA) InitGoblinn(nom string, currentlife int, maxlife int, attaque int) {
		i.nom = nom
		i.cuurentlife = currentlife
		i.maxlife = maxlife
		i.attaque = attaque
}

func (c *Personnage) TrainingFight(i *IA) { //fonction pour le combat 
	round := false
	var counter int
for i.cuurentlife > 0 && c.currentlife > 0 {
		counter++
	if !round {
		round = c.CharTurn(i)
	} else if round {
		if counter%3 == 0 {
			c.GoblinPattern(i)
		} else {
			c.currentlife -= i.attaque
			fmt.Println("")
		}
		round = false 
		}
	}
	if i.cuurentlife <= 0 {
		fmt.Println("victoire")
		c.argent += 10
	} else {
		fmt.Println("Perdu")
		c.Death()
	}
	c.Menu()
}

func (c *Personnage) GoblinPattern(i *IA) {
	attaque := i.attaque * 2
	c.currentlife -= attaque
	fmt.Println()
}

func (c *Personnage) CharTurn(i *IA) bool {
	fmt.Println("1-attaque")
	fmt.Println("2-inventaire")
	fmt.Println("3-fuir")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	w := scanner.Text()
	
	switch w {
	case "1":
		i.cuurentlife =- 10
		fmt.Println("")
	case "2":
		c.AccesInventaire()

	case "3":
		c.Menu()
	}
	return true
}

func (i *IA) DeathGoblin() { // fonction de mort du goblin
	if i.cuurentlife == 0 {
		fmt.Println(i.nom, "est mort", "\n", )
		} else {
		fmt.Println()
	}
}