package main

import (
	"todoapp/komandi"
	"todoapp/scanner"
	simpleconnection "todoapp/simple_connection"
)

func main() {
	komandiKomand := komandi.NewKomand()

	scanner := scanner.NewScanner(komandiKomand)

	scanner.Start()

	simpleconnection.CheckConnection()
}
