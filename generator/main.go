package main

import (
	"generator/generator"
	"log"
	"os"
	"strconv"
)

const BAG_SIZE_ARG = 1
const NB_ITEM_ARG = 2

func main() {
	bagSize, nbItem := handleArgs()

	newSet := generator.GenerateNewKnapSet(bagSize, nbItem)

	log.Printf("%+v\n\n", newSet)
}

func handleArgs() (bagSize int, nbItem int) {
	if len(os.Args) == 1 {
		return 0, 0
	}

	var err error

	bagSize, err = strconv.Atoi(os.Args[BAG_SIZE_ARG])
	if err != nil {
		log.Fatal("Error while reading Bag Size arg :", err)
	}
	if bagSize < 0 {
		log.Fatal("Bag Size argument cannot be < 0")
	}

	nbItem, err = strconv.Atoi(os.Args[NB_ITEM_ARG])
	if err != nil {
		log.Fatal("Error while reading NB item Size arg :", err)
	}
	if nbItem < 0 {
		log.Fatal("Nb item argument cannot be <= 0")
	}

	return bagSize, nbItem
}
