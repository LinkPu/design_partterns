package main

import "fmt"

// abstract factory
type Animal interface {
	AnimalShow()
}

// concrete animal
type Horse struct {
	Age int
}

// make sure Horse implement Animal
var _ Animal = &Horse{}

// implement Animal
func (*Horse) AnimalShow() {
	// here use *Horse implement Animal, so only its pointer object can be
	// considered to implement the interface
	fmt.Println("this is horse")
}

type Cattle struct{}

var _ Animal = &Cattle{}

func (*Cattle) AnimalShow() {
	fmt.Println("this is cattle")
}

type Plant interface {
	PlantShow()
}

type Vagetable struct{}

var _ Plant = &Vagetable{}

func (*Vagetable) PlantShow() {
	fmt.Println("this is vagetable")
}

type Fruit struct{}

var _ Plant = &Fruit{}

func (*Fruit) PlantShow() {
	fmt.Println("this is plant")
}

// abstract
type Farm interface {
	NewAnimal() Animal
	NewPlant() Plant
}

type HappyFarm struct{}

var _ Farm = &HappyFarm{}

func (*HappyFarm) NewAnimal() Animal {
	fmt.Println("new horse")
	return &Horse{}
}

func (*HappyFarm) NewPlant() Plant {
	fmt.Println("new fruit")
	return &Fruit{}
}

// you can define more Farm that you want

func main() {
	var farm Farm
	// if use "farm = HappyFarm{}" to init the Farm, you will get the error
	// cannot use (HappyFarm literal) (value of type HappyFarm) as Farm value
	// in assignment: missing method NewAnimal, because it's *Horse implement
	// the Farm rather than Horse
	farm = &HappyFarm{}
	farmAnimal := farm.NewAnimal()
	farmPlant := farm.NewPlant()
	farmAnimal.AnimalShow()
	farmPlant.PlantShow()
}
