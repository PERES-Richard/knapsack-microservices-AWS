package generator

import (
	. "generator/entities"
	"log"
	"math/rand"
	"sync"
)

const MAX_DEFAULT_SIZE = 100
const MAX_DEFAULT_NB_ITEM = 100
const MAX_ITEM_VALUE = 100

func GenerateNewKnapSet(knapSet KnapSet, nbItem int) KnapSet {
	//rand.Seed(time.Now().UnixNano()) TODO

	if knapSet.BagSize == 0 {
		knapSet.BagSize = rand.Intn(MAX_DEFAULT_SIZE) + 1 // To avoid 0
	}

	if nbItem == 0 {
		nbItem = rand.Intn(MAX_DEFAULT_NB_ITEM) + 1 // To avoid 0
	}

	knapSet.Items = generateItems(nbItem, knapSet.BagSize)

	return knapSet
}

func generateItems(nbItem int, bagSize int) []Item {
	if nbItem <= 0 {
		log.Fatal("nbItem cannot be <= 0")
	}

	itemList := make([]Item, nbItem)
	var wg sync.WaitGroup

	for itemID := 0; itemID < nbItem; itemID++ {
		wg.Add(1)

		go func(itemID int, bagSize int) {
			defer wg.Done()
			itemList[itemID] = generateItem(itemID, bagSize)
		}(itemID, bagSize)
	}

	wg.Wait()

	return itemList
}

func generateItem(itemID int, maxSize int) Item {
	if maxSize <= 0 {
		log.Fatal("Max Size cannot be <= 0")
	}
	//rand.Seed(time.Now().UnixNano()) TODO

	value := rand.Intn(MAX_ITEM_VALUE) + 1
	size := rand.Intn(maxSize) + 1

	return Item{Id: itemID, Value: value, Size: size}
}
