package main

import "fmt"

type director struct {
    builder iBuilder
}

//house описание параметров желаемого дома
type house struct {
    windowType string
    doorType   string
    floor      int
}

//normalBuilder описание предложения обычного строителя
type normalBuilder struct {
    windowType string
    doorType   string
    floor      int
}

//iglooBuilder описание предложения строителя иглу
type iglooBuilder struct {
    windowType string
    doorType   string
    floor      int
}

//iBuilder методы постройки
type iBuilder interface {
    setWindowType()
    setDoorType()
    setNumFloor()
    getHouse() house
}

func main() {
    normalBuilder := getBuilder("normal")
    iglooBuilder := getBuilder("igloo")

    director := newDirector(normalBuilder)
    normalHouse := director.buildHouse()

    fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
    fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
    fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

    director.setBuilder(iglooBuilder)
    iglooHouse := director.buildHouse()

    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}

//newDirector новый вид строителя
func newDirector(b iBuilder) *director {
    return &director{
        builder: b,
    }
}

//setBuilder выбор строителя
func (d *director) setBuilder(b iBuilder) {
    d.builder = b
}

//buildGouse построить дом
func (d *director) buildHouse() house {
    d.builder.setDoorType()
    d.builder.setWindowType()
    d.builder.setNumFloor()
    return d.builder.getHouse()
}


//newNormalBuilder новый обычный строитель
func newNormalBuilder() *normalBuilder {
    return &normalBuilder{}
}

//setWindowType выбор типа окна
func (b *normalBuilder) setWindowType() {
    b.windowType = "Wooden Window"
}

//setDoorType выбор типа двери
func (b *normalBuilder) setDoorType() {
    b.doorType = "Wooden Door"
}

//setNumFloor ыбор типа пола
func (b *normalBuilder) setNumFloor() {
    b.floor = 2
}

//getHouse построить дом
func (b *normalBuilder) getHouse() house {
    return house{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}

//newIglooBuilder новый строитель  иглу
func newIglooBuilder() *iglooBuilder {
    return &iglooBuilder{}
}

//setWindowType выбор типа окна
func (b *iglooBuilder) setWindowType() {
    b.windowType = "Snow Window"
}

//setDoorType выбор типа двери
func (b *iglooBuilder) setDoorType() {
    b.doorType = "Snow Door"
}

//setNumFloor ыбор типа пола
func (b *iglooBuilder) setNumFloor() {
    b.floor = 1
}

//getHouse построить дом
func (b *iglooBuilder) getHouse() house {
    return house{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}

//getBuilder выбор строителя
func getBuilder(builderType string) iBuilder {
    if builderType == "normal" {
        return &normalBuilder{}
    }

    if builderType == "igloo" {
        return &iglooBuilder{}
    }
    return nil
}