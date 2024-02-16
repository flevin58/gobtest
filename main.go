package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"

	hw "github.com/flevin58/gobtest/lib"
)

type Person struct {
	Name    string
	Surname string
	Birth   time.Time
}

func main() {
	log.SetFlags(0)
	birth := time.Date(1958, time.August, 8, 0, 0, 0, 0, time.UTC)
	nando := Person{"Fernando", "Levin", birth}
	gout, err := os.Create("data/person.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer gout.Close()
	genc := gob.NewEncoder(gout)
	genc.Encode(nando)

	gin, err := os.Open("data/person.gob")
	if err != nil {
		log.Fatal(err)
	}
	var person Person
	gdec := gob.NewDecoder(gin)
	err = gdec.Decode(&person)
	if err != nil {
		log.Fatal(err)
	}

	serialnum := hw.GetComputerSerialNumber()
	fmt.Println(serialnum)
	fmt.Println(nando)
	fmt.Println(person)
}
