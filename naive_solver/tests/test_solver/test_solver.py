import unittest
from solver.solver import bestPathFinder

test_cases = [
    [10, [
        {
            "id": 0,
            "size": 3,
            "value": 16
        },
        {
            "id": 1,
            "size": 4,
            "value": 15
        },
        {
            "id": 2,
            "size": 4,
            "value": 14
        },
        {
            "id": 3,
            "size": 2,
            "value": 7
        }
    ], 38, '{}'
     ],
    [10, [
        {
            "id": 0,
            "size": 9,
            "value": 26
        },
        {
            "id": 1,
            "size": 2,
            "value": 70
        }
    ], 70, '{}'
     ],
    [100,
     [{'id': 0, 'value': 6, 'size': 54}, {'id': 1, 'value': 8, 'size': 9}, {'id': 2, 'value': 2, 'size': 5},
      {'id': 3, 'value': 9, 'size': 31}, {'id': 4, 'value': 6, 'size': 20}, {'id': 5, 'value': 8, 'size': 48},
      {'id': 6, 'value': 5, 'size': 2}, {'id': 7, 'value': 4, 'size': 69}, {'id': 8, 'value': 7, 'size': 72},
      {'id': 9, 'value': 1, 'size': 11}, {'id': 10, 'value': 3, 'size': 8}, {'id': 11, 'value': 10, 'size': 60},
      {'id': 12, 'value': 8, 'size': 55}, {'id': 13, 'value': 3, 'size': 10}, {'id': 14, 'value': 11, 'size': 74},
      {'id': 15, 'value': 6, 'size': 99}, {'id': 16, 'value': 5, 'size': 2}, {'id': 17, 'value': 11, 'size': 85},
      {'id': 18, 'value': 10, 'size': 74},
      {'id': 19, 'value': 10, 'size': 5}, {'id': 20, 'value': 9, 'size': 12}, {'id': 21, 'value': 7, 'size': 67},
      {'id': 22, 'value': 10, 'size': 53}, {'id': 23, 'value': 2, 'size': 89}, {'id': 24, 'value': 2, 'size': 53},
      {'id': 25, 'value': 7, 'size': 9}, {'id': 26, 'value': 5, 'size': 21}, {'id': 27, 'value': 2, 'size': 53},
      {'id': 28, 'value': 10, 'size': 82}, {'id': 29, 'value': 2, 'size': 90}, {'id': 30, 'value': 9, 'size': 9},
      {'id': 31, 'value': 4, 'size': 64}, {'id': 32, 'value': 5, 'size': 21}, {'id': 33, 'value': 6, 'size': 12},
      {'id': 34, 'value': 6, 'size': 4}, {'id': 35, 'value': 6, 'size': 30}, {'id': 36, 'value': 4, 'size': 78},
      {'id': 37, 'value': 3, 'size': 95}, {'id': 38, 'value': 4, 'size': 77}, {'id': 39, 'value': 3, 'size': 9},
      {'id': 40, 'value': 11, 'size': 4}, {'id': 41, 'value': 3, 'size': 46}, {'id': 42, 'value': 5, 'size': 75},
      {'id': 43, 'value': 10, 'size': 73},
      {'id': 44, 'value': 4, 'size': 79}, {'id': 45, 'value': 7, 'size': 21}, {'id': 46, 'value': 1, 'size': 18},
      {'id': 47, 'value': 4, 'size': 30}, {'id': 48, 'value': 4, 'size': 68}
     ], 87, '{}']
]


class MyTestCase(unittest.TestCase):

    def test_bestPathFinder(self):
        for case in test_cases:
            bagsize = case[0]
            items = case[1]
            expectedMaxvalue = case[2]
            expectedItems = case[3]
            maxvalue, items = bestPathFinder(bagsize, items)
            self.assertEqual(maxvalue, expectedMaxvalue)
            # self.assertEqual(items, expectedItems)


if __name__ == '__main__':
    unittest.main()
