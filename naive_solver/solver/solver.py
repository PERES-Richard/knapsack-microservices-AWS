from concurrent.futures import ThreadPoolExecutor, Future
from collections import namedtuple, defaultdict

Result = namedtuple('Result', ('totalItems', 'totalvalue', 'sizeLeft', 'itemsLeft'))

def bestPathFinderRec(currentItems, currentValue, sizeLeft, itemsLeft, nb_thread=0):
    # If the size left is negative then we went 1 step to deep
    if sizeLeft < 0:
        return currentItems[:-1], currentValue - currentItems[-1]["value"], sizeLeft + currentItems[-1][
            "size"], itemsLeft

    # If there is no size or item left then we are at the end of the path
    if sizeLeft == 0 or len(itemsLeft) == 0:
        return currentItems, currentValue, sizeLeft, itemsLeft

    if nb_thread <= 0:

        # Current best combination is this current path
        totalMax = currentValue
        totalItems = currentItems[:]

        newSizeLeft = sizeLeft
        newItemsLeft = []

        for i in range(len(itemsLeft)):

            # ... we test a new combination/path using 'current' + next items
            newCurrentItems = currentItems[:]
            newCurrentItems.append(itemsLeft[i])

            newItems, newMax, sizeL, itemsL = bestPathFinderRec(newCurrentItems, currentValue + itemsLeft[i]["value"], sizeLeft - itemsLeft[i]["size"], itemsLeft[i+1:])

            # If this combination/path is better than what we already have
            if newMax > totalMax:
                totalItems = newItems
                totalMax = newMax
                newSizeLeft = sizeL
                newItemsLeft = itemsL

        return totalItems, totalMax, newSizeLeft, newItemsLeft

    with ThreadPoolExecutor(nb_thread) as executer:
        futures = defaultdict(list)
        for i in range(len(itemsLeft)):
            # ... we test a new combination/path using 'current' + next items
            newCurrentItems = currentItems[:]
            newCurrentItems.append(itemsLeft[i])

            threadsAvailable = 0
            if nb_thread - len(itemsLeft) - i > 0:
                threadsAvailable = len(itemsLeft) - i
                nb_thread -= threadsAvailable

            futures[i] = executer.submit(bestPathFinderRec, newCurrentItems, currentValue + itemsLeft[i]["value"],
                                                                sizeLeft - itemsLeft[i]["size"], itemsLeft[i + 1:], threadsAvailable)

        max_Result = ('', 0)
        for i in range(len(futures)):
            res = futures[i].result() # Wait for the thread to finish
            if res[1] > max_Result[1]: # Compare total value
                max_Result = res

        return max_Result


def bestPathFinder(initialBagSize, initialItems):
    items, max, _, _ = bestPathFinderRec([], 0, initialBagSize, initialItems, len(initialItems)*4) # nb max threads TBD
    return max, items
