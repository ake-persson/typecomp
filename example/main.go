package main

import (
	"fmt"
	"log"

	"github.com/mickep76/kvstore/cmp"
)

type Car struct {
	Manufacturer string
	Model        string
}

type Cars []*Car

func (c Car) String() string {
	return c.Manufacturer + " " + c.Model
}

func (c Car) Eq(b interface{}) (bool, error) {
	return c.String() == b.(Car).String(), nil
}

func (c Car) Lt(b interface{}) (bool, error) {
	return c.String() < b.(Car).String(), nil
}

func TestInterface(c cmp.Comparer) {
}

func main() {
	cars := Cars{
		&Car{
			Manufacturer: "Audi",
			Model:        "Q3",
		},
		&Car{
			Manufacturer: "Audi",
			Model:        "Q5",
		},
		&Car{
			Manufacturer: "Audi",
			Model:        "Q7",
		},
		&Car{
			Manufacturer: "Volvo",
			Model:        "XC40",
		},
		&Car{
			Manufacturer: "Volvo",
			Model:        "XC60",
		},
		&Car{
			Manufacturer: "Volvo",
			Model:        "XC90",
		},
	}

	TestInterface(*cars[0])

	ok, err := cmp.Eq(cars[0], cars[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ok: %v\n", ok)

	ok, err = cmp.Eq(cars[0], cars[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ok: %v\n", ok)

	ok, err = cmp.Neq(cars[0], cars[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ok: %v\n", ok)

	ok, err = cmp.Lt(cars[0], cars[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ok: %v\n", ok)
}
