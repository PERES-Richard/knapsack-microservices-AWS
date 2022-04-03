import json

def bestPathFinderRec(currentItems, currentValue, sizeLeft, itemsLeft):
    # If the size left is negative then we went 1 step to deep
    if sizeLeft < 0:
        return currentItems[:-1], currentValue-currentItems[-1]["Value"], sizeLeft+currentItems[-1]["Size"], itemsLeft

    # If there is no size or item left then we are at the end of the path
    if sizeLeft == 0 or len(itemsLeft) == 0:
        return currentItems, currentValue, sizeLeft, itemsLeft

    # Current best combination is this current path
    totalMax = currentValue
    totalItems = currentItems[:]


    newSizeLeft = sizeLeft
    newItemsLeft = []

    # For each item left in the queue...
    for i in range(len(itemsLeft)):

        # ... we test a new combination/path using 'current' + next items
        newCurrentItems = currentItems[:]
        newCurrentItems.append(itemsLeft[i])

        newItems, newMax, sizeL, itemsL = bestPathFinderRec(newCurrentItems, currentValue + itemsLeft[i]["Value"], sizeLeft - itemsLeft[i]["Size"], itemsLeft[i+1:])

        # If this combination/path is better than what we already have
        if newMax > totalMax:
            totalItems = newItems
            totalMax = newMax
            newSizeLeft = sizeL
            newItemsLeft = itemsL

    return totalItems, totalMax, newSizeLeft, newItemsLeft


def bestPathFinder(initialBagSize, initialItems):
    items, max, _, _ = bestPathFinderRec([], 0, initialBagSize, initialItems)
    return max, items

def __main__():
    f = open ('../../data-3.json', "r")
    data = json.loads(f.read())

    bagSize = data["BagSize"]
    items = data["Items"]

    f.close()

    print(bagSize)
    print(items)

    max, items = bestPathFinder(bagSize, items)

    print("-------------------")
    print("Max = ", max)
    print("With items = ", items)

if __name__ == '__main__':
    __main__()
