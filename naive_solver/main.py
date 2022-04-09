import solver.solver
import json
import time
from solver.solver import bestPathFinder

def main():
    f = open('../data-5.json', "r")
    data = json.loads(f.read())
    f.close()


    start = time.time()
    max, items = bestPathFinder(data["BagSize"], data["Items"])
    end = time.time()

    print("-------------------")
    print("Max = ", max)
    print("With items = ", items)
    print(f'Time : {end-start:.2f}s\n')


if __name__ == '__main__':
    main()
