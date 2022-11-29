package generator

import (
	"fmt"
	. "knapset-generator/entities"
	"math/rand"
	"sync"
	"time"
)

const MAX_DEFAULT_BAG_SIZE = 100
const MAX_DEFAULT_NB_ITEM = 100
const MAX_ITEM_VALUE = 100
const MAX_ITEM_SIZE = 100

func GenerateNewKnapSet(bagSize int, nbItem int) (KnapSet, error) {
	if bagSize < 0 || nbItem < 0 {
		return KnapSet{}, fmt.Errorf("error : incorrect bagSize (%d) or nbItem (%d) can't be a negative int", bagSize, nbItem)
	}

	rand.Seed(time.Now().UnixNano())

	if bagSize == 0 {
		bagSize = rand.Intn(MAX_DEFAULT_BAG_SIZE) + 1 // To avoid 0
	}

	if nbItem == 0 {
		nbItem = rand.Intn(MAX_DEFAULT_NB_ITEM) + 1 // To avoid 0
	}

	generatedItems, err := generateItems(nbItem, bagSize)
	if err != nil {
		return KnapSet{}, err
	}

	return KnapSet{
		BagSize: bagSize,
		Items:   generatedItems,
	}, nil
}

func generateItems(nbItem int, bagSize int) ([]Item, error) {
	if nbItem <= 0 || bagSize <= 0 {
		return nil, fmt.Errorf("error : incorrect bagSize (%d) or nbItem (%d) can't be a negative int", bagSize, nbItem)
	}

	itemList := make([]Item, nbItem)
	var wg sync.WaitGroup

	for itemID := 0; itemID < nbItem; itemID++ {
		wg.Add(1) // wait for 1 more group

		var err error
		go func(itemID int, bagSize int, err error) {
			defer wg.Done() // notify wg on finish
			item, newErr := generateItem(itemID)
			if newErr != nil {
				err = newErr
			}
			itemList[itemID] = item

		}(itemID, bagSize, err)

		if err != nil {
			return nil, err
		}
	}

	wg.Wait()

	return itemList, nil
}

func generateItem(itemID int) (Item, error) {
	if itemID < 0 {
		return Item{}, fmt.Errorf("error : incorrect itemID can't be a negative int (got %d)", itemID)
	}
	rand.Seed(time.Now().UnixNano())

	value := rand.Intn(MAX_ITEM_VALUE) + 1
	size := rand.Intn(MAX_ITEM_SIZE) + 1

	return Item{Id: itemID, Value: value, Size: size}, nil
}
