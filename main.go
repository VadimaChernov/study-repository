package main

import (
	"todoapp/komandi"
	"todoapp/scanner"
)

func main() {
	komandiKomand := komandi.NewKomand()

	scanner := scanner.NewScanner(komandiKomand)

	scanner.Start()
}
