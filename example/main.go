package main

import (
	"fmt"
	"log"

	"github.com/mickep76/cmp"
)

// Car type for examp;e.
type Car struct {
	Manufacturer string
	Model        string
}

// Cars multiple cars.
type Cars []*Car

// String representation of car.
func (c Car) String() string {
	return c.Manufacturer + " " + c.Model
}

// Eq interface for comparison.
func (c Car) Eq(b interface{}) (bool, error) {
	return c.String() == b.(Car).String(), nil
}

// Lt interface for comparison.
func (c Car) Lt(b interface{}) (bool, error) {
	return c.String() < b.(Car).String(), nil
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
	}

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

	ok, err = cmp.Eq("123", 123)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ok: %v\n", ok)
}
