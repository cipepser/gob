package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func Encode(f string, in interface{}) error {
	fw, err := os.Create(f)
	if err != nil {
		return err
	}
	defer fw.Close()

	enc := gob.NewEncoder(fw)

	err = enc.Encode(in)
	if err != nil {
		return err
	}

	return nil
}

func Decode(f string, out interface{}) error {
	fr, err := os.Open(f)
	if err != nil {
		return err
	}
	defer fr.Close()

	dec := gob.NewDecoder(fr)
	err = dec.Decode(out)
	if err != nil {
		return err
	}

	return nil

}

func main() {
	a := "a"
	err := Encode("save.txt", a)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	var out string
	err = Decode("save.txt", &out)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}

	fmt.Println(out)
}
