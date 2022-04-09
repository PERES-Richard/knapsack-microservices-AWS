import unittest
from solver.solver import bestPathFinder

test_cases = [
    [10, [
        {
            "Id": 0,
            "Size": 3,
            "Value": 16
        },
        {
            "Id": 1,
            "Size": 4,
            "Value": 15
        },
        {
            "Id": 2,
            "Size": 4,
            "Value": 14
        },
        {
            "Id": 3,
            "Size": 2,
            "Value": 7
        }
    ], 38, '{}'
     ],
    [10, [
        {
            "Id": 0,
            "Size": 9,
            "Value": 26
        },
        {
            "Id": 1,
            "Size": 2,
            "Value": 70
        }
    ], 70, '{}'
     ],
    [100,
     [{'Id': 0, 'Value': 6, 'Size': 54}, {'Id': 1, 'Value': 8, 'Size': 9}, {'Id': 2, 'Value': 2, 'Size': 5},
      {'Id': 3, 'Value': 9, 'Size': 31}, {'Id': 4, 'Value': 6, 'Size': 20}, {'Id': 5, 'Value': 8, 'Size': 48},
      {'Id': 6, 'Value': 5, 'Size': 2}, {'Id': 7, 'Value': 4, 'Size': 69}, {'Id': 8, 'Value': 7, 'Size': 72},
      {'Id': 9, 'Value': 1, 'Size': 11}, {'Id': 10, 'Value': 3, 'Size': 8}, {'Id': 11, 'Value': 10, 'Size': 60},
      {'Id': 12, 'Value': 8, 'Size': 55}, {'Id': 13, 'Value': 3, 'Size': 10}, {'Id': 14, 'Value': 11, 'Size': 74},
      {'Id': 15, 'Value': 6, 'Size': 99}, {'Id': 16, 'Value': 5, 'Size': 2}, {'Id': 17, 'Value': 11, 'Size': 85},
      {'Id': 18, 'Value': 10, 'Size': 74},
      {'Id': 19, 'Value': 10, 'Size': 5}, {'Id': 20, 'Value': 9, 'Size': 12}, {'Id': 21, 'Value': 7, 'Size': 67},
      {'Id': 22, 'Value': 10, 'Size': 53}, {'Id': 23, 'Value': 2, 'Size': 89}, {'Id': 24, 'Value': 2, 'Size': 53},
      {'Id': 25, 'Value': 7, 'Size': 9}, {'Id': 26, 'Value': 5, 'Size': 21}, {'Id': 27, 'Value': 2, 'Size': 53},
      {'Id': 28, 'Value': 10, 'Size': 82}, {'Id': 29, 'Value': 2, 'Size': 90}, {'Id': 30, 'Value': 9, 'Size': 9},
      {'Id': 31, 'Value': 4, 'Size': 64}, {'Id': 32, 'Value': 5, 'Size': 21}, {'Id': 33, 'Value': 6, 'Size': 12},
      {'Id': 34, 'Value': 6, 'Size': 4}, {'Id': 35, 'Value': 6, 'Size': 30}, {'Id': 36, 'Value': 4, 'Size': 78},
      {'Id': 37, 'Value': 3, 'Size': 95}, {'Id': 38, 'Value': 4, 'Size': 77}, {'Id': 39, 'Value': 3, 'Size': 9},
      {'Id': 40, 'Value': 11, 'Size': 4}, {'Id': 41, 'Value': 3, 'Size': 46}, {'Id': 42, 'Value': 5, 'Size': 75},
      {'Id': 43, 'Value': 10, 'Size': 73},
      {'Id': 44, 'Value': 4, 'Size': 79}, {'Id': 45, 'Value': 7, 'Size': 21}, {'Id': 46, 'Value': 1, 'Size': 18},
      {'Id': 47, 'Value': 4, 'Size': 30}, {'Id': 48, 'Value': 4, 'Size': 68}
     ], 87, '{}']
]


class MyTestCase(unittest.TestCase):

    def test_bestPathFinder(self):
        for case in test_cases:
            bagSize = case[0]
            items = case[1]
            expectedMaxValue = case[2]
            expectedItems = case[3]
            maxValue, items = bestPathFinder(bagSize, items)
            self.assertEqual(maxValue, expectedMaxValue)
            # self.assertEqual(items, expectedItems)


if __name__ == '__main__':
    unittest.main()
