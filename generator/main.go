package main

import (
	"encoding/json"
	"fmt"
	"generator/entities"
	"generator/generator"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const BAG_SIZE_ARG = 1
const NB_ITEM_ARG = 2
const CONSOLE_PRINT_ARG = 3
const JSON_PATH_ARG = 4

const defaultJSONPath = "data.json"

func main() {
	bagSize, nbItem, printToConsole, JSONpath := handleArgs()

	newSet := generator.GenerateNewKnapSet(bagSize, nbItem)

	saveToJSON(newSet, JSONpath, printToConsole)
}

func saveToJSON(knapSet entities.KnapSet, path string, alsoPrint bool) {
	file, err := json.MarshalIndent(knapSet, "", " ")
	if err != nil {
		log.Fatal("Error while marshalling JSON")
	}

	err = ioutil.WriteFile(path, file, 0644)
	if err != nil {
		log.Fatal("Error while writing JSON file")
	}

	if alsoPrint {
		fmt.Println(string(file))
	}
}

func handleArgs() (bagSize int, nbItem int, printResult bool, JSONpath string) {
	JSONpath = defaultJSONPath

	if len(os.Args) == 1 {
		return
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

	if len(os.Args) >= CONSOLE_PRINT_ARG+1 {
		printResult, err = strconv.ParseBool(os.Args[CONSOLE_PRINT_ARG])
		if err != nil {
			log.Fatal("Error while reading console print bool arg :", err)
		}
	}

	if len(os.Args) >= JSON_PATH_ARG+1 {
		JSONpath = os.Args[JSON_PATH_ARG]
	}

	return
}
