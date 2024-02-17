package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/flevin58/gobtest/lib"
	stream "github.com/nknorg/encrypted-stream"
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

	key := lib.GetSecretKey()
	config := &stream.Config{
		Cipher:          stream.NewXSalsa20Poly1305Cipher(&key),
		SequentialNonce: true, // only when key is unique for every stream
		Initiator:       true, // only on the dialer side
	}

	// Create the gob file stream
	gobStream, err := os.Create("data/person.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer gobStream.Close()

	// Create an encrypted stream from the gob stream
	cryptStream, err := stream.NewEncryptedStream(gobStream, config)
	if err != nil {
		log.Fatal(err)
	}
	defer cryptStream.Close()

	// Write the encrypted GOB to file
	genc := gob.NewEncoder(cryptStream)
	genc.Encode(nando)

	// Now open the encrypted file
	gin, err := os.Open("data/person.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer gin.Close()

	config.Initiator = false
	decryptStream, err := stream.NewEncryptedStream(gin, config)
	if err != nil {
		log.Fatal(err)
	}
	defer decryptStream.Close()

	var person Person
	gdec := gob.NewDecoder(decryptStream)
	err = gdec.Decode(&person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(nando)
	fmt.Println(person)
}
