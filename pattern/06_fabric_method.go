package main

import "fmt"

type musket struct {
    gun
}

type ak47 struct {
    gun
}

type iGun interface {
    setName(name string)
    setPower(power int)
    getName() string
    getPower() int
}

type gun struct {
    name  string
    power int
}

func main() {
    ak47, _ := getGun("ak47")
    musket, _ := getGun("musket")

    printDetails(ak47)
    printDetails(musket)
}

//printDetails печать подробностей
func printDetails(g iGun) {
    fmt.Printf("Gun: %s", g.getName())
    fmt.Println()
    fmt.Printf("Power: %d", g.getPower())
    fmt.Println()
}

//newMusket создание мушкета
func newMusket() iGun {
    return &musket{
        gun: gun{
            name:  "Musket gun",
            power: 1,
        },
    }
}

//getGun получить ствол
func getGun(gunType string) (iGun, error) {
    if gunType == "ak47" {
        return newAk47(), nil
    }
    if gunType == "musket" {
        return newMusket(), nil
    }
    return nil, fmt.Errorf("Wrong gun type passed")
}

//setName назначить имя
func (g *gun) setName(name string) {
    g.name = name
}

//getName получить имя
func (g *gun) getName() string {
    return g.name
}

//setPower назначить силу выстрела
func (g *gun) setPower(power int) {
    g.power = power
}

//getPower получить силу выстрела
func (g *gun) getPower() int {
    return g.power
}

//newAk47 создание АК
func newAk47() iGun {
    return &ak47{
        gun: gun{
            name:  "AK47 gun",
            power: 4,
        },
    }
}